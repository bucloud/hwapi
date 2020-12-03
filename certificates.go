package hwapi

import (
	"encoding/json"
	"fmt"
)

// Certificate A certificate used for secure delivery
type Certificate struct {
	AccountHash            string                          `json:"account_hash,omitempty"`
	ID                     int                             `json:"id"`             //The certificate's id
	CommonName             string                          `json:"commonName"`     //Primary hostname for which this certificate can serve traffic
	Certificate            string                          `json:"certificate"`    //The text of the x.509 certificate itself
	Key                    string                          `json:"key"`            //The text of the x.509 private key
	CaBundle               string                          `json:"caBundle"`       //The text of the certificate's CA bundle
	Ciphers                string                          `json:"ciphers"`        //The cipher list which should be used during the SSL handshake
	Fingerprint            string                          `json:"fingerprint"`    //The cryptographic hash of the certificate used for uniqueness checking
	Issuer                 string                          `json:"issuer"`         //The name of the organization which issued this certificate
	Requester              *Requester                      `json:"requester"`      //The user which uploaded this certificate
	CreatedDate            string                          `json:"createdDate"`    //The time at which this certificate was uploaded
	UpdatedDate            string                          `json:"updatedDate"`    //The time at which this certificate was last updated
	ExpirationDate         string                          `json:"expirationDate"` //The time at which this certificate is no longer valid
	Trusted                bool                            `json:"trusted"`        //Whether or not this certificate passes CA validation
	CertificateInformation `json:"certificateInformation"` //Information extracted from the certificate
}

// CertificateInformation certificate details
type CertificateInformation struct {
	Name string `json:"name"`

	// DC would contains array strings in some suitaions
	// Use interface{} instead of string
	Subject          map[string]interface{} `json:"subject"`
	Hash             string                 `json:"hash"`
	Issuer           map[string]string      `json:"issuer"`
	Version          int                    `json:"version"`
	SerialNumber     string                 `json:"serialNumber"`
	SerialNumberHex  string                 `json:"serialNumberHex"`
	ValidFrom        string                 `json:"validFrom"`
	ValidTo          string                 `json:"validTo"`
	ValidFromTimeT   int                    `json:"validFrom_time_t"`
	ValidToTimeT     int                    `json:"validTo_time_t"`
	SignatureTypeSN  string                 `json:"signatureTypeSN"`
	SignatureTypeLN  string                 `json:"signatureTypeLN"`
	SignatureTypeNID int                    `json:"signatureTypeNID"`
	Extensions       map[string]string      `json:"extensions"`
}

// Requester uploader
type Requester struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// CertificateResponse Certificate response
type CertificateResponse struct {
	List []*Certificate `json:"list"` //Status
}

// HostsForCertificate which host use this certificate
type HostsForCertificate map[string][]string

// GetCertificates List all certificates on an account
//Path /api/v1/accounts/{account_hash}/certificates
func (api *HWApi) GetCertificates(accountHash string) (*CertificateResponse, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/certificates", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &CertificateResponse{}
	return al, json.Unmarshal(r.body, al)
}

//UploadCertificate a new certificate
//Path /api/v1/accounts/{account_hash}/certificates
func (api *HWApi) UploadCertificate(accountHash string, certificate *Certificate) (*Certificate, error) {
	r, e := api.Request(
		&Request{
			Method: POST,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/certificates", accountHash),
			Body:   certificate,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Certificate{}
	return al, json.Unmarshal(r.body, al)
}

// DeleteCertificate Delete a certificate
//Path /api/v1/accounts/{account_hash}/certificates/{certificate_id}
func (api *HWApi) DeleteCertificate(accountHash string, certID int) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: DELETE,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/certificates/%d", accountHash, certID),
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}

// GetCertificate Get a certificate
//Path /api/v1/accounts/{account_hash}/certificates/{certificate_id}
func (api *HWApi) GetCertificate(accountHash string, certID int) (*Certificate, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/certificates/%d", accountHash, certID),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Certificate{}
	return al, json.Unmarshal(r.body, al)
}

// UpdateCertificate Update a certificate (useful for expired certs)
//Path /api/v1/accounts/{account_hash}/certificates/{certificate_id}
func (api *HWApi) UpdateCertificate(accountHash string, certID int) (*Certificate, error) {
	r, e := api.Request(
		&Request{
			Method: PUT,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/certificates/%d", accountHash, certID),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Certificate{}
	return al, json.Unmarshal(r.body, al)
}

// GetHostsForCertificate Get Hosts for Certificate
//Path /api/v1/accounts/{account_hash}/certificates/{certificate_id}/hosts
func (api *HWApi) GetHostsForCertificate(accountHash string, certID int) (*HostsForCertificate, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/certificates/%d/hosts", accountHash, certID),
		},
	)
	if e != nil {
		return nil, e
	}
	var al *HostsForCertificate
	return al, json.Unmarshal(r.body, &al)
}
