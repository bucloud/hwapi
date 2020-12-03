package hwapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// Purges used by post body
type Purges struct {
	List []*Purge `json:"list"`
}

// WarmUp alias Purge
type WarmUp struct {
	URL       string   `json:"url"`
	Recursive bool     `json:"recursive"`
	Range     bool     `json:"-"`
	FixRange  string   `json:"-"`
	Headers   []string `json:"-"`

	// RegionFilter control which pop handle warmup
	RegionFilter string `json:"-"`
}

// Purge object
type Purge struct {
	URL             string `json:"url"`
	Recursive       bool   `json:"recursive"`
	PurgeAllDynamic bool   `json:"purgeAllDynamic"`
}

// PurgeState purge response
type PurgeState struct {
	ID       string  `json:"id"`
	Progress float32 `json:"progress"`
}

// GetPurgeState check purge state by purgeID
func (api *HWApi) GetPurgeState(accountHash string, purgeID string) (float32, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/purge/%s", accountHash, purgeID),
		},
	)
	if e != nil {
		return 0, e
	}
	al := &PurgeState{}
	return al.Progress, json.Unmarshal(r.body, al)
}

func parsePurgeList(p []interface{}) (*Purges, error) {
	pl := &Purges{}
	if len(p) == 0 {
		return nil, errors.New("purgeList nil exception")
	}

	for _, v := range p {
		switch x := v.(type) {
		case string:
			pl.List = append(pl.List, &Purge{
				URL: x,
			})
		case *Purge:
			pl.List = append(pl.List, x)
		default:
			return nil, errors.New("Only string/*Purge are supported")
		}
	}

	return pl, nil
}

//Purge Url
//Usage Purge("a1b1c1d1",Purge{url:"1"},Purge{url:"2"}) or Purge("a1b1c1d1",[Purge{url:"1"},Purge{url:"2"}])
//You can Purge{url:"1"} as "1"
func (api *HWApi) Purge(accountHash string, purgeList ...interface{}) (*PurgeState, error) {
	pl, err := parsePurgeList(purgeList)
	if err != nil {
		return nil, err
	}
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/purge", accountHash),
			Body:   pl,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &PurgeState{}
	return al, json.Unmarshal(r.body, al)
}
func warmUpConvert(i interface{}) *WarmUp {
	switch i.(type) {
	case string:
		return &WarmUp{
			URL: i.(string),
		}
	case *WarmUp:
		return i.(*WarmUp)
	default:
		return &WarmUp{
			URL: "http://205.185.216.10/",
		}
	}
}

//WarmUp Url
//Usage WarmUp("a1b1c1d1",WarmUp{url:"1"},WarmUp{url:"2"}) or WarmUp("a1b1c1d1",[WarmUp{url:"1"},WarmUp{url:"2"}])
func (api *HWApi) WarmUp(accountHash string, warmList ...interface{}) (bool, error) {
	if len(warmList) == 0 {
		return false, fmt.Errorf("nothing todo")
	}
	// get pop list
	pops, e := api.GetPoPs()
	if e != nil {
		return false, e
	}
	// Get test url
	var tu *url.URL
	w0 := warmUpConvert(warmList[0])
	tu, _ = url.Parse(w0.URL)
	// check available pops
	availablePOP := []*POP{}
	for _, pop := range pops.List {
		tu.Host = "doppler." + pop.Code + ".hwcdn.net"
		_, e := api.hc.RoundTrip(&http.Request{
			Method: GET,
			URL:    tu,
		})
		if e != nil {
			// unknow error
			fmt.Printf("unknow error %s teat as failed", e.Error())
		}
		availablePOP = append(availablePOP, pop)
	}
	// for _, u := range warmList {
	// 	var tempURL *url.URL
	// 	wt := warmUpConvert(u)
	// 	tempURL, _ = url.Parse(wt.URL)
	// 	for _, p := range availablePOP {
	// 		tempURL.Host = "doppler." + p.Code + ".hwcdn.net"

	// 		_, e := api.hc.RoundTrip(&http.Request{
	// 			URL:    tempURL,
	// 			Method: "GET",
	// 			Header: http.Header{
	// 				"Host":""
	// 			},
	// 		})
	// 	}
	// }

	return true, nil
}
