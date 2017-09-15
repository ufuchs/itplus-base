package fcc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// /////////////////////////////////////////////////////////////////////////////
// helper
// /////////////////////////////////////////////////////////////////////////////

//
// GetWGFromContext
//
func GetWGFromContext(ctx context.Context) (*sync.WaitGroup, error) {
	wg, ok := ctx.Value(0).(*sync.WaitGroup)
	if !ok {
		return nil, errors.New("==> Failed to get the work group")
	}
	return wg, nil
}

//
//
//
func getFreeTCPPort() (string, error) {

	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}

	defer l.Close()

	addr := l.Addr().String()

	i := strings.LastIndex(addr, ":")

	//ip, _, err := net.SplitHostPort(req.RemoteAddr)

	return addr[i+1:], nil

}

//
//
//
func GetPortNumber(appPort int) (aPort string, iPort int, err error) {

	if appPort > 0 {
		iPort = appPort
		aPort = strconv.Itoa(appPort)
	} else {
		aPort, err = getFreeTCPPort()
		if err != nil {
			return
		}
		iPort, err = strconv.Atoi(aPort)
	}

	return

}

//
//
//
func GetHostname() (hostname string, err error) {

	for i := 0; i < 3; i++ {

		if hostname, err = os.Hostname(); err == nil {
			break
		}

		time.Sleep(1 * time.Second)

	}

	return
}

//
//
//
func Fatal(a ...interface{}) {
	fmt.Println(a...)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v...))
	os.Exit(1)
}

//
//
//
func RemoveFiles(profiles []string, dir string, t *testing.T) {

	var filterDirsGlob = func(dir, suffix string) ([]string, error) {
		return filepath.Glob(filepath.Join(dir, suffix))
	}

	for _, profile := range profiles {

		files, err := filterDirsGlob(os.TempDir(), profile+"*.yml")
		if err != nil {
			t.Error(err)
		}
		for _, file := range files {
			err := os.Remove(file)
			if err != nil {
				t.Error(err)
			}

		}

	}

}

//
//
//
func UnmarshalDeviceConfigJSON(name string) (*Configuration, error) {

	d := &Configuration{}
	if err := json.Unmarshal([]byte(name), d); err != nil {
		return nil, err
	}

	return d, nil

}

//
//
//
func UnmarshalDeviceConfigYML(name string) (*Configuration, error) {

	d := &Configuration{}
	if err := yaml.Unmarshal([]byte(name), d); err != nil {
		return nil, err
	}

	return d, nil

}
