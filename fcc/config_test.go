package fcc

import (
	"testing"
)

// /////////////////////////////////////////////////////////////////////////////
// helper
// /////////////////////////////////////////////////////////////////////////////

func newConfig() *Configuration {

	return &Configuration{
		Name: "aleta",
		OutputScheme: OutputScheme{
			ShowUnknownDevices: true,
			Interval:           60,
		},
		Devices: Devices{
			{
				Id:     1,
				Num:    8,
				Addr:   77,
				Type:   "",
				Alias:  "Unter Sofa",
				Absent: false,
				Locked: false,
				Lon:    52.13,
				Lat:    13.28,
				Alt:    33.0,
			},
		},
	}

}

var deviceExist = Device{
	Id:     1,
	Num:    8,
	Addr:   77,
	Type:   "",
	Alias:  "Unter Sofa",
	Absent: false,
	Locked: false,
	Lon:    52.13,
	Lat:    13.28,
	Alt:    33.0,
}

var updateOne = Device{
	Id:     1,
	Num:    9,
	Addr:   77,
	Type:   "",
	Alias:  "Unter Sofa",
	Absent: false,
	Locked: false,
	Lon:    52.13,
	Lat:    13.28,
	Alt:    120.0,
}

var newOne = Device{
	Id:     1,
	Num:    10,
	Addr:   78,
	Type:   "",
	Alias:  "Unter Kühlschrank",
	Absent: false,
	Locked: false,
	Lon:    52.13,
	Lat:    13.28,
	Alt:    33.0,
}

// /////////////////////////////////////////////////////////////////////////////
// tests
// /////////////////////////////////////////////////////////////////////////////

func TestCurrId(t *testing.T) {

	cfg := newConfig()

	for i := 0; i < 5; i++ {

		id := cfg.currId()

		if id != 1 {
			t.Error("currId(): should return 1")
		}

	}

}

//
//
//
func TestAddDeviceRejected(t *testing.T) {

	cfg := newConfig()

	id, err := cfg.addDevice(deviceExist)
	if err == nil {
		t.Error("AddDevice(): should throw 'A same device exists.'")
	}
	if id != -1 {
		t.Error("AddDevice(): should return -1, if device exists")
	}

}

//
//
//
func TestAddDevice(t *testing.T) {

	cfg := newConfig()

	id, err := cfg.addDevice(newOne)
	if err != nil {
		t.Error("AddDevice():", err)
	}
	if id == -1 {
		t.Error("AddDevice(): should return a 'Id' greather -1")
	}

	if len(cfg.Devices) == 1 {
		t.Error("AddDevice(): adding a new device failed")
	}

}

func TestAddDeviceByJSONRejected(t *testing.T) {

	var withSytaxError = `{
			"id":_ 1,
			"num": 200,
			"addr": 7,
			"type": "TX29TDH-IT",
			"alias": "R 2.7 - Schreibtisch",
			"absent": false,
			"locked": false,
			"lon": 19.31,
			"lat": 19.32,
			"alt": 19.33
	}`

	cfg := newConfig()
	_, err := cfg.AddDeviceByJSON([]byte(withSytaxError))
	if err == nil {
		t.Error("UpdateDeviceByJSON(): should throw => 'invalid character '_' looking for beginning of value'")
	}

}

//
//
//
func TestAddDeviceByJSON(t *testing.T) {

	var newOne = `{
			"id": 1,
			"num": 200,
			"addr": 7,
			"type": "TX29TDH-IT",
			"alias": "R 2.7 - Schreibtisch",
			"absent": false,
			"locked": false,
			"lon": 19.31,
			"lat": 19.32,
			"alt": 19.33
	}`

	cfg := newConfig()

	id, err := cfg.AddDeviceByJSON([]byte(newOne))
	if err != nil {
		t.Error("AddDeviceByJSON():", err)
	}
	if id == -1 {
		t.Error("AddDeviceByJSON(): should return a 'Id' greather -1")
	}

	if len(cfg.Devices) == 1 {
		t.Error("AddDeviceByJSON(): adding a new device failed")
	}

}

//
//
//
func TestUpdateDeviceRejected(t *testing.T) {

	var withIdEquals0 = Device{

		Id:     0,
		Num:    8,
		Addr:   77,
		Type:   "TX29TDH-IT",
		Alias:  "Unter Sofa",
		Absent: false,
		Locked: false,
		Lon:    52.13,
		Lat:    13.28,
		Alt:    33.0,
	}

	cfg := newConfig()
	err := cfg.updateDevice(withIdEquals0)
	if err == nil {
		t.Error("UpdateDeviceByJSON(): should throw => 'updating a device with ID = 0'")
	}

	var newOne = Device{

		Id:     1,
		Num:    200,
		Addr:   7,
		Type:   "TX29TDH-IT",
		Alias:  "Unter Sofa",
		Absent: false,
		Locked: false,
		Lon:    52.13,
		Lat:    13.28,
		Alt:    33.0,
	}

	cfg = newConfig()
	err = cfg.updateDevice(newOne)
	if err != nil {
		t.Errorf("UpdateDeviceByJSON(): %v", err)
	}

	var sameOne = Device{

		Id:     1,
		Num:    200,
		Addr:   7,
		Type:   "TX29TDH-IT",
		Alias:  "Unter Sofa",
		Absent: false,
		Locked: false,
		Lon:    52.13,
		Lat:    13.28,
		Alt:    33.0,
	}

	err = cfg.updateDevice(sameOne)
	if err == nil {
		t.Errorf("UpdateDeviceByJSON(): should throw ==> 'No need to update, data are equal'")
	}

}

//
//
//
func TestUpdateDevice(t *testing.T) {

	cfg := newConfig()

	err := cfg.updateDevice(updateOne)
	if err != nil {
		t.Error("UpdateDevice():", err)
	}

}

func TestUpdateDeviceByJSONRejected(t *testing.T) {

	var withSytaxError = `{
			"id":_ 1,
			"num": 200,
			"addr": 7,
			"type": "TX29TDH-IT",
			"alias": "R 2.7 - Schreibtisch",
			"absent": false,
			"locked": false,
			"lon": 19.31,
			"lat": 19.32,
			"alt": 19.33
	}`

	cfg := newConfig()
	err := cfg.UpdateDeviceByJSON([]byte(withSytaxError))
	if err == nil {
		t.Error("UpdateDeviceByJSON(): should throw => 'invalid character '_' looking for beginning of value'")
	}

}

//
//
//
func TestUpdateDeviceByJSON(t *testing.T) {

	var theUpdate = `{
			"id": 1,
			"num": 200,
			"addr": 7,
			"type": "TX29TDH-IT",
			"alias": "R 2.7 - Schreibtisch",
			"absent": false,
			"locked": false,
			"lon": 19.31,
			"lat": 19.32,
			"alt": 19.33
	}`

	cfg := newConfig()

	beforeUpdate, err := cfg.GetDevice(1)
	if err != nil {
		t.Error("UpdateDevice():", err)
	}

	err = cfg.UpdateDeviceByJSON([]byte(theUpdate))
	if err != nil {
		t.Error("UpdateDevice():", err)
	}

	var afterUpdate Device
	afterUpdate, err = cfg.GetDevice(1)
	if err != nil {
		t.Error("UpdateDevice():", err)
	}

	if beforeUpdate == afterUpdate {
		t.Error("UpdateDeviceByJSON(): should be differend")
	}

}

func TestGetDeviceRejected(t *testing.T) {

	cfg := newConfig()

	_, err := cfg.GetDevice(-1)
	if err == nil {
		t.Error("GetDevice(): failed to get device 'Id = 1'")
	}

}

//
//
//
func TestGetDevice(t *testing.T) {

	cfg := newConfig()

	_, err := cfg.GetDevice(1)
	if err != nil {
		t.Error("GetDevice(): failed to get device 'Id = 1'")
	}

}

//
//
//
func TestRemoveDevice(t *testing.T) {

	cfg := newConfig()

	err := cfg.RemoveDevice(1)
	if err != nil {
		t.Error("RemoveDevice(): failed")
	}

	err = cfg.RemoveDevice(-1)
	if err == nil {
		t.Error("RemoveDevice(): failed")
	}

	if len(cfg.Devices) != 0 {
		t.Error("RemoveDevice(): failed")
	}

}
