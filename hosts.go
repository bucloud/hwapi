package hwapi

import (
	"encoding/json"
	"fmt"
)

//HostList list of host
type HostList struct {
	List []*Host `json:"list"` //list
}

// Hosts return hosts map directly
func (hl *HostList) Hosts() map[string]*Host {
	r := map[string]*Host{}
	for _, h := range hl.List {
		r[h.HashCode] = h
	}
	return r
}

// CloneHost hostname
type CloneHost struct {
	Name      string   `json:"name"`      //The name of the host
	Hostnames []string `json:"hostnames"` //Array of hostnames
}

// Host basic host info
type Host struct {
	AccountHash string     `json:"account_hash,omitempty"` // leave blank field
	Name        string     `json:"name"`
	HashCode    string     `json:"hashCode"`
	Type        string     `json:"type"`
	CreatedDate string     `json:"createdDate"`
	UpdatedDate string     `json:"updatedDate"`
	Services    []*Service `json:"services"`
	Scopes      []*Scope   `json:"scopes"`
}

// CreateHost Create a new delivery host
//Path /api/v1/accounts/{account_hash}/hosts
func (api *HWApi) CreateHost(accountHash string, host CloneHost) (*Host, error) {
	r, e := api.Request(
		&Request{
			Method: POST,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hosts", accountHash),
			Body:   host,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Host{}
	return al, json.Unmarshal(r.body, al)
}

// GetHosts List delivery hosts for the specified account
//Path /api/v1/accounts/{account_hash}/hosts
func (api *HWApi) GetHosts(accountHash string) (*HostList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hosts", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &HostList{}
	return al, json.Unmarshal(r.body, al)
}

//Clone an existing delivery host
//Path /api/v1/accounts/{account_hash}/hosts/{host_hash}
func (api *HWApi) Clone(accountHash string, hostHash string, cloneHost CloneHost) (*Host, error) {
	r, e := api.Request(
		&Request{
			Method: POST,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hosts/%s", accountHash, hostHash),
			Body:   cloneHost,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Host{}
	return al, json.Unmarshal(r.body, al)
}

// DeleteHost Delete a delivery host
//Path /api/v1/accounts/{account_hash}/hosts/{host_hash}
func (api *HWApi) DeleteHost(accountHash string, hostHash string) (bool, error) {
	if _, e := api.Request(
		&Request{
			Method: DELETE,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hosts/%s", accountHash, hostHash),
		},
	); e != nil {
		return false, e
	}
	return true, nil

}

// GetHost Get a delivery host
//Path /api/v1/accounts/{account_hash}/hosts/{host_hash}
func (api *HWApi) GetHost(accountHash string, hostHash string) (*Host, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hosts/%s", accountHash, hostHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Host{}
	return al, json.Unmarshal(r.body, al)
}

// UpdateHost Update a delivery host
//Path /api/v1/accounts/{account_hash}/hosts/{host_hash}
func (api *HWApi) UpdateHost(accountHash string, hostHash string, host *Host) (*Host, error) {
	r, e := api.Request(
		&Request{
			Method: PUT,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hosts/%s", accountHash, hostHash),
			Body:   host,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Host{}
	return al, json.Unmarshal(r.body, al)
}

// Graph return simple configure graph
//GET /api/v1/accounts/{account_hash}/graph
func (api *HWApi) Graph(accountHash string) (*map[string]interface{}, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/graph", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &map[string]interface{}{}
	return al, json.Unmarshal(r.body, al)
}
