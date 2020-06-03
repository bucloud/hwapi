package hwapi

import (
	"net/http"
	"time"
)

type hwapi struct {
	hc          *http.Client
	AuthToken   *AuthToken
	authInfo    *authInfo
	CurrentUser *User
}

func init() {
}

//Initiation HWApi
//Default timeout is 30s and maxConns is 10
func Init(opts ...http.Transport) *hwapi {
	if opts[0].MaxIdleConns == 0 {
		opts[0].MaxIdleConns = 10
	}
	if opts[0].IdleConnTimeout == 0 {
		opts[0].IdleConnTimeout = 30
	}
	if opts[0].DisableCompression == false {
		opts[0].DisableCompression = true
	}
	tr := &http.Transport{
		MaxIdleConns:       opts[0].MaxIdleConns,
		IdleConnTimeout:    opts[0].IdleConnTimeout * time.Second,
		DisableCompression: opts[0].DisableCompression,
	}
	client := &http.Client{Transport: tr}
	return &hwapi{
		hc: client,
	}
}
