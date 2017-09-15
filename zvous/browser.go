package zvous

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ufuchs/zeroconf"
)

type (

	//
	Browser struct {
		quit        chan struct{}
		options     zeroconf.ClientOption
		timeout     time.Duration
		serviceName string
	}
)

//
//
//
func NewBrowser(serviceName string, ipType zeroconf.IPType, timeOut time.Duration) *Browser {

	fmt.Printf("==> Zeroconf   : Discovery timeout is %v\n", timeOut)

	return &Browser{
		quit:        make(chan struct{}),
		timeout:     time.Second * timeOut,
		options:     zeroconf.SelectIPTraffic(ipType),
		serviceName: serviceName,
	}
}

//
//
//
func (d *Browser) Close() {
	close(d.quit)
	return
}

//
// Browse
//
func (d *Browser) Browse() ([]zeroconf.ServiceEntry, error) {

	var (
		discovered = []zeroconf.ServiceEntry{}
		err        error
		entrie     = make(chan *zeroconf.ServiceEntry, 1)
		resolver   *zeroconf.Resolver
	)

	ctx, _ := context.WithTimeout(context.Background(), d.timeout)
	if resolver, err = zeroconf.NewResolver(d.options); err != nil {
		return nil, err
	}

	if err = resolver.Browse(ctx, d.serviceName, "local.", entrie); err != nil {
		return nil, err
	}

	for {
		select {
		case <-d.quit:
			log.Println("==> Discoverer : finalized")
			return nil, nil
		case <-ctx.Done():
			return discovered, ctx.Err()
		case e := <-entrie:
			if e != nil {
				discovered = append(discovered, *e)
			}
		}
	}

}
