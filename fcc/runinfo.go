package fcc

type (
	//
	RunInfo struct {
		Profile            string
		ShowUnknownDevices bool
		Devices            map[int]Device
		DeviceAddr         map[int]int
		Events             map[int]*Event
	}
)

//
//
//
func NewRunInfoFromConfig(c *Configuration) *RunInfo {

	return &RunInfo{
		Profile:            c.Name,
		ShowUnknownDevices: c.OutputScheme.ShowUnknownDevices,
		DeviceAddr:         deviceAddressesToMap(c.Devices),
		Devices:            devicesToMap(c.Devices),
		Events:             eventsToMap(c.Devices),
	}

}

//
//
//
func eventsToMap(events []Device) map[int]*Event {
	result := make(map[int]*Event)

	return result
}

//
//
//
func deviceAddressesToMap(devices Devices) map[int]int {
	result := make(map[int]int, len(devices))
	for _, m := range devices {
		result[m.Addr] = m.Num
	}
	return result
}

//
//
//
func devicesToMap(devices Devices) map[int]Device {
	result := make(map[int]Device, len(devices))
	for _, attr := range devices {
		result[attr.Num] = attr
	}
	return result
}
