package hwapi

import (
	"net/http"

	"github.com/VictoriaMetrics/fastcache"
)

type HWApi struct {
	hc                  *http.Transport
	AuthToken           *AuthToken
	authInfo            *authInfo
	CurrentUser         *User
	cache               *fastcache.Cache
	downloadConcurrency uint
}

const (
	// maxCacheSize mem/file cache max bytes
	maxCacheSize  int    = 128 * 1024 * 1024
	cacheFilePath string = "./.state"
)

func init() {
}

//Init HWApi
//Default timeout is 30s and maxConns is 10
func Init(tr *http.Transport) *HWApi {
	return &HWApi{
		hc:                  tr,
		cache:               fastcache.LoadFromFileOrNew(cacheFilePath, maxCacheSize),
		downloadConcurrency: 1,
	}
}
