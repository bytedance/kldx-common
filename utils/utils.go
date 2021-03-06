package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bytedance/kldx-common/constants"
	"github.com/bytedance/kldx-common/structs"
	"os"
)

func GetServiceId() string {
	return os.Getenv(constants.EnvKSvcID)
}

func GetAppidAndSecret() (string, string, error) {
	tenantName := os.Getenv(constants.EnvMicroSvcTenantName)
	namespace := os.Getenv(constants.EnvMicroSvcNamespace)
	dClientID := os.Getenv(constants.EnvMicroSvcClientID)
	dClientSecret := os.Getenv(constants.EnvMicroSvcClientSecret)
	if tenantName == "" || namespace == "" || dClientID == "" || dClientSecret == "" {
		return "", "", errors.New("system params is error")
	}

	key := paddingN([]byte(tenantName+namespace), 32)
	clientID, err := AesDecryptText(0, key, dClientID)
	if err != nil {
		return "", "", fmt.Errorf("decrypt clientID err: %v", err)
	}
	clientSecret, err := AesDecryptText(0, key, dClientSecret)
	if err != nil {
		return "", "", fmt.Errorf("decrypt clientSecret err: %v", err)
	}
	return clientID, clientSecret, nil
}

func GetTenant() structs.Tenant {
	return structs.Tenant{
		Id:        0,
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

func GetEnv() string {
	return os.Getenv(constants.EnvKENV)
}

func Decode(input interface{}, output interface{}) error {
	bytes, err := json.Marshal(input)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &output); err != nil {
		return err
	}
	return nil
}
