package hwapi

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

// SearchResult Full text search for entities tied to an account
//Path /api/v1/accounts/{account_hash}/search
type SearchResult struct {
	Hosts      []*HostName                `json:"hosts"`
	Hostnames  []*HostName                `json:"hostnames"`
	Origins    []*Origin                  `json:"origins"`
	HcsTenants []*map[string]string       `json:"hcsTenants"` //HCS is outof support
	Accounts   []*AccountInSearchResponse `json:"accounts"`
	Users      []*User                    `json:"users"`
}

// AccountInSearchResponse accountInfo in search result
type AccountInSearchResponse struct {
	DisplayName string `json:"displayName"`
	AccountName string `json:"accountName"`
	AccountHash string `json:"accountHash"`
	Parent      string `json:"parent"`
	GrandParent string `json:"grandparent"`
	ID          int    `json:"id"`
}

// UnmarshalJSON re-format string number to number
func (c *SearchResult) UnmarshalJSON(b []byte) error {
	type t SearchResult
	r := &t{}
	if err := json.Unmarshal(regexp.MustCompile(`"([0-9]+)"`).ReplaceAll(b, []byte("$1")), r); err != nil {
		return err
	}
	// setDefaultField(conf, reflect.StructTag(""))
	*c = (SearchResult)(*r)
	return nil
}

// Search accounts? hosts? origins? certificates
func (api *HWApi) Search(accountHash string, search string, maxResults int) (*SearchResult, error) {
	if maxResults == 0 {
		maxResults = 5
	}
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/search", accountHash),
			Query: map[string]string{
				"search":     search,
				"maxResults": strconv.Itoa(maxResults),
			},
		},
	)
	if e != nil {
		return nil, e
	}
	al := &SearchResult{}
	return al, json.Unmarshal(r.body, al)
}
