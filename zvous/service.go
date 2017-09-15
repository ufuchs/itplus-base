package zvous

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ufuchs/zeroconf"
)

var wait = []int{2, 10}

type InterfaceType uint8

// Options for IPType.
const (
	USB = 0x01
	NET = 0x02
)

type (

	//
	ZCBrowserService struct {
		quit       chan struct{}
		browser    *Browser
		discovered *AvailableServices
		Out        chan []*ServiceEntry
	}

	browseResult struct {
		discovered []zeroconf.ServiceEntry
		err        error
	}
)

//
// NewDiscoverTCPService
//
func NewZCBrowserService(serviceName string, t zeroconf.IPType, timeout time.Duration) *ZCBrowserService {

	d := &ZCBrowserService{
		quit:       make(chan struct{}),
		browser:    NewBrowser(serviceName, t, timeout),
		discovered: NewAvailableServices(),
		Out:        make(chan []*ServiceEntry),
	}

	go d.Browse()

	return d

}

//
// DecMissing
//
func (s *ZCBrowserService) DecMissing(missing int, i int) {
	if missing > 0 {
		s.discovered.List[i].Missing = missing - 1
	}
}

//
// Close
//
func (s *ZCBrowserService) Close() {
	close(s.quit)
	return
}

//
// random
//
func random(min, max int) time.Duration {
	rand.Seed(time.Now().Unix())
	return time.Duration(rand.Intn(max-min) + min)
}

//
// Browse
//
func (s *ZCBrowserService) Browse() error {

	var (
		next       time.Time
		browseDone chan browseResult
		fetchDelay time.Duration
		startFetch <-chan time.Time
		onStart    = true
	)

	s.discovered.Clear()

	for {

		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}

		if browseDone == nil {
			startFetch = time.After(fetchDelay)
		}

		select {

		case <-s.quit:
			s.browser.Close()
			close(s.Out)

			fmt.Println("==> Zeroconf   : finalize browser")
			return nil

		case <-startFetch:

			browseDone = make(chan browseResult, 1)

			go func() {
				discovered, err := s.browser.Browse()
				browseDone <- browseResult{discovered, err}
			}()

		case result := <-browseDone:

			browseDone = nil

			s.discovered.DropDischarged() // moved to here

			s.discovered.AddFrom(result.discovered)

			for i, entry := range s.discovered.List {

				var missing = entry.Missing
				var found *zeroconf.ServiceEntry

				// Check if the device got temporarily out of range
				if found = entry.SearchMatchingEntry(result.discovered); found == nil {
					s.discovered.List[i].Missing = missing + 1
					if missing+1 == 10 {
						s.discovered.List[i].Discharge = true
					}
					fmt.Printf("==> Zeroconf   : %v:%v - missing response %v\n", entry.ExtractHostname(), entry.Port, missing+1)
					continue
				}

				if entry.Port != found.Port {
					s.discovered.List[i].Discharge = true
					s.discovered.Add(*found)
					continue
				}

				s.DecMissing(missing, i)

			}

			//s.discovered.DropDischarged()

			s.Out <- append([]*ServiceEntry(nil), s.discovered.List...)

			var nxt time.Duration
			if !onStart {
				nxt = random(wait[0], wait[1])
			} else {
				nxt = 1
				onStart = false
			}

			nxt = random(wait[0], wait[1])
			next = time.Now().Add(nxt * time.Second)

		}

	}

}
