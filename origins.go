package hwapi

import (
	"encoding/json"
	"fmt"
)

// Origin An origin server from which the CDN pulls content
type Origin struct {
	ID                           int    `json:"id,omitempty"`                           //The origin's id
	Name                         string `json:"name,omitempty"`                         //The origin's friendly name
	Type                         string `json:"Type,omitempty"`                         //The origin's type (defaults to EXTERNAL for external origins)
	Path                         string `json:"path,omitempty"`                         //The path to prepend to requests
	CreatedDate                  string `json:"createdDate,omitempty"`                  //createdDate
	UpdatedDate                  string `json:"updatedDate,omitempty"`                  //updatedDate
	RequestTimeoutSeconds        int    `json:"requestTimeoutSeconds,omitempty"`        //requestTimeoutSeconds
	ErrorCacheTTLSeconds         int    `json:"errorCacheTTLSeconds,omitempty"`         //errorCacheTTLSeconds
	MaxRetryCount                int    `json:"maxRetryCount,omitempty"`                //maxRetryCount
	AuthenticationType           string `json:"authenticationType,omitempty"`           //authenticationType
	Hostname                     string `json:"hostname,omitempty"`                     //The hostname the origin is reached at
	Port                         int    `json:"port,omitempty"`                         //The port to communicate with the origin on
	SecurePort                   int    `json:"securePort,omitempty"`                   //securePort
	OriginPullHeaders            string `json:"originPullHeaders,omitempty"`            //originPullHeaders
	OriginCacheHeaders           string `json:"originCacheHeaders,omitempty"`           //originCacheHeaders
	VerifyCertificate            bool   `json:"verifyCertificate,omitempty"`            //verifyCertificate
	CertificateCN                string `json:"certificateCN,omitempty"`                //certificateCN
	MaximumOriginPullSeconds     int    `json:"maximumOriginPullSeconds,omitempty"`     //maximumOriginPullSeconds
	MaxConnectionsPerEdge        int    `json:"maxConnectionsPerEdge,omitempty"`        //If enabled, the maximum number of concurrent connection any single edge will make to the origin
	MaxConnectionsPerEdgeEnabled bool   `json:"maxConnectionsPerEdgeEnabled,omitempty"` //Indicates if the CDN should limit the number of connections each edge should make when pulling content
	OriginPullNegLinger          string `json:"originPullNegLinger,omitempty"`          //originPullNegLinger
	OriginDefaultKeepAlive       string `json:"originDefaultKeepAlive,omitempty"`       //originDefaultKeepAlive
	DisplayName                  string `json:"displayName,omitempty"`
	AccountName                  string `json:"accountName,omitempty"`
	AccountHash                  string `json:"accountHash,omitempty"`
}

// OriginList List of origins
type OriginList struct {
	List []*Origin `json:"list"`
}

// CreateOrigin Create a new origin
//Path /api/v1/accounts/{account_hash}/origins
func (api *HWApi) CreateOrigin(accountHash string, origin *Origin) (*Origin, error) {
	r, e := api.Request(
		&Request{
			Method: POST,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/origins", accountHash),
			Body:   origin,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Origin{}
	return al, json.Unmarshal(r.body, al)
}

// GetOrigins List all origins on an account
//Path /api/v1/accounts/{account_hash}/origins
func (api *HWApi) GetOrigins(accountHash string) (*OriginList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/origins", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &OriginList{}
	return al, json.Unmarshal(r.body, al)
}

// DeleteOrigin Delete an origin
//Path /api/v1/accounts/{account_hash}/origins/{origin_id}
func (api *HWApi) DeleteOrigin(accountHash string, originID int) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: DELETE,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/origins/%d", accountHash, originID),
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}

// UpdateOrigin Update an origin
//Path /api/v1/accounts/{account_hash}/origins/{origin_id}
func (api *HWApi) UpdateOrigin(accountHash string, originID int, origin *Origin) (*Origin, error) {
	r, e := api.Request(
		&Request{
			Method: PUT,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/origins/%d", accountHash, originID),
			Body:   origin,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Origin{}
	return al, json.Unmarshal(r.body, al)
}

// GetOrigin Get an origin
//Path /api/v1/accounts/{account_hash}/origins/{origin_id}
func (api *HWApi) GetOrigin(accountHash string, originID int) (*Origin, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/origins/%d", accountHash, originID),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Origin{}
	return al, json.Unmarshal(r.body, al)
}
