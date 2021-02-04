package yuque_sdk

import (
	"context"
	"fmt"
	"testing"
)

func TestImpl_SendText(t *testing.T) {
	sdk := NewYuQueSDK(testToken, testEndpoint)
	ctx := context.Background()
	resp, err := sdk.SendText(ctx, testDeviceID, "25378354803@chatroom", "🤩 world!")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#v\n", resp)
}
