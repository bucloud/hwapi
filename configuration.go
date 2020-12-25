package hwapi

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// AccessLogger Configure settings relevant to the global settings that AccessLogger uses when storing access logs, origin pull logs, and receipt logs.
// AllowedScope PRODUCT
// DefaultPolicy  {"uploadToHCS":true,"enableCompression":true}
type AccessLogger struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// EnableCompression Enable gzip compression of access logs for this customer.
	EnableCompression *bool `json:"enableCompression," default:"1" role:"HWADMIN" writeonly:""`

	// UploadToHCS Upload access logs for this customer directly to Highwinds Cloud Storage
	UploadToHCS *bool `json:"uploadToHCS," default:"1" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// ExpireTimeHCS Time in seconds that an accesslog is allowed to live before it is expired from HCS
	ExpireTimeHCS uint32 `json:"expireTimeHCS,omitempty" default:"3888000" role:"HWADMIN" writeonly:""`

	// ExpireTimeLocal Time in seconds that an accesslog is allowed to live before it is expired from the accesslogger local storage
	// NOTE: This is used by SysEng's script to purge old access log files and the default value is subjected to change
	ExpireTimeLocal uint32 `json:"expireTimeLocal,omitempty" default:"3888000" role:"HWADMIN" writeonly:""`
}

// AccessLogs Configure settings relevant to Access Logs.
// AllowedScope DIR
// DefaultPolicy  {"enabled":false}
type AccessLogs struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled <p>Enable flag for this configuration type.</p>
	Enabled *bool `json:"enabled," default:"false" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// AccessLogIPObfuscation Enable/Disable IP address obfuscation in access logs for GDPR compliance.
// AllowedScope DIR
// DefaultPolicy  {"enabled":true}
type AccessLogIPObfuscation struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Enable IP address obfuscation of access logs for this customer. Complies with GDPR and obfuscates IPv4 addresses using /24 and IPv6 addresses using /96 bitmasks.
	// WARNING: DO NOT TURN THIS OFF UNLESS WE GOT A CLEAR FROM LEGAL AND SECURITY TEAM
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// AccessLogsConfig Configure settings relevant to Access Log Settings.
// AllowedScope DIR
// DefaultPolicy  null
type AccessLogsConfig struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// ExtraLogFields <p>Comma delimited list of HTTP header fields to append to the standard fields in the access log. Each field must have the 'sc:' (server-to-client) or 'cs:' (client-to-server) prefix.</p>
	// <p>Example: cs:Cookie, sc:x-custom-header</p>
	ExtraLogFields string `json:"extraLogFields,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// HostnameReporting Controls analytics and billing reporting by each unique hostname that maps to your site.
// AllowedScope PRODUCT
// DefaultPolicy  {"enabled":false}
type HostnameReporting struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Enables reporting by hostname on a site.
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// NrtReporting Near Real Time File Traffic Reporting
// AllowedScope PRODUCT
// DefaultPolicy  null
type NrtReporting struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// ReportVHost Enable realtime reporting by hostname.
	ReportVHost *bool `json:"reportVHost,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// OriginPullLogs Configure settings relevant to Origin Pull Logs.
// AllowedScope DIR
// DefaultPolicy  {"enabled":false}
type OriginPullLogs struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled <p>Enable flag for this configuration type.</p>
	Enabled *bool `json:"enabled," default:"false" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// OriginPullLogsConfig Configure settings relevant to Origin Pull Log Settings.
// AllowedScope DIR
// DefaultPolicy  null
type OriginPullLogsConfig struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// ExtraLogFields Comma delimited list of HTTP header fields to append to the standard fields in the origin pull log. Each field must have the 'sc:' (server-to-client) or 'cs:' (client-to-server) prefix.
	// <p>Example: cs:Cookie,sc:x-custom-header</p>
	ExtraLogFields string `json:"extraLogFields,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// ReceiptLogs Configure settings relevant to receipt logs.
// AllowedScope DIR
// DefaultPolicy  {"enabled":false}
type ReceiptLogs struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Enables receipt logs at the edge.
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// ReceiptLogsConfig Configure settings relevant to Receipt Log Settings.
// AllowedScope DIR
// DefaultPolicy  null
type ReceiptLogsConfig struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// ExtraLogFields <p>Comma delimited list of HTTP header fields to append to the standard fields in the receipt access logs. Each field must have the 'sc:' (server-to-client) or 'cs:' (client-to-server) prefix.</p>
	// <p>Example: cs:Cookie, sc:x-custom-header</p>
	ExtraLogFields string `json:"extraLogFields,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// RequestReceipt Delivery Receipts
// AllowedScope DIR
// DefaultPolicy  null
type RequestReceipt struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// UriFormat The Full GET URL for delivery receipts.  The URL entered in this field must specify the protocol, host (port is optional), and path. Query string parameters are required UNLESS a requestReceipt/headers policy is defined. Query string parameters are in the following format:  <name>=<value> (note the equal sign) where <name> is any HTTP legal query parameter name  and <value> is either a CDN Variable or static literal.
	URIFormat string `json:"uriFormat," default:"" role:"HWADMIN" writeonly:""`

	// VerifyCertificate Enable Cert Verification while doing SSL connection to Receipt Origin
	VerifyCertificate *bool `json:"verifyCertificate," default:"1" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// AddIdToAccessLog The add receipt identifier to access logs setting allows you to track delivery receipts in your access logs. By enabling this,  the CDN caching servers will add the X-HW-Receipt Header to each receipt's corresponding Client Request Access Log entry. This is not referring to the Receipt Access Log entry. If this feature is enabled, the customer must have access logging enabled (see the Customer conf type).
	AddIDToAccessLog *bool `json:"addIdToAccessLog,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// ClientResponseCodeFilter
	ClientResponseCodeFilter string `json:"clientResponseCodeFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// ClientResponseHeaderFilter
	ClientResponseHeaderFilter string `json:"clientResponseHeaderFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Headers A pipe delimited list of strings describing HTTP header fields to insert into  the delivery receipt. Each string should be in the form of a legal HTTP header with the following format: <name>: <value> (note the colon) where <name> is any HTTP legal header name and <value> is either a CDN Variable or static literal.
	Headers string `json:"headers,omitempty" default:"" role:"HWADMIN" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// CertificateCN Verify Origin Certificate's Common Name with Host specified in this policy. Hardcoded string (www.foo.com) or variable %origin.request.host% - Host header we send to Receipt Origin. If empty or not defined than CDN  will skip CN Verification.
	CertificateCN string `json:"certificateCN,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// MaxAge The receipt backlog TTL is the maximum age of a pending receipt. Receipts older than the specified amount are dropped. A value of zero  indicates that the receipt does not expire and to only try delivering the receipt one time unless MaxRetry is defined.
	MaxAge uint32 `json:"maxAge,omitempty" default:"0" role:"HWADMIN" writeonly:""`

	// MaxRetry The retry count is the maximum number of times to retry the delivery of a single receipt before discarding it. This count is in addition to the initial delivery attempt.  For example, a value of 3 means that a delivery edge will try  to deliver a receipt up to 4 times.  NOTE: if a MaxAge is also defined, then a receipt will be discarded if it expires  before the maximum number of retries has been reached.
	MaxRetry uint32 `json:"maxRetry,omitempty" default:"0" role:"HWADMIN" writeonly:""`
}

// RequestReceiptReportPercentage The delivery receipts report percentage policy allows you to configure the percentage of requests to provide delivery confirmation receipts.
// AllowedScope PRODUCT
// DefaultPolicy  {"dedupReportPercentage":100,"cacheHitReportPercentage":100,"originPullReportPercentage":100}
type RequestReceiptReportPercentage struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// DedupReportPercentage Percentage of OriginPull dedup queue request to report to the receipt server.
	DedupReportPercentage uint16 `json:"dedupReportPercentage,omitempty" default:"100" role:"HWADMIN" writeonly:"" range:"0,100"`

	// OriginPullReportPercentage Percentage of OriginPull request to report to the receipt server.
	OriginPullReportPercentage uint16 `json:"originPullReportPercentage,omitempty" default:"100" role:"HWADMIN" writeonly:"" range:"0,100"`

	// CacheHitReportPercentage Percentage of cache hit request to report to the receipt server.
	CacheHitReportPercentage uint16 `json:"cacheHitReportPercentage,omitempty" default:"100" role:"HWADMIN" writeonly:"" range:"0,100"`
}

// AwsSignedS3PostV4 Defines how to pre/sign post requests to be made by the CDN to an AWS origin.
// Note, even though this policy is groupable, if more than one policy is defined, only one policy will ever be applied.
// The CDN iterates over each policy until it finds the first match or applicable policy based on scope and/or filter.
// The CDN does not failover or attempt other policies if the chosen one failed. The Groupability was added with the
// primary intent to provide flexibilty when needing to define different AccessKeyId/SecretAccessKey combinations, such
// as using a popFilter to use one AccessKeyId/SecretAccessKey pair for a particular AWSRegion and another AccessKeyId/SecretAccessKey
// pair for a different AWSRegion.  Likewise, the site may only use one AccessKeyId/SecretAccessKey across multiple AWSRegions.
// AllowedScope DIR
// DefaultPolicy  null
type AwsSignedS3PostV4 struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// SecretAccessKey Shared secret key assigned to the AccessKeyID.
	SecretAccessKey string `json:"secretAccessKey," default:"" role:"normal" writeonly:"" advancedType:"PASSWORD"`

	// Enabled Set to true to enable policy.
	Enabled *bool `json:"enabled," default:"false" role:"normal" writeonly:""`

	// AwsRegion AWS region scope the access key.
	// see: https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
	AwsRegion string `json:"awsRegion," default:"" role:"normal" writeonly:""`

	// AccessKeyId Identifies the shared access key to be used to presign the request.
	AccessKeyID string `json:"accessKeyId," default:"" role:"normal" writeonly:""`

	// AuthenticationType Specifies what type of AWS authentication to use.
	// query: Provide authentication information using query string parameters. Using query parameters to authenticate requests is
	//   useful when you want to express a request entirely in a URL. This method is also referred as presigning a URL.
	// header: Use the HTTP Authorization header.
	// The query and header algorithms are identical except that the expireTimeSeconds policy only is applicable to the query authentication type.
	AuthenticationType string `json:"authenticationType,omitempty" default:"query" role:"normal" writeonly:"" enum:"[query,header]"`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// MethodFilter Method Filter is used to determine if this type should be applied or not based on List of HTTP Methods provided
	// Optionally, you may use an exclamation point in the list to describe the subset of HTTP methods excluded from this policy and all
	// other requests method will be included.
	// WARNING: You should not mix include and exclude in the same list.
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PopFilter POP filter is list of pattern to match POPs where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of POPs excluded from this policy.
	// Use lower case or use '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// RegionFilter Region filter is list of pattern to match Region where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of Regions excluded from this policy.
	// Use lower case or '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	RegionFilter string `json:"regionFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// SignedHeaders List of additional headers to be used when calculating the signature.
	// The headers "Host" and "x-amz-*" (customer AWS headers internall generated) are required and included by default.
	// Headers not permitted and invalidate the policy if set are "user-agent" and "x-amzn-trace-id".
	SignedHeaders string `json:"signedHeaders,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// AwsService AWS service scope of the access key.
	AwsService string `json:"awsService,omitempty" default:"s3" role:"normal" writeonly:""`

	// BucketIdentifier Specify to scope this policy to a particular bucket.  Note this value is directly coupled with the Host header, which is not always the origin hostname.
	// See https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html#VirtualHostingSpecifyBucket, Buckets are accessed primarily through the Host header
	// except when using SSL 'When you use virtual hosted–style buckets with Secure Sockets Layer (SSL), the SSL wildcard certificate only matches buckets that
	// don't contain periods. To work around this, use HTTP or write your own certificate verification logic. We recommend that you do not use periods (".") in
	// bucket names when using virtual hosted–style buckets.'
	// Buckets are identified in one of three ways:
	// 1) bucketname.s3.amazonaws.com
	// 2) s3.amazonaws.com/bucketname
	// 3) <custom.hostname>, such as www.myhost.com, where the host name is the bucketname
	// To match an origin request to the correct policy, the CDN appends the path of the URL in the origin request to the value in the Host header of the request.
	// The CDN checks if the constructed string "starts with" the value set in this policy.  The key factor is that the Host header is used, which may not
	// equal the origin hostname, such as the case a StaticHeader/OriginPull policy or proxying the Host header from a client request.
	// Leaving this blank/unset indicates it is the default policy to use for all origin pulls when a specific AwsSignedOriginPullV4 policy has not been matched.
	// If a default policy is used with one or more specific policies, the default needs to be listed last.
	BucketIdentifier string `json:"bucketIdentifier,omitempty" default:"" role:"normal" writeonly:""`

	// ExpireTimeSeconds Time period, in seconds, for which the generated presigned URL is valid.
	// Note, this policy only is applicable to the 'query' authentication type (see awsSignedOriginPullV4/authenticationType).
	ExpireTimeSeconds uint32 `json:"expireTimeSeconds,omitempty" default:"5" role:"normal" writeonly:""`
}

// AuthACL Enable access to content based on a customizable list of IP addresses.
// AllowedScope DIR
// DefaultPolicy  null
type AuthACL struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// AccessCode Access code that indicates whether to allow or deny the IP access to the requested content.
	AccessCode string `json:"accessCode," default:"" role:"normal" writeonly:"" enum:"[allow,deny]"`

	// IpList The list of IP addresses (or CIDR blocks) to that apply to this policy.  The IP addresses listed in this field will be allowed or denied based on the access directive provided in "Access Directive" field.
	IPList string `json:"ipList," default:"" role:"normal" writeonly:"" list:"IP"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Protocol Protocol for which this policy applies.
	Protocol string `json:"protocol,omitempty" default:"both" role:"normal" writeonly:"" enum:"[http,https,both]"`

	// ClientIPSrc <p>Source for the client IP to match against this policy.  Valid values are:</p>
	// <ul>
	// <li><b>socket</b>: IP address from the client connection is used.</li>
	// <li><b>header</b>: IP address from the specified header is used.</li>
	// </ul>
	ClientIPSrc string `json:"clientIPSrc,omitempty" default:"socket" role:"normal" writeonly:"" enum:"[socket,header]"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// Header Name of the http request header from which to obtain the client IP address when Client IP Source is set to header.
	Header string `json:"header,omitempty" default:"" role:"normal" writeonly:""`
}

// AuthGeo Restrict access to content based on the geographic location of the end-user.
// AllowedScope DIR
// DefaultPolicy  null
type AuthGeo struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Code The geographic code from MaxMind to apply.
	Code string `json:"code," default:"" role:"normal" writeonly:"" enum:"[countryCode,region,subdivisionCodes,city,postalCode,continentCode,timeZone,dmaCode,areaCode]"`

	// Values <p>Comma separated list of geographic codes for the region type selected. For an exclusion, use ! (exclamation). </p>
	// <p>You should not use both inclusions and exclusions in this list.  If you want to include the continent of Europe but exclude France, you must use two different types to express that.  If a request matches any of the include rules (or if there are no include rules), and that client does not match any exclude rules, they will be granted access.</p>
	Values string `json:"values," default:"" role:"normal" writeonly:"" list:"string"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// AuthHTTPBasic Require authentication in the form of a username and password from within an HTTP user agent, or web browser.
// AllowedScope DIR
// DefaultPolicy  null
type AuthHTTPBasic struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// BindingPoint <p>The URL to the authorization endpoint.</p>
	// <p><em>HTTPS URLs are currently not supported by this policy.</em></p>
	BindingPoint string `json:"bindingPoint," default:"" role:"normal" writeonly:""`

	// Realm The name of the authentication realm given back to the user on requests which don't contain credentials. For HTTP Basic Authentication, this value is usually displayed to the user when they are prompted for their login information.
	Realm string `json:"realm," default:"" role:"normal" writeonly:""`

	// Ttl Session timeout that an edge uses to avoid making an auth binding point call for each HTTP request. When it successfully authenticates a user, it will ask the user agent to set a cookie containing an encrypted authentication token and the TTL for the token.  Effectively, a given user should only be authenticated against the configured binding point once within the tokens TTL.
	TTL uint32 `json:"ttl," default:"" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// ConnectCount The maximum number of connections an edge server will make to the authentication binding point. This is an integer value not to exceed 99.
	ConnectCount uint32 `json:"connectCount,omitempty" default:"4096" role:"HWADMIN" writeonly:""`
}

// AuthReferer Restrict access to content based on a customizable list of websites or domains, or "referrers."
// AllowedScope DIR
// DefaultPolicy  null
type AuthReferer struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Referer The list of domains authorized to access the content requested.
	Referer string `json:"referer," default:"" role:"normal" writeonly:"" list:"GLOB"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// AuthSignUrlsInPlaylist Automatically apply my URL Signing policy to URLs inside my HLS playlists.
// AllowedScope DIR
// DefaultPolicy  null
type AuthSignUrlsInPlaylist struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// FilenamePatterns A list of glob pattern for files containing URL that needs to be signed.
	FilenamePatterns string `json:"filenamePatterns," default:"" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// UseCookie When signing the playlist, put the token in Set-Cookie of the response instead of in the URL's inside the m3u8 file.
	// NOTE: Currently, only the AKv2 signing is supported using Cookie, all other signing method will ignore this setting
	UseCookie *bool `json:"useCookie,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// CookieName When Cookie Authentication is enabled to sign the playlist, this is the name of the Cookie to be used to store the token. If Cookie is not enabled, this will be the name of the query string parameter used to store the whole signing token. For using with the AKv2 signing, this name should be one of the following: hdnea, hdnts or hdntl.
	// NOTE: Currently, only the AKv2 signing is supported using Cookie or storing the whole signing token inside a single QS param, all other signing method will ignore this setting
	CookieName string `json:"cookieName,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// ExtendTTL Sign the URL in the playlist with a diffrent TTL n seconds from the time of master playlist request.  No extending or re-signing by default when the value is set to 0 second.
	// NOTE: Because of the nature of the short life and long life token, only the AKv2 algorithm supports this feature.
	ExtendTTL uint32 `json:"extendTTL,omitempty" default:"0" role:"HWADMIN" writeonly:""`
}

// AuthURLSign Protect files from unauthorized access with an encrypted key.
// AllowedScope DIR
// DefaultPolicy  null
type AuthURLSign struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// TokenField Query string parameter name which contains the URL's MD5 token signature.
	TokenField string `json:"tokenField," default:"" role:"normal" writeonly:""`

	// PassPhraseField The name of the query string parameter to use when constructing the URL to input into the md5 hash function.
	PassPhraseField string `json:"passPhraseField," default:"" role:"normal" writeonly:""`

	// PassPhrase The shared secret used when signing URLs.
	PassPhrase string `json:"passPhrase," default:"" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// IgnoreFieldsAfterToken Ignore the fields after the Token field when verifying the URL signature. (Default: false)
	IgnoreFieldsAfterToken *bool `json:"ignoreFieldsAfterToken,omitempty" default:"false" role:"normal" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// IpAddressField <p>If present this will be a query string parameter containing an IP address of the client peer. The edge server will match the IP address in this query string parameter against the client requesting content for authorization.</p>
	// <p><b>NOTE:</b> Only IPv4 address are supported.</p>
	IPAddressField string `json:"ipAddressField,omitempty" default:"" role:"normal" writeonly:""`

	// UriLengthField <p>If present this will restrict the number of bytes in the path to consider for URL signing. For example, if this value is 10 and the request is for http://mydomain.com/this/is/my/path/to/a/file?queryStringStuff, then the MD5 will be calculated using the first 10 bytes of the path and the query string:</p>
	// <p>MD5("this/is/my?queryStringStuff")</p>
	// <p>A length value of 0 means it will strip off the filename and use directory only (with trailing '/') plus the query string parameters.</p>
	URILengthField string `json:"uriLengthField,omitempty" default:"" role:"normal" writeonly:""`

	// UserAgentField If present this will restrict access based on the user agent. It is not required that that user agent be added to the field on the original request, just that the user agent parameter be present. The user agent will automatically be taken from the request header.
	UserAgentField string `json:"userAgentField,omitempty" default:"" role:"normal" writeonly:""`

	// ExpiresField The query string parameter which contains Unix epoch time after which this link is considered invalid.
	ExpiresField string `json:"expiresField,omitempty" default:"" role:"normal" writeonly:""`
}

// AuthURLSignAliCloudA Ali Cloud Type-A URL Signing
// AllowedScope DIR
// DefaultPolicy  null
type AuthURLSignAliCloudA struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// PassPhrase Specify the shared passphrase, or sequence of words or other text, to use when generating the signature when authenticating a request made to the CDN.
	PassPhrase string `json:"passPhrase," default:"" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// IncludeParamsBeforeToken Set to true when query string parameters listed before the token should be included when generating the signature hash.
	IncludeParamsBeforeToken *bool `json:"includeParamsBeforeToken,omitempty" default:"false" role:"normal" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// TokenField Set to override the default name of the query string parameter that will be used by the publisher to specify the signature for the URL.
	TokenField string `json:"tokenField,omitempty" default:"auth_key" role:"normal" writeonly:""`

	// ExpirationExtension Number of seconds to add to the expiration time given in a request, which extends the life of the signature. This value does not affect the expiration value in the request nor does it affect the signature itself.
	ExpirationExtension uint32 `json:"expirationExtension,omitempty" default:"0" role:"normal" writeonly:""`
}

// AuthURLSignAliCloudB Ali Cloud Type-B URL Signing
// AllowedScope PRODUCT
// DefaultPolicy  null
type AuthURLSignAliCloudB struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// PassPhrase Specify the shared passphrase, or sequence of words or other text, to use when generating the signature when authenticating a request made to the CDN.
	PassPhrase string `json:"passPhrase," default:"" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// ExpirationExtension Number of seconds to add to the expiration time given in a request, which extends the life of the signature. This value does not affect the expiration value in the request nor does it affect the signature itself.
	ExpirationExtension uint32 `json:"expirationExtension,omitempty" default:"1800" role:"normal" writeonly:""`
}

// AuthURLSignAliCloudC Ali Cloud Type-C URL Signing
// AllowedScope PRODUCT
// DefaultPolicy  null
type AuthURLSignAliCloudC struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// PassPhrase Specify the shared passphrase, or sequence of words or other text, to use when generating the signature when authenticating a request made to the CDN.
	PassPhrase string `json:"passPhrase," default:"" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// TokenField Set to use query string parameter to specify signing signature instead of putting it in the path of the URL.
	TokenField string `json:"tokenField,omitempty" default:"" role:"normal" writeonly:""`

	// ExpireField Set to use query string parameter to specify the expire time instead of putting it in the path of the URL.
	ExpireField string `json:"expireField,omitempty" default:"" role:"normal" writeonly:""`

	// ExpirationExtension Number of seconds to add to the expiration time given in a request, which extends the life of the signature. This value does not affect the expiration value in the request nor does it affect the signature itself.
	ExpirationExtension uint32 `json:"expirationExtension,omitempty" default:"1800" role:"normal" writeonly:""`
}

// AuthURLSignHmacTlu URL Signing HMAC TLU
// AllowedScope DIR
// DefaultPolicy  null
type AuthURLSignHmacTlu struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// AlgorithmIdMap One or more key-value pairs, where the key is an ID and the value is a predetermined HMAC algorithm name maps. The ID is given in a signed URL and specifies which HMAC algorithm to use for authorization.
	AlgorithmIDMap string `json:"algorithmIdMap," default:"" role:"normal" writeonly:"" hashMap:"string,enum[hmacsha1|hmacsha256]"`

	// SymmetricKeyIdMap One or more key-value pairs, where the key is an ID and the value is shared symmetric key. The value can only printable ASCII characters and HTML encoded. The ID is given in a signed URL and specifies which symmetric key to use for authorization.
	SymmetricKeyIDMap string `json:"symmetricKeyIdMap," default:"" role:"normal" writeonly:"1" hashMap:"string,string"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// AlgorithmIdParameterName Name of the query string parameter that contains the HMAC algorithm identifier for the signed URL.
	AlgorithmIDParameterName string `json:"algorithmIdParameterName,omitempty" default:"P3" role:"normal" writeonly:""`

	// DigestParameterName Name of the query string parameter that contains the HMAC digest (hash) for the signed URL.
	DigestParameterName string `json:"digestParameterName,omitempty" default:"P4" role:"normal" writeonly:""`

	// ExpireParameterName Name of the query string parameter that contains the expiration time for the signed URL.
	ExpireParameterName string `json:"expireParameterName,omitempty" default:"P1" role:"normal" writeonly:""`

	// KeyIdParameterName Name of the query string parameter that contains the shared symmetric key identifier for the signed URL.
	KeyIDParameterName string `json:"keyIdParameterName,omitempty" default:"P2" role:"normal" writeonly:""`
}

// AuthURLSignIq The IQIYI signing policy allows you to restrict access to your content using various query parameters. Client requests to the CDN supply parameters that specifiy how to generate the secure token. Since the shared token and details of the algorithm are only known by the publisher and Stackpath, URL  signatures cannot be generated by unauthorized users.
// WARNING: This needs to have a script set up in order to work properly.
// AllowedScope DIR
// DefaultPolicy  null
type AuthURLSignIq struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// MethodFilter Method Filter is used to determine if this type should be applied or not based on List of HTTP Methods provided
	// Optionally, you may use an exclamation point in the list to describe the subset of HTTP methods excluded from this policy and all
	// other requests method will be included.
	// WARNING: You should not mix include and exclude in the same list.
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// SecretKey Security token used for signing in IQIYI's unique URL signing method.
	SecretKey string `json:"secretKey,omitempty" default:"" role:"normal" writeonly:""`
}

// AuthURLAsymmetricSignTlu The ASYMMETRIC Time Limited URL (TLU) signing policy allow you to restrict access to your content by by use of an expiration time and Asymmetric Key based signed alglorithm that utilizes RSA private/public keys. Client requests to the CDN supply IDs that specifiy the shared public key and specific algorithm to apply to validate the signature that is also supplied in the request.  Since the private asymmetric key are only known by the publisher, URL signatures cannot be generated by unauthorized users.
// AllowedScope DIR
// DefaultPolicy  null
type AuthURLAsymmetricSignTlu struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// AlgorithmIdMap One or more key-value pairs, where the key is an ID and the value is a predetermined HMAC algorithm name maps. The ID is given in a signed URL and specifies which HMAC algorithm to use for authorization.
	AlgorithmIDMap string `json:"algorithmIdMap," default:"" role:"normal" writeonly:"" hashMap:"string,enum[hmacsha1|hmacsha256]"`

	// PublicKeyIdMap One or more key-value pairs, where the key is an ID and the value is shared public key. The ID is given in a signed URL and specifies which asymmetric key to use for authorization. Key is expected to be in Modulus and Exponent format delimited by Pipe (|).  Example: modulus: base64_value|exponent: base64_value
	PublicKeyIDMap string `json:"publicKeyIdMap," default:"" role:"normal" writeonly:"1" hashMap:"string,string"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// MethodFilter Method Filter is used to determine if this type should be applied or not based on List of HTTP Methods provided
	// Optionally, you may use an exclamation point in the list to describe the subset of HTTP methods excluded from this policy and all
	// other requests method will be included.
	// WARNING: You should not mix include and exclude in the same list.
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// AlgorithmIdParameterName Name of the query string parameter that contains the HMAC algorithm identifier for the signed URL.
	AlgorithmIDParameterName string `json:"algorithmIdParameterName,omitempty" default:"P3" role:"normal" writeonly:""`

	// DigestParameterName Name of the query string parameter that contains the HMAC digest (hash) for the signed URL.
	DigestParameterName string `json:"digestParameterName,omitempty" default:"P4" role:"normal" writeonly:""`

	// ExpireParameterName Name of the query string parameter that contains the expiration time for the signed URL.
	ExpireParameterName string `json:"expireParameterName,omitempty" default:"P1" role:"normal" writeonly:""`

	// KeyIdParameterName Name of the query string parameter that contains the shared symmetric key identifier for the signed URL.
	KeyIDParameterName string `json:"keyIdParameterName,omitempty" default:"P2" role:"normal" writeonly:""`
}

// AuthURLSignL3 The Level 3 URL Signing policy allows you to create a signed URL that implements the same signing method used by Level 3; therefore, published URLs from an Level 3 CDN network can be transitioned to the Highwinds network without you having to change your signing methods.
// AllowedScope DIR
// DefaultPolicy  null
type AuthURLSignL3 struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// SharedSecretTable An ordered list of shared secrets. The order is CRITICAL and it MUST be identical to the ordered table used by the Client.
	SharedSecretTable string `json:"sharedSecretTable," default:"" role:"HWADMIN" writeonly:"" list:"string"`

	// TokenField This is the name of the query string parameter that will be used by the publisher to specify the signature for the URL.
	TokenField string `json:"tokenField," default:"" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// IncludeProtocolAndHost Indicates whether or not to include both the Protocol and Host when calculating the signature.
	IncludeProtocolAndHost *bool `json:"includeProtocolAndHost,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// InjectClientIPAddress Indicates whether or not to include the Client's IP address when calculating the signature.
	InjectClientIPAddress *bool `json:"injectClientIPAddress,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// IncludeHostOnly Indicates whether or not to include the Host without the request Protocol when calculating the signature.
	IncludeHostOnly *bool `json:"includeHostOnly,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// TimeFormat Used to describe the format of expireField and startField. The CDN currently supports two formats.
	//   1. epoch: An integer representing the number of seconds since January 1, 1970 on a UNIX/POSIX system.
	//   2. datetime: A numerical representation of a date and time in GMT in the format yyyymmddHHMMSS.
	TimeFormat string `json:"timeFormat,omitempty" default:"epoch" role:"HWADMIN" writeonly:"" enum:"[epoch,datetime]"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// ExcludedParameters A list of patterns that are used to describe query string parameters that should be omitted from the hashing algorithm if contained in the URL. A asterisk '*' by itself indicates to exclude all query string parameters from the hashing algorithm. The tokenField is always excluded. On the other hand, the startField and/or expireField are always included in the hashing algorithm if present in the request even if listed here. Users may explicitly specify parameters to keep (not exclude) by preceding the glob with an exclamation "!". This may be useful if a User wants to exclude all query string parameters except one ore more known parameters.  For example, a value of '*,!version' means exclude all parameters except "version".
	ExcludedParameters string `json:"excludedParameters,omitempty" default:"" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// ClientIPAddressField Provides the capability to rename the query string parameter that is used to inject the Client's IP address into the hashing algorithm. This configuration is only applicable when injectClientIPAddress is set to true.
	ClientIPAddressField string `json:"clientIPAddressField,omitempty" default:"clientip" role:"HWADMIN" writeonly:""`

	// StartField The name of the query string parameter that contains the start time when the request is considered valid.
	StartField string `json:"startField,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// ExpireField This is the name of the query string parameter that contains the time after which the URL is considered invalid. If defined, requests must contain the parameter, and its value must be in the future.
	ExpireField string `json:"expireField,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// AuthURLSignAKv1 The Akamai URL Signing v1 policy allows you to create a signed URL that implements the same signing  method used by Akamai; therefore, published URLs from an Akamai CDN network can be transitioned to the Highwinds network without you having to change your signing methods.
// AllowedScope DIR
// DefaultPolicy  null
type AuthURLSignAKv1 struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Salt The salt is used as a shared secret in the signing process. This value should only be known by Highwinds and by systems authorized to sign your content.
	Salt string `json:"salt," default:"" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// Param The authentication parameter defines the query string parameter in the request URL that contains the authentication information.
	Param string `json:"param,omitempty" default:"__gda__" role:"HWADMIN" writeonly:""`

	// Extract This indicates a component to extract from the request.  If specified, it must exist in the request to pass authentication. If present in the request, its value is used to generate the authorization hash. The format is componentType:componentName. Currently, the only supported componentType is "header".
	Extract string `json:"extract,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// AuthURLSignAKv2 The Akamai URL Signing v2 policy allows you to create a signed URL that implements the same signing  method used by Akamai; therefore, published URLs from an Akamai CDN network can be transitioned to the Highwinds network without you having to change your signing methods.
// AllowedScope DIR
// DefaultPolicy  null
type AuthURLSignAKv2 struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// PassPhrase This is the shared secret used to sign the URL.  This value must be set to a hexadecimal value padded to a byte boundary.  This value should only be known by Highwinds and by personnel authorized to sign your content.
	PassPhrase string `json:"passPhrase," default:"" role:"HWADMIN" writeonly:""`

	// MatchURL Add the path portion of the URL (e.g., /path/to/file.txt) into the token before hashing.
	MatchURL *bool `json:"matchURL," default:"1" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// EnableACLWildcard This allows you to enable the use of wildcard matches in your ACL list.
	EnableACLWildcard *bool `json:"enableACLWildcard," default:"1" role:"HWADMIN" writeonly:""`

	// AclDelimiter This is the delimiter used to separate the IP addresses in the ACL list.
	ACLDelimiter string `json:"aclDelimiter,omitempty" default:"!" role:"HWADMIN" writeonly:"" advancedType:"char"`

	// FieldDelimiter This is the field delimiter used to separate the parts of your token.
	FieldDelimiter string `json:"fieldDelimiter,omitempty" default:"~" role:"HWADMIN" writeonly:"" advancedType:"char"`

	// HashStrategy This is the hashing algorithm used to sign the URLs.
	HashStrategy string `json:"hashStrategy,omitempty" default:"sha256" role:"HWADMIN" writeonly:"" enum:"[sha1,sha256,md5]"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// TokenField The token parameter is the name of the query string parameter or cookies that contains the value of the token used in the signing algorithm.
	TokenField string `json:"tokenField,omitempty" default:"hdntl" role:"HWADMIN" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// Salt This is random data used as additional input to the hashing algorithm.
	Salt string `json:"salt,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// AuthURLSignLMV The Limelight Networks URL signing policy allows you to create a signed URL that implements the same signing  method used by Limelight Networks; therefore, published URLs from a Limelight CDN network can be transitioned to the Highwinds network without you having to change your URLs (or the signing process).
// AllowedScope DIR
// DefaultPolicy  null
type AuthURLSignLMV struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Secret This is the shared secret used to sign the URL.  This value should only be known by Highwinds and by personnel authorized to sign your content.
	Secret string `json:"secret," default:"" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// ParameterMatchPatterns A list of glob match pattern to indicate whether a query parameter will be included when generating the MD5 for validation.  This will also preserve the order of all query parameters as they are presented in the request.  The signing paraterms (s, e, ip, ri and rs) are always included if present.  To include all other parameters as well, set this to '*'.  When set to empty (default), it will only include the signing parameters and the signing parameters must be in this order: s, e, p, ip, ri, rs.
	ParameterMatchPatterns string `json:"parameterMatchPatterns,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// EndTimeFieldName The parameter name to specify the end time value.
	EndTimeFieldName string `json:"endTimeFieldName,omitempty" default:"e" role:"HWADMIN" writeonly:""`

	// IpFieldName The parameter name to specify the ip mask value.
	IPFieldName string `json:"ipFieldName,omitempty" default:"ip" role:"HWADMIN" writeonly:""`

	// LengthFieldName The parameter name to specify the length value.
	LengthFieldName string `json:"lengthFieldName,omitempty" default:"p" role:"HWADMIN" writeonly:""`

	// RiFieldName The parameter name to specify the ri value.
	RiFieldName string `json:"riFieldName,omitempty" default:"ri" role:"HWADMIN" writeonly:""`

	// RsFieldName The parameter name to specify the rs value.
	RsFieldName string `json:"rsFieldName,omitempty" default:"rs" role:"HWADMIN" writeonly:""`

	// StartTimeFieldName The parameter name to specify the start time value.
	StartTimeFieldName string `json:"startTimeFieldName,omitempty" default:"s" role:"HWADMIN" writeonly:""`

	// TokenFieldName The parameter name to specify the token value.
	TokenFieldName string `json:"tokenFieldName,omitempty" default:"h" role:"HWADMIN" writeonly:""`
}

// AuthVhostLockout The Hostname Access policy allows you to restrict delivery of your content to your configured Hostnames.  Any request for your content that is not using one of your configured hostnames will be denied.
// AllowedScope DIR
// DefaultPolicy  {"lockout":false}
type AuthVhostLockout struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Lockout By enabling this checkbox your content is only accessible through one of your configured Hostnames.
	Lockout *bool `json:"lockout," default:"false" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// BandWidthLimit Limit the transfer rate of files by extension.
// AllowedScope DIR
// DefaultPolicy  null
type BandWidthLimit struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Rule Rule to apply bandwidth limiting on responses. Each rule is separated by '|' and consists of a glob match for the user agent and a coma separated glob match list of file extensions separated by ':' [/Pattern/User-Agent 1/Pattern/]:[comma separated Path/FileExtension]|[/Pattern/User-Agent 2/Pattern/]:[comma separated Path/FileExtension]
	Rule string `json:"rule," default:"" role:"normal" writeonly:"" list:"string"`

	// Values Values of Initial Bytes to send at full speed and SustainRate-KiloBitsPerSecs for the rest of the file(RI and RS value) Ex: ri=100,rs=1000 - send the first 100 bytes at full speed and then throttle to 1000kbps. ri has default unit of Bytes and rs has default unit of KiloBitsPerSec unless they are changed in bandWidthRateLimitUnits
	Values string `json:"values," default:"" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// BandwidthRateLimit Limit the transfer rate of files in general, as opposed to by extension like Pattern Based Bandwidth Rate Limiting.
// AllowedScope DIR
// DefaultPolicy  null
type BandwidthRateLimit struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// InitialBurstUnits The units used by the initial burst parameter on the URL.
	InitialBurstUnits string `json:"initialBurstUnits,omitempty" default:"byte" role:"normal" writeonly:"" enum:"[byte,kilobyte]"`

	// SustainedRateUnits The units used by the Sustained Rate parameter on the URL.
	SustainedRateUnits string `json:"sustainedRateUnits,omitempty" default:"kilobit" role:"normal" writeonly:"" enum:"[kilobit,kilobyte]"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// InitialBurstName The name of the query string parameter that establishes the initial burst rate to use when delivering content. Currently optional, however, this policy will become required in order to enable rate limiting support. The default name until required is ri.
	InitialBurstName string `json:"initialBurstName,omitempty" default:"" role:"normal" writeonly:""`

	// SustainedRateName The name of the query string paramter that establishes the sustained rate to use when delivering content. Currently optional, however, this policy will become required in order to enable rate limiting support. The default name until required is rs.
	SustainedRateName string `json:"sustainedRateName,omitempty" default:"" role:"normal" writeonly:""`
}

// BandWidthRateLimitUnits Override the default units used by the CDN when processing the bandwidth throttling policies.
// AllowedScope DIR
// DefaultPolicy  {"sustained":"kilobit","initial":"byte"}
type BandWidthRateLimitUnits struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Initial Deprecated By bandwidthRateLimit/initialBurstUnits
	Initial string `json:"initial,omitempty" default:"byte" role:"normal" writeonly:"" enum:"[byte,kilobyte]"`

	// Sustained Deprecated By bandwidthRateLimit/sustainedRateUnits
	Sustained string `json:"sustained,omitempty" default:"kilobit" role:"normal" writeonly:"" enum:"[kilobit,kilobyte]"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// ClientAccess This allows you to override the default client access policy file (clientaccesspolicy.xml) delivered by the CDN caching servers.
// AllowedScope DIR
// DefaultPolicy  {"policy":"PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4KPGFjY2Vzcy1wb2xpY3k+CiAgPGNyb3NzLWRvbWFpbi1hY2Nlc3M+CiAgICA8cG9saWN5PgogICAgICA8YWxsb3ctZnJvbSBodHRwLXJlcXVlc3QtaGVhZGVycz0iU09BUEFjdGlvbiI+CiAgICAgICAgPGRvbWFpbiB1cmk9IioiLz4KICAgICAgPC9hbGxvdy1mcm9tPgogICAgICA8Z3JhbnQtdG8+CiAgICAgICAgPHJlc291cmNlIHBhdGg9Ii8iIGluY2x1ZGUtc3VicGF0aHM9InRydWUiLz4KICAgICAgPC9ncmFudC10bz4KICAgIDwvcG9saWN5PgogIDwvY3Jvc3MtZG9tYWluLWFjY2Vzcz4KPC9hY2Nlc3MtcG9saWN5PgoK"}
type ClientAccess struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Policy The contents of the cross domain file you want to serve instead of the default
	Policy string `json:"policy," default:"" role:"HWADMIN" writeonly:"" advancedType:"text"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// Compression Speed up your websites or web apps by making certain files smaller before they're delivered to end-users.
// AllowedScope DIR
// DefaultPolicy  null
type Compression struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Level
	// The level of compression for requested compressed files. Acceptable values are 0 to 6. 0 is no compression, 1 is high speed, 6 is high compression. The default is 6. If an invalid number is chosen, we will use the default.
	Level string `json:"level,omitempty" default:"6" role:"normal" writeonly:"" enum:"[1,2,3,4,5,6]"`

	// Gzip Comma separated list of extensions to apply gzip compression. Content is only gziped after the first cache response.
	// <p><b>WARNING:</b> This setting is being deprecated. All new config should use the GzipOriginPull/Enabled setting</p>
	Gzip string `json:"gzip,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Mime Provide a comma-separated list of strings where each string is a rule. All rules are applied from left to right. The last matching rule is used. It is recommended that rules are written in order of generic to specific. An asterisk may be supplied for the type, subtype, or both to represent all values. Negative rules are defined by an exclamation point prefix. Content is only gziped after a request for a compressed version.
	Mime string `json:"mime,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// ContentDispositionByURL Control the Content-Disposition header on the response from the Origin via the request URL of end-user clients.
// AllowedScope DIR
// DefaultPolicy  null
type ContentDispositionByURL struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// DispositionNameQSParam The Query String parameter name which will specify the filename used in the Content-Disposition header
	DispositionNameQSParam string `json:"dispositionNameQSParam," default:"" role:"normal" writeonly:""`

	// DispositionTypeQSParam The Query String parameter name which will specify the type (inline or attachment) used in the Content-Disposition header.
	DispositionTypeQSParam string `json:"dispositionTypeQSParam," default:"" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// OverrideOriginHeader This setting allow the CDN-generated Content-Disposition header to override the one preserved from the origin using the OriginPullPolicy/HttpHeaders or Origin/OriginCacheHeaders settings.
	OverrideOriginHeader *bool `json:"overrideOriginHeader," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// DispositionOverrideQSParam The Query String parameter name which will override the whole value in the Content-Disposition header. If this is present in the request URL, the DispositionNameQSParam and DispositionTypeQSParam will be ignored. If the value of the parameter in the URL is empty, it will remove the Content-Disposition header completely.
	DispositionOverrideQSParam string `json:"dispositionOverrideQSParam,omitempty" default:"" role:"normal" writeonly:""`
}

// ContentDispositionByHeader Control the Content-Disposition header on the responses from the Origin using a pattern matched against the value of any HTTP header present in an end-user's request for content.
// AllowedScope DIR
// DefaultPolicy  null
type ContentDispositionByHeader struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// HeaderValueMatch A list of glob patterns to enable the Content-Disposition header if any of them matches the value in the header field specified in the HeaderFieldName setting.
	HeaderValueMatch string `json:"headerValueMatch," default:"" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFieldName The name of the Http header field used to get the value to match the pattern specified in the HeaderValueMatch setting
	HeaderFieldName string `json:"headerFieldName," default:"" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// OverrideOriginHeader This allow the CDN-generated Content-Disposition header to override the one preserved from the origin using the OriginPullPolicy/HttpHeaders or Origin/OriginCacheHeaders settings.
	OverrideOriginHeader *bool `json:"overrideOriginHeader," default:"1" role:"normal" writeonly:""`

	// DefaultType The default Content-Disposition type when the header is enabled
	DefaultType string `json:"defaultType,omitempty" default:"attachment" role:"normal" writeonly:"" enum:"[attachment,inline]"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// CookieBehavior The setting controls how the CDN deal with Cookie (from client) and Set-Cookie (from origin) headers
// AllowedScope DIR
// DefaultPolicy  {"allowCachingSetCookie":false}
type CookieBehavior struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// AllowCachingSetCookie The actual preserving header is controlled by OriginPullPolicy/HttpHeaders. This option is mainly for when the OriginPullPolicy/HttpHeaders is set to '*', we want to treat the Set-Cookie hearder differently from other headers as OriginpullPolicy doesn't allow generic glob or negative matching. Set-Cookie can be proxied to the client when dudupping is disabled when we proxy all the headers through; otherwise, we will not store the Set-Cookie header in edgefile, so it cannot be sent from cache or deduping when this setting is disabled.
	AllowCachingSetCookie *bool `json:"allowCachingSetCookie," default:"false" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// MethodFilter Method Filter is used to determine if this type should be applied or not based on List of HTTP Methods provided
	// Optionally, you may use an exclamation point in the list to describe the subset of HTTP methods excluded from this policy and all
	// other requests method will be included.
	// WARNING: You should not mix include and exclude in the same list.
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// CrossDomain Enable and configure the crossdomain.xml file required to enable the Dynamic Files policy.
// AllowedScope DIR
// DefaultPolicy  {"file":"PD94bWwgdmVyc2lvbj0iMS4wIj8+Cjxjcm9zcy1kb21haW4tcG9saWN5PgogICA8IS0tIFRoaXMgaXMgYSBtYXN0ZXItcG9saWN5IGZpbGUgLS0+CiAgIDxzaXRlLWNvbnRyb2wgcGVybWl0dGVkLWNyb3NzLWRvbWFpbi1wb2xpY2llcz0iYWxsIiAvPgogICA8YWxsb3ctYWNjZXNzLWZyb20gZG9tYWluPSIqIiB0by1wb3J0cz0iODAsNDQzIiAvPgo8L2Nyb3NzLWRvbWFpbi1wb2xpY3k+Cgo="}
type CrossDomain struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// File User specified crossdomain file base 64 encoded.
	// <p><b>NOTE:</b> This is a legacy policy, new sites should host the cross domain file on the origin.</p>
	File string `json:"file," default:"" role:"HWADMIN" writeonly:"" advancedType:"text"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// CustomMimeType Map file extensions directly to mime types.
// AllowedScope DIR
// DefaultPolicy  null
type CustomMimeType struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// ExtensionMap Comma separated list of extension maps (extension:mimetype). No glob match is supported but a single "*" can be used to match all extension. A * extension at the end of the list can be used to catch everything else not listed in the map.
	ExtensionMap string `json:"extensionMap," default:"" role:"normal" writeonly:"" list:"string"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// Code Comma separated list of status codes to match. This is optional and defaults to 200,206 when not set. No glob match is supported but a single "*" can be use to match all.
	Code string `json:"code,omitempty" default:"200,206" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// DNSIPv6 DNS Configuration for Ipv6
// AllowedScope PRODUCT
// DefaultPolicy  null
type DNSIPv6 struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enable Enable IPv6 DNS records for this host.
	Enable *bool `json:"enable," default:"false" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// DNSOverride DNS Configuration
// AllowedScope PRODUCT
// DefaultPolicy  null
type DNSOverride struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Type The record type. Use A for IPv4, AAAA for IPv6, and CNAME to forward to a foreign DNS system. You may have multiple DNS overrides of type A or AAAA, but you cannot combine a CNAME override with an A or AAAA record.
	Type string `json:"type," default:"" role:"HWADMIN" writeonly:"" enum:"[A,AAAA,CNAME]"`

	// Answer The answer that the DNS server will give. This should be an IPv4 address for records of type A, an IPv6 address for records of type AAAA, or a valid domain for type CNAME
	Answer string `json:"answer," default:"" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// Pop The pop for which this DNS override will apply. If left out, this override will apply to all pops (for dedicated IP addresses). If an account has overrides both with and without the pop field, the override with the pop field will squash the one without it (in that pop).
	Pop string `json:"pop,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// Weight Weight for this particlar override. Use lowest possible postive integer between 1-24.
	Weight uint32 `json:"weight,omitempty" default:"1" role:"HWADMIN" writeonly:"" range:"1,24"`

	// Ttl The time to live to present to caching name servers.
	TTL uint32 `json:"ttl,omitempty" default:"300" role:"HWADMIN" writeonly:"" range:"1,600"`
}

// DynamicCacheRule Trigger specific status codes for precise URLs or domains.
// AllowedScope DIR
// DefaultPolicy  null
type DynamicCacheRule struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// StatusCode The status code that applies to this policy.
	StatusCode uint32 `json:"statusCode," default:"" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// Headers List of headers to include on the response.  This field may use server variables (e.g. %server.ip%)
	Headers string `json:"headers,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// Flv The flash initial bytes policy allows you to force the CDN to send the initial bytes of a FLV file which contains the header information that is used when jumping to different offsets in the file.
// AllowedScope DIR
// DefaultPolicy  null
type Flv struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// InitByteSize This setting is typically set to 13 bytes.
	InitByteSize uint32 `json:"initByteSize," default:"" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// FlvPseudoStreaming Define how the CDN delivers Flash media.
// AllowedScope DIR
// DefaultPolicy  null
type FlvPseudoStreaming struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// JumpToByteInitialBytesParam The name of the query string parameter used to specify the initial bytes of a video that should be returned before sending the requested byte offset. Typically players us the "ib" for this query string parameter.
	JumpToByteInitialBytesParam string `json:"jumpToByteInitialBytesParam,omitempty" default:"" role:"normal" writeonly:""`

	// JumpToByteStartOffsetParam The name of the query string parameter used to specify the specific byte offset into the requested video. Typically players use the "fs" for this query string parameter.
	JumpToByteStartOffsetParam string `json:"jumpToByteStartOffsetParam,omitempty" default:"" role:"normal" writeonly:""`

	// InitialByteSize Configures a default initial bytes that should be delivered to the client. Typically the player wants the initial 13 bytes of the FLV, if you leave this value as zero then you should set an Initial Bytes value on each request. This default value only applies to requests with a MimeType of "video/x-flv".
	InitialByteSize uint32 `json:"initialByteSize,omitempty" default:"0" role:"HWADMIN" writeonly:""`
}

// General The zero byte file support policy enables the CDN to cache zero length files.  By default, the CDN proxies zero length files without caching them.
// AllowedScope DIR
// DefaultPolicy  {"allowZeroByteFile":false}
type General struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// AllowZeroByteFile The allow zero byte files policy enables the caching of zero byte files (empty files) on the  caching servers.
	AllowZeroByteFile *bool `json:"allowZeroByteFile," default:"false" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// HTTPMethods Selectively enable additional HTTP methods you'd like the CDN to process.
// AllowedScope DIR
// DefaultPolicy  {"passThru":"POST"}
type HTTPMethods struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// PassThru A comma separated list of HTTP methods for no-store like pass thru behavior. GET and HEAD is always excluded from this list. A "*" can also be used to include all methods. GET and HEAD are always excluded from this list even when "*" is used.
	PassThru string `json:"passThru," default:"" role:"normal" writeonly:"" list:"string"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// LegacyXdomain The legacy cross domain policy allows you to override the default cross domain file delivered by the  CDN.  This policy is being deprecated, and you should ensure that any custom cross domain file you wish the CDN to deliver can be requested from your origin.
// AllowedScope DIR
// DefaultPolicy  {"enabled":false}
type LegacyXdomain struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled When enabled, the CDN supports the use of the Client Access and Cross Domain policies.
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// LiveStreaming Live Streaming Optimization
// AllowedScope DIR
// DefaultPolicy  null
type LiveStreaming struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// EnablePlaylistOptimization If enabled, extra check will be performat to make sure playlist file doesn't get cached for too long
	EnablePlaylistOptimization *bool `json:"enablePlaylistOptimization," default:"false" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// PreserveRedirectHost Preserve Redirect Host
// AllowedScope DIR
// DefaultPolicy  null
type PreserveRedirectHost struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// StatusCodes Comma separated list of redirect HTTP status codes to apply this policy. GFS only preserves the Host header if the status code of the origin response is listed.
	StatusCodes string `json:"statusCodes," default:"" role:"HWADMIN" writeonly:"" list:"uint32"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// QueryStrParam Define special customer query string parameters the CDN will use to alter responses.
// AllowedScope DIR
// DefaultPolicy  null
type QueryStrParam struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// DispositionName The query string parameter name which will include the filename used in the Content-Disposition header.
	DispositionName string `json:"dispositionName,omitempty" default:"" role:"normal" writeonly:""`

	// DispositionOverride The query string parameter name which will override the whole value in the Content-Disposition header.
	DispositionOverride string `json:"dispositionOverride,omitempty" default:"" role:"normal" writeonly:""`

	// DispositionType The query string parameter name which will specify the type (inline or attachment) used in the Content-Disposition header.
	DispositionType string `json:"dispositionType,omitempty" default:"" role:"normal" writeonly:""`

	// RateLimitInitial <p>Deprecated by bandwidthRateLimit/initialBurstName
	// </p>
	// <p><b>Deprecated:</b> This key will be removed in a future version</p>
	RateLimitInitial string `json:"rateLimitInitial,omitempty" default:"ri" role:"normal" writeonly:""`

	// RateLimitSustained <p>Deprecated by bandwidthRateLimit/sustainedRateName
	// </p>
	// <p><b>Deprecated:</b> This key will be removed in a future version</p>
	RateLimitSustained string `json:"rateLimitSustained,omitempty" default:"rs" role:"normal" writeonly:""`

	// JumpToByteStartOffset <p>Deprecated by flvPseudoStreaming/jumpToByteStartOffsetParam
	// </p>
	// <p><b>Deprecated:</b> This key will be removed in a future version</p>
	JumpToByteStartOffset string `json:"jumpToByteStartOffset,omitempty" default:"" role:"normal" writeonly:""`

	// JumpToByteInitialBytes <p>This key is used by legacy sites, new sites should use the flvPseudoStreaming/jumpToByteInitialBytesParam</p>
	// <p><b>Deprecated:</b> This key will be removed in a future version</p>
	JumpToByteInitialBytes string `json:"jumpToByteInitialBytes,omitempty" default:"" role:"normal" writeonly:""`

	// JumpToTimeEnd <p>This key is used by legacy sites, new sites should use the timePseudoStreaming/jumpToTimeEndParam.</p>
	// <p><b>Deprecated:</b> This key will be removed in a future version</p>
	JumpToTimeEnd string `json:"jumpToTimeEnd,omitempty" default:"" role:"normal" writeonly:""`

	// JumpToTimeStart <p>This key is used by legacy sites, new sites should use the timePseudoStreaming/jumpToTimeStartParam</p>
	// <p><b>Deprecated:</b> This key will be removed in a future version</p>
	JumpToTimeStart string `json:"jumpToTimeStart,omitempty" default:"" role:"normal" writeonly:""`
}

// RedirectExceptions Make exceptions for which web browsers or user agents see the custom redirect response URLs based on a customizable list.
// AllowedScope DIR
// DefaultPolicy  null
type RedirectExceptions struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// RedirectAgentCode <p>Comma separated list of user agent and redirect code pairs. The pair is separated by ':' and user agent can be a wildcard matching pattern. Redirect Code cannot be empty. Must put the more specific matching patterns first as once a match is found, the rest of the codes are ignored. If no match is found, default redirect code is 301.</p>
	RedirectAgentCode string `json:"redirectAgentCode," default:"" role:"normal" writeonly:"" list:"string"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// RedirectMappings Redirect users to a custom response URL based on the error response code they encounter.
// AllowedScope DIR
// DefaultPolicy  null
type RedirectMappings struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// RedirectURL HTTP address to redirect to when HTTP response code is encountered. Any replacement tokens in the HTTP address will be replaced with the URL address that caused the error to be generated.
	RedirectURL string `json:"redirectURL," default:"" role:"normal" writeonly:""`

	// Code HTTP response code which will be redirected instead of returned to client.
	Code uint32 `json:"code," default:"" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// ReplacementToken An arbitrary token name used to substitute the URL that caused the error in the redirect URL specified in the policy.
	ReplacementToken string `json:"replacementToken,omitempty" default:"" role:"normal" writeonly:""`
}

// ResponseHeader Enable and bypass certain Origin headers that affect the delivery of content.
// AllowedScope DIR
// DefaultPolicy  {"enableETag":true}
type ResponseHeader struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// EnableETag Enables the e-tag header on client responses from the CDN.
	EnableETag *bool `json:"enableETag," default:"1" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// Http <p>Enables attachment type for the content-disposition on all requests that match the specified user agents.</p>
	// <p><b>NOTE:</b> This key is used by legacy sites, all new sites should use ClientResponseModification to achieve this behavior.</p>
	HTTP string `json:"http,omitempty" default:"" role:"normal" writeonly:""`
}

// RobotsTxt Define how to the CDN delivers the Robots.txt file.
// AllowedScope PRODUCT
// DefaultPolicy  null
type RobotsTxt struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// File <p>The contents of the robots.txt file you want delivered instead of the default.  This value must be base64 encoded.</p>
	File string `json:"file," default:"" role:"normal" writeonly:"" advancedType:"text"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// CacheControlHeader <p>The cache control header send with the Robots.txt file</p>
	CacheControlHeader string `json:"cacheControlHeader,omitempty" default:"max-age=86400" role:"normal" writeonly:""`
}

// StaticHeader Insert HTTP headers into the CDN request and response process.
// AllowedScope DIR
// DefaultPolicy  null
type StaticHeader struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// OriginPull <p>The full HTTP header, including the value(s), to insert into the HTTP request from the CDN to the origin.  This field allows use of client and server variables (e.g. %server.ip% or %client.ip%).</p>
	// <p>>b>NOTE:</b> Using client variables with request collapsing is not recommended because the origin will not see all client requests to the cache.</p>
	OriginPull string `json:"originPull,omitempty" default:"" role:"normal" writeonly:""`

	// ClientRequest The full HTTP header, including the value(s), to insert into the HTTP request to the CDN.  This field allows use of client and server variables (e.g. %server.ip% or %client.ip%)
	ClientRequest string `json:"clientRequest,omitempty" default:"" role:"normal" writeonly:""`

	// Http The full HTTP header, including the value(s), to insert into the HTTP response from the CDN.  This field allows use of client and server variables (e.g. %server.ip% or %client.ip%)
	HTTP string `json:"http,omitempty" default:"" role:"normal" writeonly:""`
}

// StreamChunkedEncodingResponse Stream Chunked-Encoding Response in Dedup Queue
// AllowedScope DIR
// DefaultPolicy  {"enabled":false}
type StreamChunkedEncodingResponse struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Enabled Chunked-Encoding response in dedup queue if needed.
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// TimePseudoStreaming Enable Flash based video players to support seeking to random locations within an MP4 or FLV file without having to download the entire video.
// AllowedScope DIR
// DefaultPolicy  null
type TimePseudoStreaming struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// JumpToTimeEndParam Defines the end parameter used for pseudo-streaming in the request URL.
	JumpToTimeEndParam string `json:"jumpToTimeEndParam,omitempty" default:"" role:"normal" writeonly:""`

	// JumpToTimeStartParam Defines the start parameter used for pseudo-streaming in the request URL.
	JumpToTimeStartParam string `json:"jumpToTimeStartParam,omitempty" default:"" role:"normal" writeonly:""`
}

// HTTP2Support Enable support of HTTP2
// AllowedScope PRODUCT
// DefaultPolicy  {"enabled":false}
type HTTP2Support struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Enable support of HTTP/2.
	Enabled *bool `json:"enabled," default:"false" role:"normal" writeonly:""`

	// PopFilter POP filter is list of pattern to match POPs where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of POPs excluded from this policy.
	// Use lower case or use '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// RegionFilter Region filter is list of pattern to match Region where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of Regions excluded from this policy.
	// Use lower case or '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	RegionFilter string `json:"regionFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// OcspParsing Enable OCSP Parsing
// AllowedScope DIR
// DefaultPolicy  null
type OcspParsing struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Enable OCSP Support
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PopFilter POP filter is list of pattern to match POPs where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of POPs excluded from this policy.
	// Use lower case or use '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// RegionFilter Region filter is list of pattern to match Region where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of Regions excluded from this policy.
	// Use lower case or '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	RegionFilter string `json:"regionFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// Hostname Specifiy the unique domains end-users use to access your content, and the CDN uses to identify your content.
// AllowedScope DIR
// DefaultPolicy  null
type Hostname struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Domain
	Domain string `json:"domain," default:"" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// BlockingOriginPullMode Block all responses until the full file has been downloaded in the background.
// AllowedScope DIR
// DefaultPolicy  {"enabled":false}
type BlockingOriginPullMode struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Enable blocking origin pull feature.
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// CustomHeader Override the name of the X-Forwarded-For header the CDN sends to the Origin.
// AllowedScope DIR
// DefaultPolicy  null
type CustomHeader struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// XForwardedForAuth String to be used in place of "X-Forwarded-For" when making requests to the basic authentication server.
	XForwardedForAuth string `json:"xForwardedForAuth,omitempty" default:"" role:"normal" writeonly:""`

	// XForwardedForOrigin String to be used in place of "X-Forwarded-For" when making requests to the origin server.
	XForwardedForOrigin string `json:"xForwardedForOrigin,omitempty" default:"" role:"normal" writeonly:""`
}

// DynamicOrigin Override the default Origin domain set for the Scope by passing a different Origin as a query string parameter in a URL.
// AllowedScope DIR
// DefaultPolicy  null
type DynamicOrigin struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// QueryParam The query string parameter name that is used to set the origin. The value should be the full url in the form of http://foo.com:80/dir/. https is not supported at this point
	QueryParam string `json:"queryParam," default:"" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// AllowedDomains Provide an exclusive list of domains allowed as dynamic origins in the policy.  When this is not set, all domain names are allowed.
	AllowedDomains string `json:"allowedDomains,omitempty" default:"" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// FailSafeOriginPull Fail safe origin pull is when we get a negative response (4xx and 5xx) from the origin, we will try to fallback to secondary origin if available
// AllowedScope DIR
// DefaultPolicy  {"enabled":true,"statusCodeMatch":"4*,5*"}
type FailSafeOriginPull struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Flag for enabling the Fail Safe Origin Pull Feature.
	Enabled *bool `json:"enabled," default:"false" role:"normal" writeonly:""`

	// StatusCodeMatch Comma separated status code glob patterns to indicate which status code this policy apply to.
	// This list will only accept error response code (4xx and 5xx).
	StatusCodeMatch string `json:"statusCodeMatch,omitempty" default:"4*,5*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// FarAheadRangeProxy Configuring Far Ahead Range Proxy value with threshold bytes
// AllowedScope DIR
// DefaultPolicy  {"enabled":true,"thresholdBytes":2097152}
type FarAheadRangeProxy struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Flag for enabling the Far Ahead Range Proxy Feature
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// ThresholdBytes When a range request is requesting a byte range that is beyond the threshold bytes from the current full download offset. The range request will get proxy straight to the origin to provide better user experience. This feature is irrelevant when FileSegmentation is enabled as we will pull the required segment sized range and cache the segment to full fill the range request.
	ThresholdBytes uint32 `json:"thresholdBytes," default:"" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// FileSegmentation Enable the CDN to download and store files in small parts rather than as whole, and potentially large assets.
// AllowedScope DIR
// DefaultPolicy  {"enabled":false}
type FileSegmentation struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled
	Enabled *bool `json:"enabled," default:"false" role:"normal" writeonly:""`

	// InitialOriginRequestBehavior Specifies how the CDN requests assets from the origin by default for assets in this scope when fileSegmentation/enabled is true. Range: The CDN makes a GET request with a Range for the first segment, bytes 0 through <segment_size> - 1. The CDN retries with a full-file request if the range request fails.  Full: The CDN makes a GET request for the full/entire asset (i.e. no Range).  If the asset can be segmented, the CDN halts the download once <segment_size> bytes have been ingested and closes the connection.   The variable segment_size is globally defined by the CDN.  Once the CDN determines an asset is segmentable, it uses Range requests for all additonal segments regardless of the initial origin request type.
	InitialOriginRequestBehavior string `json:"initialOriginRequestBehavior,omitempty" default:"full" role:"HWADMIN" writeonly:"" enum:"[range,full]"`

	// InitialRangeRetryFilter All inclusive or exclusive list of HTTP status code glob patterns that filters what responses trigger a retry.   When the initialOriginRequestBehavior is set to range and the origin response is not 200/206, the CDN immediately  makes a request for the full-file (no Range) when the status code matches this filter. Otherwise, the CDN waits  until the asset expires before trying again.   By default, the CDN does not immediately remake retry after receiving an internal server (5XX) or a not found (404) response. Note that the CDN handles redirect origin responses (301, 304, etc) based on relavent policies and configuration for this site,  such as originPull/redirectAction.  Therefore those related status codes are not required and should not be included unless  there is a legit reason.
	InitialRangeRetryFilter string `json:"initialRangeRetryFilter,omitempty" default:"!404,!5*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// CustomSegmentSizeBytes The number of bytes the CDN uses to segment a new asset into parts while ingesting it. This overrides the CDN default size.   The maximum size is defined by the CDN, and it cannot be overridden.  As of 01 Nov 2017, the maximum size is 8 MB (8388608). Note, this policy applies to new assets, which includes new versions, and it has no affect on segments/assets already cached by the CDN.
	CustomSegmentSizeBytes uint32 `json:"customSegmentSizeBytes,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// VaryHeaderField Policy for configuring how the CDN handles a Vary field header delivered from an origin.
// AllowedScope DIR
// DefaultPolicy  null
type VaryHeaderField struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// ValuesFilter List of one or more glob patterns (see fnmatch) that limit what Vary field values (header names) are acceptable when given in an origin response.  All patterns are treated as case-insensitive regardless if case-insensitivity is specified. The CDN performs a case-insentive match using each glob pattern provided to either keep or discard values given in the Vary field.  The list given must be entirely inclusive or entirely exclusive.  A single inclusive wildcard (all values) pattern is not permitted, and it will be ignored by the CDN.
	ValuesFilter string `json:"valuesFilter," default:"" role:"HWADMIN" writeonly:"" list:"IGLOB"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// ProxyBehavior Define how the CDN proxies the Vary header from the origin to an end-user. whole: send the entire Vary header as-is regardless if the CDN used it for dynamic caching (similar to OriginPullPolicy/HttpHeaders)  filtered: send only the values from the origin Vary header that the CDN actually used (which is a product from the valuesFilter, the Vary origin header and the headers in the client request) none: do not proxy any of the Vary header regardless if it was used for dynamic caching Note, this policy takes precendence over OriginPullPolicy/HttpHeaders
	ProxyBehavior string `json:"proxyBehavior,omitempty" default:"filtered" role:"HWADMIN" writeonly:"" enum:"[whole,filtered,none]"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// GzipOriginPull Enable the CDN to request and accept Gzipped content from the Origin.
// AllowedScope DIR
// DefaultPolicy  {"enabled":false}
type GzipOriginPull struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled
	Enabled *bool `json:"enabled," default:"false" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// OriginPersistentConnections Enable Origin persistent connections.
// AllowedScope PRODUCT
// DefaultPolicy  {"enabled":false}
type OriginPersistentConnections struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled
	Enabled *bool `json:"enabled," default:"false" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// OriginPull Control the behavior of Origin pull requests.
// AllowedScope DIR
// DefaultPolicy  {"redirectAction":"follow","noQSParams":false,"defaultBehavior":"dedup","transparentMode":false,"passAllHeadersOnDedup":false}
type OriginPull struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// NoQSParams Enables the CDN to send only a path without query string parameters when making external origin requests.
	NoQSParams *bool `json:"noQSParams,omitempty" default:"false" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// PassAllHeadersOnDedup Enabling this feature instructs the CDN to proxy all the headers from the client HTTP request the  CDN is using to dedup an origin pull. Consider static header injection and other similar features which may cause duplicates or conflicts.
	PassAllHeadersOnDedup *bool `json:"passAllHeadersOnDedup,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// TransparentMode When set to true, it will preserve the original Host header and Path from the client's request when going to the shielding GFS and the origin.
	TransparentMode *bool `json:"transparentMode,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// DefaultBehavior <p>This setting changes the default origin pull behavior which by default is to always DEDUP.</p>
	// <p><b>NOCACHE:</b> assume the no-cache/no-store behavior until we found out otherwise from the origin.
	// <p><b>NOSTORE:</b> force everything to be uncachable regardless of what the origin returns
	DefaultBehavior string `json:"defaultBehavior,omitempty" default:"dedup" role:"HWADMIN" writeonly:"" enum:"[dedup,nocache,nostore]"`

	// RedirectAction <p>Dictates the behavior upon a HTTP redirect response from origin. Possible values are:</p>
	// <p><b>follow:</b> For following the redirects from the origin</p>
	// <p><b>proxy:</b> For proxying the redirect response back to the client without following it</p>
	RedirectAction string `json:"redirectAction,omitempty" default:"follow" role:"normal" writeonly:"" enum:"[follow,proxy]"`

	// RetryMethods Comma-delimited list of HTTP Methods that define types of origin pull requests that can be retried if a failure occurs after sending a previous request. List must be entirely inclusion or exclusion.
	RetryMethods string `json:"retryMethods,omitempty" default:"" role:"normal" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// ShieldResponseTimeoutOverride Enter the maximum number of seconds an edge GFS may wait for a shielding GFS to respond after a connection is made and the request is sent. It is recommended to take into account polices that affect origin pull requests, such as Origin/OriginTimeoutDuration and Origin/OriginPullRetries. A value of zero has the special meaning that instructs the CDN to use the server's local default internal server (shield) timeout value instead of calculating a timeout.
	ShieldResponseTimeoutOverride uint32 `json:"shieldResponseTimeoutOverride,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// OriginPullProtocol Configure whether the CDN should use secured or non-secured connections when communicating with the Origin.
// AllowedScope DIR
// DefaultPolicy  {"protocol":"http"}
type OriginPullProtocol struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Protocol Set the origin pull protocol to use on requests to the origin.  Valid values are HTTP Only, HTTPs Only, Match Request Protocol (CDN caches asset separately for HTTP and HTTPs).
	Protocol string `json:"protocol," default:"" role:"normal" writeonly:"" enum:"[http,https,match]"`

	// EnableSNI If enabled, CDN will use SNI while making secured connection to origin
	EnableSNI *bool `json:"enableSNI," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// OriginPullPops You should not be using the region filter on this policy before 975-1 goes CDN-wide.
// AllowedScope DIR
// DefaultPolicy  null
type OriginPullPops struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// PopList This is a list of CDN data centers that are used to shield the customer's origin.
	PopList string `json:"popList," default:"" role:"HWADMIN" writeonly:"" list:"string"`

	// PopFilter
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// RegionFilter
	RegionFilter string `json:"regionFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// OriginPullShield Origin shielding reduces the load on your origin by routing all origin pull requests through a specific data center on the network instead of having multiple data centers across the network request the same file from origin.
// AllowedScope DIR
// DefaultPolicy  {"enabled":false,"behavior":"redirect"}
type OriginPullShield struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled This enables origin shielding.
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// PermissibleShieldInternalErrors Control how the Edge PoP handles errors experienced by a shielding PoP before it receives a full response from an external origin.
	PermissibleShieldInternalErrors string `json:"permissibleShieldInternalErrors,omitempty" default:"CONNECTION_ONLY" role:"HWADMIN" writeonly:"" enum:"[NONE,CONNECTION_ONLY,WRITE_ONLY,WRITE_READ,ALL]"`

	// Behavior Changes how doppler response to shielding request. "redirect" will use a 301 redirect response to direct the client facing gfs to the correct shielding gfs for the file. "tlb" will serve the file  directly through TLB.
	Behavior string `json:"behavior,omitempty" default:"redirect" role:"HWADMIN" writeonly:"" enum:"[redirect,tlb]"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// OriginRoundRobinDNS The CDN can use a round-robin algorithm when selecting the IP address returned by the Domain Name Server  for the origin hostname specified. By default, the CDN utilizes its application level DNS caching where a single IP address is used until the next DNS refresh.
// AllowedScope PRODUCT
// DefaultPolicy  null
type OriginRoundRobinDNS struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Flag for turning on the Round Robin DNS feature.
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// Host Origin from OriginPullHost for which this policy applies. '*' indicates all origin hosts.
	Host string `json:"host," default:"" role:"HWADMIN" writeonly:""`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// MethodFilter Method Filter is used to determine if this type should be applied or not based on List of HTTP Methods provided
	// Optionally, you may use an exclamation point in the list to describe the subset of HTTP methods excluded from this policy and all
	// other requests method will be included.
	// WARNING: You should not mix include and exclude in the same list.
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// DnsRefreshSeconds Set the frequency of how often to refresh the IP addresses.  A value of zero indicates the CDN ought to use the global-conf default.   Note, in the case of a DNS failure, the CDN will continue to use the last know IPs for twice  this value - a grace period.  In the mean time, it will continue attempts to refresh the IPs.  If the CDN is unable to refresh the IPs after the grace period.  It will fall back to its  default behavior. The value cannot exceed GFS_MAXROUNDROBINDNSREFRESHSECONDS.
	DNSRefreshSeconds uint32 `json:"dnsRefreshSeconds,omitempty" default:"0" role:"HWADMIN" writeonly:""`
}

// AwsSignedOriginPullV4 Defines how to pre/sign requests to be made by the CDN to an AWS origin.
// Note, even though this policy is groupable, if more than one policy is defined, only one policy will ever be applied.
// The CDN iterates over each policy until it finds the first match or applicable policy based on scope and/or filter.
// The CDN does not failover or attempt other policies if the chosen one failed. The Groupability was added with the
// primary intent to provide flexibilty when needing to define different AccessKeyId/SecretAccessKey combinations, such
// as using a popFilter to use one AccessKeyId/SecretAccessKey pair for a particular AWSRegion and another AccessKeyId/SecretAccessKey
// pair for a different AWSRegion.  Likewise, the site may only use one AccessKeyId/SecretAccessKey across multiple AWSRegions.
// AllowedScope DIR
// DefaultPolicy  null
type AwsSignedOriginPullV4 struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// SecretAccessKey Shared secret key assigned to the AccessKeyID.
	SecretAccessKey string `json:"secretAccessKey," default:"" role:"normal" writeonly:"" advancedType:"PASSWORD"`

	// Enabled Set to true to enable policy.
	Enabled *bool `json:"enabled," default:"false" role:"normal" writeonly:""`

	// AwsRegion AWS region scope the access key.
	// see: https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
	AwsRegion string `json:"awsRegion," default:"" role:"normal" writeonly:""`

	// AccessKeyId Identifies the shared access key to be used to presign the request.
	AccessKeyID string `json:"accessKeyId," default:"" role:"normal" writeonly:""`

	// AuthenticationType Specifies what type of AWS authentication to use.
	// query: Provide authentication information using query string parameters. Using query parameters to authenticate requests is
	//   useful when you want to express a request entirely in a URL. This method is also referred as presigning a URL.
	// header: Use the HTTP Authorization header.
	// The query and header algorithms are identical except that the expireTimeSeconds policy only is applicable to the query authentication type.
	AuthenticationType string `json:"authenticationType,omitempty" default:"query" role:"normal" writeonly:"" enum:"[query,header]"`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// MethodFilter Method Filter is used to determine if this type should be applied or not based on List of HTTP Methods provided
	// Optionally, you may use an exclamation point in the list to describe the subset of HTTP methods excluded from this policy and all
	// other requests method will be included.
	// WARNING: You should not mix include and exclude in the same list.
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PopFilter POP filter is list of pattern to match POPs where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of POPs excluded from this policy.
	// Use lower case or use '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// RegionFilter Region filter is list of pattern to match Region where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of Regions excluded from this policy.
	// Use lower case or '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	RegionFilter string `json:"regionFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// SignedHeaders List of additional headers to be used when calculating the signature.
	// The headers "Host" and "x-amz-*" (customer AWS headers internall generated) are required and included by default.
	// Headers not permitted and invalidate the policy if set are "user-agent" and "x-amzn-trace-id".
	SignedHeaders string `json:"signedHeaders,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// AwsService AWS service scope of the access key.
	AwsService string `json:"awsService,omitempty" default:"s3" role:"normal" writeonly:""`

	// BucketIdentifier Specify to scope this policy to a particular bucket.  Note this value is directly coupled with the Host header, which is not always the origin hostname.
	// See https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html#VirtualHostingSpecifyBucket, Buckets are accessed primarily through the Host header
	// except when using SSL 'When you use virtual hosted–style buckets with Secure Sockets Layer (SSL), the SSL wildcard certificate only matches buckets that
	// don't contain periods. To work around this, use HTTP or write your own certificate verification logic. We recommend that you do not use periods (".") in
	// bucket names when using virtual hosted–style buckets.'
	// Buckets are identified in one of three ways:
	// 1) bucketname.s3.amazonaws.com
	// 2) s3.amazonaws.com/bucketname
	// 3) <custom.hostname>, such as www.myhost.com, where the host name is the bucketname
	// To match an origin request to the correct policy, the CDN appends the path of the URL in the origin request to the value in the Host header of the request.
	// The CDN checks if the constructed string "starts with" the value set in this policy.  The key factor is that the Host header is used, which may not
	// equal the origin hostname, such as the case a StaticHeader/OriginPull policy or proxying the Host header from a client request.
	// Leaving this blank/unset indicates it is the default policy to use for all origin pulls when a specific AwsSignedOriginPullV4 policy has not been matched.
	// If a default policy is used with one or more specific policies, the default needs to be listed last.
	BucketIdentifier string `json:"bucketIdentifier,omitempty" default:"" role:"normal" writeonly:""`

	// ExpireTimeSeconds Time period, in seconds, for which the generated presigned URL is valid.
	// Note, this policy only is applicable to the 'query' authentication type (see awsSignedOriginPullV4/authenticationType).
	ExpireTimeSeconds uint32 `json:"expireTimeSeconds,omitempty" default:"5" role:"normal" writeonly:""`
}

// UploadLimit Use to configure limits for client upload requests via POST or PUT.
// AllowedScope PRODUCT
// DefaultPolicy  null
type UploadLimit struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// MethodFilter Method Filter is used to determine if this type should be applied or not based on List of HTTP Methods provided
	// Optionally, you may use an exclamation point in the list to describe the subset of HTTP methods excluded from this policy and all
	// other requests method will be included.
	// WARNING: You should not mix include and exclude in the same list.
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// RequestLimitBytes Set the maximum number of bytes a server will accept and process for each request for this site. Note, this value should be less than or equal to ConcurrentLimitBytes.  If greater, ConcurrentLimitBytes the CDN will override and apply 5 times this value.
	RequestLimitBytes uint32 `json:"requestLimitBytes,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// ConcurrentLimitBytes Set the maximum total number of bytes a server will accept and process across all requests for this site at any given time. Note, this value should be greater than or equal to RequestLimitBytes. If less than, the CDN will override and apply 5 times  the RequestLimitBytes value.
	ConcurrentLimitBytes uint32 `json:"concurrentLimitBytes,omitempty" default:"524288000" role:"HWADMIN" writeonly:""`
}

// Waf Web Application Firewall
// WARNING: This setting is directory scope for flexiblity. However, enabling WAF on scope other that product scope will break WAF IF 1) the Vhost/Domain is not also at the same scope AND 2) the OriginPullHost/OriginUrl is not also at the same scope. The reason is that WAF will often inject iframe into html pages to request asset /sbbi/?sbbg=.... If /sbbi doesn't have WAF enabled, it will not work.
// WARNING: When path/url filter is used to enable WAF, only negative match should be used because if the filter would disable WAF for /sbbi/, the website will not load correctly.
// AllowedScope DIR
// DefaultPolicy  null
type Waf struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled If enabled, origin requests are proxied through Web Application Firewall before reaching the real origin.
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// StandAloneMode If enabled, all traffic going through the CDN will be in straight proxy mode. This is for customer who wants WAF only without CDN;
	// however, we want to protect the WAF infrastructure behind the CDN.
	StandAloneMode *bool `json:"standAloneMode,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// FailoverToOrigin This setting determines what should be the behavior for the CDN when WAF is not available.  When it's set to true,
	// the CDN will try to fulfill the request by falling back to the origin directly without going through WAF. Otherwise,
	// the CDN will return error directly to the end user without going the origin.
	FailoverToOrigin *bool `json:"failoverToOrigin,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// MethodFilter Method Filter is used to determine if this type should be applied or not based on List of HTTP Methods provided
	// Optionally, you may use an exclamation point in the list to describe the subset of HTTP methods excluded from this policy and all
	// other requests method will be included.
	// WARNING: You should not mix include and exclude in the same list.
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PopFilter POP filter is list of pattern to match POPs where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of POPs excluded from this policy.
	// Use lower case or use '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// CanonicalName This is the host name for WAF to identify the customer in the WAF cluster.  The Host header in the Request from
	// the CDN will contain this host name.  If this is left empty, the Host header from the request will be pass through
	// in the WAF requests.
	CanonicalName string `json:"canonicalName,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// WafClustersOverride Web Application Firewall
// AllowedScope PRODUCT
// DefaultPolicy  null
type WafClustersOverride struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// ServerList This is a comma seperated list of server:port to override the global-config setting for which WAF cluster to go to. The first server on the list is primary
	// and the rest of the list will be used as failover only when the first server is unreachable. The port of the server is configured in global-conf and cannot
	// be overwritten.
	ServerList string `json:"serverList," default:"" role:"HWADMIN" writeonly:"" list:"string"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// MethodFilter Method Filter is used to determine if this type should be applied or not based on List of HTTP Methods provided
	// Optionally, you may use an exclamation point in the list to describe the subset of HTTP methods excluded from this policy and all
	// other requests method will be included.
	// WARNING: You should not mix include and exclude in the same list.
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PopFilter POP filter is list of pattern to match POPs where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of POPs excluded from this policy.
	// Use lower case or use '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// XForwardedForBehavior Use to set or change how the CDN handles the X-Forwarded-For, which may affect how or what IP address the CDN associates to the end-user.
// AllowedScope DIR
// DefaultPolicy  {"enabled":false,"followHttpSpec":false}
type XForwardedForBehavior struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// FollowHttpSpec The legacy behavior for the CDN is to ignore the X-Forwarded-For from clients.   The CDN assumes/treats the client as the end-user. The CDN uses the IP address of  the TCP connection used by the client, which may or may not be the actual end-user.  It could be a a proxy instead. The CDN uses this IP address to generate its outgoing  X-Forwarded-For in origin and/or auth requests. The CDN also uses this IP address  for reporting, Geo Location, etc.   By setting this to true, the CDN processes and creates the X-Forwarded-For in accordance  with the HTTP spec, where the left-most IP address in the list is that of the end-user.  The CDN proxies a received X-Forwarded-For after appending the IP address of the  connected client/proxy.
	FollowHTTPSpec *bool `json:"followHttpSpec,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// MethodFilter Method Filter is used to determine if this type should be applied or not based on List of HTTP Methods provided
	// Optionally, you may use an exclamation point in the list to describe the subset of HTTP methods excluded from this policy and all
	// other requests method will be included.
	// WARNING: You should not mix include and exclude in the same list.
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// WebSocket WebSocket support (For SP 2.0 customers only)
// AllowedScope DIR
// DefaultPolicy  null
type WebSocket struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Enable WebSocket support for the scope.
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// PopFilter POP filter is list of pattern to match POPs where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of POPs excluded from this policy.
	// Use lower case or use '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// RegionFilter Region filter is list of pattern to match Region where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of Regions excluded from this policy.
	// Use lower case or '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	RegionFilter string `json:"regionFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// WsCanonicalName This is the host name to be used in the HTTP Upgrade request to the WebSocket origin. The Host header in the Request from
	// the CDN will contain this host name.  If this is left empty, the Host header from the client request will be pass through
	// in the WebSocket origin.
	WsCanonicalName string `json:"wsCanonicalName,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// WsMaxConnections This is the maximum number of connections each H2Proxy instance with make to each IP address of the WebSocket origin.
	WsMaxConnections uint16 `json:"wsMaxConnections,omitempty" default:"30000" role:"HWADMIN" writeonly:""`

	// WsOriginIdleTimeoutDuration Number of seconds to time out an idle connection to a WebSocket origin.
	WsOriginIdleTimeoutDuration uint32 `json:"wsOriginIdleTimeoutDuration,omitempty" default:"21600" role:"HWADMIN" writeonly:""`
}

// CacheControl Apply custom browser caching behaviors.
// AllowedScope DIR
// DefaultPolicy  [{"synchronizeMaxAge":true,"statusCodeMatch":"2*,301,302,303,304,305,307"},{"synchronizeMaxAge":false,"statusCodeMatch":"*"}]
type CacheControl struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// SynchronizeMaxAge Enables the synchronization of the edge cache with browser cache, such that content will expire from the browser at the same time it does from the edge.
	SynchronizeMaxAge *bool `json:"synchronizeMaxAge," default:"1" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// MustRevalidate If set to true, cacheControl must-revalidate header is added by CDN if sending its own cache-control headers.
	MustRevalidate *bool `json:"mustRevalidate,omitempty" default:"false" role:"normal" writeonly:""`

	// MaxAge Time in seconds browsers should cache content. Sets the Max-Age: HTTP header in cache responses. No default if not set (Max-Age will be set according to the cache policy). Accepts 'd' for day, 'h' for hour and 'm' for minute at the end of the number.
	MaxAge int32 `json:"maxAge,omitempty" default:"-1" role:"normal" writeonly:"" range:"-1,31536000"`

	// StatusCodeMatch Comma separated status code glob pattern to indicate which status code(s) this policy applies too.
	StatusCodeMatch string `json:"statusCodeMatch,omitempty" default:"" role:"normal" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// Override Override the Cache-Control header with the response. This takes precedent over the maxAge setting.
	Override string `json:"override,omitempty" default:"" role:"normal" writeonly:""`
}

// CacheKeyModification The Cache Key Modification policy allows for manipulation of the way the cache uniquely stores assets.
// AllowedScope PRODUCT
// DefaultPolicy  {"normalizeKeyPathToLowerCase":false}
type CacheKeyModification struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// NormalizeKeyPathToLowerCase When set, the path portion of the cache key will be normalized to lower case. So a client request or a purge for /PaTh/To/File.TxT will have a cache key of /path/to/file.txt. We will still pass the original request through to the origin, unaltered. This only applies to the path NOT the query string.
	NormalizeKeyPathToLowerCase *bool `json:"normalizeKeyPathToLowerCase," default:"false" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// DynamicContent Specify which parts of the end-user request should be used to build additional cache keys.
// AllowedScope DIR
// DefaultPolicy  null
type DynamicContent struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFields Provide a glob pattern for each HTTP request header that should be included in the cache key generation. More than one pattern must be separated by a comma. Use a colon in the pattern to define patterns for both the header name and value. A pattern without a colon ':' is treated as a header name ONLY match. If the pattern matches the header, then all values are used as a part of the cache key. If the pattern contains a colon, the CDN uses the pattern on the left of the colon to match the header. The pattern to the right of the colon is used to match the value. The CDN only uses the header/value as a part of the cache key if both patterns result in a match. Note, no pattern after a colon indicates an empty header (no value).
	// See the fnmatch manpage for acceptable glob patterns.
	HeaderFields string `json:"headerFields,omitempty" default:"" role:"normal" writeonly:"" list:"GLOB"`

	// QueryParams Comma separated list of query string parameters to treat as dynamic content parameters. This is not a glob match but a single '*' means all parameters will be used.
	QueryParams string `json:"queryParams,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// OriginPullCacheExtension Tell the CDN how to treat stale content.
// AllowedScope DIR
// DefaultPolicy  null
type OriginPullCacheExtension struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// ExpiredCacheExtension Number of seconds to extend file in cache if edge can not refresh the cache from the origin. Defaults to number of seconds original in  HTTP cache-control headers 0. This is the setting which determine how often we will try to go back to the origin to get the file again.
	ExpiredCacheExtension int32 `json:"expiredCacheExtension," default:"" role:"normal" writeonly:"" range:"0,31536000"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// OriginUnreachableCacheExtension This is the max time period we will serve from cache when we cannot get the file from the origin.
	OriginUnreachableCacheExtension int32 `json:"originUnreachableCacheExtension,omitempty" default:"86400" role:"normal" writeonly:"" range:"0,31536000"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// OriginPullPolicy Define how and when content stored specifically in the CDN cache expires and is replaced with new content from your Origin.
// AllowedScope DIR
// DefaultPolicy  [{"expireSeconds":86400,"statusCodeMatch":"2*,301,302,303,304,305,307","expirePolicy":"CACHE_CONTROL"},{"expireSeconds":60,"statusCodeMatch":"*","expirePolicy":"INGEST"}]
type OriginPullPolicy struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// ExpirePolicy <strong>Origin Controlled:</strong><br />Cache-Control headers on content from your Origin will determine expiration.
	// <br /><br />
	// <strong>Relative to Ingest:</strong><br />The time of ingest plus CDN TTL will determine expiration.
	// <br /><br />
	// <strong>Relative to Last Modified:</strong><br />CDN TTL will be used to check the Origin for modified assets, and if assets are modified the CDN will pull and cache them.
	// <br /><br />
	// <strong>Never Expire:</strong><br />Content in cache will remain in cache eternally.
	// <br /><br />
	// <strong>Do Not Cache:</strong><br />Content will not be cached.
	ExpirePolicy string `json:"expirePolicy," default:"" role:"normal" writeonly:"" enum:"[CACHE_CONTROL,INGEST,LAST_MODIFY,NEVER_EXPIRE,DO_NOT_CACHE]"`

	// ExpireSeconds If expirePolicy is INGEST or LAST_MODIFY, then this is the number of seconds since ingest, last access or last modify to expire the file. If expirePolicy is CACHE_CONTROL and there is no Cache-Control header, this is the default caching max-age for positive response, 0 in this case means cache as long as possible (until LRU remove the file). For negative response without a statusCodeMatch matching the status code, the originPullNegLinger value will be used.
	ExpireSeconds int32 `json:"expireSeconds," default:"" role:"normal" writeonly:"" range:"0,31536000"`

	// ForceBypassCache <p>Force this asset to bypass the cache. Typical use case is to turn this on for certain status code like 403, 404 ...etc using the statusCodeMatch key, so it won't bust our cache.</p>
	// <p>NOTE: This feature only applies for no-cache asset or OriginPull/DefaultBehavior is set to NOCACHE.</p>
	ForceBypassCache *bool `json:"forceBypassCache,omitempty" default:"false" role:"normal" writeonly:""`

	// MaxAgeZeroToNoCache Add on no-cache when maxage=0.
	MaxAgeZeroToNoCache *bool `json:"maxAgeZeroToNoCache,omitempty" default:"false" role:"normal" writeonly:""`

	// MustRevalidateToNoCache Behave like no-cache when must-revalidate is present in the Cache-Control header.
	MustRevalidateToNoCache *bool `json:"mustRevalidateToNoCache,omitempty" default:"false" role:"normal" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// HonorMustRevalidate Honor must-revalidate or proxy-revalidate cache-control from origin.
	HonorMustRevalidate *bool `json:"honorMustRevalidate,omitempty" default:"false" role:"normal" writeonly:""`

	// HonorNoCache Honor no-cache cache-control from origin.
	HonorNoCache *bool `json:"honorNoCache,omitempty" default:"false" role:"normal" writeonly:""`

	// HonorNoStore Honor no-store cache-control from origin.
	HonorNoStore *bool `json:"honorNoStore,omitempty" default:"false" role:"normal" writeonly:""`

	// HonorPrivate Honor private cache-control from origin.
	HonorPrivate *bool `json:"honorPrivate,omitempty" default:"false" role:"normal" writeonly:""`

	// HonorSMaxAge Honor s-maxage cache-control from origin.
	HonorSMaxAge *bool `json:"honorSMaxAge,omitempty" default:"false" role:"normal" writeonly:""`

	// UpdateHttpHeadersOn304Response Update the cached origin headers in cache when there is a newer value in a 304 response from the origin. To remove existing cached headers, the response from origin must contain the header key with an empty value.
	UpdateHTTPHeadersOn304Response *bool `json:"updateHttpHeadersOn304Response,omitempty" default:"false" role:"normal" writeonly:""`

	// EnableOPShieldForNoCache <p>Set to false to override the origin shielding setting when the cache-control is no-cache.</p>
	// <p>NOTE: This feature only applies for no-cache asset.</p>
	EnableOPShieldForNoCache *bool `json:"enableOPShieldForNoCache," default:"1" role:"HWADMIN" writeonly:""`

	// NoCacheBehavior Specify between the legacy or the spec behavior. The legacy behavior is to just treat the asset as always expired. The spec behavior is un-dedup the queue and proxy the headers from the client directly to the origin.
	NoCacheBehavior string `json:"noCacheBehavior,omitempty" default:"legacy" role:"normal" writeonly:"" enum:"[legacy,spec]"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// StatusCodeMatch Comma separated status code glob patterns to indicate which status code this policy applies to. If the key doesn't exist the policy will apply to all positive response codes (2xx, 301, 302 and 307).
	StatusCodeMatch string `json:"statusCodeMatch,omitempty" default:"" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// CachedHeadersOverride Comma-separated list of Glob patterns indicating what HTTP headers to cache from this origin. The list must be entirely inclusive (whitelist) or exclusive (blacklist). The '(?i)' prefix indicates patterns are case insensitive.
	CachedHeadersOverride string `json:"cachedHeadersOverride,omitempty" default:"" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HttpHeaders <p>Comma delimited list of HTTP response headers to proxy from the origin back to the edge for future delivery to the client. The list of headers ought to be a subset of Origin/CachedHeadersOverride or OriginPullPolicy/CachedHeadersOverride (or the default setting).</p>
	HTTPHeaders string `json:"httpHeaders,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// BypassCacheIdentifier <p>If not empty and the specified string appears in the Cache-Control header, the response from the origin will be proxied without caching.</p>
	// <p>NOTE: This feature only applies for no-cache asset.</p>
	BypassCacheIdentifier string `json:"bypassCacheIdentifier,omitempty" default:"" role:"normal" writeonly:""`
}

// ClientRequestQueue The script engine client request queue provides access to requests received by the CDN, giving you the ability to alter the request before the host processes it.  Access to the client’s request also provides you the ability to dynamically change your host's configuration policies based on business rules that require visibility to the client request.
// AllowedScope DIR
// DefaultPolicy  null
type ClientRequestQueue struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// ScriptPath This is the path of the script on the CDN caching server.
	ScriptPath string `json:"scriptPath," default:"" role:"HWADMIN" writeonly:""`

	// SendRequestBody Boolean flag for indicating whether the CDN Caching Server needs to send the body provided in a Client request to the Script Engine.
	SendRequestBody *bool `json:"sendRequestBody,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// QuitOnError Indicates whether GFS should quit processing a request if this script fails or is unable to execute.
	QuitOnError *bool `json:"quitOnError,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// ProvideIPGeoInfo This is a flag for indicating whether or not the script requires Geographic IP information. (Default: false)
	ProvideIPGeoInfo *bool `json:"provideIPGeoInfo,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// IpListAccessCode This is used as a variable in a script for controlling how an ACL is applied. If set to allow, the ACL is treated as a white-list.  If set to deny, the ACL is treated as a black-list (Default: allow).
	// NOTE: This is NOT a base functionality provided by all scripts.  It must be written into each customer script.  Check with the script author prior to using this key.
	IPListAccessCode string `json:"ipListAccessCode,omitempty" default:"allow" role:"HWADMIN" writeonly:"" enum:"[allow,deny]"`

	// LogLevel This determines the level of script engine logging done for the queue. It is useful for debugging script engine  failures when developing new scripts.
	LogLevel string `json:"logLevel,omitempty" default:"error" role:"HWADMIN" writeonly:"" enum:"[debug,info,warning,error,crit]"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// IpList This is a comma separated list of IP address and/or CIDR blocks that limit the execution of the script.
	IPList string `json:"ipList,omitempty" default:"" role:"HWADMIN" writeonly:"" list:"IP"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// IpListStageDir This is a subdirectory path that must exist in the request path in order to apply the script. For example,  if the value is "foo/bar" (with or without slashes on the ends), and /foo/bar/ exists in the request path, then a dynamic IP-based content protection policy will be applied. Otherwise, the request will be allowed to pass through untouched.
	IPListStageDir string `json:"ipListStageDir,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// RequestBodyMaximumSize Integer that defines the maximum size in bytes of a Client request's body that can be sent by the CDN Caching Server to the Script Engine when SendRequestBody has been set to true (enabled).
	RequestBodyMaximumSize uint32 `json:"requestBodyMaximumSize,omitempty" default:"1024" role:"HWADMIN" writeonly:""`
}

// ClientResponseQueue The script engine client response queue policy allows you to register a PHP script to execute on the CDN caching server prior to the server returning a response to a client.  Scripts defined in this queue can modify, add and/or  delete HTTP headers in the CDN repsonse.
// AllowedScope DIR
// DefaultPolicy  null
type ClientResponseQueue struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// ScriptPath This is the path of the script on the CDN caching server.
	ScriptPath string `json:"scriptPath," default:"" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// QuitOnError This value instructs the CDN to terminate the processing of the request on all script failures.
	QuitOnError *bool `json:"quitOnError,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// LogLevel This determines the level of script engine logging done for the queue. It is useful for debugging script  engine failures when developing new scripts.
	LogLevel string `json:"logLevel,omitempty" default:"error" role:"HWADMIN" writeonly:"" enum:"[debug,info,warning,error,crit]"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// ClientKeepAlive The clientKeepAlive policy allows you to specify how long you want the CDN caching server to keep an  client connection open after serving a request.
// AllowedScope DIR
// DefaultPolicy  {"timeout":-1}
type ClientKeepAlive struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Timeout Timeout in seconds for idle client connections of GFS.
	Timeout int32 `json:"timeout," default:"" role:"HWADMIN" writeonly:"" range:"-1,300"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// MethodFilter Method Filter is used to determine if this type should be applied or not based on List of HTTP Methods provided
	// Optionally, you may use an exclamation point in the list to describe the subset of HTTP methods excluded from this policy and all
	// other requests method will be included.
	// WARNING: You should not mix include and exclude in the same list.
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PopFilter POP filter is list of pattern to match POPs where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of POPs excluded from this policy.
	// Use lower case or use '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// RegionFilter Region filter is list of pattern to match Region where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of Regions excluded from this policy.
	// Use lower case or '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	RegionFilter string `json:"regionFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// ConsistentHashing The consistent hashing policy allows you to customize the consistent hashing algorithm used by Doppler.
// AllowedScope FILE
// DefaultPolicy  {"defaultLoadBalanceHosts":"0"}
type ConsistentHashing struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// DefaultLoadBalanceHosts The number of additional hosts you would like Doppler to use when selecting a host for content.  The  additional hosts will be evenly load balanced by Doppler and will cause content to be duplicated for the customer on each server.  The value entered in this field will be limited to the number of hosts  available in the data center.
	DefaultLoadBalanceHosts string `json:"defaultLoadBalanceHosts," default:"" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// EnableSidewayPulling Pull the file from the primary edge in the same pop
	EnableSidewayPulling *bool `json:"enableSidewayPulling," default:"1" role:"HWADMIN" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PopFilter
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// RegionFilter
	RegionFilter string `json:"regionFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// H2proxyCaching This policy is used to override the memoryCacheable policy derived from other policies.
// AllowedScope DIR
// DefaultPolicy  {"enabled":true}
type H2proxyCaching struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled This is a HWADMIN control variable, if disabled, to force H2Proxy to NOT cache assets in memory even if memoryCacheable policy indicates it should.
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// Customer settings
// AllowedScope ROOT
// DefaultPolicy  {"suspended":false}
type Customer struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// CompressAccessLogs Enable gzip compression of access logs for this customer.
	CompressAccessLogs *bool `json:"compressAccessLogs," default:"1" role:"HWADMIN" writeonly:""`

	// OpLogs Should origin pull logging be enabled for this customer.
	OpLogs *bool `json:"opLogs,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// ReceiptLogs Should receipt access logging be enabled for this customer (see RequestReceipt).
	ReceiptLogs *bool `json:"receiptLogs,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// UploadAccessLogsToHCS Upload access logs for this customer directly to Highwinds Cloud Storage
	UploadAccessLogsToHCS *bool `json:"uploadAccessLogsToHCS," default:"1" role:"HWADMIN" writeonly:""`

	// ReceiptLogFields Comma delimited list of HTTP header fields to append to the standard fields in the Receipt Access log. Each field must have the 'sc:' (server-to-client) or 'cs:' (client-to-server) prefix.
	// NOTE: the colon (:) is required.
	ReceiptLogFields string `json:"receiptLogFields,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// AccessLogFields Comma delimited list of HTTP header fields to append to the standard fields in the access log. Each field must have the 'sc:' (server-to-client) or 'cs:' (client-to-server) prefix.
	// NOTE: the colon (:) is required.
	AccessLogFields string `json:"accessLogFields,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// OpLogFields Comma delimited list of HTTP header fields to append to the standard fields in the origin pull log. Each field must have the 'sc:' (server-to-client) or 'cs:' (client-to-server) prefix.
	// NOTE: the colon (:) is required.
	OpLogFields string `json:"opLogFields,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// AccessLogExpireTimeHCS Time in seconds that an accesslog is allowed to live before it is expired from HCS
	AccessLogExpireTimeHCS uint32 `json:"accessLogExpireTimeHCS,omitempty" default:"3888000" role:"HWADMIN" writeonly:""`

	// AccessLogExpireTimeLocal Time in seconds that an accesslog is allowed to live before it is expired from the accesslogger local storage
	// NOTE: This is used by SysEng's script to purge old access log files and the default value is subjected to change
	AccessLogExpireTimeLocal uint32 `json:"accessLogExpireTimeLocal,omitempty" default:"3888000" role:"HWADMIN" writeonly:""`
}

// DeviceBasedDynamicContent Extends dynamic content by rewriting the "DEVICE" parameter and header based on the User-Agent in the Client Request.
// AllowedScope DIR
// DefaultPolicy  null
type DeviceBasedDynamicContent struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// MobileDevicePattern The regular expression used to determine if a User-Agent is a mobile device.  The CDN considers a positive match using this pattern as being a mobile device.
	MobileDevicePattern string `json:"mobileDevicePattern," default:"" role:"HWADMIN" writeonly:"" advancedType:"REGEX"`

	// PassToOrigin A flag that tells the CDN to pass through the DEVICE parameter to the origin.
	PassToOrigin *bool `json:"passToOrigin,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// NameOverride Provides the ability to change the name of the parameter/header of interest.
	NameOverride string `json:"nameOverride,omitempty" default:"device" role:"HWADMIN" writeonly:""`
}

// HashType The type of the hash
// AllowedScope ROOT
// DefaultPolicy  {"class":"HOST"}
type HashType struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`
}

// InternalError The CDN internal error caching policy allows you to control the TTL for internally generated errors in the caching servers.
// AllowedScope DIR
// DefaultPolicy  {"maxAge":10}
type InternalError struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// MaxAge The default caching time for internally generated errors.
	MaxAge string `json:"maxAge," default:"" role:"HWADMIN" writeonly:"" advancedType:"ttl"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// LanguageRedirect The language code origin request rewrite policy allows you to rewrite responses from your origin to 301 response codes such that you can re-issue the request to your origin with a new request URL.  This policy was created to specifically map language codes in a origin request URL to default languages when a resource was not found on the origin.  NOTE: This policy requires a custom script to be configured on the script engine.
// AllowedScope DIR
// DefaultPolicy  null
type LanguageRedirect struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// PathRegex A regular expression that identifies the paths or specific resource that applies to this policy
	PathRegex string `json:"pathRegex," default:"" role:"HWADMIN" writeonly:"" advancedType:"REGEX"`

	// Mapping This is a list of language code mappings that maps one or more requested codes to a code to use in the redirect request. The mapping must be separated by an equals sign. If more than one mapping is provided, the mappings will be applied in the order listed until a match is made. Note that an asterisk is permitted to represent all requested codes.
	Mapping string `json:"mapping," default:"" role:"HWADMIN" writeonly:"" list:"string"`

	// HttpCode The origin HTTP response code that applies to this policy
	HTTPCode uint16 `json:"httpCode," default:"" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// MidTierCaching Mid Tier Caching Configuration
// AllowedScope DIR
// DefaultPolicy  {"enabled":false}
type MidTierCaching struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Enable Mid Tier Caching
	// WARNING: When using this feature to pull content from a sister pop, popFilter has to be set! WARNING: DO NOT ENABLE THIS SETTING BEFORE CONSULTING WITH THE DEV AND SYSENG TEAM!
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// Protocol The protocol used to communicate with the Mid Tier Cache server. 'match' will use the same protocol from the end user request.
	Protocol string `json:"protocol," default:"" role:"HWADMIN" writeonly:"" enum:"[http,https,match]"`

	// RequestFormat This setting determine how the request should be constructed when origin pulling from the Mid Tier Cache server. "internal" will construct the request like we are communicating with a shield. "external" will construct the request like we are communicating with an external origin.
	// WARNING: only "internal" is support for now as of 974-1
	RequestFormat string `json:"requestFormat," default:"" role:"HWADMIN" writeonly:"" enum:"[internal,external]"`

	// ServerAndPort The mid tier caching server endpoint in the format of server:port.  If port is not present, it will be default to port 80
	ServerAndPort string `json:"serverAndPort," default:"" role:"HWADMIN" writeonly:""`

	// PopFilter
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// RegionFilter
	RegionFilter string `json:"regionFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// OriginRequestQueue The script engine origin pull request queue policy allows you to register a PHP script to execute on the CDN caching server prior to the server making an origin pull request  to your origin.  Scripts defined in this queue can modify, add, and/or delete HTTP headers on the origin pull request.
// AllowedScope DIR
// DefaultPolicy  null
type OriginRequestQueue struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// ScriptPath This is the path of the script on the CDN caching server.
	ScriptPath string `json:"scriptPath," default:"" role:"HWADMIN" writeonly:""`

	// SendRequestBody Boolean flag for indicating whether the CDN Caching Server needs to send the body of the Origin request to the Script Engine.
	SendRequestBody *bool `json:"sendRequestBody,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// QuitOnError This indicates the behavior of the caching server on script errors.  When selected, the caching server issues a 500 HTTP response code to an end user if a request failed to properly execute the script  designated by this policy.  If this option is unselected, the caching server ignores the script errors and continues processing the request.
	QuitOnError *bool `json:"quitOnError,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// LogLevel This determines the level of script engine logging done for the queue.  It is useful for debugging script engine failures when developing new scripts.
	LogLevel string `json:"logLevel,omitempty" default:"error" role:"HWADMIN" writeonly:"" enum:"[debug,info,warning,error,crit]"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// RequestBodyMaximumSize Integer that defines the maximum size in bytes of a Origin request's body that can be sent by the CDN Caching Server to the Script Engine when SendRequestBody has been set to true (enabled).
	RequestBodyMaximumSize uint32 `json:"requestBodyMaximumSize,omitempty" default:"1024" role:"HWADMIN" writeonly:""`
}

// OriginResponseQueue The script engine origin pull response queue policy allows you to register a PHP script to execute on the CDN caching server prior to the server proxying or caching the response from  your origin.  Scripts defined in this queue can modify, add, and/or delete HTTP headers on the response from your origin.
// AllowedScope DIR
// DefaultPolicy  null
type OriginResponseQueue struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// ScriptPath This is the path of the script on the CDN caching server.
	ScriptPath string `json:"scriptPath," default:"" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// QuitOnError This indicates the behavior of the caching server on script errors.  When selected, the caching server issues a 500 HTTP response code to an end user if a request failed to properly execute the script  designated by this policy.  If this option is unselected, the caching server ignores script errors and continues processing the request.
	QuitOnError *bool `json:"quitOnError,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// LogLevel This determines the level of script engine logging done for the queue.  This is useful for debugging script engine failures when developing new scripts.
	LogLevel string `json:"logLevel,omitempty" default:"error" role:"HWADMIN" writeonly:"" enum:"[debug,info,warning,error,crit]"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// PathModification Request URL rewriting policies can be used to modify the URL path of a CDN request.  This policy requires a custom script to be configured.
// This policy requires the Script Engine service enabled on your account. If you do not have this service enabled, please contact your sales representative for more information.
// AllowedScope DIR
// DefaultPolicy  null
type PathModification struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// RegEx A regular expression pattern used to identify a portion of the path that is targeted by this policy for  replacement.  This should not include the beginning and trailing '/' character.  All other '/' characters will be escaped  by default.
	RegEx string `json:"regEx," default:"" role:"HWADMIN" writeonly:"" advancedType:"REGEX"`

	// Header A regular expression that is matched against incoming request headers to determine which requests are targeted for rewriting. For example, User-Agent: /(?=.*Opera)(?=.*Mobi).*|(.*Android.*)/i will match all Opera Mobile and Android requests. To match all requests, put * here.
	Header string `json:"header," default:"" role:"HWADMIN" writeonly:"" advancedType:"REGEX"`

	// Replacement The string used to replace the portion of the URL targeted by the search pattern
	Replacement string `json:"replacement," default:"" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// IncludeQSParamInPath Include the Query String parameters in the path for matching.
	IncludeQSParamInPath *bool `json:"includeQSParamInPath,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// CaseInsensitiveMatch Use case insensitive regex match
	CaseInsensitiveMatch *bool `json:"caseInsensitiveMatch,omitempty" default:"false" role:"HWADMIN" writeonly:""`

	// EscapeSlashCharacter Whether or not to escape the slash character in the regular expression.  Turn this off if you will escape this character yourself in the regular expression below.
	EscapeSlashCharacter *bool `json:"escapeSlashCharacter," default:"1" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// ScriptNegCaching The legacy negative response code caching policy allowed the CDN to cache the body of non-200 responses.  This policy is no longer required now that the CDN supports the caching of all response codes from an origin.  Please consider removing this policy and configuring this behavior using a CDN Caching policy.
// AllowedScope DIR
// DefaultPolicy  null
type ScriptNegCaching struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// StatusCodes This is a comma separated list of negative HTTP response codes to cache.  Wildcards, such as 4**, are supported to represent all 400 level status codes.
	StatusCodes string `json:"statusCodes," default:"" role:"HWADMIN" writeonly:"" list:"string"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// OriginHeaderOverride Pipe delimited ('|') list of headers to overwrite or insert in the origin response prior to the processing it on the caching server
	OriginHeaderOverride string `json:"originHeaderOverride,omitempty" default:"" role:"HWADMIN" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// ServerlessScripting Serverless Script Processing
// AllowedScope DIR
// DefaultPolicy  null
type ServerlessScripting struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// ScriptId This is the Id of the script to be loaded.
	ScriptID string `json:"scriptId," default:"" role:"HWADMIN" writeonly:""`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"HWADMIN" writeonly:""`

	// PopFilter POP filter is list of pattern to match POPs where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of POPs excluded from this policy.
	// Use lower case or use '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// RegionFilter Region filter is list of pattern to match Region where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of Regions excluded from this policy.
	// Use lower case or '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	RegionFilter string `json:"regionFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`

	// ProcessorAddress Location of server to process the scripts.
	ProcessorAddress string `json:"processorAddress,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// TossbackBypass Instructs the CDN caching server to continue serving a pipeline request without tossing the connection back to Doppler.
// AllowedScope PRODUCT
// DefaultPolicy  {"enabled":false}
type TossbackBypass struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Enables the bypass of BTP tossbacks between the caching server and Doppler.
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// CloseHalfOpenConnections Instructs the CDN caching server to fully close the connection immediately after receiving a TCP FIN from the client.
// AllowedScope PRODUCT
// DefaultPolicy  {"enabled":false}
type CloseHalfOpenConnections struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Force the close of client connections upon receiving TCP FIN from clients.
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// TossbackAlways Instructs the CDN caching server to always toss a connection back to Doppler to choose the edge for next pipeline request.
// AllowedScope PRODUCT
// DefaultPolicy  {"enabled":false}
type TossbackAlways struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Force the tossback of connections from the caching server to Doppler.
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// Rti Used to set or change commodity routing versus other types of routing for files in certain file paths (2: commodity routing, 0: premium (default))
// AllowedScope DIR
// DefaultPolicy  null
type Rti struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Enable use of premium versus commodity (versus other-future-hybrid-approaches) routing.
	Enabled *bool `json:"enabled," default:"false" role:"HWADMIN" writeonly:""`

	// TableNumber What table number will be used by gfs to deliver certain file.
	TableNumber uint32 `json:"tableNumber," default:"" role:"HWADMIN" writeonly:""`

	// HeaderFilter Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	// headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	// WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	// WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PopFilter POP filter is list of pattern to match POPs where Policy needs to applied.
	// Optionally, you may use an exclamation point in the list to describe the subset of POPs excluded from this policy.
	// Use lower case or use '(?i)' prefix which indicates patterns are case insensitive.
	// WARNING: You should not mix include and exclude in the same list.
	PopFilter string `json:"popFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"HWADMIN" writeonly:"" list:"GLOB"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"HWADMIN" writeonly:""`
}

// ClientRequestModification Configure options for modifying client requests.
// AllowedScope DIR
// DefaultPolicy  null
type ClientRequestModification struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// FlowControl <p>When multiple policies are defined, flow control describe how we want to process next policies. If client request was modified we apply other policies based on flow control.</p>
	// <p><b>Next</b> - This is default, continue processing next policy</p>
	// <p><b>Break</b> - Don't process any other policy</p>
	FlowControl string `json:"flowControl,omitempty" default:"next" role:"normal" writeonly:"" enum:"[next,break]"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// AddHeaders <p>The http header that will be injected into the response. By default, the header specified here will be appended to the response.  You can specify different behaviors as follows:</p>
	// <p><b>append:</b> add the header to the response without checking for its existence first (default)</p>
	// <p><b>replace:</b> replaces the header if and only if it exists.</p>
	// <p><b>create:</b> adds the header if it doesn't exist</p>
	AddHeaders string `json:"addHeaders,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// HeaderPattern <p>The pattern to match against client response headers.  This value supports the following types of matches:</p>
	// <p><b>Wildcard Match:</b>  For wildcard matches use the key word wildcard (e.g. wildcard:/dir/*.html)</p>
	// <p><b>Glob Match:</b> For a glob match use the keyword glob (e.g. glob:/dir/*.html)</p>
	// <p><b>Regex Match:</b> For regular expression match use the keyword regex (e.g. regex:/dir/*.html)</p>
	// <p><b>RegexGlobal:</b> For regular expression global match use the keyword regexglobal (e.g. regexglobal:/dir/*.hmtl)</p>
	// <p> Regex and Regexglobal use RE2 Syntax: https://github.com/google/re2/wiki/Syntax).
	HeaderPattern string `json:"headerPattern,omitempty" default:"" role:"normal" writeonly:""`

	// UrlPattern The URL pattern that applies to this policy.
	URLPattern string `json:"urlPattern,omitempty" default:"" role:"normal" writeonly:""`

	// UrlRewrite The replacement URL used in conjunction with the url pattern. This key can be used with the client and server variables (e.g. %server.ip%)
	URLRewrite string `json:"urlRewrite,omitempty" default:"" role:"normal" writeonly:""`

	// HeaderRewrite The replacement header used in conjunction with the header pattern.  This key can be used with the client and server variables (e.g. %server.ip%)
	HeaderRewrite string `json:"headerRewrite,omitempty" default:"" role:"normal" writeonly:""`
}

// ClientResponseModification Configure options for client response modification.
// AllowedScope DIR
// DefaultPolicy  null
type ClientResponseModification struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// FlowControl <p>When multiple policies are defined, flow control describe how we want to process next policies. If client request was modified we apply other policies based on flow control.</p>
	// <p><b>Next</b> - This is default, continue processing next policy</p>
	// <p><b>Break</b> - Don't process any other policy</p>
	FlowControl string `json:"flowControl,omitempty" default:"next" role:"normal" writeonly:"" enum:"[next,break]"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// AddHeaders <p>The http header that will be injected into the response. By default, the header specified here will be appended to the response.  You can specify different behaviors as follows:</p>
	// <p><b>append:</b> add the header to the response without checking for its existence first (default)</p>
	// <p><b>replace:</b> replaces the header if and only if it exists.</p>
	// <p><b>create:</b> adds the header if it doesn't exist</p>
	AddHeaders string `json:"addHeaders,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// HeaderPattern <p>The pattern to match against client response headers.  This value supports the following types of matches:</p>
	// <p><b>Wildcard Match:</b>  For wildcard matches use the key word wildcard (e.g. wildcard:/dir/*.html)</p>
	// <p><b>Glob Match:</b> For a glob match use the keyword glob (e.g. glob:/dir/*.html)</p>
	// <p><b>Regex Match:</b> For regular expression match use the keyword regex (e.g. regex:/dir/*.html)</p>
	// <p><b>RegexGlobal:</b> For regular expression global match use the keyword regexglobal (e.g. regexglobal:/dir/*.hmtl)</p>
	// <p> Regex and Regexglobal use RE2 Syntax: https://github.com/google/re2/wiki/Syntax).
	HeaderPattern string `json:"headerPattern,omitempty" default:"" role:"normal" writeonly:""`

	// HeaderRewrite The replacement header used in conjunction with the header pattern.  This key can be used with the client and server variables (e.g. %server.ip%)
	HeaderRewrite string `json:"headerRewrite,omitempty" default:"" role:"normal" writeonly:""`

	// StatusCodeRewrite The new client response code to issue.
	StatusCodeRewrite uint32 `json:"statusCodeRewrite,omitempty" default:"" role:"normal" writeonly:""`
}

// OriginRequestModification Configure options for modifying Origin requests.
// AllowedScope DIR
// DefaultPolicy  null
type OriginRequestModification struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// AddHeaders
	AddHeaders string `json:"addHeaders,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// FlowControl
	FlowControl string `json:"flowControl,omitempty" default:"" role:"normal" writeonly:""`

	// HeaderPattern
	HeaderPattern string `json:"headerPattern,omitempty" default:"" role:"normal" writeonly:""`

	// HeaderRewrite
	HeaderRewrite string `json:"headerRewrite,omitempty" default:"" role:"normal" writeonly:""`

	// UrlPattern
	URLPattern string `json:"urlPattern,omitempty" default:"" role:"normal" writeonly:""`

	// UrlRewrite
	URLRewrite string `json:"urlRewrite,omitempty" default:"" role:"normal" writeonly:""`
}

// OriginResponseModification Configure options for Origin response modification.
// AllowedScope DIR
// DefaultPolicy  null
type OriginResponseModification struct {
	// ID configurationID, used when update configuration
	ID int64 `json:"id,omitempty"`

	// Enabled Generic Enabled Flag for all Config Types
	Enabled *bool `json:"enabled," default:"1" role:"normal" writeonly:""`

	// FlowControl <p>When multiple policies are defined, flow control describe how we want to process next policies. If client request was modified we apply other policies based on flow control.</p>
	// <p><b>Next</b> - This is default, continue processing next policy</p>
	// <p><b>Break</b> - Don't process any other policy</p>
	FlowControl string `json:"flowControl,omitempty" default:"next" role:"normal" writeonly:"" enum:"[next,break]"`

	// MethodFilter
	MethodFilter string `json:"methodFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// HeaderFilter
	HeaderFilter string `json:"headerFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// PathFilter Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	// This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	// patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	// Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	// other requests will be included.
	// Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	// it also supports Protocol and Host Name match.
	// pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: dir/*.html).
	// Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	// Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	// Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	// WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter,omitempty" default:"*" role:"normal" writeonly:"" list:"GLOB"`

	// AddHeaders <p>The http header that will be injected into the response. By default, the header specified here will be appended to the response.  You can specify different behaviors as follows:</p>
	// <p><b>append:</b> add the header to the response without checking for its existence first (default)</p>
	// <p><b>replace:</b> replaces the header if and only if it exists.</p>
	// <p><b>create:</b> adds the header if it doesn't exist</p>
	AddHeaders string `json:"addHeaders,omitempty" default:"" role:"normal" writeonly:"" list:"string"`

	// Comment Explain to other users why you are making this change
	Comment string `json:"comment,omitempty" default:"" role:"normal" writeonly:""`

	// HeaderPattern <p>The pattern to match against origin response headers.  This value supports the following types of matches:</p>
	// <p><b>Wildcard Match:</b>  For wildcard matches use the key word wildcard (e.g. wildcard:/dir/*.html)</p>
	// <p><b>Glob Match:</b> For a glob match use the keyword glob (e.g. glob:/dir/*.html)</p>
	// <p><b>Regex Match:</b> For regular expression match use the keyword regex (e.g. regex:/dir/*.html)</p>
	// <p><b>RegexGlobal:</b> For regular expression global match use the keyword regexglobal (e.g. regexglobal:/dir/*.hmtl)</p>
	// <p> Regex and Regexglobal use RE2 Syntax: https://github.com/google/re2/wiki/Syntax).
	HeaderPattern string `json:"headerPattern,omitempty" default:"" role:"normal" writeonly:""`

	// HeaderRewrite The replacement header used in conjunction with the header pattern.  This key can be used with the client and server variables (e.g. %server.ip%)
	HeaderRewrite string `json:"headerRewrite,omitempty" default:"" role:"normal" writeonly:""`

	// StatusCodeRewrite Origin response code
	StatusCodeRewrite uint32 `json:"statusCodeRewrite,omitempty" default:"" role:"normal" writeonly:""`
}

// OriginPullHost doesn't included in configuration DOC
// Just contains primary&secondary configure ID
type OriginPullHost struct {
	ID int `json:"id"` // configure ID

	// Primary configure ID
	// This field include number by default, but set to interface{} in order to relace with origin struct
	Primary   interface{} `json:"primary"` // This field include number by default, but set to
	Secondary interface{} `json:"secondary"`

	// Default path used for origin pull request
	Path string `json:"path"`

	//Header Filter is used to determine if this type should be applied or not based on Expression Provide. Expressions are match against request headers.
	//This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy. By default the
	//patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	//Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	//other requests will be included.  Please note that you should not mix include and exclude patterns in the same list.
	//headerFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	//Wildcard match - '*' will match all characters including '/'. (Ex: wildcard: User-Agent: Mozilla* - will match User-Agent: Mozilla/Firefox 6.0 or Mozilla 8.0).
	//Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:User-Agent: Mozilla* - will match Mozilla 6.0. Won't match Mozilla/Firefox 6.0)
	//Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/User-Agent:.*(iphone|android).*/,/EXP/).
	//WARNING: You should not mix include and exclude patterns in the same list.
	//WARNING: Header Filter might not work for originPullPolicy unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	//WARNING: Header Filter might not work for originRequestQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	//WARNING: Header Filter might not work for OriginResponseQueue unless if it is Dynamic Cache based on Header or if it is non-cacheable asset.
	HeaderFilter string `json:"headerFilter" default:"*" role:"HWADMIN"`

	//Method Filter is used to determine if this type should be applied or not based on List of HTTP Methods provided
	//Optionally, you may use an exclamation point in the list to describe the subset of HTTP methods excluded from this policy and all
	//other requests method will be included.
	//WARNING: You should not mix include and exclude in the same list.
	MethodFilter string `json:"methodFilter" default:"*" role:"HWADMIN"`

	//Path Filter is used to determine if this type should be applied or not based on Expression Provide.
	//This is a list of patterns that are used to describe a subset of requests that are included (or optionally excluded) by this policy.  By default the
	//patterns you add to this list are interpreted as described in the subset of requests included in this policy and all others will be ignored.
	//Optionally, you may use an exclamation point on each element in the list to describe the subset of requests excluded from this policy and all
	//other requests will be included.
	//Expression can either be used as Path Filter or URL Filter. If expression starts with [protocol]:// it is consider as URL Filter. In URL filter along with Path Match
	//it also supports Protocol and Host Name match.
	//pathFilter support three types of Match - Wildcard Match, Glob Match, Regex Match. Filter expression should start with Match Type (Ex: wildcard: /dir/*.html or glob: /dir/*.html).
	//Wildcard match - '*' will match all characters including '/'. (Ex: wildcard:/DIR/*.html - will match any HTML file under DIR or any Sub-directory under DIR. Will match DIR/FOO/index.html).
	//Glob match - Its Path("/") Match. '*' will match all characters except '/'. (Ex: glob:/DIR/*.html - will match all HTML file under DIR and not HTML file under sub directory of DIR. Won't match DIR/FOO/index.html)
	//Regex match, it will use RE2 rules for regular expression match (RE2 Syntax: https://github.com/google/re2/wiki/Syntax). Expression should be sorruned by "/" (Ex: regex:/.*DIR/\d/.*file.txt/,/EXP/).
	//WARNING: You should not mix include and exclude patterns in the same list.
	PathFilter string `json:"pathFilter" default:"*" role:"HWADMIN"`
}

// ConfigScopeObject Host configuration for given scope
/*

 */
type ConfigScopeObject struct {
	Scope `json:"scope"` //scope
}

// ConfigScope Configuration scope
type ConfigScope struct {
	ID          int    `json:"id"`             //Id
	Platform    string `json:"platform"`       //platform
	Path        string `json:"path"`           //path
	Name        string `json:"name,omitempty"` //name
	CreatedDate string `json:"createdDate"`    //createdDate
	UpdatedDate string `json:"updatedDate"`    //updatedDate
}

// ConfigStatus Configuration status
type ConfigStatus struct {
	Progress float32 `json:"progress"` //Percentage of edges which have received the specified configuration update
}

// ConfigScopeList ConfigScope list
type ConfigScopeList struct {
	List []*ConfigScope `json:"list"` //list
}

// ConfigurationHostNamesList Hostnames list
type ConfigurationHostNamesList struct {
	List []*HostName `json:"list"` //Hostnames
}

// HostName Hostname
type HostName struct {
	ID          int         `json:"id"`     //Unique id used for editing hostnames
	Domain      string      `json:"domain"` //Domain of the hostname entry
	Scope       ConfigScope `json:"scope"`  //Scope at which this hostname applies
	AccountHash string      `json:"accountHash"`
	HostHash    string      `json:"hostHash"`
	AccountName string      `json:"accountName"`
	Name        string      `json:"name"`
	ScopeID     int         `json:"scopeId"`
}

// Graph Configuration graph
type Graph struct {
	// value could been empty array, means this account doesn't contains any vaild hosts
	Hosts   map[string]*HostAttributes `json:"hosts,omitempty"`
	Origins []*OriginAttributes        `json:"origins,omitempty"`
}

// OriginsUsed return origins list used by hosthash's scopeID
func (g *Graph) OriginsUsed(hostHash, scopeID string) []*OriginAttributes {
	if !(g.Hosts[hostHash] == nil || g.Hosts[hostHash].Scopes == nil || g.Hosts[hostHash].Scopes[scopeID] == nil) {
		return []*OriginAttributes{}
	}
	return g.Hosts[hostHash].Scopes[scopeID].Origins
}

// HostAttributes Host attributes
type HostAttributes struct {
	Name   string                  `json:"name"`   //name
	Scopes map[string]*GraphScopes `json:"scopes"` //scopes
}

// OriginAttributes used by Graph
type OriginAttributes struct {
	ID   string `json:"id,omitempty"`   // origin ID
	Name string `json:"name,omitempty"` // origin name
	URL  string `json:"url,omitempty"`  // origin hostname
	Type string `json:"type,omitempty"` // origin type
}

// GraphScopes contains hostnames&origins
type GraphScopes struct {
	Path      string              `json:"path"` // /cds
	Hostnames []string            `json:"hostnames"`
	Origins   []*OriginAttributes `json:"origins"`
}

// Scopes configurations
type Scopes struct {
	List []*Scope `json:"list"` //scopes
}

// Scope scope
type Scope struct {
	Path string `json:"path"` ///cds
	ID   int    `json:"id"`

	//CDS == CDN ; ALL Reference to default platform such like stream
	Platform string `json:"platform"`

	CreatedDate string `json:"createdDate"`
	UpdatedDate string `json:"updatedDate"`
	Name        string `json:"name,omitempty"`
}

// OriginInfo origin information
type OriginInfo struct {
	URL  string `json:"url"`  //url
	Type string `json:"type"` //type
	Name string `json:"name"` //name
	ID   string `json:"id"`   //id
}

// Configuration A container for configuration on a scope
type Configuration struct {
	ID                             string                          `json:"id,omitempty"` //For updates, this is the configuration receipt id that can be used to poll for status
	Scope                          Scope                           `json:"scope"`        //The scope at which this configuration is set
	AccessLogger                   *AccessLogger                   `json:"accessLogger,omitempty"`
	AccessLogs                     *AccessLogs                     `json:"accessLogs,omitempty"`
	AccessLogIPObfuscation         *AccessLogIPObfuscation         `json:"accessLogIpObfuscation,omitempty"`
	AccessLogsConfig               *AccessLogsConfig               `json:"accessLogsConfig,omitempty"`
	HostnameReporting              *HostnameReporting              `json:"hostnameReporting,omitempty"`
	NrtReporting                   *NrtReporting                   `json:"nrtReporting,omitempty"`
	OriginPullLogs                 *OriginPullLogs                 `json:"originPullLogs,omitempty"`
	OriginPullLogsConfig           *OriginPullLogsConfig           `json:"originPullLogsConfig,omitempty"`
	ReceiptLogs                    *ReceiptLogs                    `json:"receiptLogs,omitempty"`
	ReceiptLogsConfig              *ReceiptLogsConfig              `json:"receiptLogsConfig,omitempty"`
	RequestReceipt                 []*RequestReceipt               `json:"requestReceipt,omitempty"`
	RequestReceiptReportPercentage *RequestReceiptReportPercentage `json:"requestReceiptReportPercentage,omitempty"`
	AwsSignedS3PostV4              []*AwsSignedS3PostV4            `json:"awsSignedS3PostV4,omitempty"`
	AuthACL                        []*AuthACL                      `json:"authAcl,omitempty"`
	AuthGeo                        []*AuthGeo                      `json:"authGeo,omitempty"`
	AuthHTTPBasic                  *AuthHTTPBasic                  `json:"authHttpBasic,omitempty"`
	AuthReferer                    *AuthReferer                    `json:"authReferer,omitempty"`
	AuthSignUrlsInPlaylist         *AuthSignUrlsInPlaylist         `json:"authSignURLsInPlaylist,omitempty"`
	AuthURLSign                    []*AuthURLSign                  `json:"authUrlSign,omitempty"`
	AuthURLSignAliCloudA           []*AuthURLSignAliCloudA         `json:"authUrlSignAliCloudA,omitempty"`
	AuthURLSignAliCloudB           []*AuthURLSignAliCloudB         `json:"authUrlSignAliCloudB,omitempty"`
	AuthURLSignAliCloudC           []*AuthURLSignAliCloudC         `json:"authUrlSignAliCloudC,omitempty"`
	AuthURLSignHmacTlu             []*AuthURLSignHmacTlu           `json:"authUrlSignHmacTlu,omitempty"`
	AuthURLSignIq                  []*AuthURLSignIq                `json:"authUrlSignIq,omitempty"`
	AuthURLAsymmetricSignTlu       []*AuthURLAsymmetricSignTlu     `json:"authUrlAsymmetricSignTlu,omitempty"`
	AuthURLSignL3                  []*AuthURLSignL3                `json:"authUrlSignL3,omitempty"`
	AuthURLSignAKv1                []*AuthURLSignAKv1              `json:"authUrlSignAKv1,omitempty"`
	AuthURLSignAKv2                []*AuthURLSignAKv2              `json:"authUrlSignAKv2,omitempty"`
	AuthURLSignLMV                 []*AuthURLSignLMV               `json:"authUrlSignLMV,omitempty"`
	AuthVhostLockout               *AuthVhostLockout               `json:"authVhostLockout,omitempty"`
	BandWidthLimit                 *BandWidthLimit                 `json:"bandWidthLimit,omitempty"`
	BandwidthRateLimit             *BandwidthRateLimit             `json:"bandwidthRateLimit,omitempty"`
	BandWidthRateLimitUnits        *BandWidthRateLimitUnits        `json:"bandWidthRateLimitUnits,omitempty"`
	ClientAccess                   *ClientAccess                   `json:"clientAccess,omitempty"`
	Compression                    *Compression                    `json:"compression,omitempty"`
	ContentDispositionByURL        *ContentDispositionByURL        `json:"contentDispositionByUrl,omitempty"`
	ContentDispositionByHeader     []*ContentDispositionByHeader   `json:"contentDispositionByHeader,omitempty"`
	CookieBehavior                 []*CookieBehavior               `json:"cookieBehavior,omitempty"`
	CrossDomain                    *CrossDomain                    `json:"crossDomain,omitempty"`
	CustomMimeType                 []*CustomMimeType               `json:"customMimeType,omitempty"`
	DNSIpv6                        *DNSIPv6                        `json:"dnsIpv6,omitempty"`
	DNSOverride                    []*DNSOverride                  `json:"dnsOverride,omitempty"`
	DynamicCacheRule               []*DynamicCacheRule             `json:"dynamicCacheRule,omitempty"`
	Flv                            *Flv                            `json:"flv,omitempty"`
	FlvPseudoStreaming             *FlvPseudoStreaming             `json:"flvPseudoStreaming,omitempty"`
	General                        *General                        `json:"general,omitempty"`
	HTTPMethods                    *HTTPMethods                    `json:"httpMethods,omitempty"`
	LegacyXdomain                  *LegacyXdomain                  `json:"legacyXdomain,omitempty"`
	LiveStreaming                  []*LiveStreaming                `json:"liveStreaming,omitempty"`
	PreserveRedirectHost           *PreserveRedirectHost           `json:"preserveRedirectHost,omitempty"`
	QueryStrParam                  *QueryStrParam                  `json:"queryStrParam,omitempty"`
	RedirectExceptions             *RedirectExceptions             `json:"redirectExceptions,omitempty"`
	RedirectMappings               []*RedirectMappings             `json:"redirectMappings,omitempty"`
	ResponseHeader                 *ResponseHeader                 `json:"responseHeader,omitempty"`
	RobotsTxt                      []*RobotsTxt                    `json:"robotsTxt,omitempty"`
	StaticHeader                   []*StaticHeader                 `json:"staticHeader,omitempty"`
	StreamChunkedEncodingResponse  *StreamChunkedEncodingResponse  `json:"streamChunkedEncodingResponse,omitempty"`
	TimePseudoStreaming            *TimePseudoStreaming            `json:"timePseudoStreaming,omitempty"`
	HTTP2Support                   *HTTP2Support                   `json:"http2Support,omitempty"`
	OcspParsing                    []*OcspParsing                  `json:"ocspParsing,omitempty"`
	Hostname                       []*Hostname                     `json:"hostname,omitempty"`
	BlockingOriginPullMode         *BlockingOriginPullMode         `json:"blockingOriginPullMode,omitempty"`
	CustomHeader                   *CustomHeader                   `json:"customHeader,omitempty"`
	DynamicOrigin                  *DynamicOrigin                  `json:"dynamicOrigin,omitempty"`
	FailSafeOriginPull             *FailSafeOriginPull             `json:"failSafeOriginPull,omitempty"`
	FarAheadRangeProxy             *FarAheadRangeProxy             `json:"farAheadRangeProxy,omitempty"`
	FileSegmentation               *FileSegmentation               `json:"fileSegmentation,omitempty"`
	VaryHeaderField                *VaryHeaderField                `json:"varyHeaderField,omitempty"`
	GzipOriginPull                 *GzipOriginPull                 `json:"gzipOriginPull,omitempty"`
	OriginPersistentConnections    *OriginPersistentConnections    `json:"originPersistentConnections,omitempty"`
	OriginPull                     *OriginPull                     `json:"originPull,omitempty"`
	OriginPullProtocol             *OriginPullProtocol             `json:"originPullProtocol,omitempty"`
	OriginPullPops                 []*OriginPullPops               `json:"originPullPops,omitempty"`
	OriginPullShield               *OriginPullShield               `json:"originPullShield,omitempty"`
	OriginPullHost                 *OriginPullHost                 `json:"originPullHost,omitempty"`
	OriginRoundRobinDNS            []*OriginRoundRobinDNS          `json:"originRoundRobinDns,omitempty"`
	AwsSignedOriginPullV4          []*AwsSignedOriginPullV4        `json:"awsSignedOriginPullV4,omitempty"`
	UploadLimit                    []*UploadLimit                  `json:"uploadLimit,omitempty"`
	Waf                            []*Waf                          `json:"waf,omitempty"`
	WafClustersOverride            []*WafClustersOverride          `json:"wafClustersOverride,omitempty"`
	XForwardedForBehavior          *XForwardedForBehavior          `json:"xForwardedForBehavior,omitempty"`
	WebSocket                      []*WebSocket                    `json:"webSocket,omitempty"`
	CacheControl                   []*CacheControl                 `json:"cacheControl,omitempty"`
	CacheKeyModification           *CacheKeyModification           `json:"cacheKeyModification,omitempty"`
	DynamicContent                 []*DynamicContent               `json:"dynamicContent,omitempty"`
	OriginPullCacheExtension       *OriginPullCacheExtension       `json:"originPullCacheExtension,omitempty"`
	OriginPullPolicy               []*OriginPullPolicy             `json:"originPullPolicy,omitempty"`
	ClientRequestQueue             *ClientRequestQueue             `json:"clientRequestQueue,omitempty"`
	ClientResponseQueue            *ClientResponseQueue            `json:"clientResponseQueue,omitempty"`
	ClientKeepAlive                []*ClientKeepAlive              `json:"clientKeepAlive,omitempty"`
	ConsistentHashing              []*ConsistentHashing            `json:"consistentHashing,omitempty"`
	H2proxyCaching                 *H2proxyCaching                 `json:"h2proxyCaching,omitempty"`
	Customer                       *Customer                       `json:"customer,omitempty"`
	DeviceBasedDynamicContent      *DeviceBasedDynamicContent      `json:"deviceBasedDynamicContent,omitempty"`
	HashType                       *HashType                       `json:"hashType,omitempty"`
	InternalError                  *InternalError                  `json:"internalError,omitempty"`
	LanguageRedirect               []*LanguageRedirect             `json:"languageRedirect,omitempty"`
	MidTierCaching                 []*MidTierCaching               `json:"midTierCaching,omitempty"`
	OriginRequestQueue             *OriginRequestQueue             `json:"originRequestQueue,omitempty"`
	OriginResponseQueue            *OriginResponseQueue            `json:"originResponseQueue,omitempty"`
	PathModification               []*PathModification             `json:"pathModification,omitempty"`
	ScriptNegCaching               *ScriptNegCaching               `json:"scriptNegCaching,omitempty"`
	ServerlessScripting            []*ServerlessScripting          `json:"serverlessScripting,omitempty"`
	TossbackBypass                 *TossbackBypass                 `json:"tossbackBypass,omitempty"`
	CloseHalfOpenConnections       *CloseHalfOpenConnections       `json:"closeHalfOpenConnections,omitempty"`
	TossbackAlways                 *TossbackAlways                 `json:"tossbackAlways,omitempty"`
	Rti                            []*Rti                          `json:"rti,omitempty"`
	ClientRequestModification      []*ClientRequestModification    `json:"clientRequestModification,omitempty"`
	ClientResponseModification     []*ClientResponseModification   `json:"clientResponseModification,omitempty"`
	OriginRequestModification      []*OriginRequestModification    `json:"originRequestModification,omitempty"`
	OriginResponseModification     []*OriginResponseModification   `json:"originResponseModification,omitempty"`
}

// // MarshalJSON fill default value
// func (c Configuration) MarshalJSON() ([]byte, error) {
// 	type t Configuration
// 	setDefault(c)
// 	return json.Marshal((t)(c))
// }

// UnmarshalJSON custom unmarshal configuration
func (c *Configuration) UnmarshalJSON(b []byte) error {
	type t Configuration
	conf := &t{}
	if err := json.Unmarshal(b, conf); err != nil {
		return err
	}
	setDefault(conf)
	// setDefaultField(conf, reflect.StructTag(""))
	*c = (Configuration)(*conf)
	return nil
}

// setDefault must provides non-reflect.Value paraments, due to field tag must parse-able
func setDefault(v interface{}, d ...interface{}) {
	var rc reflect.Value
	var rt reflect.Type
	// switch v.(type) {
	// case reflect.Value:
	// default:
	rc = reflect.ValueOf(v)
	rt = reflect.TypeOf(v)
	// }
	if rc.Kind() == reflect.Ptr {
		rc = rc.Elem()
		rt = rt.Elem()
	}
	var fieldsCount int = 0

	if rc.IsValid() {
		fieldsCount = rc.NumField()
	}
	for r := 0; r < fieldsCount; r++ {
		switch rc.Field(r).Kind() {
		case reflect.String:
			if b, f := rt.Field(r).Tag.Lookup("default"); f && rc.Field(r).IsZero() {
				rc.Field(r).SetString(b)
			}
		case reflect.Interface:
			// include nil value
			// fmt.Printf("%v", rc.Field(r).Elem())
		case reflect.Ptr:
			if b, f := rt.Field(r).Tag.Lookup("default"); f && rc.Field(r).Type().String() == "*bool" {
				if fv, fe := strconv.ParseBool(b); fe == nil && rc.Field(r).IsNil() {
					rc.Field(r).Set(reflect.ValueOf(&fv))
				}
			} else {
				setDefault(rc.Field(r).Interface())
			}
		case reflect.Struct:
			setDefault(rc.Field(r).Interface())
		case reflect.Float32:
			if b, f := rt.Field(r).Tag.Lookup("default"); f && rc.Field(r).IsZero() {
				if fv, fe := strconv.ParseFloat(b, 32); fe == nil {
					rc.Field(r).SetFloat(fv)
				}
			}
		case reflect.Float64:
			if b, f := rt.Field(r).Tag.Lookup("default"); f && rc.Field(r).IsZero() {
				if fv, fe := strconv.ParseFloat(b, 64); fe == nil {
					rc.Field(r).SetFloat(fv)
				}
			}
		case reflect.Uint32:
			if b, f := rt.Field(r).Tag.Lookup("default"); f && rc.Field(r).IsZero() {
				if fv, fe := strconv.ParseUint(b, 10, 32); fe == nil {
					rc.Field(r).SetUint(fv)
				}
			}
		case reflect.Uint16:
			if b, f := rt.Field(r).Tag.Lookup("default"); f && rc.Field(r).IsZero() {
				if fv, fe := strconv.ParseUint(b, 10, 16); fe == nil {
					rc.Field(r).SetUint(fv)
				}
			}
		case reflect.Int:
			if b, f := rt.Field(r).Tag.Lookup("default"); f && rc.Field(r).IsZero() {
				if fv, fe := strconv.ParseInt(b, 10, 0); fe == nil {
					rc.Field(r).SetInt(fv)
				}
			}
		case reflect.Int32:
			if b, f := rt.Field(r).Tag.Lookup("default"); f && rc.Field(r).IsZero() {
				if fv, fe := strconv.ParseInt(b, 10, 32); fe == nil {
					rc.Field(r).SetInt(fv)
				}
			}
		// ignore bool value, due to false could been re-write to true
		// case reflect.Bool:
		// 	if b, f := rt.Field(r).Tag.Lookup("default"); f && rc.Field(r).IsZero() {
		// 		if fv, fe := strconv.ParseBool(b); fe == nil {
		// 			rc.Field(r).SetBool(fv)
		// 		}
		// 	}
		case reflect.Array, reflect.Slice:
			for ari := 0; ari < rc.Field(r).Len(); ari++ {
				setDefault(rc.Field(r).Index(ari).Interface())
			}
		}
	}
}

// ConfigurationScope A uniquely addressable path on the CDN to which configuration can be written
type ConfigurationScope struct {
	ID       int    `json:"id"`       //The unique identifier for this ConfigurationScope
	Path     string `json:"path"`     //The path at which this configuration applies on the host
	Platform string `json:"platform"` //The platform to which this configuration applies
	Name     string `json:"name"`     //The friendly name for this Configuration Scope
}

// GetConfigurationGraph Get configuration graph
//Path /api/v1/accounts/{account_hash}/graph
//Get configuration graph
func (api *HWApi) GetConfigurationGraph(accountHash string) (*Graph, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/graph", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Graph{}
	return al, json.Unmarshal(r.body, al)
}

// GetHostNames List the hostnames that exist for an account
//Path /api/v1/accounts/{account_hash}/hostnames
//List the hostnames that exist for an account
func (api *HWApi) GetHostNames(accountHash string) (*ConfigurationHostNamesList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hostnames", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &ConfigurationHostNamesList{}
	return al, json.Unmarshal(r.body, al)
}

// CreateScope Create a new configuration scope for a given host
//Path /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/scopes
//Create a new configuration scope for a given host
func (api *HWApi) CreateScope(accountHash string, hostHash string, scope *Scope) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: POST,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hosts/%s/configuration/scopes", accountHash, hostHash),
			Body:   scope,
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}

// GetScopes List the scopes at which configuration exists for a given host
//Path /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/scopes
//List the scopes at which configuration exists for a given host
func (api *HWApi) GetScopes(accountHash string, hostHash string) (*ConfigScopeList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hosts/%s/configuration/scopes", accountHash, hostHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &ConfigScopeList{}
	return al, json.Unmarshal(r.body, al)
}

// DeleteScope Delete a configuration scope
//Path /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id}
//Delete a configuration scope
func (api *HWApi) DeleteScope(accountHash string, hostHash string, scopeID int) (bool, error) {
	if _, e := api.Request(
		&Request{
			Method: DELETE,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hosts/%s/configuration/%d", accountHash, hostHash, scopeID),
		},
	); e != nil {
		return false, e
	}
	return true, nil
}

// String Convert configuration to JSON
func (c *Configuration) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}

// GetConfiguration Get host configuration at a certain scope
//Path /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id}
//Get host configuration at a certain scope
func (api *HWApi) GetConfiguration(accountHash string, hostHash string, scopeID int) (*Configuration, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hosts/%s/configuration/%d", accountHash, hostHash, scopeID),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Configuration{}
	return al, json.Unmarshal(r.body, al)
}

// UpdateConfiguration Update host configuration at a certain scope
//Path /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id}
//Update host configuration at a certain scope
func (api *HWApi) UpdateConfiguration(accountHash string, hostHash string, scopeID int, configuration *Configuration) (*Configuration, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hosts/%s/configuration/%d", accountHash, hostHash, scopeID),
			Body:   configuration,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Configuration{}
	return al, json.Unmarshal(r.body, al)
}

func (c *ConfigStatus) String() (string, error) {
	s, e := json.Marshal(c)
	if e != nil {
		return "", e
	}
	return string(s), nil
}

// CheckConfigUpdateStatus Check on configuration update status
//Path /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id}/{configuration_receipt_id}
func (api *HWApi) CheckConfigUpdateStatus(accountHash string, hostHash string, scopeID int, configurationID string) (*ConfigStatus, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/accounts/%s/hosts/%s/configuration/%d/%s", accountHash, hostHash, scopeID, configurationID),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &ConfigStatus{}
	return al, json.Unmarshal(r.body, al)
}

// GetConfigurationDoc List the configuration types that this API supports
//Path /api/v1/configuration
func (api *HWApi) GetConfigurationDoc() (string, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    fmt.Sprintf("/api/v1/configuration"),
		},
	)
	if e != nil {
		return "", e
	}
	return string(r.body), nil
}
