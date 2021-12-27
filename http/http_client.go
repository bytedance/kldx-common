package http

import (
	"bytes"
	"code.byted.org/apaas/goapi_common/constants"
	"code.byted.org/apaas/goapi_common/exceptions"
	"code.byted.org/apaas/goapi_common/structs"
	"code.byted.org/apaas/goapi_common/utils"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type HttpClient struct {
	http.Client
	Url string
}

var (
	openapiClientOnce sync.Once
	openapiClient     *HttpClient

	openapiClientConf = &structs.HttpConfig{
		Url:                 utils.GetOpenapiUrl(),
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     10 * time.Second,
	}
)

func GetOpenapiClient() *HttpClient {
	openapiClientOnce.Do(func() {
		openapiClient = &HttpClient{
			Client: http.Client{
				Transport: &http.Transport{
					MaxIdleConns:        openapiClientConf.MaxIdleConns,
					MaxIdleConnsPerHost: openapiClientConf.MaxIdleConnsPerHost,
					IdleConnTimeout:     openapiClientConf.IdleConnTimeout,
				},
			},
			Url: openapiClientConf.Url,
		}
	})
	return openapiClient
}

func (c *HttpClient) doRequest(req *http.Request, headers map[string][]string, mids []ReqMiddleWare) ([]byte, error) {
	for _, mid := range mids {
		err := mid(req)
		if err != nil {
			return nil, err
		}
	}

	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// timeout
	ctx, cancel := getTimeoutCtx()
	defer cancel()

	resp, err := c.Do(req.WithContext(ctx))
	if err != nil {
		return nil, exceptions.InternalError("doRequest failed, err: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, exceptions.InternalError("doRequest failed: statusCode is %d", resp.StatusCode)
	}

	// http resp body
	datas, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, exceptions.InternalError("doRequest failed, err: %v", err)
	}

	return datas, nil
}

func (c *HttpClient) Get(path string, headers map[string][]string, mids ...ReqMiddleWare) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.Url+path, nil)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.Get failed, err: %v", err)
	}

	return c.doRequest(req, headers, mids)
}

func (c *HttpClient) PostJson(path string, headers map[string][]string, data interface{}, mids ...ReqMiddleWare) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PostJson failed, err: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.Url+path, bytes.NewReader(body))
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PostJson failed, err: %v", err)
	}

	if headers == nil {
		headers = map[string][]string{}
	}
	headers[constants.HttpHeaderKey_ContentType] = []string{constants.HttpHeaderValue_Json}
	return c.doRequest(req, headers, mids)
}

func (c *HttpClient) PostFormData(path string, headers map[string][]string, body *bytes.Buffer, mids ...ReqMiddleWare) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, c.Url+path, body)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PostFormData failed, err: %v", err)
	}
	return c.doRequest(req, headers, mids)
}

func (c *HttpClient) PatchJson(path string, headers map[string][]string, data interface{}, mids ...ReqMiddleWare) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PatchJson failed, err: %v", err)
	}

	req, err := http.NewRequest(http.MethodPatch, c.Url+path, bytes.NewReader(body))
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PatchJson failed, err: %v", err)
	}

	if headers == nil {
		headers = map[string][]string{}
	}
	headers[constants.HttpHeaderKey_ContentType] = []string{constants.HttpHeaderValue_Json}
	return c.doRequest(req, headers, mids)
}

func (c *HttpClient) DeleteJson(path string, headers map[string][]string, data interface{}, mids ...ReqMiddleWare) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.DeleteJson failed, err: %v", err)
	}

	req, err := http.NewRequest(http.MethodDelete, c.Url+path, bytes.NewReader(body))
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.DeleteJson failed, err: %v", err)
	}

	if headers == nil {
		headers = map[string][]string{}
	}
	headers[constants.HttpHeaderKey_ContentType] = []string{constants.HttpHeaderValue_Json}
	return c.doRequest(req, headers, mids)
}

func getTimeoutCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
