package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/bytedance/kldx-common/constants"
	"github.com/bytedance/kldx-common/exceptions"
	"github.com/bytedance/kldx-common/utils"
)

type HttpClient struct {
	http.Client
	Url string
}

var (
	openapiClientOnce sync.Once
	openapiClient     *HttpClient
)

func GetOpenapiClient() *HttpClient {
	openapiClientOnce.Do(func() {
		openapiClient = &HttpClient{
			Client: http.Client{
				Transport: &http.Transport{
					MaxIdleConns:        100,
					MaxIdleConnsPerHost: 10,
					IdleConnTimeout:     10 * time.Second,
				},
			},
			Url: utils.GetOpenapiUrl(),
		}
	})
	return openapiClient
}

func (c *HttpClient) doRequest(ctx context.Context, req *http.Request, headers map[string][]string, mids []ReqMiddleWare) ([]byte, error) {
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
	ctx, cancel := getTimeoutCtx(ctx)
	defer cancel()

	resp, err := c.Do(req.WithContext(ctx))
	if err != nil {
		return nil, exceptions.InternalError("doRequest failed, err: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	logid, _ := resp.Header[constants.HttpHeaderKey_Logid]
	fmt.Printf("logid: %s\n", logid)

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

func (c *HttpClient) Get(ctx context.Context, path string, headers map[string][]string, mids ...ReqMiddleWare) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.Url+path, nil)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.Get failed, err: %v", err)
	}

	return c.doRequest(ctx, req, headers, mids)
}

func (c *HttpClient) PostJson(ctx context.Context, path string, headers map[string][]string, data interface{}, mids ...ReqMiddleWare) ([]byte, error) {
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
	return c.doRequest(ctx, req, headers, mids)
}

func (c *HttpClient) PostBson(ctx context.Context, path string, headers map[string][]string, data interface{}, mids ...ReqMiddleWare) ([]byte, error) {
	body, err := bson.Marshal(data)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PostBson failed, err: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.Url+path, bytes.NewReader(body))
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PostBson failed, err: %v", err)
	}

	if headers == nil {
		headers = map[string][]string{}
	}
	headers[constants.HttpHeaderKey_ContentType] = []string{constants.HttpHeaderValue_Bson}
	return c.doRequest(ctx, req, headers, mids)
}

func (c *HttpClient) PostFormData(ctx context.Context, path string, headers map[string][]string, body *bytes.Buffer, mids ...ReqMiddleWare) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, c.Url+path, body)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PostFormData failed, err: %v", err)
	}
	return c.doRequest(ctx, req, headers, mids)
}

func (c *HttpClient) PatchJson(ctx context.Context, path string, headers map[string][]string, data interface{}, mids ...ReqMiddleWare) ([]byte, error) {
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
	return c.doRequest(ctx, req, headers, mids)
}

func (c *HttpClient) DeleteJson(ctx context.Context, path string, headers map[string][]string, data interface{}, mids ...ReqMiddleWare) ([]byte, error) {
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
	return c.doRequest(ctx, req, headers, mids)
}

func getTimeoutCtx(ctx context.Context) (context.Context, context.CancelFunc) {
	timeoutMap, ok1 := ctx.Value(constants.ApiTimeoutMapKey).(map[string]int64)
	method, ok2 := ctx.Value(constants.ApiTimeoutMethodKey).(string)
	if ok1 && ok2 {
		timeout, ok := timeoutMap[method]
		if ok {
			return context.WithTimeout(ctx, time.Duration(timeout)*time.Millisecond)
		}
	}
	return context.WithTimeout(ctx, constants.ApiTimeoutDefault)
}
