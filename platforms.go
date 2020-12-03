package hwapi

import (
	"encoding/json"
	"fmt"
)

// PlatformList List of platforms
type PlatformList struct {
	List []*Platform `json:"list"` //List
}

// Platform A CDN platform which provides an end-user service
type Platform struct {
	ID           int       `json:"id"`           //The platform id
	Code         string    `json:"code"`         //The platforms product code
	Capabilities []*string `json:"capabilities"` //capabilities
	Name         string    `json:"name"`         //The friendly name of the platform
	Type         string    `json:"type"`         //The type of platform this is (DELIVERY, INGEST, etc.)
	Available    bool      `json:"available"`    //whether or not the platform is available for the given account
}

// GetPlatforms List all platforms enabled on the specified account
//Path /api/v1/accounts/{account_hash}/platforms
func (api *HWApi) GetPlatforms(accountHash string) (*PlatformList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/platforms", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &PlatformList{}
	return al, json.Unmarshal(r.body, al)
}
