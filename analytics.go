package hwapi

import (
	"encoding/json"
	"fmt"
	"time"
)

//Analytics list of series
type Analytics struct {
	Series []*Series `json:"series"` //Status
}

// Series tsdb data
type Series struct {
	Type    string      `json:"type"`    //type
	Key     string      `json:"key"`     //Key
	Metrics []string    `json:"metrics"` //Metrics
	Data    [][]float64 `json:"data"`    //Data
}

// StorageData A user enters a start time and end time (the timeframe) for transfer data. When the API returns the transfer data, the timeframe of the answer may not exactly match the timeframe of the request. This happens because the API returns data in terms of time buckets defined by the requested granularity:
// Five minute granularity begins at the top of the hour and at each five-minute interval and ends after five minutes.
// Hourly granularity begins at the top of the hour and ends just before the end of the hour.
// Daily granularity begins at midnight and ends just before the next midnight.
// Monthly granularity begins at midnight of the first day of the month and ends just before midnight of the last day of the month.
// When a user asks the API for transfer data beginning at a particular minute/hour/day/month, the API displays data rounded to the beginning of the time bucket. For example, a user's:
// Minute-specific request rounds back to the nearest prior :00, :05, :10, :15, etc.
// Hour-specific request rounds back to the nearest prior hour, such as 1 am, 2 am, 3 am, etc.
// Day-specific request rounds back to the prior midnight.
// Month-specific request rounds back to midnight of the first day of the month.
// When a user asks the API for transfer data ending at a particular minute/hour/day/month, the API displays data rounded to the end of the next bucket and then subtracts five minutes. The reason the API subtracts five minutes is so that the next bucket is not included in the timeframe that displays. For example, a user's:
// Minute-specific request rounds forward to the next :00, :05, :10, :15, etc. and then subtracts five minutes.
// Hour-specific request rounds forward to the next hour, such as 1 am, 2 am, 3 am, etc. (which is part of the next hour bucket) and then subtracts five minutes so that only data from the the last hour in the timeframe displays.
// Day-specific request rounds forward to midnight (which is part of the next day bucket) and then subtracts five minutes so that only data from the last day in the timeframe displays.
// Month-specific request rounds forward to midnight of the first day of the next month and then subtracts five minutes so that only data from the last month in the timeframe displays.
// Applying the principles above, a user who requests:
// Minute data from 1:13 pm to 2:47 pm would receive data from 1:10 pm to 2:45 pm
// Hourly data from 2:17 am to 4:43 am would receive data from 2 am to 4:55 am
// Hourly data from 4 pm to 5 pm would receive data from 4 pm to 5:55 pm
// Daily data from 7 pm on 5/19 to 6 am to 5/20 would receive data from midnight on 5/19 to 11:55 pm on 5/20
// Daily data from midnight on 6/26 to midnight 6/27 would receive data from midnight on 6/26 to 11:55 pm on 6/27
// Monthly data from 8/13 to 10/17 would receive data from midnight on 8/1 to 11:55 pm on 10/31
// Monthly data from 4/1 to 5/1 would receive data from midnight on 4/1 to 11:55 pm on 5/31
// Requesting a large number of accounts in the filter or grouping by accounts with no filter may return a bad request exception. To alleviate, specify an account filter with fewer accounts per request.
type StorageData struct {
	usageTime                 time.Time // the timestamp at which this bucket started
	edgeStorageTotalB         float32   // total bytes stored on every edge, may contain duplicate assets
	edgeStorageMaxB           float32   // maximum five minute bucket of bytes stored
	edgeStorageMaxUsageTime   float32   // timestamp at which maximum bytes stored occurred
	edgeStorageMinB           float32   // minimum five minute bucket of bytes stored
	edgeStorageMinUsageTime   float32   // timestamp at which minimum bytes stored occurred
	edgeStorageMeanB          float32   // average of five minute bucket bytes stored
	edgeFileCountTotal        float32   // total assets stored on every edge, may contain duplicate assets
	edgeFileCountMax          float32   // maximum five minute bucket of files stored
	edgeFileCountMaxUsageTime float32   // timestamp at which maximum files stored occurred
	edgeFileCountMin          float32   // minimum five minute bucket of files stored
	edgeFileCountMinUsageTime float32   // timestamp at which minimum files stored occurred
	edgeFileCountMean         float32   // average of five minute bucket files stored
	edgeFileSizeMeanB         float32   // average of five minute bucket bytes stored
	lastUpdatedTime           float32   // last time this bucket was updated
}

// TransferData A user enters a start time and end time (the timeframe) for transfer data. When the API returns the transfer data, the timeframe of the answer may not exactly match the timeframe of the request. This happens because the API returns data in terms of time buckets defined by the requested granularity:
// Five minute granularity begins at the top of the hour and at each five-minute interval and ends after five minutes.
// Hourly granularity begins at the top of the hour and ends just before the end of the hour.
// Daily granularity begins at midnight and ends just before the next midnight.
// Monthly granularity begins at midnight of the first day of the month and ends just before midnight of the last day of the month.
// When a user asks the API for transfer data beginning at a particular minute/hour/day/month, the API displays data rounded to the beginning of the time bucket. For example, a user's:
// Minute-specific request rounds back to the nearest prior :00, :05, :10, :15, etc.
// Hour-specific request rounds back to the nearest prior hour, such as 1 am, 2 am, 3 am, etc.
// Day-specific request rounds back to the prior midnight.
// Month-specific request rounds back to midnight of the first day of the month.
// When a user asks the API for transfer data ending at a particular minute/hour/day/month, the API displays data rounded to the end of the next bucket and then subtracts five minutes. The reason the API subtracts five minutes is so that the next bucket is not included in the timeframe that displays. For example, a user's:
// Minute-specific request rounds forward to the next :00, :05, :10, :15, etc. and then subtracts five minutes.
// Hour-specific request rounds forward to the next hour, such as 1 am, 2 am, 3 am, etc. (which is part of the next hour bucket) and then subtracts five minutes so that only data from the the last hour in the timeframe displays.
// Day-specific request rounds forward to midnight (which is part of the next day bucket) and then subtracts five minutes so that only data from the last day in the timeframe displays.
// Month-specific request rounds forward to midnight of the first day of the next month and then subtracts five minutes so that only data from the last month in the timeframe displays.
// Applying the principles above, a user who requests:
// Minute data from 1:13 pm to 2:47 pm would receive data from 1:10 pm to 2:45 pm
// Hourly data from 2:17 am to 4:43 am would receive data from 2 am to 4:55 am
// Hourly data from 4 pm to 5 pm would receive data from 4 pm to 5:55 pm
// Daily data from 7 pm on 5/19 to 6 am to 5/20 would receive data from midnight on 5/19 to 11:55 pm on 5/20
// Daily data from midnight on 6/26 to midnight 6/27 would receive data from midnight on 6/26 to 11:55 pm on 6/27
// Monthly data from 8/13 to 10/17 would receive data from midnight on 8/1 to 11:55 pm on 10/31
// Monthly data from 4/1 to 5/1 would receive data from midnight on 4/1 to 11:55 pm on 5/31
// Requesting a large number of accounts in the filter or grouping by accounts with no filter may return a bad request exception. To alleviate, specify an account filter with fewer accounts per request.
type TransferData struct {
	usageTime               time.Time // the timestamp at which this bucket started
	xferUsedTotalMB         float32   // total MB transferred
	xferUsedMinMB           float32   // minimum five minute bucket of MB transferred
	xferUsedMaxMB           float32   // maximum five minute bucket of MB transferred
	xferUsedMeanMB          float32   // average of five minute buckets of MB transferred
	xferAttemptedTotalMB    float32   // total MB attempted
	durationTotal           float32   // total transfer time of all requests
	xferRateMaxMbps         float32   // maximum transfer rate
	xferRateMaxUsageTime    float32   // timestamp at which maximum transfer rate occurred
	xferRateMinMbps         float32   // minimum transfer rate
	xferRateMinUsageTime    float32   // timestamp at which minimum transfer rate occurred
	xferRateMeanMbps        float32   // average transfer rate
	requestsCountTotal      float32   // total requests
	requestsCountMin        float32   // minimum five minute bucket of requests per second
	requestsCountMax        float32   // maximum five minute bucket of requests per second
	requestsCountMean       float32   // average of five minute bucket requests per second
	rpsMax                  float32   // maximum requests per second
	rpsMaxUsageTime         float32   // timestamp at which maximum requests per second occurred
	rpsMin                  float32   // minimum requests per second
	rpsMinUsageTime         float32   // timestamp at which minimum requests per second occurred
	rpsMean                 float32   // mean requests per second
	lastUpdatedTime         float32   // last time this bucket was updated
	xferRateMbps            float32   // average transfer rate in Mbps
	userXferRateMbps        float32   // total transfer divided by duration
	rps                     float32   // requests per second, calculated as total requests divided by number of seconds in bucket
	completionRatio         float32   // completed requests divided by attempted requests
	responseSizeMeanMB      float32   // total MB transferred divided by number of requests
	peakToMeanMBRatio       float32   // maximum transfer rate divided by mean transfer rate
	peakToMeanRequestsRatio float32   // maximum requests per second divided by mean requests per second
}

// AnalyticsQuery query parameters
type AnalyticsQuery struct {
	StartDate        string `json:"startDate,omitempty"`        //The start date of the query range in ISO8601 format (e.g. 2013-11-01T00:00:00Z)
	EndDate          string `json:"endDate,omitempty"`          //The end date of the query range in ISO8601 format (e.g. 2013-11-02T00:00:00Z)
	Granularity      string `json:"granularity,omitempty"`      //The granularity of the data series in ISO8601 format
	Platforms        string `json:"platforms,omitempty"`        //Comma separated list of platforms to filter by (accepts ids, codes or types)
	POPs             string `json:"pops,omitempty"`             //Comma separated list of pops to filter by. Cannot be used in conjunction with the billingRegions filter.
	BillingRegions   string `json:"billingRegions,omitempty"`   //Comma separated list of billing regions to filter by. Cannot be used in conjunction with the pops filter.
	Accounts         string `json:"accounts,omitempty"`         //Comma separated list of account hashes to filter by. Cannot be used in conjunction with the hosts filter.
	Hosts            string `json:"hosts,omitempty"`            //Comma separated list of host hashes to filter by. Cannot be used in conjunction with the accounts filter.
	StatusCodes      string `json:"statusCode,omitempty"`       //Only usable when query status data Comma separated list of 3 digit http status codes to filter by.
	StatusCategories string `json:"statusCategories,omitempty"` //Only usable when query status data Comma separated list of 1 digit http status codes categories to filter by.
	GroupBy          string `json:"groupBy,omitempty"`          //Groups the results by the specified entity
}

// BatchJob Returns batch job status, or results if completed for specified job ID
//Path /api/v1/accounts/{account_hash}/analytics/poll
//Doesn't implemented
func (api *HWApi) BatchJob(accountHash string) {
	// r, e := api.Request(
	// &Request{
	// Method: GET,
	// Url:    fmt.Sprintf("/api/v1/accounts/%s/analytics/poll", accountHash),
	// },
	// )

	// if e != nil {
	// return nil, e
	// }
	// al := &SearchResult{}
	// return al, json.Unmarshal(r.body, al)
}

// GetStatusData Returns http status code analytics for the specified account hash
//Path /api/v1/accounts/{account_hash}/analytics/status
func (api *HWApi) GetStatusData(accountHash string, q *AnalyticsQuery) (*Analytics, error) {
	return api.GetAnalytics("status", accountHash, q)
}

// GetStorageData Returns account storage analytics for the specified account hash
//Path /api/v1/accounts/{account_hash}/analytics/storage
func (api *HWApi) GetStorageData(accountHash string, q *AnalyticsQuery) (*Analytics, error) {
	return api.GetAnalytics("storage", accountHash, q)
}

// GetTransferData Returns account transfer analytics for the specified account hash
//Path /api/v1/accounts/{account_hash}/analytics/transfer
func (api *HWApi) GetTransferData(accountHash string, q *AnalyticsQuery) (*Analytics, error) {
	return api.GetAnalytics("transfer", accountHash, q)
}

// GetAnalytics Get analytics Data wrap
func (api *HWApi) GetAnalytics(dt string, accountHash string, query interface{}) (*Analytics, error) {

	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/analytics/%s", accountHash, dt),
			Query:  query.(map[string]string),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Analytics{}
	return al, json.Unmarshal(r.body, al)
}
