package fcc

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

var testConfigDir = os.TempDir()

var TestMemberNames = []string{"aleta", "home"}

// /////////////////////////////////////////////////////////////////////////////
// tests
// /////////////////////////////////////////////////////////////////////////////

//
//
//
func TestWrite(t *testing.T) {

	defer RemoveFiles(TestMemberNames, testConfigDir, t)

	dao := NewConfigDAO(testConfigDir)

	for k, v := range AletaYML {

		cfg, err := UnmarshalDeviceConfigYML(v)
		if err != nil {
			t.Error("unmarshalDeviceConfigYML():", err)
		}

		if err = dao.WriteConfig(cfg, k); err != nil {
			t.Error("dao.writeConfig():", err)
		}

	}

}

//
//
//
func TestRetrieve(t *testing.T) {

	defer RemoveFiles(TestMemberNames, testConfigDir, t)

	dao := NewConfigDAO(testConfigDir)

	////
	// test, if file not exists
	////

	if _, err := dao.Retrieve("aleta"); err == nil {
		t.Error("Retrieve():", errors.New("should throw 'no such file or directory'"))
	}

	////
	// test, first create the files,
	//       second read them
	////

	for k, v := range AletaYML {

		cfg, err := UnmarshalDeviceConfigYML(v)
		if err != nil {
			t.Error("unmarshalDeviceConfigYML():", err)
		}

		// fmt.Println("---------------------")
		// fmt.Println(v)
		// fmt.Println(config)

		if err = dao.WriteConfig(cfg, k); err != nil {
			t.Error("dao.writeConfig():", err)
		}

	}

	// second read them
	for k, _ := range AletaYML {

		filename := fmt.Sprintf("aleta-%v", k)

		_, err := dao.Retrieve(filename)
		if err != nil {
			t.Error("dao.Retrieve():", err)
		}

	}

}
