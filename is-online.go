package yuque_sdk

import (
	"context"
	"encoding/json"
	"net/url"
)

// IsOnlineResponse ...
type IsOnlineResponse struct {
	Message string
	Code    json.Number
}

func (i *impl) IsOnline(ctx context.Context, deviceID string) (*IsOnlineResponse, error) {
	data := url.Values{}
	data.Add("deviceId", deviceID)
	var resp IsOnlineResponse
	err := YuQueDo(ctx, i.httpClient, i.yuqueEndpoint+"/open/isOnline", &resp, data, WithAuthorization(i.defaultToken))
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
