package oops

import "fmt"

type (

	// Used in
	Err struct {
		ErrNo int
		Msg   string
	}

	// Used in REST
	AppErrorDTO struct {
		Message     string // map Error object to string
		Description string
	}
)

const APP = "ITPLUS-%0.4d: %v"
const UNKNOWN_ERRNO = "Unknown error code"

//
//
//
func NewErr(errno int, msg string) error {

	return &Err{
		ErrNo: errno,
		Msg:   msg,
	}

}

//
//
//
func (e *Err) Error() string {
	return fmt.Sprintf(APP, e.ErrNo, e.Msg)
}

func (err Err) errorNumber() int {
	return err.ErrNo
}

// Split multiple (variadic) return values into a slice of values
// in this case, where [0] = value and [1] = the error message
func split(args ...interface{}) []interface{} {
	return args
}
