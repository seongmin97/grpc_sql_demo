package handler

import (
	"context"
	pb "exercise/gRPC_test/gen"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserInfoHandler_QueryUserInfo(t *testing.T) {
	ctx := context.Background()
	req := &pb.QueryUserInfoReq{
		Name: "jsm",
	}
	resp, err := NewUserInfoHandler().QueryUserInfo(ctx, req)
	assert.Nil(t, err)
	t.Logf("%v", resp)
}
