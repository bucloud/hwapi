package hwapi

import "encoding/json"

//Billing region list
type BillingRegionList struct {
	List []*BRegion `json:"list"` //List
}

//Billing region
type BRegion struct {
	ID   int    `json:"id"`   //Id
	Code string `json:"code"` //Code
	Name string `json:"name"` //Name
}

//Returns the list of billing regions in the CDN.
//Path /api/v1/billingRegions
func (api *hwapi) GetBillingRegions() (*BillingRegionList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    "/api/v1/billingRegions",
		},
	)
	if e != nil {
		return nil, e
	}
	al := &BillingRegionList{}
	return al, json.Unmarshal(r.body, al)
}
