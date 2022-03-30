package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	ID        int            `gorm:"id" json:"ID"`
	Name      string         `gorm:"name" json:"Name"`
	Gender    int            `gorm:"gender" json:"Gender"`
	Hobby     string         `gorm:"hobby" json:"Hobby"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt" ` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt" ` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt" ` // 删除时间
}

type UserInfoDao struct{}

func (UserInfo) TableName() string {
	return "user_infos"
}

func NewUserInfoDaoInstance() *UserInfoDao {
	return &UserInfoDao{}
}

// CreateUserInfo create user info into database
func (d *UserInfoDao) CreateUserInfo(ctx context.Context, userInfo *UserInfo) error {
	return testDB.WithContext(ctx).Model(&UserInfo{}).Create(&userInfo).Error
}

// QueryUserByName query user info by name
func (d *UserInfoDao) QueryUserByName(ctx context.Context, name string) (*UserInfo, error) {
	users := make([]*UserInfo, 0)
	err := testDB.WithContext(ctx).Model(&UserInfo{}).
		Where("name = ?", name).Find(&users).Limit(1).Error
	if err != nil {
		return nil, err
	}
	if len(users) > 0 {
		return users[0], nil
	}
	return nil, nil
}
