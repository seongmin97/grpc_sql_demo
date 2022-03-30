package model

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserInfoDao_CreateUserInfo(t *testing.T) {
	t.Skip()
	err := InitDatabase()
	if err != nil {
		t.Logf("init database error: %v\n", err)
		return
	}
	ctx := context.Background()
	userInfo := &UserInfo{
		ID:        0,
		Name:      "zq",
		Gender:    2,
		Hobby:     "sleep",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = NewUserInfoDaoInstance().CreateUserInfo(ctx, userInfo)
	assert.Nil(t, err)
}

func TestUserInfoDao_QueryUserByName(t *testing.T) {
	ctx := context.Background()
	err := InitDatabase()
	userInfo, err := NewUserInfoDaoInstance().QueryUserByName(ctx, "jsm")
	assert.Nil(t, err)
	t.Log(userInfo)
}
