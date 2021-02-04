package yuque_sdk

import (
	"context"
	"encoding/json"
	"net/url"
)

// CheckLoginShortResponse ...
type CheckLoginShortResponse struct {
	Message string                      `json:"message"`
	Code    json.Number                 `json:"code"`
	Data    checkLoginShortResponseData `json:"data"`
}

type checkLoginShortResponseData struct {
	WcID              string `json:"wcId"`
	WAccount          string `json:"wAccount"`
	Country           string `json:"country"`
	City              string `json:"city"`
	Signature         string `json:"signature"`
	NickName          string `json:"nickName"`
	Sex               int    `json:"sex"`
	HeadURL           string `json:"headUrl"`
	SmallHeadImageURL string `json:"smallHeadImageUrl"`
	Status            int    `json:"status"`
	Type              int    `json:"type"`
}

func (i *impl) CheckLoginShort(ctx context.Context, deviceID string) (*CheckLoginShortResponse, error) {
	data := url.Values{}
	data.Add(fieldDeviceID, deviceID)
	var resp CheckLoginShortResponse
	err := YuQueDo(ctx, i.httpClient, i.yuqueEndpoint+"/open/checkLoginShort", &resp, data, WithAuthorization(i.defaultToken))
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
