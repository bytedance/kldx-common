package utils

import (
	"github.com/bytedance/kldx-common/constants"
	"github.com/bytedance/kldx-common/structs"
	"os"
)

func GetMicroserviceId() string {
	return os.Getenv(constants.EnvMicroSvcID)
}

func GetAppidAndSecret() (string, string) {
	return os.Getenv(constants.EnvMicroSvcClientID), os.Getenv(constants.EnvMicroSvcClientSecret)
}

func GetTenant() structs.Tenant {
	return structs.Tenant{
		Id:        6187,
		Name:      os.Getenv(constants.EnvMicroSvcTenantName),
		Namespace: os.Getenv(constants.EnvMicroSvcNamespace),
		Type:      1,
	}
}

func GetOpenapiUrl() string {
	return os.Getenv(constants.EnvOpenApiDomain)
}

func GetFaasinfraUrl() string {
	return os.Getenv(constants.EnvFaaSInfraDomain)
}
