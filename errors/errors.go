package errors

import (
	"encoding/json"
	"strconv"
)

const (
	CODE_FATAL_ERROR = iota
	CODE_GENERAL_ERROR
	CODE_GENERAL_MISSING_PARAMETER
	CODE_GENERAL_INVALID_HASH_CODE
	CODE_GENERAL_ACCOUNT_CONTEXT_NOT_FOUND
	CODE_GENERAL_INVALID_JSON
	CODE_GENERAL_FEATURE_DISABLED
	CODE_GENERAL_INVALID_CONFIG
)
const (
	CODE_ANALYTICS_INVALID_GROUPBY = iota + 100
	CODE_ANALYTICS_INVALID_GRANULARITY
	CODE_ANALYTICS_INVALID_FILTER
	CODE_ANALYTICS_INVALID_DATE_RANGE
	CODE_ANALYTICS_END_DATE_TOO_FAR
	CODE_ANALYTICS_QUERY_RESULTS_EXPIRED
	CODE_ANALYTICS_QUERY_RESULTS_UNAVAILABLE
	CODE_ANALYTICS_BULK_ANALYTICS_SERVICE_NOT_ENABLED
)

const (
	CODE_AUTH_INVALID_GRANT_TYPE = iota + 200
	CODE_AUTH_INVALID_ACCOUNT_CONTEXT
	CODE_AUTH_NOT_AUTHENTICATED
	CODE_AUTH_PASSWORD_EXPIRED
	CODE_AUTH_INVALID_CONFIRMATION
	CODE_IP_WHITELIST_VIOLATION
)

const (
	CODE_ACL_INSUFFICIENT_PERMISSIONS = iota + 300
	CODE_ACL_USER_SUSPENDED
	CODE_ACL_ACCOUNT_SUSPENDED
	CODE_ACL_USER_NO_ACCOUNT
)

const (
	CODE_VALIDATION_FAILED = iota + 400
	CODE_VALIDATION_DUPLICATE_ORIGIN
	CODE_VALIDATION_RECORD_NOT_FOUND = iota + 402
	CODE_VALIDATION_ENDPOINT_NOT_FOUND
	CODE_VALIDATION_CONFLICT = iota + 405
	CODE_VALIDATION_GONE
	CODE_VALIDATION_CONFLICT_WILDCARD
	CODE_RESOURCE_LOCK = 423
	CODE_RATE_LIMIT    = 429
)

const (
	CODE_INFRASTRUCTURE_DATABASE_UNAVAILABLE = iota + 503
	CODE_INFRASTRUCTURE_CDN_UNAVAILABLE
	CODE_INFRASTRUCTURE_API_UNAVAILABLE
	CODE_INFRASTRUCTURE_HCS_UNAVAILABLE
	CODE_INFRASTRUCTURE_SOLR_UNAVAILABLE
	CODE_INFRASTRUCTURE_NRT_UNAVAILABLE
	CODE_INFRASTRUCTURE_ANALYTICS_TIMEOUT
	CODE_INFRASTRUCTURE_ANALYTICS_RESOURCE_UNAVAILABLE
	CODE_INFRASTRUCTURE_ANALYTICS_UNAVAILABLE
)

const (
	CODE_HCS_VALIDATION_FAILED = iota + 600
	CODE_HCS_INVALID_AUTH_TOKEN
	CODE_HCS_SERVICE_NOT_ENABLED
	CODE_HCS_MAX_TENANTS_EXCEEDED
)

const (
	CODE_EVERYSTREAM_SERVICE_NOT_ENABLED = iota + 700
	CODE_EVERYSTREAM_ACCOUNT_NOT_FOUND
	CODE_EVERYSTREAM_ACCOUNT_SUSPENDED
	CODE_EVERYSTREAM_JOB_IN_PROGRESS
	CODE_EVERYSTREAM_ENCODE_QUOTA_EXCEEDED
	CODE_EVERYSTREAM_TRANSMUX_SERVICE_NOT_ENABLED
	CODE_EVERYSTREAM_TRANSMUX_NOT_PROVISIONED
	CODE_EVERYSTREAM_TRANSMUX_SUSPENDED
)

//403,201 error code
type HCConnectError struct {
	code        int
	category    string
	description string
}
type HCRequestError struct {
	code        int
	category    string
	description string
}

type APIError struct {
	code        int
	category    string
	description string
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

var errorDescription map[int]string = map[int]string{
	CODE_FATAL_ERROR:                                   "An unhandled fatal error has occurred",
	CODE_GENERAL_ERROR:                                 "Generic error",
	CODE_GENERAL_MISSING_PARAMETER:                     "A required parameter is missing from a request",
	CODE_GENERAL_INVALID_HASH_CODE:                     "A hash code given is not of a valid format",
	CODE_GENERAL_ACCOUNT_CONTEXT_NOT_FOUND:             "The account request in the account context was not found",
	CODE_GENERAL_INVALID_JSON:                          "The request data is not valid JSON",
	CODE_GENERAL_FEATURE_DISABLED:                      "This feature is temporarily disabled",
	CODE_GENERAL_INVALID_CONFIG:                        "An expected configuration value was not set",
	CODE_ANALYTICS_INVALID_GROUPBY:                     "The groupBy parameter on an analytics request is missing or invalid",
	CODE_ANALYTICS_INVALID_GRANULARITY:                 "An invalid granularity parameter was provided for an analytics request",
	CODE_ANALYTICS_INVALID_FILTER:                      "A required filter is missing or a filter contains a bad value",
	CODE_ANALYTICS_INVALID_DATE_RANGE:                  "The date range provided for an analytics request is not valid",
	CODE_ANALYTICS_END_DATE_TOO_FAR:                    "The end date provided is too far in the future (past the end of the current month)",
	CODE_ANALYTICS_QUERY_RESULTS_EXPIRED:               "The query results for the requested job ID have expired",
	CODE_ANALYTICS_QUERY_RESULTS_UNAVAILABLE:           "The query results are not yet available (probably because the query is still running). @var int",
	CODE_ANALYTICS_BULK_ANALYTICS_SERVICE_NOT_ENABLED:  "The account sending the request does not have the Bulk Analytics service enabled.",
	CODE_AUTH_INVALID_GRANT_TYPE:                       "The grant_type parameter for authentication is missing or an invalid value",
	CODE_AUTH_INVALID_ACCOUNT_CONTEXT:                  "The specified account context is invalid (account is suspended or deleted)",
	CODE_AUTH_NOT_AUTHENTICATED:                        "Resource requires authentication but user is not authenticated",
	CODE_AUTH_PASSWORD_EXPIRED:                         "Resource requires authentication but the user's password is expired",
	CODE_AUTH_INVALID_CONFIRMATION:                     "Confirmation code is invalid or expired",
	CODE_IP_WHITELIST_VIOLATION:                        "The Client IP is not in the Users IP Whitelist",
	CODE_ACL_INSUFFICIENT_PERMISSIONS:                  "The authenticated user does not have permissions to perform the requested action",
	CODE_ACL_USER_SUSPENDED:                            "The authenticated user is suspended",
	CODE_ACL_ACCOUNT_SUSPENDED:                         "The account is suspended",
	CODE_ACL_USER_NO_ACCOUNT:                           "The authenticated user does not have an associated account",
	CODE_VALIDATION_FAILED:                             "A resource has failed validation",
	CODE_VALIDATION_DUPLICATE_ORIGIN:                   "A duplicate origin exists with the same hostname, port, and path",
	CODE_VALIDATION_RECORD_NOT_FOUND:                   "The requested resource was not found",
	CODE_VALIDATION_ENDPOINT_NOT_FOUND:                 "The requested endpoint was not found, please check your url",
	CODE_VALIDATION_CONFLICT:                           "This resource already exists",
	CODE_VALIDATION_GONE:                               "This resource has been deleted or is expired",
	CODE_VALIDATION_CONFLICT_WILDCARD:                  "There is a conflict with a wildcard hostname",
	CODE_RESOURCE_LOCK:                                 "The requested resource is locked",
	CODE_RATE_LIMIT:                                    "Your use of this resource exceeds specified rate limit",
	CODE_INFRASTRUCTURE_DATABASE_UNAVAILABLE:           "Unable to reach the database. Please try again later.",
	CODE_INFRASTRUCTURE_CDN_UNAVAILABLE:                "Unable to send configuration to the CDN. Please try again later.",
	CODE_INFRASTRUCTURE_API_UNAVAILABLE:                "Application is currently down for maintenance. Please try again later.",
	CODE_INFRASTRUCTURE_HCS_UNAVAILABLE:                "Unable to reach HCS. Please try again later.",
	CODE_INFRASTRUCTURE_SOLR_UNAVAILABLE:               "Unable to reach the SOLR API. Please try again later.",
	CODE_INFRASTRUCTURE_NRT_UNAVAILABLE:                "Unable to retrieve current analytics data",
	CODE_INFRASTRUCTURE_ANALYTICS_TIMEOUT:              "Analytics request timed out. Please try again later.",
	CODE_INFRASTRUCTURE_ANALYTICS_RESOURCE_UNAVAILABLE: "Unable to find analytics resource. Please try again later.",
	CODE_INFRASTRUCTURE_ANALYTICS_UNAVAILABLE:          "Unable to reach analytics database. Please try again later.",
	CODE_HCS_VALIDATION_FAILED:                         "HCS validation failed",
	CODE_HCS_INVALID_AUTH_TOKEN:                        "HCS Invalid Auth Token",
	CODE_HCS_SERVICE_NOT_ENABLED:                       "HCS Service Not Enabled",
	CODE_HCS_MAX_TENANTS_EXCEEDED:                      "HCS Max tenants exceeded",
	CODE_EVERYSTREAM_SERVICE_NOT_ENABLED:               "EveryStream Service Not Enabled",
	CODE_EVERYSTREAM_ACCOUNT_NOT_FOUND:                 "EveryStream Account Not Found",
	CODE_EVERYSTREAM_ACCOUNT_SUSPENDED:                 "EveryStream Account Suspended",
	CODE_EVERYSTREAM_JOB_IN_PROGRESS:                   "Encoding Job already in progress",
	CODE_EVERYSTREAM_ENCODE_QUOTA_EXCEEDED:             "EveryStream encoded quota exceeded",
	CODE_EVERYSTREAM_TRANSMUX_SERVICE_NOT_ENABLED:      "Transmux Service is not enabled on this account",
	CODE_EVERYSTREAM_TRANSMUX_NOT_PROVISIONED:          "Transmux Service has not been provisioned",
	CODE_EVERYSTREAM_TRANSMUX_SUSPENDED:                "Transmux Service is suspended",
}

// UnmarshalJSON custom unmarshal configuration
func (c *ErrorResponse) UnmarshalJSON(b []byte) error {
	type t ErrorResponse
	response := &t{}
	if err := json.Unmarshal(b, response); err != nil {
		return err
	}
	if errorDescription[response.Code] != "" {
		response.Error = errorDescription[response.Code]
	}
	// setDefaultField(conf, reflect.StructTag(""))
	*c = (ErrorResponse)(*response)
	return nil
}

func (e *HCConnectError) Error() string {
	return "Connect error " + strconv.Itoa(e.code) + " : " + e.description
}

func (e *HCRequestError) Error() string {
	return "Request error " + strconv.Itoa(e.code) + " : " + e.description
}

func (e *APIError) Error() string {
	return "API error " + strconv.Itoa(e.code) + " : " + e.description
}
