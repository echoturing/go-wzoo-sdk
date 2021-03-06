package wzoo_sdk

import (
	"context"
	"fmt"
	"testing"
)

func TestImpl_GetAddressList(t *testing.T) {
	sdk := NewWZooSDK(testToken, testEndpoint)
	ctx := context.Background()
	resp, err := sdk.GetAddressList(ctx, testDeviceID)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#v\n", resp)
}
