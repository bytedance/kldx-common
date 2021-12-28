package conf

var (
	OpenapiEnvToUrl = map[string]string{
		"staging::extranet":        "https://oapi-kunlun-staging.bytedance.com",
		"gray::extranet":           "https://oapi-kunlun-gray.kundou.cn",
		"online::extranet":         "https://oapi-kunlun.kundou.cn",
	}
	FaasinfraEnvToUrl = map[string]string{
		"staging::extranet":        "https://apaas-faasinfra-staging.bytedance.com",
		"gray::extranet":           "https://apaas-faasinfra-gray.kundou.cn",
		"online::extranet":         "https://apaas-faasinfra.kundou.cn",
	}
)
