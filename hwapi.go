package hwapi

import (
	"net/http"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/aws/aws-sdk-go/aws"
)

// HWApi highwinds API struct
type HWApi struct {
	hc             *http.Transport
	AuthToken      *AuthToken
	authInfo       *authInfo
	hcsCredentials *HCSCredentials
	remoteS3       map[string]*aws.Config
	CurrentUser    *User
	cache          *fastcache.Cache
	workers        uint
}

const (
	// maxCacheSize mem/file cache max bytes
	maxCacheSize  int    = 128 * 1024 * 1024
	cacheFilePath string = "./.state"
)

//Init HWApi
//Default timeout is 30s and maxConns is 10
func Init(tr *http.Transport) *HWApi {
	return &HWApi{
		hc:      tr,
		cache:   fastcache.LoadFromFileOrNew(cacheFilePath, maxCacheSize),
		workers: 1,
	}
}
