package hwapi

import "encoding/json"

// BillingRegionList Billing region list
type BillingRegionList struct {
	List []*BRegion `json:"list"` //List
}

// BRegion Billing region
type BRegion struct {
	ID   int    `json:"id"`   //Id
	Code string `json:"code"` //Code
	Name string `json:"name"` //Name
}

// GetBillingRegions Returns the list of billing regions in the CDN.
//Path /api/v1/billingRegions
func (api *HWApi) GetBillingRegions() (*BillingRegionList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    "/api/v1/billingRegions",
		},
	)
	if e != nil {
		return nil, e
	}
	al := &BillingRegionList{}
	return al, json.Unmarshal(r.body, al)
}
