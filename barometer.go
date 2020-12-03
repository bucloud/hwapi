package hwapi

import (
	"encoding/json"
	"strings"
)

// TraceRoute result list
type TraceRoute struct {
	List []*TraceRouteResponse `json:"list"`
}

// TraceRouteResponse traceroute info
type TraceRouteResponse struct {
	DataCenter  string //POP COde
	IPAddress   string //Hostname's IPAddress
	RoundTripMS int    //RTT
	HopCount    int    //hop
}

// BarometerList list contains multiple BarometerResponse
type BarometerList struct {
	List []*BarometerResponse `json:"list"`
}

// BarometerResponse http barometer response
type BarometerResponse struct {
	Datacenter  string            //Datacenter
	DNSMs       int               //dnsMS
	ConnectMS   int               //connectMS
	SslMS       int               //sslMS
	FirstByteMS int               //firstByteMS
	DownloadMS  int               //downloadMS
	TotalMS     int               //totalMS
	Status      int               //Status
	Headers     map[string]string //Header
}

// BarometerTrace Get the list of traceroute responses for a given url
//hostName is Address to test
//pops is List of pops to test
func (api *HWApi) BarometerTrace(hostName string, pops ...string) (*TraceRouteResponse, error) {
	pop := ""
	for _, k := range pops {
		pop += k + ","
	}
	var q map[string]string
	q["hostname"] = hostName
	if pop != "" {
		q["pop"] = strings.TrimRight(pop, ",")
	}
	if r, e := api.Request(&Request{
		URL: "/api/v1/barometer/traceroute",
		Query: map[string]string{
			"hostname": hostName,
			"pop":      pop,
		},
	}); e == nil {
		al := &TraceRouteResponse{}
		return al, json.Unmarshal(r.body, al)
	} else {
		return nil, e
	}
}

// BarometerRequest test performance between POPs and destionation
// hostname is Url to test
// pops is List of pops to test
func (api *HWApi) BarometerRequest(hostName string, pops ...string) (*BarometerResponse, error) {
	pop := ""
	for _, k := range pops {
		pop += k + ","
	}
	var q map[string]string
	q["hostname"] = hostName
	if pop != "" {
		q["pop"] = strings.TrimRight(pop, ",")
	}
	if r, e := api.Request(&Request{
		URL: "/api/v1/barometer/traceroute",
		Query: map[string]string{
			"hostname": hostName,
			"pop":      pop,
		},
	}); e == nil {
		al := &BarometerResponse{}
		return al, json.Unmarshal(r.body, al)
	} else {
		return nil, e
	}
}
