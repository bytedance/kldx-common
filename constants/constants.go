package constants

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
	HttpHeaderKey_Tenant         = "Tenant"
	HttpHeaderKey_User           = "User"
	HttpHeaderKey_MicroserviceId = "X-Kunlun-MicroService-Id"
	HttpHeaderKey_Authorization  = "Authorization"
	HttpHeaderKey_ContentType    = "Content-Type"
	HttpHeaderKey_Logid          = "X-Tt-Logid"
)

const (
	HttpHeaderValue_Json     = "application/json"
)

const (
	EnvMicroSvcID           = "KMicroSvcID"
	EnvMicroSvcTenantName   = "KMicroSvcTenantName"
	EnvMicroSvcNamespace    = "KMicroSvcNamespace"
	EnvMicroSvcClientID     = "KMicroSvcClientID"
	EnvMicroSvcClientSecret = "KMicroSvcClientSecret"
	EnvOpenApiDomain        = "KOpenApiDomain"
	EnvFaaSInfraDomain      = "KFaaSInfraDomain"
)