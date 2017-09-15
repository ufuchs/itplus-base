package fcc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type (
	//
	ConfigDAO struct {
		configDir string
	}
)

//
//
//
func NewConfigDAO(configDir string) *ConfigDAO {
	return &ConfigDAO{
		configDir: configDir,
	}
}

//
//
//
func (dc *ConfigDAO) Retrieve(configName string) (Configuration, error) {

	filename := path.Join(dc.configDir, configName)

	config := Configuration{}
	err := ReadYML(&config, filename+".yml")

	return config, err

}

func (dc *ConfigDAO) RetrieveJSON(configName string) (Configuration, error) {

	filename := path.Join(dc.configDir, configName)

	config := Configuration{}
	err := ReadJSON(&config, filename+".json")

	return config, err

}


//
//
//
func (dc *ConfigDAO) WriteConfig(d *Configuration, timeStamp string) error {

	if len(timeStamp) == 0 {
		timeStamp = time.Now().UTC().Format("2006-01-02T150405Z") // ISO 8601
		d.ModifiedOn = timeStamp
	}

	filename := path.Join(
		dc.configDir,
		fmt.Sprintf(
			"%v-%v.yml",
			d.Name,
			d.ModifiedOn))

	return WriteYML(d, filename)
}

//
//
//
func WriteYML(v interface{}, filename string) error {

	var (
		buf []byte
		err error
	)

	if buf, err = yaml.Marshal(v); err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf, 0644)

}

//
//
//
func WriteJSON(v interface{}, filename string) error {

	var (
		buf []byte
		err error
	)

	if buf, err = json.Marshal(v); err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf, 0644)

}


//
//
//
func ReadYML(v interface{}, filename string) error {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(raw, v)

}


//
//
//
func ReadJSON(v interface{}, filename string) error {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, v)

}


