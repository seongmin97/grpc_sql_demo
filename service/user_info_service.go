package service

import (
	"context"
	"errors"
	"exercise/gRPC_test/common"
	"exercise/gRPC_test/model"
	"log"
	"sync"
)

type UserInfoService struct {
	userInfoDao *model.UserInfoDao
}

var (
	userInfoServiceInstance *UserInfoService
	userInfoServiceOnce     sync.Once
)

func NewUserInfoService() *UserInfoService {
	userInfoServiceOnce.Do(func() {
		userInfoServiceInstance = &UserInfoService{
			userInfoDao: model.NewUserInfoDaoInstance(),
		}
	})
	return userInfoServiceInstance
}

// QueryUser query user info by name
func (s *UserInfoService) QueryUser(ctx context.Context, name string) (*model.UserInfo, error) {
	userInfo, err := s.userInfoDao.QueryUserByName(ctx, name)
	if err != nil {
		log.Printf("[UserInfoService.QueryUser] query database error: %v", err.Error())
		return nil, err
	}
	if userInfo == nil {
		log.Printf("[UserInfoService.QueryUser] record not found, name = %s", name)
		return nil, errors.New(common.RecordNotFound)
	}
	return userInfo, nil
}
