package handler

import (
	"context"
	"errors"
	"exercise/gRPC_test/common"
	pb "exercise/gRPC_test/gen"
	"exercise/gRPC_test/service"
	"log"
	"sync"
)

type UserInfoHandler struct {
	userInfoService *service.UserInfoService
}

var (
	userInfoHandlerOnce     sync.Once
	userInfoHandlerInstance *UserInfoHandler
)

// NewUserInfoHandler create instance
func NewUserInfoHandler() *UserInfoHandler {
	userInfoHandlerOnce.Do(func() {
		userInfoHandlerInstance = &UserInfoHandler{
			userInfoService: service.NewUserInfoService(),
		}
	})
	return userInfoHandlerInstance
}

// QueryUserInfo query user information by name form database
func (h *UserInfoHandler) QueryUserInfo(ctx context.Context, req *pb.QueryUserInfoReq) (*pb.QueryUserInfoResp, error) {
	resp := &pb.QueryUserInfoResp{
		BaseResp: &pb.BaseResp{
			StatusMessage: common.OKStatusMessage,
			StatusCode:    common.OKStatusCode,
		},
	}

	// check parameter
	if len(req.GetName()) == 0 {
		log.Printf("[UserInfoHandler.QueryUserInfo] Parameter check error: %s", common.InputNameEmpty)
		resp.BaseResp.StatusMessage = common.InputNameEmpty
		resp.BaseResp.StatusCode = common.MissParaStatusCode
		return resp, errors.New(common.InputNameEmpty)
	}

	// query information
	userInfo, err := h.userInfoService.QueryUser(ctx, req.Name)
	if err != nil {
		log.Printf("[UserInfoHandler.QueryUserInfo] query user info error: %v", err.Error())
		resp.BaseResp.StatusMessage = err.Error()
		resp.BaseResp.StatusCode = common.ErrStatusCode
		return resp, nil
	}

	// construct resp
	resp.UserInfo = &pb.UserInfo{
		ID:     int64(userInfo.ID),
		Name:   userInfo.Name,
		Gender: int32(userInfo.Gender),
		Hobby:  userInfo.Hobby,
	}
	return resp, nil
}
