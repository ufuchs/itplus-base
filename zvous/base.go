package zvous

const (
	SUCCESS         = 1
	TIMEDOUT        = 2
	FETCH_FAILED    = 3
	RESOLVER_FAILED = 4
)

//const SERVICENAME = "_measurement._itplus._tcp"

const (
	//	AVAHI_CONFIGURATION = "_configuration._itplus._tcp"
	//	AVAHI_CONF_TEXT     = "{ \"default\": \"%v\", \"offs\": %v }"
	AVAHI_DATA         = "_data1._tcp"
	AVAHI_MEASUREMENT  = "_itplus._measurement._tcp"
	AVAHI_STORAGE_TEXT = "ws://%v/ws"
)
