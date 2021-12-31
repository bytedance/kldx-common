package http

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/kldx-common/constants"
	"github.com/bytedance/kldx-common/exceptions"
	"github.com/bytedance/kldx-common/structs"
	"github.com/bytedance/kldx-common/utils"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var (
	appToken           atomic.Value
	appTokenExpireTime atomic.Value // 单位 s
	tokenRemainingTime int64 = 1200 // 单位 s，20min
)

type ReqMiddleWare func(req *http.Request) error

func AppTokenMiddleware(req *http.Request) error {
	if req == nil || req.Header == nil {
		return nil
	}
	token, err := GetAppToken()
	if err != nil {
		return err
	}
	req.Header.Add(constants.HttpHeaderKey_Authorization, token)
	return nil
}

func GetAppToken() (string, error) {
	// 1.get token from memory
	token := getAppTokenFromMem()
	if token != "" {
		return token, nil
	}

	// 2.get token from remote
	token, err := refreshAppToken()
	if err != nil {
		return "", err
	}
	return token, nil
}

func getAppTokenFromMem() string {
	expireTime, ok := appTokenExpireTime.Load().(int64)
	if !ok {
		return ""
	}

	token, ok := appToken.Load().(string)
	if !ok {
		return ""
	}

	// token 为空 或 20分钟内过期，不再使用
	if expireTime-time.Now().Unix() < tokenRemainingTime || token == "" {
		return ""
	}

	return token
}

func refreshAppToken() (string, error) {
	// 1.get lock
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	// 2.recheck
	token := getAppTokenFromMem()
	if token != "" {
		return token, nil
	}

	// 3.refresh token
	appid, secret, err := utils.GetAppidAndSecret()
	if err != nil {
		return "", fmt.Errorf("refresh app token err: %s", err.Error())
	}
	data := map[string]string{
		"clientId":     appid,
		"clientSecret": secret,
	}

	body, err := GetOpenapiClient().PostJson(OpenapiPath_GetToken, nil, data)
	if err != nil {
		return "", err
	}

	tokenResult := structs.OpenapiTokenResult{}
	err = json.Unmarshal(body, &tokenResult)
	if err != nil {
		return "", exceptions.InternalError("unmarshal OpenapiTokenResult failed, err: %v", err)
	}

	if tokenResult.Code != "0" {
		return "", exceptions.InternalError("refresh app token failed, code: %v, msg: %v", tokenResult.Code, tokenResult.Msg)
	}

	if tokenResult.Data.AccessToken == "" {
		return "", exceptions.InternalError("openapi accessToken is empty")
	}

	appToken.Store(tokenResult.Data.AccessToken)
	appTokenExpireTime.Store(tokenResult.Data.ExpireTime / 1000)
	return tokenResult.Data.AccessToken, nil
}
