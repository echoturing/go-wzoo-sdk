package yuque_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// YuQueDo ...
func YuQueDo(ctx context.Context, client *http.Client, url string, resp interface{}, reqData url.Values, options ...Option) error {
	req, err := newYuQueSDKRequest(ctx, "POST", url, bytes.NewBufferString(reqData.Encode()), options...)
	if err != nil {
		return err
	}
	return do(ctx, client, req, resp)
}

// newYuQueSDKRequest ...
func newYuQueSDKRequest(ctx context.Context, method, url string, body io.Reader, options ...Option) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	for _, option := range options {
		option(req)
	}
	return req, nil
}

// do ...
func do(ctx context.Context, client *http.Client, req *http.Request, ret interface{}) error {
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("debug resp:", string(respBytes))
	err = json.Unmarshal(respBytes, ret)
	if err != nil {
		return err
	}
	return nil
}

// Option ...
type Option func(req *http.Request)

// WithAuthorization ...
func WithAuthorization(token string) Option {
	return func(req *http.Request) {
		req.Header.Add("Authorization", token)
	}
}
