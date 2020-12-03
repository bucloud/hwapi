package hwapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
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

var (
	downloadWorker chan downloadJob
	wg             sync.WaitGroup
)

// SearchLogs search logs
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
		return fmt.Errorf("Parse interface{} to []byte failed")
	}
	api.cache.Set([]byte(k), b)
	return api.cache.SaveToFile(cacheFilePath)
}

// SetWorkers accesslogs download concurrent count
// maxConcurrent should less than 100
func (api *HWApi) SetWorkers(n uint) {
	if n >= 100 {
		api.workers = 100
	} else if n <= 0 {
		api.workers = 1
	} else {
		api.workers = n
	}
}

// Downloads wrap accesslogs download
// urls need to download while store in disk, you can re-call this method when error returned
func (api *HWApi) Downloads(destDir string, urls ...string) (bool, error) {
	// store this job and history urls in local temp file with logToken as fileName
	// md5 := md5.New()
	if api.workers == 1 {
		for _, u := range urls {
			if _, e := api.download(destDir, u); e != nil {
				return false, e
			}
		}
	}
	// reset channel
	if downloadWorker != nil {
		downloadWorker = nil
	}
	defer func() { downloadWorker = nil }()
	downloadWorker = make(chan downloadJob, api.workers)
	// start worker
	for i := uint(1); i <= api.workers; i++ {
		go api.downloadConcurrently()
	}
	for _, u := range urls {
		downloadWorker <- downloadJob{
			URL:  u,
			Dest: destDir,
		}
	}
	close(downloadWorker)
	wg.Wait()
	return true, nil
}

// DownloadCurrently currently download logs,
func (api *HWApi) downloadConcurrently() {
	// store this job and history urls in local temp file with logToken as fileName
	// md5 := md5.New()
	wg.Add(1)
	for j := range downloadWorker {
		destPath := j.Dest + "/"
		u := j.URL
		if !strings.HasPrefix(u, "http") {
			u = storageURL + "/" + u
		}
		t := api.getCacheData(u)
		if t.State == 1 {
			continue
		}
		t.StartedDate = time.Now().UTC()
		defer api.saveState(u, t)
		url, e := url.Parse(strings.Trim(u, "\r"))
		if e != nil {
			t.State = 10
			fmt.Printf("parse accesslog url %s failed, %s", u, e.Error())
			api.saveState(u, t)
			continue
		}
		destPath += url.Path[:strings.LastIndex(url.Path, "/")]
		r, e2 := api.Fetch(&http.Request{
			Method: GET,
			URL:    url,
		})
		if e2 != nil {
			t.State = 11
			fmt.Printf("download accesslogs %s failed, %s", u, e2.Error())
			api.saveState(u, t)
			continue
		}
		t.Size = r.Headers.Get("Content-Length")
		// store body to dest
		if mkdirError := os.MkdirAll(destPath, 0755); mkdirError != nil {
			t.State = 20
			fmt.Printf("create dir %s failed, %s", destPath, mkdirError.Error())
			api.saveState(u, t)
			continue
		}
		f, fe := os.OpenFile(destPath+url.Path[strings.LastIndex(url.Path, "/"):], os.O_WRONLY|os.O_CREATE, 0755)
		if fe != nil {
			t.State = 12
			fmt.Printf("open file %s failed, %s", destPath+url.Path[strings.LastIndex(url.Path, "/"):], fe.Error())
			api.saveState(u, t)
			continue
		}
		if _, fwe := f.Write(r.body); fwe != nil {
			t.State = 13
			fmt.Printf("write logdata to file %s failed, %s", destPath+url.Path[strings.LastIndex(url.Path, "/"):], fwe.Error())
			api.saveState(u, t)
			continue
		}
		if closeError := f.Close(); closeError != nil {
			t.State = 14
			fmt.Printf("close file %s failed, %s", destPath+url.Path[strings.LastIndex(url.Path, "/"):], fe.Error())
			api.saveState(u, t)
			continue
		}
		t.State = 1
		api.saveState(u, t)
	}
	wg.Done()
}

// Download accesslogs
func (api *HWApi) download(destDir, u string) (bool, error) {
	// store this job and history urls in local temp file with logToken as fileName
	// md5 := md5.New()
	destPath := destDir + "/"
	if !strings.HasPrefix(u, "http") {
		u = storageURL + "/" + u
	}
	t := api.getCacheData(u)
	if t.State == 1 {
		return true, nil
	}
	t.StartedDate = time.Now().UTC()
	defer api.saveState(u, t)
	url, e := url.Parse(strings.Trim(u, "\r"))
	if e != nil {
		t.State = 10
		return false, e
	}
	destPath += url.Path[:strings.LastIndex(url.Path, "/")]
	r, e2 := api.Fetch(&http.Request{
		Method: GET,
		URL:    url,
	})
	if e2 != nil {
		t.State = 11
		return false, e2
	}
	t.Size = r.Headers.Get("Content-Length")
	// store body to dest
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
