package hwapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type AuthToken struct {
	AccessToken  string `json:"access_token,omitepty"`  //Access token
	TokenType    string `json:"token_type"`             //Token type
	ExpiresIn    int    `json:"expires_in"`             //Time to expire
	RefreshToken string `json:"refresh_token,omitepty"` //Refresh token
	UserAgent    string `json:"user_agent"`             //User agent
	Application  string `json:"application"`            //Application
	Ip           string `json:"ip"`                     //IP
}

type Authentication struct {
	token       string //Token
	application string //Application
	ip          string //IP
}

//Request model for an API token request
type ApiTokenRequest struct {
	password    string //The user's current password
	application string //The name of the application this token will be used for
}

type authInfo struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Grant_type    string `json:"grant_type"`
	Refresh_token string `json:"refresh_token"`
}

type AccessTokenList struct {
	list []*AccessToken
}

type AccessToken struct {
	id          string //Unique ID for this AccessToken
	token       string //The token used to authenticate with the API
	ip          string //The remote address of the client which created this token
	application string //The name of the application which will use this token
	expiration  string //Expiration date of the token
	active      int    //Whether or not the token is active
	refresh     bool   //Whether or not the token can be used to refresh an access token
}

//Create an API token with infinite expiration
func (api *hwapi) CreateToken(accountHash string, uid int, tokenRequest ...*ApiTokenRequest) (*Authentication, error) {
	if (accountHash == "" || uid == 0) && api.CurrentUser.AccountHash == "" {
		return nil, errors.New("accountHash must supplied or use AboutMe to generate current user info")
	}
	if accountHash == "" {
		accountHash = api.CurrentUser.AccountHash
	}
	if uid == 0 {
		uid = api.CurrentUser.Id
	}

	//Create create token request
	// r := http.NewRequest("POST", "/api/v1/accounts/"+accountHash+"/"+uid+"/tokens",)
	r, e := api.Request(
		&Request{
			Method: POST,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/users/%d/tokens", accountHash, uid),
			Body:   &tokenRequest,
		},
	)
	if e != nil {
		return nil, e
	}
	a := &Authentication{}
	return a, json.Unmarshal(r.body, a)
}

//Fetch all tokens associated with this user
func (api *hwapi) GetTokens(accountHash string, uid int) (*AccessTokenList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/users/%d/tokens", accountHash, uid),
		},
	)
	if e != nil {
		return nil, e
	}
	atl := &AccessTokenList{}
	return atl, json.Unmarshal(r.body, atl)
}

//Delete token
func (api *hwapi) DeleteToken(a string, uid int, tokenId int) (bool, error) {
	_, e := api.Request(&Request{
		Url:    fmt.Sprintf("/api/v1/accounts/%s/users/%d/tokens/%d", a, uid, tokenId),
		Method: DELETE,
	})
	if e != nil {
		return true, nil
	}
	return false, e
}

//Authenticate user or refresh an access token
func (api *hwapi) Auth(u string, p string) (*AuthToken, error) {
	r, e := api.Request(&Request{
		Method: POST,
		Url:    "/auth/token",
		Body: &authInfo{
			Grant_type: "password",
			Username:   u,
			Password:   p,
		},
	})
	if e != nil {
		return nil, e
	}
	return api.AuthToken, json.Unmarshal(r.body, &api.AuthToken)
}

//Set token
//Check if token available, than set to AuthToken if available
func (api *hwapi) SetToken(t string) {
	api.AuthToken.AccessToken = t
}

//Use /api/v1/users/me to check accesstoken vaildation
func (api *hwapi) checkToken() (bool, error) {
	_, e := api.AboutMe()
	if e != nil {
		return false, e
	}
	return true, nil
}

//Automate refresh token when token is not available.
//Note, this function had deprecated.
func (api *hwapi) RefreshToken(refreshT ...string) (*AuthToken, error) {
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
			Url:    "/auth/token",
			Body: &authInfo{
				Grant_type:    "refresh_token",
				Refresh_token: t,
			},
		},
	)
	if e != nil {
		return nil, e
	}
	return api.AuthToken, json.Unmarshal(r.body, &api.AuthToken)
}
