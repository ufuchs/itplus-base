package oops

// app
const (
	_ = iota + 100
	EC_PARAM_CONFIG_DIR_EMPTY
	EC_NO_CONFIGFILES_FOUND
	EC_MISSING_TIMELINE_NAME
	EC_TIMELINE_FOR_XY_DOESNT_EXIST
	EC_MISSING_ENV_ITPLUS_HOME
)

// base
const (
	_ = iota + 200
	EC_MISSING_SEP
	EC_TO_MANY_SEP
	EC_MISSPLACED_SEP_L
	EC_MISSPLACED_SEP_R
	EC_MISSING_VALUE
)

// device
const (
	_ = iota + 300
	EC_SAME_DEVICE_EXISTS
	EC_DEVICE_DOESNT_EXISTS
	EC_UPDATE_DEVICE_HAS_NO_ID
	EC_DEVICE_DATA_ARE_EQUAL
)

const (
	_ = iota + 400
	EC_1
	EC_2
)

// discovery
const (
	_ = iota + 500
	EC_MISSING_NETWORK_INTERFACE
)

var EC_Messages = map[int]string{
	EC_MISSING_ENV_ITPLUS_HOME: "Environment 'export ITPLUS_HOME=/usr/local/itplus' is missing.",
}

var ErrMissingEnvironmentITPLUS_HOME = func() error {
	errno := EC_MISSING_ENV_ITPLUS_HOME
	msg := getErrorDescription(errno)
	return &Err{errno, msg}
}

// GetDescription returns the corresponding verbal description of 'code'
func getErrorDescription(errno int) string {
	desc, ok := EC_Messages[errno]
	if !ok {
		desc = UNKNOWN_ERRNO
	}
	return desc
}
