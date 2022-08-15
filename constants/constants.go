package constants

import "time"

const (
	BoeTag         = "boe"
	IntranetNetTag = "intranet"
	ExtranetNetTag = "extranet"
)

const (
	ReplaceNamespace     = ":namespace"
	ReplaceObjectApiName = ":objectApiName"
	ReplaceFieldApiName  = ":fieldApiName"
	ReplaceRecordId      = ":recordId"
	ReplaceFileId        = ":fileId"
)

const (
	HttpHeaderKey_Tenant        = "Tenant"
	HttpHeaderKey_User          = "User"
	HttpHeaderKey_ServiceID     = "X-Kunlun-Service-Id"
	HttpHeaderKey_Authorization = "Authorization"
	HttpHeaderKey_ContentType   = "Content-Type"
	HttpHeaderKey_Logid         = "X-Tt-Logid"
)

const (
	HttpHeaderValue_Json = "application/json"
	HttpHeaderValue_Bson = "application/bson"
)

const (
	EnvKENV                 = "ENV"
	EnvKSvcID               = "KSvcID"
	EnvMicroSvcTenantName   = "KTenantName"
	EnvMicroSvcNamespace    = "KNamespace"
	EnvMicroSvcClientID     = "KClientID"
	EnvMicroSvcClientSecret = "KClientSecret"
	EnvOpenApiDomain        = "KOpenApiDomain"
	EnvFaaSInfraDomain      = "KFaaSInfraDomain"
)

// key in ctx
const (
	ApiTimeoutMapKey = "KApiTimeoutMap"
	ApiTimeoutMethodKey = "KApiTimeoutMap"
)

const (
	ApiTimeoutDefault = 12 * time.Second
)
