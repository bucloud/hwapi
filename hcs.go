package hwapi

import (
	"encoding/json"
	"fmt"
)

//HcsContainer list
type HcsContainerList struct {
	List []*HcsContainer `json:"list"` //list
}

//HcsContainer
type HcsContainer struct {
	Name            string            `json:"name,omitempty"`            //(POST only) The container's name
	Region          string            `json:"region"`                    //(read only) The container's region
	Tenant          string            `json:"tenant"`                    //(read only) The container's tenant
	Count           int               `json:"count"`                     //(read only) The number of objects in this container
	Bytes           int               `json:"bytes"`                     //(read only) The total size (in bytes) of all objects in this container
	Quota           int               `json:"quota,omitempty"`           //(PUT only) The container's size quota (in bytes)
	ReadPermissions string            `json:"readPermissions,omitempty"` //(PUT only) The container's read permissions ('.r:*,.rlistings' will make it public)
	Meta            map[string]string `json:"meta,omitempty"`            //(PUT only) An indexed-array (json object) of key-value pairs for setting custom meta-data on containers. (Key must follow this form: 'X-Container-Meta-{name})
}

//HcsTenant list
type HcsTenantList struct {
	List []*HcsTenant `json:"list"` //list
}

//HcsTenant
type HcsTenant struct {
	Name            string `json:"name"`            //The tenants's friendly name
	HCSUser         string `json:"hcsUser"`         //The username to be used for the hcs account the system creates
	HCSUserPassword string `json:"hcsUserPassword"` //The password for hcsUser
	HCSRegion       string `json:"hcsRegion"`       //The HCS global region to assign
	ID              int    `json:"id,omitepty"`
	AccountID       string `json:"accountId,omitepty"`
	HCSTenant       string `json:"hcsTenant,omitepty"`
	DeltedDate      string `json:"deletedDate,omitepty"`
	CreatedDate     string `json:"createdDate,omitepty"`
	UpdatedDate     string `json:"updatedDate,omitepty"`
}

//HcsObject
type HcsObject struct {
	Etag         string            `json:"etag"`          //(read only) The object's eTag
	Hash         string            `json:"hash"`          //(read only) The object's eTag
	LastModified string            `json:"last_modified"` //(read only) The last time this object was modified
	DeleteAt     string            `json:"delete-at"`     //Setting this value will mark the object for automatic deletion for the given timestamp
	Meta         map[string]string `json:"meta"`          //An array of key-value pairs for setting custom meta-data on objects. (Key must follow this form: 'X-Object-Meta-{name})
	Bytes        int32             `json:"bytes"`
	Name         string            `json:"name"` //filename
}

//Create a new HcsTenant
//Path /api/v1/accounts/{account_hash}/hcs/tenants
//JSON representation of the HcsTenant to create. The structure should match the response class model minus the id property as that is not allowed when creating a new HcsTenant.
func (api *hwapi) CreateHCSTenant(accountHash string, hcsT HcsTenant) (*HcsTenant, error) {
	r, e := api.Request(
		&Request{
			Method: POST,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/tenants", accountHash),
			Body:   hcsT,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &HcsTenant{}
	return al, json.Unmarshal(r.body, al)
}

//Returns a list of HcsTenants that belong to the given account.
//Path /api/v1/accounts/{account_hash}/hcs/tenants
func (api *hwapi) GetHCSTenants(accountHash string) (*HcsTenantList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/tenants", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &HcsTenantList{}
	return al, json.Unmarshal(r.body, al)
}

//Delete a host name
//Path /api/v1/accounts/{account_hash}/hcs/tenants/{tenant_id}
func (api *hwapi) DeleteHCSTenant(accountHash string, tenantID int) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: DELETE,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/tenants/%d", accountHash, tenantID),
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}

//Returns specific HcsTenant on the specified account.
//Path /api/v1/accounts/{account_hash}/hcs/tenants/{tenant_id}
func (api *hwapi) GetHCSTenant(accountHash string, tenantID int) (*HcsTenant, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/tenants/%d", accountHash, tenantID),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &HcsTenant{}
	return al, json.Unmarshal(r.body, al)
}

//Update an existing HcsTenant for an account
//Path /api/v1/accounts/{account_hash}/hcs/tenants/{tenant_id}
//JSON representation of the HcsTenant to update. The structure should match the response class model.
func (api *hwapi) UpdateHCSTenant(accountHash string, tenantID int, t HcsTenant) (*HcsTenant, error) {
	r, e := api.Request(
		&Request{
			Method: PUT,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/tenants/%d", accountHash, tenantID),
			Body:   t,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &HcsTenant{}
	return al, json.Unmarshal(r.body, al)
}

//Return a list of containers that belong to the given account
//Path /api/v1/accounts/{account_hash}/hcs/containers
func (api *hwapi) GetHCSContainers(accountHash string) (*HcsContainerList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/containers", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &HcsContainerList{}
	return al, json.Unmarshal(r.body, al)
}

//Create a new container
//Path /api/v1/accounts/{account_hash}/hcs/containers/{tenant_name}
//JSON representation of the HcsContainer to create. The only field accepted on HcsContainer creation is 'name'.
func (api *hwapi) CreateHCSContainer(accountHash string, tenantName string, containerName string) (*HcsContainer, error) {
	r, e := api.Request(
		&Request{
			Method: POST,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/containers/%s", accountHash, tenantName),
			Body: &HcsContainer{
				Name: containerName,
			},
		},
	)
	if e != nil {
		return nil, e
	}
	al := &HcsContainer{}
	return al, json.Unmarshal(r.body, al)
}

//Delete a container
//Path /api/v1/accounts/{account_hash}/hcs/containers/{tenant_name}/{container_name}
func (api *hwapi) DeleteHCSContainer(accountHash string, tenantName string, containerName string) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: DELETE,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/containers/%s/%s", accountHash, tenantName, containerName),
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}

//Returns specific container on the specified account and tenant.
//Path /api/v1/accounts/{account_hash}/hcs/containers/{tenant_name}/{container_name}
func (api *hwapi) getHCSContainer(accountHash string, tenantName string, containerName string) (*HcsContainer, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/containers/%s/%s", accountHash, tenantName, containerName),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &HcsContainer{}
	return al, json.Unmarshal(r.body, al)
}

//Update an existing container
//Path /api/v1/accounts/{account_hash}/hcs/containers/{tenant_name}/{container_name}
//JSON representation of the HcsContainer to create. The only fields accepted on HcsContainer updates are 'quota, readPermissions, meta'.
func (api *hwapi) UpdateHCSContainer(accountHash string, tenantName string, containerName string, container HcsContainer) (*HcsContainer, error) {
	r, e := api.Request(
		&Request{
			Method: PUT,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/containers/%s/%s", accountHash, tenantName, containerName),
			Body: &HcsContainer{
				Quota:           container.Quota,
				ReadPermissions: container.ReadPermissions,
				Meta:            container.Meta,
			},
		},
	)
	if e != nil {
		return nil, e
	}
	al := &HcsContainer{}
	return al, json.Unmarshal(r.body, al)
}

//Get the objects
//Path /api/v1/accounts/{account_hash}/hcs/objects/{tenant_name}/{container_name}
func (api *hwapi) getHCSObjects(accountHash string, tenantName string, containerName string, prefix ...string) ([]*HcsObject, error) {
	if prefix == nil {
		prefix[0] = ""
	}
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/containers/%s/%s", accountHash, tenantName, containerName),
			Query: map[string]string{
				"prefix": prefix[0],
			},
		},
	)
	if e != nil {
		return nil, e
	}
	var al []*HcsObject
	return al, json.Unmarshal(r.body, &al)
}

//Delete HCS object
//Path /api/v1/accounts/{account_hash}/hcs/objects/{tenant_name}/{container_name}/{object_name}
func (api *hwapi) DeleteHCSObject(accountHash string, tenantName string, containerName string, objectName string) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: DELETE,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/containers/%s/%s/%s", accountHash, tenantName, containerName, objectName),
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}

//Returns specific object on the specified account.
//Path /api/v1/accounts/{account_hash}/hcs/objects/{tenant_name}/{container_name}/{object_name}
func (api *hwapi) GetHCSObject(accountHash string, tenantName string, containerName string, objectName string) (*HcsObject, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/containers/%s/%s/%s", accountHash, tenantName, containerName, objectName),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &HcsObject{}
	return al, json.Unmarshal(r.body, al)
}

//Update an existing object
//Path /api/v1/accounts/{account_hash}/hcs/objects/{tenant_name}/{container_name}/{object_name}
//JSON representation of the Hcs Object to update. The structure should match the response class model without the 'etag' and 'lastUpdated' fields.
func (api *hwapi) UpdateHCSObject(accountHash string, tenantName string, containerName string, objectName string, h *HcsObject) (*HcsObject, error) {
	r, e := api.Request(
		&Request{
			Method: PUT,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/hcs/containers/%s/%s/%s", accountHash, tenantName, containerName, objectName),
			Body:   h,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &HcsObject{}
	return al, json.Unmarshal(r.body, al)
}
