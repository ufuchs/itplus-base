package fcc

import "strings"

type (
	Device struct {
		Id     int     `yaml:"id" json:"id"`
		Num    int     `yaml:"num" json:"num"`
		Addr   int     `yaml:"addr" json:"addr"`
		Type   string  `yaml:"type" json:"type"`
		Alias  string  `yaml:"alias" json:"alias"`
		Absent bool    `yaml:"absent" json:"absent"`
		Locked bool    `yaml:"locked" json:"locked"`
		Lon    float64 `yaml:"lon" json:"lon"`
		Lat    float64 `yaml:"lat" json:"lat"`
		Alt    float64 `yaml:"alt" json:"alt"`
	}

	Devices []Device

	deviceIterFunc func(Device) bool

	MeasurementDTO struct {
		Host           string  `json:"host"`
		Num            float32 `json:"num"`
		Alias          string  `json:"alias"`
		PhenomenonTime int64   `json:"phenomenontime"`
		Lon            float64 `json:"lon"`
		Lat            float64 `json:"lat"`
		Alt            float64 `json:"alt"`
		Temp           float64 `json:"temp"`
		Pressure       float64 `json:"pressure"`
		Humidity       float64 `json:"humidity"`
		LowBattery     bool    `json:"lowbattery"`
	}
)

func (slice Devices) Len() int           { return len(slice) }
func (slice Devices) Less(i, j int) bool { return slice[i].Id < slice[j].Id }
func (slice Devices) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

//
//
//
func FilterDevices(in Devices, filterFunc deviceIterFunc) Devices {

	//res := []Device{}
	// b := a[:0]
	filtered := in[:0]

	for _, attr := range in {

		if filterFunc(attr) {
			filtered = append(filtered, attr)
		}

	}

	return filtered

}

//
//
//
func DiffDevices(devices1, devices2 Devices) Devices {

	var diff = Devices{}
	var found bool

	for _, d1 := range devices1 {

		found = false

		for _, d2 := range devices2 {

			if d1.Id == d2.Id {
				found = true
				break
			}

		}

		if !found {
			diff = append(diff, d1)
		}

	}

	return diff

}

//
//
//
func (d *Device) Validate_TX25TP_Alias() (err error) {

	var sep string = ":"
	var part []string

	if len(d.Alias) == 0 {
		return ErrMissingValue()
	}

	if deli := strings.Index(d.Alias, sep); deli == -1 {
		return ErrMissingSep(d.Alias)
	}

	if part = strings.Split(d.Alias, sep); len(part) != 2 {
		return ErrToManySep(d.Alias)
	}

	if len(part[0]) == 0 {
		return ErrMissplacedSepL(d.Alias)
	}

	if len(part[1]) == 0 {
		return ErrMissplacedSepR(d.Alias)
	}

	return nil

}

//
//
//
func (d *Device) Validate() (err error) {
	if d.Type == "TX25TP-IT" {
		err = d.Validate_TX25TP_Alias()
	}
	return
}

//
//
//
// func (d *Device) ValidateEx() (result int32) {
//
// 	result = 0
//
// 	switch {
// 	case d.Num < 0:
// 		result |= 1
// 		break
// 	case len(d.Type) == 0:
// 		result |= 2
// 		break
// 	case len(d.Alias) == 0:
// 		result |= 4
// 		break
// 	}
//
// 	return result
// }
