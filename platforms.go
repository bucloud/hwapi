package hwapi

import (
	"encoding/json"
	"fmt"
)

//List of platforms
type PlatformList struct {
	List []*Platform `json:"list"` //List
}

//A CDN platform which provides an end-user service
type Platform struct {
	ID           int       `json:"id"`           //The platform id
	Code         string    `json:"code"`         //The platforms product code
	Capabilities []*string `json:"capabilities"` //capabilities
	Name         string    `json:"name"`         //The friendly name of the platform
	Type         string    `json:"type"`         //The type of platform this is (DELIVERY, INGEST, etc.)
	Available    bool      `json:"available"`    //whether or not the platform is available for the given account
}

//List all platforms enabled on the specified account
//Path /api/v1/accounts/{account_hash}/platforms
func (api *hwapi) GetPlatforms(accountHash string) (*PlatformList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/platforms", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &PlatformList{}
	return al, json.Unmarshal(r.body, al)
}
