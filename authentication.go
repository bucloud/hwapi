package hwapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AuthToken tokens contains log and API
type AuthToken struct {
	AccessToken  string `json:"access_token,omitepty"`  //Access token
	TokenType    string `json:"token_type"`             //Token type
	ExpiresIn    int    `json:"expires_in"`             //Time to expire
	RefreshToken string `json:"refresh_token,omitepty"` //Refresh token
	UserAgent    string `json:"user_agent"`             //User agent
	Application  string `json:"application"`            //Application
	IP           string `json:"ip"`                     //IP

	// LogTokens deprecated after 2021-01-01
	LogTokens string `json:"-"` //token used to access accesslogs,
}

// Authentication simple token
type Authentication struct {
	token       string //Token
	application string //Application
	ip          string //IP
}

// APITokenRequest Request model for an API token request
type APITokenRequest struct {
	password    string //The user's current password
	application string //The name of the application this token will be used for
}

type authInfo struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

// AccessTokenList list of accesstoken
type AccessTokenList struct {
	list []*AccessToken
}

// AccessToken current accesstoken
type AccessToken struct {
	id          string //Unique ID for this AccessToken
	token       string //The token used to authenticate with the API
	ip          string //The remote address of the client which created this token
	application string //The name of the application which will use this token
	expiration  string //Expiration date of the token
	active      int    //Whether or not the token is active
	refresh     bool   //Whether or not the token can be used to refresh an access token
}

//CreateToken Create an API token with infinite expiration
func (api *HWApi) CreateToken(accountHash string, uid int, tokenRequest ...*APITokenRequest) (*Authentication, error) {
	if (accountHash == "" || uid == 0) && api.CurrentUser.AccountHash == "" {
		return nil, errors.New("accountHash must supplied or use AboutMe to generate current user info")
	}
	if accountHash == "" {
		accountHash = api.CurrentUser.AccountHash
	}
	if uid == 0 {
		uid = api.CurrentUser.ID
	}

	//Create create token request
	// r := http.NewRequest("POST", "/api/v1/accounts/"+accountHash+"/"+uid+"/tokens",)
	r, e := api.Request(
		&Request{
			Method: POST,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/users/%d/tokens", accountHash, uid),
			Body:   &tokenRequest,
		},
	)
	if e != nil {
		return nil, e
	}
	a := &Authentication{}
	return a, json.Unmarshal(r.body, a)
}

// GetTokens Fetch all tokens associated with this user
func (api *HWApi) GetTokens(accountHash string, uid int) (*AccessTokenList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/users/%d/tokens", accountHash, uid),
		},
	)
	if e != nil {
		return nil, e
	}
	atl := &AccessTokenList{}
	return atl, json.Unmarshal(r.body, atl)
}

// DeleteToken Delete token
func (api *HWApi) DeleteToken(a string, uid int, tokenID int) (bool, error) {
	_, e := api.Request(&Request{
		URL:    fmt.Sprintf("/api/v1/accounts/%s/users/%d/tokens/%d", a, uid, tokenID),
		Method: DELETE,
	})
	if e != nil {
		return true, nil
	}
	return false, e
}

// Auth Authenticate user or refresh an access token
// if you want to get accesslogs token, set  accesslog to true
func (api *HWApi) Auth(u, p string, accesslog ...bool) (*AuthToken, error) {
	if len(accesslog) > 0 && accesslog[0] {

		r, e := api.Request(&Request{
			Method: GET,
			URL:    authURL,
			Headers: map[string]string{
				"X-Auth-User":   "hwcdn-logstore:" + u,
				"X-Auth-Key":    p,
				"X-Application": "HWAPI-go",
			},
		})
		if e != nil {
			return nil, fmt.Errorf("get accesslog token for failed, %s", e.Error())
		}

		if api.AuthToken == nil {
			api.AuthToken = &AuthToken{}
		}
		api.AuthToken.LogTokens = r.Headers.Get("X-Auth-Token")
		return api.AuthToken, nil
	}
	r, e := api.Request(&Request{
		Method: POST,
		URL:    "/auth/token",
		Body: &authInfo{
			GrantType: "password",
			Username:  u,
			Password:  p,
		},
	})
	if e != nil {
		return nil, e
	}
	return api.AuthToken, json.Unmarshal(r.body, &api.AuthToken)

}

//SetToken Check if token available, than set to AuthToken if available
func (api *HWApi) SetToken(t string) {
	if api.AuthToken == nil {
		api.AuthToken = &AuthToken{
			AccessToken: t,
		}
	} else {
		api.AuthToken.AccessToken = t
	}
}

//Use /api/v1/users/me to check accesstoken vaildation
func (api *HWApi) checkToken() (bool, error) {
	_, e := api.AboutMe()
	if e != nil {
		return false, e
	}
	return true, nil
}

// RefreshToken Automate refresh token when token is not available.
//Note, this function had deprecated.
func (api *HWApi) RefreshToken(refreshT ...string) (*AuthToken, error) {
	if api.AuthToken.RefreshToken == "" && refreshT == nil {
		return nil, errors.New("RefreshToken not exists, try Auth() or provide refreshtoken")
	}
	t := ""
	if refreshT == nil {
		t = api.AuthToken.RefreshToken
	} else {
		t = refreshT[0]
	}
	//create request
	r, e := api.Request(
		&Request{
			Method: POST,
			URL:    "/auth/token",
			Body: &authInfo{
				GrantType:    "refresh_token",
				RefreshToken: t,
			},
		},
	)
	if e != nil {
		return nil, e
	}
	return api.AuthToken, json.Unmarshal(r.body, &api.AuthToken)
}
