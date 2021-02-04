package wzoo_sdk

import (
	"context"
	"net/http"
)

// WZooSDK https://www.yuque.com/wechatpro/wxapi/intro
type WZooSDK interface {
	weChatAPI
	// https://www.yuque.com/wechatpro/wxapi/ycplgp
	Login(ctx context.Context, username, password string) (*LoginResponse, error)
	// https://www.yuque.com/wechatpro/wxapi/yfvfhx
	CreateWorkstation(ctx context.Context, deviceID string) (*CreateWorkstationResponse, error)
	// https://www.yuque.com/wechatpro/wxapi/ceb63h
	GetLoginQRCode(ctx context.Context, deviceID string) (*GetLoginQRCodeResponse, error)
	// https://www.yuque.com/wechatpro/wxapi/on17he
	CheckLoginShort(ctx context.Context, deviceID string) (*CheckLoginShortResponse, error)
	// https://www.yuque.com/wechatpro/wxapi/ig4ufh
	IsOnline(ctx context.Context, deviceID string) (*IsOnlineResponse, error)
	// https://www.yuque.com/wechatpro/wxapi/wciltn
	InitAddressList(ctx context.Context, deviceID string) (*InitAddressListResponse, error)
	// https://www.yuque.com/wechatpro/wxapi/grk3np
	GetAddressList(ctx context.Context, deviceID string) (*GetAddressListResponse, error)
	// https://www.yuque.com/wechatpro/wxapi/ppbpv1
	LoginAgainAsync(ctx context.Context, deviceID string) (*LoginAgainAsyncResponse, error)
}

// https://www.yuque.com/wechatpro/wxapi/goyfdx
type weChatAPI interface {
	SendText(ctx context.Context, deviceID, wxID, content string) (*SendTextResponse, error)
	SendImage(ctx context.Context, deviceID, wxID, imageURL string) (*SendImageResponse, error)
}

type impl struct {
	httpClient   *http.Client
	defaultToken string
	wzooEndpoint string
}

var _ WZooSDK = (*impl)(nil)

func NewWZooSDK(token string, endpoint string) WZooSDK {
	return &impl{
		httpClient:   http.DefaultClient,
		defaultToken: token,    // token 从login获取一次就行
		wzooEndpoint: endpoint, // wzoo api的地址
	}
}
