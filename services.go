package hwapi

import (
	"encoding/json"
	"fmt"
)

//A billable service enabled on an account or host
type Service struct {
	Id          int    `json:"id"`          //The service id, used for setting services on an entity
	Name        string `json:"name"`        //The friendly name of the service
	Description string `json:"description"` //Description of the service
	Type        string `json:"type"`        //type of service
}

type Services struct {
	List []*Service
}

//Get available services list
func (api *hwapi) GetServices(accountHash string) (*Services, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/services", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Services{}
	return al, json.Unmarshal(r.body, al)
}
