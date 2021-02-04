package wzoo_sdk

import (
	"context"
	"fmt"
	"testing"
)

func TestImpl_CheckLoginShort(t *testing.T) {
	sdk := NewWZooSDK(testToken, testEndpoint)
	ctx := context.Background()
	resp, err := sdk.CheckLoginShort(ctx, "test-device-id-shm")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#v\n", resp)
}
