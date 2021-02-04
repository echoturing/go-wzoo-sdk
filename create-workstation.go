package wzoo_sdk

import (
	"context"
	"encoding/json"
	"net/url"
)

// CreateWorkstationResponse ...
type CreateWorkstationResponse struct {
	Message string
	Code    json.Number
	Data    string
}

func (i *impl) CreateWorkstation(ctx context.Context, deviceID string) (*CreateWorkstationResponse, error) {
	path := "/open/workstation"
	data := url.Values{}
	data.Add(fieldDeviceID, deviceID)
	var resp CreateWorkstationResponse
	err := Do(ctx, i.httpClient, i.wzooEndpoint+path, &resp, data, WithAuthorization(i.defaultToken))
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
