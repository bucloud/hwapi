package hwapi

import (
	"encoding/json"
	"fmt"
)

//Delete all sessions associated with a user
func (api *HWApi) DeleteSessions(accountHash string, uid int) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: DELETE,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/users/%d/sessions", accountHash, uid),
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}

//Fetch all sessions associated with a user
func (api *HWApi) GetSessions(accountHash string, uid int) (*AccessTokenList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/users/%d/sessions", accountHash, uid),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &AccessTokenList{}
	return al, json.Unmarshal(r.body, al)
}

//Delete a session associated with a user
func (api *HWApi) DeleteSession(accountHash string, uid int, tokenID string) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: DELETE,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/users/%d/sessions/%s", accountHash, uid, tokenID),
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}
