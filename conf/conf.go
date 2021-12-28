package conf

var (
	OpenapiEnvToUrl = map[string]string{
		"development:boe:intranet": "http://boe-apaas-oapi-dev.byted.org",
		"staging:boe:intranet":     "http://oapi-kunlun-staging-boe.byted.org",
		"staging::intranet":        "https://https://oapi-kunlun-staging.bytedance.net",
		"staging::extranet":        "https://oapi-kunlun-staging.bytedance.com",
		"gray::intranet":           "https://oapi-kunlun-gray.bytedance.net",
		"gray::extranet":           "https://oapi-kunlun-gray.kundou.cn",
		"online::intranet":         "https://oapi-kunlun.bytedance.net",
		"online::extranet":         "https://oapi-kunlun.kundou.cn",
	}
	FaasinfraEnvToUrl = map[string]string{
		"development:boe:intranet": "http://apaas-faasinfra-dev.byted.org",
		"staging:boe:intranet":     "http://apaas-faasinfra-staging-boe.bytedance.net",
		"staging::intranet":        "https://apaas-faasinfra-staging.bytedance.net",
		"staging::extranet":        "https://apaas-faasinfra-staging.bytedance.com",
		"gray::intranet":           "https://apaas-faasinfra-gray.bytedance.net",
		"gray::extranet":           "https://apaas-faasinfra-gray.kundou.cn",
		"online::intranet":         "https://apaas-faasinfra.bytedance.net",
		"online::extranet":         "https://apaas-faasinfra.kundou.cn",
	}
)
