package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var testDB *gorm.DB

func InitDatabase() error {
	var err error
	// password need to be changed
	dsn := "root:*****@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	testDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("connect data base error: %v", err)
	}
	return nil
}
