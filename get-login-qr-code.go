package wzoo_sdk

import (
	"context"
	"encoding/json"
	"net/url"
)

// GetLoginQRCodeResponse ...
type GetLoginQRCodeResponse struct {
	Message string
	Code    json.Number
	Data    getLoginQRCodeResponseData
}

type getLoginQRCodeResponseData struct {
	QRCodeURL string `json:"qrCodeUrl"`
}

func (i *impl) GetLoginQRCode(ctx context.Context, deviceID string) (*GetLoginQRCodeResponse, error) {
	data := url.Values{}
	data.Add("deviceId", deviceID)
	var resp GetLoginQRCodeResponse
	err := Do(ctx, i.httpClient, i.wzooEndpoint+"/open/getLoginQrCode", &resp, data, WithAuthorization(i.defaultToken))
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
