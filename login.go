package wzoo_sdk

import (
	"context"
	"encoding/json"
	"net/url"
)

type loginResponseData struct {
	Authorization string  `json:"authorization"`
	Username      string  `json:"username"`
	Balance       float64 `json:"balance"`
	CallbackURL   string  `json:"callbackUrl"`
}

type LoginResponse struct {
	Message string
	Code    json.Number
	Data    loginResponseData
}

func (i *impl) Login(ctx context.Context, username, password string) (*LoginResponse, error) {
	data := url.Values{}
	data.Add("username", username)
	data.Add("password", password)
	var loginResponse LoginResponse
	err := Do(ctx, i.httpClient, i.wzooEndpoint+"/auth/login", &loginResponse, data)
	if err != nil {
		return nil, err
	}
	return &loginResponse, nil
}
