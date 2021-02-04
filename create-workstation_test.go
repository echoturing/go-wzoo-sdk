package yuque_sdk

import (
	"context"
	"fmt"
	"testing"
)

func TestImpl_CreateWorkstation(t *testing.T) {
	sdk := NewYuQueSDK(testToken, testEndpoint)
	ctx := context.Background()
	resp, err := sdk.CreateWorkstation(ctx, "test-device-id-shm")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#v\n", resp)
}
