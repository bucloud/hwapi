// Package hwapi HCS deprecated since 2021-01-01, Use GCS as Object Storage service
package hwapi

import (
	"encoding/json"
	"fmt"
)

// GCSAccounts contains a list of created ServiceAccount
type GCSAccounts struct {
	List []*GCSAccount `json:"list"`
}

// GCSAccount used for access GCS
type GCSAccount struct {
	CreatedDate string `json:"createdDate,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"displayName,omitempty"`

	// Email read-only GCS identify email
	Email string `json:"email,omitempty"`

	// ID used to miantain serviceAccount via striketracker
	ID string `json:"id,omitempty"`

	// Status aims to descrip whether GCS account created
	// Available value PENDING,READY
	Status string `json:"status,omitempty"`
}

// GCSPrivateKeys contains a list of GCSPrivateKey
type GCSPrivateKeys struct {
	List []*GCSPrivateKey `json:"list"`
}

// GCSPrivateKey describe the privatekey
type GCSPrivateKey struct {
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`

	// PrivateKeyData base64 encoded json data used by GCS, Note decode this to string before use it
	// Only accessable after the first create privateKey and un-recoverable
	PrivateKeyData string `json:"privateKeyData,omitempty"`
	GCSAccountID   string `json:"serviceAccountId"`
}

// GCSHMacKeys contains a list of GCSHMacKey
type GCSHMacKeys struct {
	List []*GCSHMacKey `json:"list"`
}

// GCSHMacKey describe the HMacKey
type GCSHMacKey struct {
	CreatedAt    string `json:"createdAt"`
	ID           string `json:"id"`
	GCSAccountID string `json:"serviceAccountId"`

	// AccessID similar to AWS clientID
	AccessID string `json:"accessId,omitempty"`
	// PrivateKeyData base64 encoded json data used by GCS, Note decode this to string before use it
	// Only accessable after the first create privateKey and un-recoverable
	PrivateKeyData string `json:"privateKeyData,omitempty"`
}

// CreateGCSAccount create a new GCSAccount/service_account
// POST api/v2/accounts/{account_hash}/service_accounts
func (api *HWApi) CreateGCSAccount(accountHash, description, name string) (*GCSAccount, error) {
	r, e := api.Request(
		&Request{
			Method: POST,
			URL:    fmt.Sprintf("/api/v2/accounts/%s/service_accounts", accountHash),
			Body: map[string]string{
				"description": description,
				"name":        name,
			},
		},
	)
	if e != nil {
		return nil, e
	}
	al := &GCSAccount{}
	return al, json.Unmarshal(r.body, al)
}

// GetGCSAccounts list all created service accounts
// GET api/v2/accounts/{account_hash}/service_accounts
func (api *HWApi) GetGCSAccounts(accountHash string) (*GCSAccounts, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v2/accounts/%s/service_accounts", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &GCSAccounts{}
	return al, json.Unmarshal(r.body, al)
}

// GetGCSAccount get service account info
// GET api/v2/accounts/{account_hash}/service_accounts/{id}
func (api *HWApi) GetGCSAccount(accountHash, GCSAccountID string) (*GCSAccount, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v2/accounts/%s/service_accounts/%s", accountHash, GCSAccountID),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &GCSAccount{}
	return al, json.Unmarshal(r.body, al)
}

// GetGCSPrivateKeys list all created privateKeys
// GET api/v2/accounts/{account_hash}/service_accounts/keys
func (api *HWApi) GetGCSPrivateKeys(accountHash, GCSAccountID string) (*GCSPrivateKeys, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v2/accounts/%s/service_accounts/%s/keys", accountHash, GCSAccountID),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &GCSPrivateKeys{}
	return al, json.Unmarshal(r.body, al)
}

// CreateGCSPrivateKey create a new privateKey
// POST api/v2/accounts/{account_hash}/service_accounts/keys
func (api *HWApi) CreateGCSPrivateKey(accountHash, GCSAccountID string) (*GCSPrivateKey, error) {
	r, e := api.Request(
		&Request{
			Method: POST,
			URL:    fmt.Sprintf("/api/v2/accounts/%s/service_accounts/%s/keys", accountHash, GCSAccountID),
			Body: map[string]string{
				"accountHash":      accountHash,
				"serviceAccountId": GCSAccountID,
			},
		},
	)
	if e != nil {
		return nil, e
	}
	al := &GCSPrivateKey{}
	return al, json.Unmarshal(r.body, al)
}

// DeleteGCSPrivateKey create a new privateKey
// POST api/v2/accounts/{account_hash}/service_accounts/{service_accounts_id}/keys/{private_key_id}
func (api *HWApi) DeleteGCSPrivateKey(accountHash, GCSAccountID, GCSPrivateKeyID string) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: DELETE,
			URL:    fmt.Sprintf("/api/v2/accounts/%s/service_accounts/%s/keys/%s", accountHash, GCSAccountID, GCSPrivateKeyID),
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}

// GetGCSHMacKeys list all created service accounts
// GET api/v2/accounts/{account_hash}/service_accounts/{service_accounts_id}/hmac_keys
func (api *HWApi) GetGCSHMacKeys(accountHash, serviceAccountID string) (*GCSHMacKeys, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v2/accounts/%s/service_accounts/%s/keys", accountHash, serviceAccountID),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &GCSHMacKeys{}
	return al, json.Unmarshal(r.body, al)
}

// CreateGCSHMacKey create a new privateKey
// POST api/v2/accounts/{account_hash}/service_accounts/{service_accounts_id}/hmac_keys
func (api *HWApi) CreateGCSHMacKey(accountHash, GCSAccountID string) (*GCSHMacKey, error) {
	r, e := api.Request(
		&Request{
			Method: POST,
			URL:    fmt.Sprintf("/api/v2/accounts/%s/service_accounts/%s/hmac_keys", accountHash, GCSAccountID),
			Body: map[string]string{
				"accountHash":      accountHash,
				"serviceAccountId": GCSAccountID,
			},
		},
	)
	if e != nil {
		return nil, e
	}
	al := &GCSHMacKey{}
	return al, json.Unmarshal(r.body, al)
}

// DeleteGCSHMacKey create a new privateKey
// POST api/v2/accounts/{account_hash}/service_accounts/{service_accounts_id}/hmac_keys
func (api *HWApi) DeleteGCSHMacKey(accountHash, GCSAccountID, GCSHMacKeyID string) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: DELETE,
			URL:    fmt.Sprintf("/api/v2/accounts/%s/service_accounts/%s/hmac_keys/%s", accountHash, GCSAccountID, GCSHMacKeyID),
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}
