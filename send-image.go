package wzoo_sdk

import (
	"context"
	"net/url"
)

type SendImageResponse struct {
}

func (i *impl) SendImage(ctx context.Context, deviceID, wxID, imageURL string) (*SendImageResponse, error) {
	path := "/open/sendImage"
	data := url.Values{}
	data.Add(fieldDeviceID, deviceID)
	data.Add("wxId", wxID)
	data.Add("content", imageURL)
	var resp SendImageResponse
	err := Do(ctx, i.httpClient, i.wzooEndpoint+path, &resp, data, WithAuthorization(i.defaultToken))
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
