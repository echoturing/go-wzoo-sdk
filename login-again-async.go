package yuque_sdk

import (
	"context"
	"encoding/json"
	"net/url"
)

// LoginAgainAsyncResponse ...
type LoginAgainAsyncResponse struct {
	Message string
	Code    json.Number
	Data    loginAgainAsyncResponseData
}

type loginAgainAsyncResponseData struct {
	WcID              string `json:"wcId"`
	WID               string `json:"wId"`
	WAccount          string `json:"wAccount"`
	Country           string `json:"country"`
	City              string `json:"city"`
	Signature         string `json:"signature"`
	NickName          string `json:"nickName"`
	Sex               int    `json:"sex"`
	HeadURL           string `json:"headUrl"`
	SmallHeadImageURL string `json:"smallHeadImageUrl"`
	Status            int    `json:"status"`
}

func (i *impl) LoginAgainAsync(ctx context.Context, deviceID string) (*LoginAgainAsyncResponse, error) {
	path := "/open/loginAgainAsync"
	data := url.Values{}
	data.Add(fieldDeviceID, deviceID)
	var resp LoginAgainAsyncResponse
	err := Do(ctx, i.httpClient, i.yuqueEndpoint+path, &resp, data, WithAuthorization(i.defaultToken))
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
