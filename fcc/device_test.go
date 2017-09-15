package fcc

import (
	"encoding/json"
	"sort"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

var TX25TP_IT = `
{
    "id": 4,
    "num": 7,
    "addr": 24,
    "type": "TX25TP-IT",
    "alias": "R 2.7 - Fenster:Draußen - AUCOTEAM Nordseite",
    "absent": false,
    "locked": false,
    "lon": 7.1,
    "lat": 7.2,
    "alt": 7.3
}
`

var sortTest = `
devices:
- id: 2
  num: 2
  addr: 84
  type: JeeLink
  alias: Schreibtisch
  absent: false
  locked: true
  lon: 0
  lat: 0
  alt: 0
- id: 1
  num: 1
  addr: 46
  type: TX29TDH-IT
  alias: Bad
  absent: false
  locked: false
  lon: 1.1
  lat: 1.2
  alt: 1.3
- id: 3
  num: 2
  addr: 42
  type: TX29TDH-IT
  alias: Küche
  absent: false
  locked: false
  lon: 2.1
  lat: 2.2
  alt: 2.3
`

func unmarshalDevice() (*Device, error) {

	d := &Device{}
	if err := json.Unmarshal([]byte(TX25TP_IT), d); err != nil {
		return nil, err
	}

	return d, nil

}

func unmarshalDevices() (Devices, error) {

	type X struct {
		Devices Devices
	}

	var x = &X{}

	if err := yaml.Unmarshal([]byte(sortTest), x); err != nil {
		return nil, err
	}

	return x.Devices, nil

}

//
//
//
func TestSortDevices(t *testing.T) {

	devices, err := unmarshalDevices()
	if err != nil {
		t.Error("unmarshalDevice()", err)
	}

	sort.Sort(devices)

}

//
//
//
func TestValidate_TX25TP_Alias(t *testing.T) {

	device, err := unmarshalDevice()
	if err != nil {
		t.Error("unmarshalDevice()", err)
	}

	err = device.Validate_TX25TP_Alias()
	if err != nil {
		t.Error("Validate_TX25TP_Alias()", err)
	}

	var params = map[string]string{
		"MissingValue":   "",
		"MissingSep":     "test",
		"ToManySep":      "test::test",
		"MissplacedSepL": ":test",
		"MissplacedSepR": "test:",
	}

	for k, v := range params {
		switch k {
		case "MissingValue":
			device.Alias = v
			err = device.Validate_TX25TP_Alias()
		case "MissingSep":
			device.Alias = v
			err = device.Validate_TX25TP_Alias()
		case "ToManySep":
			device.Alias = v
			err = device.Validate_TX25TP_Alias()
		case "MissplacedSepL":
			device.Alias = v
			err = device.Validate_TX25TP_Alias()
		case "MissplacedSepR":
			device.Alias = v
			err = device.Validate_TX25TP_Alias()

		}
		if err == nil {
			t.Errorf("Validate_TX25TP_Alias() should throw ==> '%v'", err)
		}
	}

}

//
//
//
func TestValidate(t *testing.T) {

	// setup

	device, err := unmarshalDevice()
	if err != nil {
		t.Error("unmarshalDevice()", err)
	}

	err = device.Validate()
	if err != nil {
		t.Error("Validate()", err)
	}

}

//
//
//
func TestDiffDevices_A_gt_B(t *testing.T) {

	// setup

	a, err := unmarshalDevices()
	if err != nil {
		t.Error("unmarshalDevice()", err)
	}

	sameAsA, err := unmarshalDevices()
	if err != nil {
		t.Error("unmarshalDevice()", err)
	}

	var id = 1
	b := FilterDevices(sameAsA, func(d Device) bool {
		return id != d.Id
	})

	if len(a) == len(b) {
		t.Error("DiffDevices(): failed on setup")
	}

	// test

	diff := DiffDevices(a, b)

	// check

	if len(diff) != 1 {
		t.Error("DiffDevices(): expected one device difference slice")
	}

}

func TestDiffDevices_A_lt_B(t *testing.T) {

	// setup

	sameAsB, err := unmarshalDevices()
	if err != nil {
		t.Error("unmarshalDevice()", err)
	}

	b, err := unmarshalDevices()
	if err != nil {
		t.Error("unmarshalDevice()", err)
	}

	var id = 1
	a := FilterDevices(sameAsB, func(d Device) bool {
		return id != d.Id
	})

	if len(a) == len(b) {
		t.Error("DiffDevices(): failed on setup")
	}

	// test

	diff := DiffDevices(a, b)

	// check

	if len(diff) != 0 {
		t.Error("DiffDevices(): expected no difference")
	}

}

//
//
//
func TestFilterDevice(t *testing.T) {

	device, err := unmarshalDevice()
	if err != nil {
		t.Error("unmarshalDevice()", err)
	}

	devices := []Device{}
	devices = append(devices, *device)

	// id 8 has no element
	var id = 8
	res := FilterDevices(devices, func(d Device) bool {
		return id != d.Id
	})

	if len(res) != 1 {
		t.Error("FilterDevices(): Failed. Element should be left")
	}

	// id 8 has no element
	id = 4
	res = FilterDevices(devices, func(d Device) bool {
		return id != d.Id
	})

	if len(res) != 0 {
		t.Error("FilterDevices(): Failed. Still one element left")
	}

}
