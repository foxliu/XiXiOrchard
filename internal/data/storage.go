// Package data XiXiOrchard/internal/data/storage.go
package data

import (
	"XiXiOrchard/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() (*gorm.DB, error) {
	// PostgreSQL 连接字符串
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		config.Cfg.DBHost, config.Cfg.DBUser, config.Cfg.DBPassword, config.Cfg.DBName)

	// 使用 GORM 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL using GORM")

	// 自动迁移数据库模型
	err = db.AutoMigrate(&MarketData{}, &Order{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
