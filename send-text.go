package yuque_sdk

import (
	"context"
	"encoding/json"
	"net/url"
)

type SendTextResponse struct {
	Message string
	Code    json.Number
}

func (i *impl) SendText(ctx context.Context, deviceID, wxID, content string) (*SendTextResponse, error) {
	path := "/open/sendText"
	data := url.Values{}
	data.Add(fieldDeviceID, deviceID)
	data.Add("wxId", wxID)
	data.Add("content", content)
	var resp SendTextResponse
	err := YuQueDo(ctx, i.httpClient, i.yuqueEndpoint+path, &resp, data, WithAuthorization(i.defaultToken))
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
