package core_api_status

/*
CoreApiStatus

Possible enum values:

1=Private

2=Public
*/

type CoreApiStatus int

const (
	Private CoreApiStatus = 1
	Public  CoreApiStatus = 2
)
