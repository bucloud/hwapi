package hwapi

import (
	"strconv"
)

//403,201 error code
type HCConnectError struct {
	code        int
	category    string
	description string
}
type HCRequestError struct {
	code        int
	category    string
	description string
}

type APIError struct {
	code        int
	category    string
	description string
}

func (e *HCConnectError) Error() string {
	return "Connect error " + strconv.Itoa(e.code) + " : " + e.description
}

func (e *HCRequestError) Error() string {
	return "Request error " + strconv.Itoa(e.code) + " : " + e.description
}

func (e *APIError) Error() string {
	return "API error " + strconv.Itoa(e.code) + " : " + e.description
}
