package hwapi

import (
	"net/http"
)

type HWApi struct {
	hc          *http.Transport
	AuthToken   *AuthToken
	authInfo    *authInfo
	CurrentUser *User
}

func init() {
}

//Initiation HWApi
//Default timeout is 30s and maxConns is 10
func Init(tr *http.Transport) *HWApi {
	return &HWApi{
		hc: tr,
	}
}
