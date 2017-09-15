package zvous

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ufuchs/zeroconf"
)

type (
	//
	ServiceEntry struct {
		ID        int
		Discharge bool
		Itype     InterfaceType
		Missing   int
		zeroconf.ServiceEntry
	}

	//
	AvailableServices struct {
		MaxID int
		List  []*ServiceEntry
	}
)

//
//
//
func NewServiceEntry(entry zeroconf.ServiceEntry, id int) *ServiceEntry {
	return &ServiceEntry{
		ID:           id,
		ServiceEntry: entry,
		Itype:        NET,
	}
}

//
// Return a matching discovered ServiceEntry by hostname
//
func (e *ServiceEntry) SearchMatchingEntry(discovered []zeroconf.ServiceEntry) *zeroconf.ServiceEntry {
	for _, entry := range discovered {
		if e.HostName == entry.HostName {
			return &entry
		}
	}
	return nil

}

//
// extractInstance
//
func (e *ServiceEntry) ExtractInstance() string {

	instance := e.Instance
	idx := strings.LastIndex(instance, "\\") + 1
	return instance[idx+1:]

}

//
// extractConn
//
func (e *ServiceEntry) ExtractConn() string {

	var addr string

	if len(e.AddrIPv4) > 0 {
		addr = e.AddrIPv4[0].String()
	}

	return addr + ":" + strconv.Itoa(e.Port)

}

//
// extractHostname
//
func (e *ServiceEntry) ExtractHostname() string {

	// hostname := entrie.HostName
	// i := strings.Index(hostname, ".")
	// return hostname[:i]

	instance := e.Instance
	idx := strings.LastIndex(instance, "\\") + 1
	return instance[idx+1:]

}

//
//
//
func (e *ServiceEntry) GetIdentifier() string {
	return fmt.Sprintf("%v:%v | %v", e.ExtractHostname(), e.Port, e.ID)
}

////////////////////////////////////////////////////////////////////////////////

//
//
//
func NewAvailableServices() *AvailableServices {
	return &AvailableServices{
		MaxID: 1,
		List:  []*ServiceEntry{},
	}
}

//
//
//
func (l *AvailableServices) Add(entry zeroconf.ServiceEntry) {
	e := NewServiceEntry(entry, l.MaxID)
	fmt.Printf("==> Zeroconf   : %v:%v - added\n", e.ExtractHostname(), e.Port)
	l.List = append(l.List, e)
	l.MaxID++
}

//
// Adds only new discovered devices.
// If a device got lost and has recovered with a new port number
// it will be registered later.
//
func (l *AvailableServices) AddFrom(discovered []zeroconf.ServiceEntry) {
	for _, entrie := range discovered {
		if l.Contains(&entrie) == nil {
			l.Add(entrie)
		}
	}
}

//
//
//
func (l *AvailableServices) Clear() {
	l.List = l.List[:0]
}

//
//
//
func (l *AvailableServices) Contains(entrie *zeroconf.ServiceEntry) *ServiceEntry {

	for i, a := range l.List {
		if equals(&a.ServiceEntry, entrie) {
			return l.List[i]
		}
	}
	return nil
}

//
//
//
func (l *AvailableServices) DropDischarged() {

	var newList = l.List[:0]
	for _, service := range l.List {
		if service.Discharge {
			fmt.Printf("==> Zeroconf   : %v:%v - removed\n", service.ExtractHostname(), service.Port)
			continue
		}
		newList = append(newList, service)
	}
	l.List = newList

}

//
//
//
func equals(s, o *zeroconf.ServiceEntry) bool {
	if o == nil {
		return false
	}
	if o == s {
		return true
	}
	return s.HostName == o.HostName &&
		s.Service == o.Service

}
