package yuque_sdk

import (
	"context"
	"fmt"
	"testing"
)

func TestImpl_SendImage(t *testing.T) {
	sdk := NewYuQueSDK(testToken, testEndpoint)
	ctx := context.Background()
	resp, err := sdk.SendImage(ctx, testDeviceID, "25378354803@chatroom", "https://t7.baidu.com/it/u=1962848802,1705699489&fm=193&f=GIF")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#v\n", resp)
}
