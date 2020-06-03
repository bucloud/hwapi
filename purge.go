package hwapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Purges struct {
	List []*Purge `json:"list"`
}

type Purge struct {
	Url             string `json:"url"`
	Recursive       bool   `json:"recursive"`
	PurgeAllDynamic bool   `json:"purgeAllDynamic"`
}

type PurgeState struct {
	Id       string  `json:"id"`
	Progress float32 `json:"progress"`
}

//Get purge status
func (api *hwapi) GetPurgeState(accountHash string, purgeID string) (float32, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/purge/%s", accountHash, purgeID),
		},
	)
	if e != nil {
		return 0, e
	}
	al := &PurgeState{}
	return al.Progress, json.Unmarshal(r.body, al)
}

func parsePugeList(p []interface{}) (*Purges, error) {
	pl := &Purges{}
	if len(p) == 0 {
		return nil, errors.New("purgeList nil exception")
	} else {
		for _, v := range p {
			switch x := v.(type) {
			case string:
				pl.List = append(pl.List, &Purge{
					Url: x,
				})
			case *Purge:
				pl.List = append(pl.List, x)
			default:
				return nil, errors.New("Only string/Purge are supported")
			}
		}
	}
	return pl, nil
}

//Purge Url
//Usage Purge("a1b1c1d1",Purge{url:"1"},Purge{url:"2"}) or Purge("a1b1c1d1",[Purge{url:"1"},Purge{url:"2"}])
//You can Purge{url:"1"} as "1"
func (api *hwapi) Purge(accountHash string, purgeList ...interface{}) (*PurgeState, error) {
	pl, err := parsePugeList(purgeList)
	if err != nil {
		return nil, err
	}
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/purge", accountHash),
			Body:   pl,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &PurgeState{}
	return al, json.Unmarshal(r.body, al)
}

//Prewarm Url
//Usage Purge("a1b1c1d1",Purge{url:"1"},Purge{url:"2"}) or Purge("a1b1c1d1",[Purge{url:"1"},Purge{url:"2"}])
//You can Purge{url:"1"} as "1"
func (api *hwapi) WarmUpUrl(accountHash string, purgeList ...interface{}) (*PurgeState, error) {
	pl, err := parsePugeList(purgeList)
	if err != nil {
		return nil, err
	}
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/purge", accountHash),
			Body:   pl,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &PurgeState{}
	return al, json.Unmarshal(r.body, al)
}
