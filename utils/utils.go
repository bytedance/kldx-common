package utils

import (
	"github.com/kldx/common/conf"
	"github.com/kldx/common/constants"
	"github.com/kldx/common/structs"
	"fmt"
)

func GetMicroserviceId() string {
	// TODO now mock，will get from env
	return "ujpa81"
}

func GetAppidAndSecret() (string, string) {
	// TODO now mock，will get from env
	return "c_fb79be28fae349ca90c0", "cd8fc6cb3c0a423d985e918e8019ec77"
}

func GetEnv() string {
	// TODO now mock，will get from env
	return "development"
}

func GetBoeTag() string {
	// TODO now mock，will get from env
	return constants.BoeTag
}

func GetInExtranetTag() string {
	// TODO now mock，will get from env
	return constants.IntranetNetTag
}

func GetTenant() structs.Tenant {
	// TODO now mock，will get from env
	return structs.Tenant{
		Id:        6187,
		Name:      "zwx_01",
		Namespace: "microService__c",
		Type:      1,
	}
}

func GetOpenapiUrl() string {
	key := fmt.Sprintf("%s:%s:%s", GetEnv(), GetBoeTag(), GetInExtranetTag())
	url, _ := conf.OpenapiEnvToUrl[key]
	return url
}

func GetFaasinfraUrl() string {
	key := fmt.Sprintf("%s:%s:%s", GetEnv(), GetBoeTag(), GetInExtranetTag())
	url, _ := conf.FaasinfraEnvToUrl[key]
	return url
}
