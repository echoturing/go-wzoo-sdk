package yuque_sdk

import (
	"context"
	"encoding/json"
	"net/url"
)

// GetAddressListResponse ...
type GetAddressListResponse struct {
	Message string                     `json:"message"`
	Code    json.Number                `json:"code"`
	Data    getAddressListResponseData `json:"data"`
}

type getAddressListResponseData struct {
	ChatRooms []string `json:"chatRooms"`
	Friends   []string `json:"friends"`
	Ghs       []string `json:"ghs"`
	Others    []string `json:"others"`
}

func (i *impl) GetAddressList(ctx context.Context, deviceID string) (*GetAddressListResponse, error) {
	path := "/open/getAddressList"
	data := url.Values{}
	data.Add(fieldDeviceID, deviceID)
	var resp GetAddressListResponse
	err := Do(ctx, i.httpClient, i.yuqueEndpoint+path, &resp, data, WithAuthorization(i.defaultToken))
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
