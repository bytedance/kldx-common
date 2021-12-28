package structs

import "time"

type Tenant struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Type      int64  `json:"type"`
}

type HttpConfig struct {
	Url                 string
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}

type OpenapiTokenResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccessToken string `json:"accessToken"`
		ExpireTime  int64  `json:"expireTime"`
	} `json:"data"`
}
