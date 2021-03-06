package hwapi

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var (
	downloadWorker chan downloadJob

	// ErrRemoteConfigNotFound remote configure not found
	ErrRemoteConfigNotFound = errors.New("AWS S3 remote configure not found")
	// ErrRemoteDestFormat remote dest format error
	ErrRemoteDestFormat = errors.New("remote dest format error, must been remoteConfigName:bucketName:/path")
	// ErrGCSCredentialsMissed configure error
	ErrGCSCredentialsMissed = errors.New("credentials missed, either privateKey or accessKey/secretKey pair should provided")
	// ErrConvertStateTOByte marshall state to byte failed
	ErrConvertStateTOByte = errors.New("Parse interface{} to []byte failed")
)

/*
Accesslogs are available via http&ftp but have different effection
1. http can search logs by hosthash and filename prefix, note time frame is available in prefix and useable, but http interface is just a method to substitute for ftp
2. login ftp with accounthash only, root dir is a list contains almost all hosthash, but it's not all hosthash!!!
3. login ftp with accounthash and hosthash, root dir is a list contains only cds/cdi or other logtype dir
NOTE, accesslogs only accessable by username&password
*/
type downloadJob struct {
	URL  string
	Dest string
}

// SearchLogsOptions used to search logs
type SearchLogsOptions struct {
	HostHash    string
	AccountHash string
	StartDate   time.Time
	EndDate     time.Time
	LogType     string
	*HCSCredentials
}

// HCSCredentials credential used to create HCS client
type HCSCredentials struct {
	// PrivateKeyJSON base64 encoded string
	PrivateKeyJSON string

	// accesskeyID pair
	AccessKeyID string
	SecretKey   string
}

// SetCredentials set global credentials
func (api *HWApi) SetCredentials(c *HCSCredentials) {
	api.hcsCredentials = c
}

// SearchLogsV2 search logs
// Search log file list, accountHash should supplied, if $end-$start > 1day, search action would act as multiple request, in order to avoid 10000 lines limitation
// Note this search method would search files according to ctime(create time)
// filename sample cds/2020/08/27/cds_20200827-210002-61686853007ch4.log.gz
func (api *HWApi) SearchLogsV2(opt *SearchLogsOptions) ([]string, error) {
	res := []string{}
	if opt.HCSCredentials == nil && api.hcsCredentials != nil {
		opt.HCSCredentials = api.hcsCredentials
	}
	if opt.PrivateKeyJSON == "" && (opt.AccessKeyID == "" || opt.SecretKey == "") {
		return []string{}, ErrGCSCredentialsMissed
	}
	if opt.LogType == "" {
		opt.LogType = "cds"
	}
	if opt.AccountHash == "" {
		opt.AccountHash = api.CurrentUser.AccountHash
	}
	bucketName := "sp-cdn-logs-" + opt.AccountHash
	markerStart := opt.HostHash + "/" + opt.StartDate.Format(opt.LogType+"/2006/01/02/"+opt.LogType+"_20060102-150405")
	markerEnd := opt.HostHash + "/" + opt.EndDate.Format(opt.LogType+"/2006/01/02/"+opt.LogType+"_20060102-150405")
	if opt.PrivateKeyJSON != "" {
		ctx := context.Background()
		jsonString, err := base64.RawStdEncoding.DecodeString(opt.PrivateKeyJSON)
		if err != nil {
			return res, err
		}
		client, err := storage.NewClient(ctx, option.WithCredentialsJSON(jsonString), option.WithScopes(storage.ScopeReadOnly))
		if err != nil {
			return res, err
		}
		conf, _ := google.JWTConfigFromJSON(jsonString)

		objects := client.Bucket(bucketName).Objects(ctx, &storage.Query{
			// Prefix:      opt.HostHash + "/",
			StartOffset: markerStart,
			EndOffset:   markerEnd,
		})
		for {
			object, err := objects.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return res, err
			}
			// generate v4 signed url
			newURL, err := storage.SignedURL(bucketName, object.Name, &storage.SignedURLOptions{
				Scheme:         storage.SigningSchemeV4,
				Method:         "GET",
				GoogleAccessID: conf.Email,
				PrivateKey:     conf.PrivateKey,
				Expires:        time.Now().Add(24 * time.Hour),
			})
			if err != nil {
				return res, err
			}
			res = append(res, newURL)
		}
		return res, nil
	}
	// try use S3 as handler
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(opt.AccessKeyID, opt.SecretKey, ""),
		Endpoint:    aws.String("http://storage.googleapis.com"),
		Region:      aws.String("us-east-4"),
	}))

	// Create a downloader with the session and default options
	svc := s3.New(sess)
	// downloader := s3manager.NewDownloader(sess)
	for {
		r, err := svc.ListObjects(&s3.ListObjectsInput{
			// Prefix: aws.String(opt.HostHash),
			Bucket: aws.String(bucketName),
			Marker: &markerStart,
			// StartAfter: aws.String("f6g4s8v3/cds/2020/12/15/cds_20201215-222826"),
		})
		if err != nil {
			return res, err
		}
		for j := 0; j < len(r.Contents); j++ {
			if *r.Contents[j].Key > markerEnd {
				return res, nil
			}
			req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String(*r.Contents[j].Key),
			})
			newURL, err := req.Presign(24 * time.Hour)
			if err != nil {
				return res, err
			}
			res = append(res, newURL)
		}
		if *r.IsTruncated {
			markerStart = *r.NextMarker
		} else {
			return res, nil
		}
	}
}

// SearchLogs search logs in HCS, Note deprecated after 2021-01-01
// Search log file list, accountHash should supplied, if $end-$start > 1day, search action would act as multiple request, in order to avoid 10000 lines limitation
// Note this search method would search files according to ctime(create time)
// filename sample cds/2020/08/27/cds_20200827-210002-61686853007ch4.log.gz
func (api *HWApi) SearchLogs(hosthash, logtype string, startDate, endDate time.Time) ([]string, error) {
	// // regexp replacer
	// rp := regexp.MustCompile(`^`)
	// create empty slice
	res := []string{}
	// if startdate > enddate reversal them
	if endDate.Before(startDate) {
		tempDate := endDate
		endDate = startDate
		startDate = tempDate
	}
	if api.AuthToken == nil || api.AuthToken.LogTokens == "" {
		return res, fmt.Errorf("log token not exists, make sure you have called auth method before using SearchLogs")
	}
	// split request to multiple request if end - start > 1day due to limitation of response
	// test response lines, if gt 10000, seperate request to two request
	r, e := api.Request(&Request{
		Method: GET,
		URL:    storageURL + "/" + hosthash,
		// timelayout 2006-01-02T15:04:05Z
		Query: map[string]string{
			"marker":     startDate.Format(logtype + "/2006/01/02/" + logtype + "_20060102-150405"),
			"end_marker": endDate.Format(logtype + "/2006/01/02/" + logtype + "_20060102-150405"),
		},
	})
	if e != nil {
		return res, fmt.Errorf("Search accesslogs failed, %s", e.Error())
	}

	lines := string(r.body)
	if len(strings.Split(lines, "\n")) >= 10000 {
		nd := endDate.Sub(startDate)
		r1, e1 := api.SearchLogs(hosthash, logtype, startDate, startDate.Add(nd/2))
		r2, e2 := api.SearchLogs(hosthash, logtype, startDate.Add(nd/2), endDate)
		if e1 != nil {
			return []string{}, e1
		}
		if e2 != nil {
			return []string{}, e2
		}
		res = append(res, r1...)
		res = append(res, r2...)
	} else {
		res = append(res, strings.Split(strings.ReplaceAll(hosthash+"/"+(strings.TrimSuffix(lines, "\n")), "\n", ";"+hosthash+"/"), ";")...)
	}

	return res, nil
}

type downState struct {
	StartedDate time.Time
	EndedDate   time.Time
	State       int
	Size        string
}

func (api *HWApi) getCacheData(key string) *downState {
	r := &downState{}
	d := api.cache.Get(nil, []byte(key))
	if len(d) == 0 {
		return r
	}
	if e := json.Unmarshal(d, r); e != nil {
		return r
	}
	return r
}
func (api *HWApi) saveState(k string, v *downState) error {
	v.EndedDate = time.Now().UTC()
	b, e := json.Marshal(v)
	if e != nil {
		return ErrConvertStateTOByte
	}
	api.cache.Set([]byte(k), b)
	return nil
	// return api.cache.SaveToFile(cacheFilePath)
}

// SetWorkers accesslogs download concurrent count
// maxConcurrent should less than 100
func (api *HWApi) SetWorkers(n int) {
	if n >= 100 {
		api.workers = 100
	} else if n <= 0 {
		api.workers = 1
	} else {
		api.workers = n
	}
}

// SetRemoteS3Conf set AWS s3 conf, used when download raw logs to AWS s3
func (api *HWApi) SetRemoteS3Conf(remoteName string, conf *aws.Config) {
	if api.remoteS3 == nil {
		api.remoteS3 = map[string]*aws.Config{}
	}
	api.remoteS3[remoteName] = conf
}

// Downloads wrap accesslogs download
// destDir support cloud storage, format remoteConfigName:bucketName:path1/path2
// urls need to download while store in disk, you can re-call this method when error returned
func (api *HWApi) Downloads(destDir string, urls ...string) (bool, error) {
	// store this job and history urls in local temp file with logToken as fileName
	// reset channel
	if downloadWorker != nil {
		downloadWorker = nil
	}
	defer func() {
		_, ok := <-downloadWorker
		if ok {
			if api.Log != nil {
				api.Log.Error().Msg("concurrent downloads failed")
			}
			close(downloadWorker)
		}
		if e := api.cache.SaveToFile(cacheFilePath); e != nil && api.Log != nil {
			api.Log.Error().Str("dest", cacheFilePath).Err(e).Msg("save cache file failed")
		}
	}()
	downloadWorker = make(chan downloadJob, api.workers)
	var wg sync.WaitGroup
	wg.Add(api.workers)
	// start worker
	for i := 1; i <= api.workers; i++ {
		go func() {
			api.downloadConcurrently()
			wg.Done()
		}()
	}
	var remoteName, remotePath, bucketName string
	var S3 *s3.S3
	if strings.Index(destDir, ":") > 0 {
		a := strings.Split(destDir, ":")
		if len(a) != 3 {
			return false, ErrRemoteDestFormat
		}
		remoteName = a[0]
		bucketName = a[1]
		remotePath = a[2]
		if api.remoteS3 == nil || api.remoteS3[remoteName] == nil {
			return false, ErrRemoteConfigNotFound
		}
		sess, err := session.NewSession(api.remoteS3[remoteName])
		if err != nil {
			return false, err
		}
		S3 = s3.New(sess)
	} else {
		remoteName = ""
		remotePath = destDir
	}
	for _, u := range urls {
		if strings.Index(destDir, ":") > 0 {
			resp, _ := S3.PutObjectRequest(&s3.PutObjectInput{
				Bucket: &bucketName,
				Key:    aws.String(remotePath + regexp.MustCompile(`.*([0-9]{4}\/[0-9]{2}\/[0-9]{2}\/[^\/]+)\?.*$`).ReplaceAllString(u, "$1")),
			})
			dstURL, err := resp.Presign(24 * time.Hour)
			if err != nil {
				if api.Log != nil {
					api.Log.Debug().Str("url", u).Err(err).Msg("presign url failed")
				}
				return false, err
			}
			downloadWorker <- downloadJob{
				URL:  u,
				Dest: dstURL,
			}
		} else {
			downloadWorker <- downloadJob{
				URL:  u,
				Dest: destDir,
			}
		}
	}
	close(downloadWorker)
	wg.Wait()
	return true, nil
}

// DownloadCurrently currently download logs,
func (api *HWApi) downloadConcurrently() {
	// store this job and history urls in local temp file with logToken as fileName
	for j := range downloadWorker {
		if _, e := api.download(j.Dest, j.URL); e != nil {
			if api.Log != nil {
				api.Log.Error().Str("url", j.URL).Str("dest", j.Dest).Err(e).Msg("download failed")
			}
		}
		if api.Log != nil {
			api.Log.Debug().Str("url", j.URL).Str("dest", j.Dest).Msg("download success")
		}
	}
}

func md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Download accesslogs
func (api *HWApi) download(destDir, u string) (bool, error) {
	// store this job and history urls in local temp file with logToken as fileName
	// md5 := md5.New()
	if !strings.HasPrefix(u, "http") {
		u = storageURL + "/" + u
	}
	url, _ := url.Parse(strings.Trim(u, "\r"))
	t := api.getCacheData(md5String(url.Path))
	defer api.saveState(md5String(url.Path), t)

	if t.State == 1 {
		return true, nil
	}
	t.StartedDate = time.Now().UTC()
	defer func() {
		api.saveState(md5String(url.Path), t)
		if api.Log != nil {
			api.Log.Debug().Str("path", url.Path).Int("result_code", t.State).Time("started", t.StartedDate).Float64("spent", time.Since(t.StartedDate).Seconds()).Msg("download file ended")
		}
	}()
	r, e2 := api.Fetch(&http.Request{
		Method: GET,
		URL:    url,
	})
	if e2 != nil {
		t.State = 11
		return false, e2
	}
	t.Size = r.Headers.Get("Content-Length")
	// try upload to remote
	if strings.HasPrefix(destDir, "http") {
		putRequest, _ := http.NewRequest(PUT, destDir, bytes.NewReader(r.body))
		_, e2 := api.Fetch(putRequest)
		if e2 != nil {
			t.State = 15
			return false, e2
		}
		t.State = 1
		return true, nil
	}
	// store body to dest

	var destPath string
	if destDir == "" || destDir == "." || destDir == "./" {
		destPath = "./"
		// destPath += strings.Replace(url.Path[:strings.LastIndex(url.Path, "/")], "v1/AUTH_hwcdn-logstore", "", 1)
		destPath += regexp.MustCompile(`.*([0-9]{4}\/[0-9]{2}\/[0-9]{2}\/)[^\/]+$`).ReplaceAllString(url.Path, "$1")
	} else {
		destPath = destDir + "/"
	}

	if mkdirError := os.MkdirAll(destPath, 0755); mkdirError != nil {
		t.State = 20
		return false, mkdirError
	}
	f, fe := os.OpenFile(destPath+url.Path[strings.LastIndex(url.Path, "/"):], os.O_WRONLY|os.O_CREATE, 0755)
	if fe != nil {
		t.State = 12
		return false, fe
	}
	if _, fwe := f.Write(r.body); fwe != nil {
		t.State = 13
		return false, fwe
	}
	if closeError := f.Close(); closeError != nil {
		t.State = 14
		return false, closeError
	}
	t.State = 1
	return true, nil
}
