package yuque_sdk

import (
	"context"
	"encoding/json"
	"net/url"
)

type InitAddressListResponse struct {
	Message string
	Code    json.Number
}

func (i *impl) InitAddressList(ctx context.Context, deviceID string) (*InitAddressListResponse, error) {
	path := "/open/initAddressList"
	data := url.Values{}
	data.Add("deviceId", deviceID)
	var resp InitAddressListResponse
	err := Do(ctx, i.httpClient, i.yuqueEndpoint+path, &resp, data, WithAuthorization(i.defaultToken))
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
