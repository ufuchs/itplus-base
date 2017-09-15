package oops

import (
	"errors"
	"fmt"
	"testing"
)

// AppError ////////////////////////////////////////////////////////////////////

func TestNewAppErrorDTO(t *testing.T) {

	err := errors.New("open cafebabe: no such file or directory")

	e := &AppErrorDTO{
		Message:     err.Error(),
		Description: "GetErrorDescription(EC_MISSING_VALUE)",
	}

	fmt.Println(e)

}

// Err /////////////////////////////////////////////////////////////////////////

//
//
//
func TestNewErr(t *testing.T) {

	actual := "abc"

	e := NewErr(EC_MISSING_VALUE, actual)

	fmt.Println(e.Error())

}
