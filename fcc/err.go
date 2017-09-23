package fcc

import (
	"fmt"

	"github.com/ufuchs/itplus/base/oops"
)

var EC_Messages map[int]string = map[int]string{
	oops.EC_MISSING_SEP:             "TX25TP, missing separator ':' in alias '%v'",
	oops.EC_TO_MANY_SEP:             "TX25TP, to many separator ':' in alias '%v'",
	oops.EC_MISSPLACED_SEP_L:        "TX25TP, missing value left of ':' in alias '%v'",
	oops.EC_MISSPLACED_SEP_R:        "TX25TP, missing value right of ':' in alias '%v'",
	oops.EC_MISSING_VALUE:           "TX25TP, parameter 'alias' is empty",
	oops.EC_SAME_DEVICE_EXISTS:      "A same device exists.",
	oops.EC_DEVICE_DOESNT_EXISTS:    "a device with 'Id = %v' doesn't exist.",
	oops.EC_UPDATE_DEVICE_HAS_NO_ID: "updating a device with ID = 0",
	oops.EC_DEVICE_DATA_ARE_EQUAL:   "No need to update, data are equal",
}

var (
	ErrMissingSep = func(alias string) error {
		errno := oops.EC_MISSING_SEP
		msg := fmt.Sprintf(EC_Messages[errno], alias)
		return &oops.Err{errno, msg}
	}

	ErrToManySep = func(alias string) error {
		errno := oops.EC_TO_MANY_SEP
		msg := fmt.Sprintf(EC_Messages[errno], alias)
		return &oops.Err{errno, msg}
	}

	ErrMissplacedSepL = func(alias string) error {
		errno := oops.EC_MISSPLACED_SEP_L
		msg := fmt.Sprintf(EC_Messages[errno], alias)
		return &oops.Err{errno, msg}
	}

	ErrMissplacedSepR = func(alias string) error {
		errno := oops.EC_MISSPLACED_SEP_R
		msg := fmt.Sprintf(EC_Messages[errno], alias)
		return &oops.Err{errno, msg}
	}

	ErrMissingValue = func() error {
		errno := oops.EC_MISSING_VALUE
		msg := EC_Messages[errno]
		return &oops.Err{errno, msg}
	}

	ErrSameDeviceExists = func() error {
		errno := oops.EC_SAME_DEVICE_EXISTS
		msg := getErrorDescription(errno)
		return &oops.Err{errno, msg}
	}

	ErrDeviceDoesntExists = func(value int) error {
		errno := oops.EC_DEVICE_DOESNT_EXISTS
		msg := fmt.Sprintf(getErrorDescription(errno), value)
		return &oops.Err{errno, msg}
	}

	ErrUpdateDeviceWithoutAnyID = func() error {
		errno := oops.EC_UPDATE_DEVICE_HAS_NO_ID
		msg := fmt.Sprint(getErrorDescription(errno))
		return &oops.Err{errno, msg}
	}

	ErrUpdateDeviceDataAreEqual = func() error {
		errno := oops.EC_DEVICE_DATA_ARE_EQUAL
		msg := fmt.Sprint(getErrorDescription(errno))
		return &oops.Err{errno, msg}
	}
)

// GetDescription returns the corresponding verbal description of 'code'
func getErrorDescription(errno int) string {
	desc, ok := EC_Messages[errno]
	if !ok {
		desc = oops.UNKNOWN_ERRNO
	}
	return desc
}
