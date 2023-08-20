package config

import (
	"vincentcore_interview/pkg/helpers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(helpers.LoadEnv("DSN")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
