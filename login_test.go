package wzoo_sdk

import (
	"context"
	"fmt"
	"testing"
)

func TestImpl_Login(t *testing.T) {
	sdk := NewWZooSDK(testToken, testEndpoint)
	ctx := context.Background()
	loginResponse, err := sdk.Login(ctx, testUsername, testPassword)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#v\n", loginResponse)
}
