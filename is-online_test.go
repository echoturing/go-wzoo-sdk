package wzoo_sdk

import (
	"context"
	"fmt"
	"testing"
)

func TestImpl_IsOnline(t *testing.T) {
	sdk := NewWZooSDK(testToken, testEndpoint)
	ctx := context.Background()
	resp, err := sdk.IsOnline(ctx, "test-device-id-shm")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#v\n", resp)
}
