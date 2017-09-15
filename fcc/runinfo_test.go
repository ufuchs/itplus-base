package fcc

import (
	"encoding/json"
	"path"
	"testing"
)

var TestMemberNames = []string{"aleta", "home"}

////////////////////////////////////////////////////////////////////////////////

//
//
//
func prepareRunInfo(dir string) error {

	for _, v := range AletaYML {

		// if k != "aleta" {
		// 	continue
		// }

		cfg, err := UnmarshalDeviceConfigYML(v)
		if err != nil {
			return err
		}

		filename := path.Join(testConfigDir, "config.json")
		if err = WriteJSON(cfg, filename); err != nil {
			return err
		}

	}

	return nil

}

//
//
//
func TestDeviceAddressesToMap(t *testing.T) {

	// setup
	var config = &Configuration{}
	json.Unmarshal([]byte(HomePureJSON), config)

	devices := config.Devices

	// test

	devAddr := deviceAddressesToMap(devices)

	// check

	for _, d := range devices {

		var num, ok = devAddr[d.Addr]
		if !ok {
			t.Errorf("deviceAddressesToMap(): failed => 'devAddr[%v]'", d.Addr)
			break
		}
		if num != d.Num {
			t.Errorf("deviceAddressesToMap(): failed => found '%v', expected '%v'", num, d.Num)
			break
		}

	}

}

//
//
//
func TestDevicesToMap(t *testing.T) {

	// setup

	var config = &Configuration{}
	json.Unmarshal([]byte(HomePureJSON), config)

	devices := config.Devices

	// test

	mapped := devicesToMap(devices)

	// check

	for _, da := range devices {

		var de, ok = mapped[da.Num]
		if !ok {
			t.Errorf("devicesToMap(): failed => 'mapped[%v]'", da.Num)
			break
		}
		if de != da {
			t.Errorf("devicesToMap(): failed => found '%v', expected '%v'", da, de)
			break
		}

	}

}

func TestNewRunInfo(t *testing.T) {

	//defer fcc.RemoveFiles(TestMemberNames, testConfigDir, t)

	if err := prepareRunInfo(testConfigDir); err != nil {
		t.Errorf("prepareRunInfo() failed => %v", err)
	}

	ri := NewRunInfo(testConfigDir)

	if ri == nil {
		t.Error("NewRunInfo() failed => shouldn't be nil")
	}

}
