package hwapi

import (
	"net"
	"net/http"
	"time"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/rs/zerolog"
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
	workers        int
	Log            *zerolog.Logger
}

var (
	// maxCacheSize mem/file cache max bytes
	maxCacheSize           int    = 128 * 1024 * 1024
	cacheFilePath          string = "./.state"
	forceSaveCacheInterval time.Duration
)

// Init HWApi
// support several options
//
// *http.Transport  use custom transport instead of defaultTransport
//
// *fastcache.Cache local cache, mainly used to store downloads state
//
// int  default download workers
//
// *zerolog.Logger log handler
//
// *hwapi.User  Current userinfo
//
// *hwapi.AuthToken  set default token
func Init(options ...interface{}) *HWApi {
	api := &HWApi{
		hc: &http.Transport{
			Proxy: nil,
			DialContext: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 60 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          10,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		workers: 1,
	}
	for _, opt := range options {
		switch opt.(type) {
		case int:
			api.workers = opt.(int)
		case *fastcache.Cache:
			api.cache = opt.(*fastcache.Cache)
		case *User:
			api.CurrentUser = opt.(*User)
		case *AuthToken:
			api.AuthToken = opt.(*AuthToken)
		case *zerolog.Logger:
			api.Log = opt.(*zerolog.Logger)
		case *LocalCacheConfig:
			cc := opt.(*LocalCacheConfig)
			if cc.FilePath != "" {
				cacheFilePath = cc.FilePath
			}
			if cc.MaxSize == 0 {
				maxCacheSize = cc.MaxSize
			}
			forceSaveCacheInterval = cc.ForceSaveInterval
		}
	}
	if api.cache == nil {
		api.cache = fastcache.LoadFromFileOrNew(cacheFilePath,maxCacheSize)
	}
	return api
}

// LocalCacheConfig config localCache
type LocalCacheConfig struct {
	FilePath          string
	MaxSize           int
	ForceSaveInterval time.Duration
}
