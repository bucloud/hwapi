package hwapi

import (
	"encoding/json"
	"fmt"
)

//A certificate used for secure delivery
type Certificate struct {
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

type CertificateInformation struct {
	Name             string            `json:"name"`
	Subject          map[string]string `json:"subject"`
	Hash             string            `json:"hash"`
	Issuer           map[string]string `json:"issuer"`
	Version          int               `json:"version"`
	SerialNumber     string            `json:"serialNumber"`
	SerialNumberHex  string            `json:"serialNumberHex"`
	ValidFrom        string            `json:"validFrom"`
	ValidTo          string            `json:"validTo"`
	ValidFrom_time_t int               `json:"validFrom_time_t"`
	ValidTo_time_t   int               `json:"validTo_time_t"`
	SignatureTypeSN  string            `json:"signatureTypeSN"`
	SignatureTypeLN  string            `json:"signatureTypeLN"`
	SignatureTypeNID int               `json:"signatureTypeNID"`
	Extensions       map[string]string `json:"extensions"`
}

type Requester struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

//Certificate response
type CertificateResponse struct {
	List []*Certificate `json:"list"` //Status
}

type HostsForCertificate map[string][]string

//List all certificates on an account
//Path /api/v1/accounts/{account_hash}/certificates
func (api *HWApi) GetCertificates(accountHash string) (*CertificateResponse, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/certificates", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &CertificateResponse{}
	return al, json.Unmarshal(r.body, al)
}

//Upload a new certificate
//Path /api/v1/accounts/{account_hash}/certificates
func (api *HWApi) UploadCertificate(accountHash string, certificate *Certificate) (*Certificate, error) {
	r, e := api.Request(
		&Request{
			Method: POST,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/certificates", accountHash),
			Body:   certificate,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Certificate{}
	return al, json.Unmarshal(r.body, al)
}

//Delete a certificate
//Path /api/v1/accounts/{account_hash}/certificates/{certificate_id}
func (api *HWApi) DeleteCertificate(accountHash string, cert_id int) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: DELETE,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/certificates/%d", accountHash, cert_id),
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}

//Get a certificate
//Path /api/v1/accounts/{account_hash}/certificates/{certificate_id}
func (api *HWApi) GetCertificate(accountHash string, cert_id int) (*Certificate, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/certificates/%d", accountHash, cert_id),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Certificate{}
	return al, json.Unmarshal(r.body, al)
}

//Update a certificate (useful for expired certs)
//Path /api/v1/accounts/{account_hash}/certificates/{certificate_id}
func (api *HWApi) UpdateCertificate(accountHash string, cert_id int) (*Certificate, error) {
	r, e := api.Request(
		&Request{
			Method: PUT,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/certificates/%d", accountHash, cert_id),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Certificate{}
	return al, json.Unmarshal(r.body, al)
}

//Get Hosts for Certificate
//Path /api/v1/accounts/{account_hash}/certificates/{certificate_id}/hosts
func (api *HWApi) GetHostsForCertificate(accountHash string, cert_id int) (*HostsForCertificate, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/certificates/%d/hosts", accountHash, cert_id),
		},
	)
	if e != nil {
		return nil, e
	}
	var al *HostsForCertificate
	return al, json.Unmarshal(r.body, &al)
}
