package hwapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	apiErrors "github.com/bucloud/hwapi/errors"
)

// Request used by API requests only, different from http.Request
type Request struct {
	URL     string //request url
	Method  string //request method
	Query   map[string]string
	Headers map[string]string
	Body    interface{}
	Options map[string]string
}

// Response simple response
type Response struct {
	StatusCode int
	StatusText string
	body       []byte
	Headers    http.Header
}

const (
	apiBase    = "https://striketracker.highwinds.com"
	authURL    = "https://hcs.hwcdn.net/stauth/v1.0"
	storageURL = "https://hcs.hwcdn.net/v1/AUTH_hwcdn-logstore"

	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	UPDATE = "PUT"
	PUT    = "PUT"
)

//Request wrap
func (api *HWApi) Request(req *Request) (*Response, error) {
	if !strings.HasPrefix(req.URL, "http") {
		if strings.HasPrefix(req.URL, "/") {
			req.URL = apiBase + req.URL
		} else {
			req.URL = apiBase + "/" + req.URL
		}
	}
	//parse body
	buf := &bytes.Buffer{}
	switch t := req.Body.(type) {
	case int:
		buf.WriteString(string(t))
	case int64:
		buf.WriteString(string(t))
	case int32:
		buf.WriteString(string(t))
	case string:
		buf.WriteString(t)
	case bool:
	case nil:
	default:
		j, e := json.Marshal(t)
		if e != nil {
			return nil, errors.New("Convert body to []byte error")
		}
		buf.Write(j)
	}
	//parse request method
	if req.Method == "" {
		if buf.Len() == 0 {
			req.Method = "GET"
		} else {
			req.Method = "POST"
		}
	}
	//parse request query strings
	queryString := []string{}
	for qk, qv := range req.Query {
		switch strings.ToLower(qk) {
		// used for analytics
		case "groupby":
			queryString = append(queryString, "groupBy="+qv)
		case "startdate":
			queryString = append(queryString, "startDate="+qv)
		case "enddate":
			queryString = append(queryString, "endDate="+qv)
		case "billingregions":
			queryString = append(queryString, "billingRegions="+qv)
		case "statuscodes":
			queryString = append(queryString, "statusCodes="+qv)
		case "statuscategories":
			queryString = append(queryString, "statusCategories="+qv)
		case "includemessage":
			queryString = append(queryString, "includeMessage="+qv)
		case "maxresults":
			queryString = append(queryString, "maxResults="+qv)
		default:
			queryString = append(queryString, strings.ToLower(qk)+"="+qv)
		}
	}
	if len(queryString) != 0 {
		if strings.Index(req.URL, "?") < 0 {
			req.URL += "?"
		} else {
			req.URL += "&"
		}
		req.URL += strings.Join(uniqueSlice(queryString), "&")
	}
	//parse request headers
	r, ee := http.NewRequest(req.Method, req.URL, buf)
	if ee != nil {
		panic(ee)
	}
	for k, v := range req.Headers {
		r.Header.Set(k, v)
	}
	//Call fetch
	return api.Fetch(r)
}

func (api *HWApi) addHeaders(req *http.Request) {
	// if !strings.Contains(req.URL.Host, "hcs.hwcdn") && !strings.HasSuffix(req.URL.Path, "auth/token") && api.AuthToken == nil {
	// 	return nil, errors.New("This endpoint requires authentication")
	// }
	if req.Header == nil {
		req.Header = http.Header{}
	}
	req.Header.Set("X-Application", "GO-HWApi")
	req.Header.Set("X-Application-Id", "GO-HWApi")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 GO-HWApi/0.1")

	if api.AuthToken == nil {
		api.AuthToken = &AuthToken{}
	}
	if strings.Contains(req.URL.Host, "hcs.hwcdn") && api.AuthToken.LogTokens != "" {
		req.Header.Set("X-Auth-Token", api.AuthToken.LogTokens)
	} else if api.AuthToken.AccessToken != "" {
		req.Header.Set("Authorization", strFirstToUpper(api.AuthToken.TokenType)+" "+api.AuthToken.AccessToken)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json, text/plain, */*")
	}
}

//Fetch wrap http.Request add required headers and parse response
func (api *HWApi) Fetch(req *http.Request) (*Response, error) {
	api.addHeaders(req)
	rep, err := api.hc.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	defer rep.Body.Close()
	d, ioerr := ioutil.ReadAll(rep.Body)
	if ioerr != nil {
		return nil, errors.New("parse response failed")
	}

	if rep.StatusCode > 300 || rep.StatusCode < 200 {
		//tre parse error info in response
		errorInfo := &apiErrors.ErrorResponse{}
		e := json.Unmarshal(d, errorInfo)
		if e != nil {
			return nil, errors.New(req.URL.String() + " : " + rep.Status)
		}
		return nil, errors.New(req.URL.String() + " : " + errorInfo.Error)
	}
	return &Response{
		StatusCode: rep.StatusCode,
		StatusText: rep.Status,
		body:       d,
		Headers:    rep.Header,
	}, nil
}

func strFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}
