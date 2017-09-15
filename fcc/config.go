package fcc

import (
	"encoding/json"
)

type (
	OutputScheme struct {
		ShowUnknownDevices bool `yaml:"show_unknown_devices" json:"show_unknown_devices"`
		Interval           int  `yaml:"interval" json:"interval"`
	}

	Configuration struct {
		Name         string       `yaml:"name" json:"name"`
		ModifiedOn   string       `json:"modifiedOn"`
		OutputScheme OutputScheme `json:"outputscheme"`
		Devices      Devices      `json:"devices"`
	}

	IConfiguration interface {
		AddDeviceByJSON([]byte) (int, error)
		GetDevice(id int) (Device, error)
		UpdateDeviceByJSON([]byte) error
		RemoveDevice(int) error
	}

	Timestamped struct {
		When int64  `json:"when"`
		Addr int    `json:"addr"`
		Num  int    `json:"num"`
		Data string `json:"data"`
	}
)

//
//
//
func (p *Configuration) addDevice(newOne Device) (int, error) {

	if err := newOne.Validate(); err != nil {
		return -1, err
	}

	newOne.Id = p.currId()

	for _, d := range p.Devices {
		if d == newOne {
			return -1, ErrSameDeviceExists()
		}
	}

	newOne.Id++

	devices := append([]Device(nil), p.Devices...)
	p.Devices = append(devices, newOne)

	return newOne.Id, nil

}

//
//
//
func (p *Configuration) AddDeviceByJSON(newOne []byte) (int, error) {

	d := Device{}
	if err := json.Unmarshal(newOne, &d); err != nil {
		return -1, err
	}

	return p.addDevice(d)

}

//
//
//
func (p *Configuration) GetDevice(id int) (Device, error) {

	for _, d := range p.Devices {
		if d.Id == id {
			return d, nil
		}
	}

	return Device{}, ErrDeviceDoesntExists(id)

}

//
//
//
func (p *Configuration) updateDevice(updateOne Device) error {

	if updateOne.Id == 0 {
		return ErrUpdateDeviceWithoutAnyID()
	}

	if err := updateOne.Validate(); err != nil {
		return err
	}

	// this code INSERTS A NEW ID!

	// err := p.RemoveDevice(updateOne.Id)
	// if err != nil {
	// 	return err
	// }
	//
	// p.addDevice(updateOne)

	for i, d := range p.Devices {

		if d.Id == updateOne.Id {
			if p.Devices[i] == updateOne {
				return ErrUpdateDeviceDataAreEqual()
			}
			p.Devices[i] = updateOne
			break
		}

	}

	return nil

}

//
//
//
func (p *Configuration) UpdateDeviceByJSON(updateOne []byte) error {

	d := Device{}
	if err := json.Unmarshal(updateOne, &d); err != nil {
		return err
	}

	return p.updateDevice(d)

}

//
// Remove
//
func (p *Configuration) RemoveDevice(id int) error {

	if _, err := p.GetDevice(id); err != nil {
		return err
	}

	res := FilterDevices(p.Devices, func(d Device) bool {
		return id != d.Id
	})

	p.Devices = res

	return nil

}

//
// NewId
//
func (p *Configuration) currId() int {

	var id = -1

	for _, d := range p.Devices {
		if d.Id > id {
			id = d.Id
		}
	}

	return id

}
