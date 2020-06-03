package hwapi

import (
	"encoding/json"
	"fmt"
)

type POPs struct {
	List []*POP `json:"list"` //list
}

//A representation of a point of presence, also called a data center
type POP struct {
	ID     int    `json:"id"`     //The id of the pop
	Code   string `json:"code"`   //The code use to refer to the POP
	Name   string `json:"name"`   //The friendly name of the POP
	Group  string `json:"group"`  //group
	Region string `json:"region"` //region
}

//IPs
type IPs struct {
	List []*string `json:"list"` //list
}

//Returns the list of IPs used by the CDN
//Path /api/v1/ips
func (api *hwapi) GetIPs() (*IPs, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/ips"),
		},
	)
	if e != nil {
		return nil, e
	}
	i := &IPs{}
	return i, json.Unmarshal(r.body, i)
}

//Returns the list of POPs (Point of Presence) on the CDN.
//Path /api/v1/pops
func (api *hwapi) GetPoPs() (*POPs, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/pops"),
		},
	)
	if e != nil {
		return nil, e
	}
	i := &POPs{}
	return i, json.Unmarshal(r.body, i)
}
