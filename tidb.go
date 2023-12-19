package main

import (
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TiDBMessage struct {
	ID            int    `json:"message_id" gorm:"primaryKey;column:message_id"`
	ToUser        int    `json:"to_user" gorm:"column:to_User"`
	FromUser      int    `json:"from_user" gorm:"column:from_User"`
	Title         string `json:"title" gorm:"column:title"`
	Message       string `json:"message" gorm:"column:message"`
	Image         string `json:"image" gorm:"column:image"`
	PhotoOrignURL string `json:"photo_origin_url" gorm:"column:photo_origin_url"`
	PhotoURL      string `json:"photo_url" gorm:"column:photo_url"`
}

func FindMessage(id int) (*TiDBMessage, error) {

	var m = TiDBMessage{}
	db, err := CreateDB()
	if err != nil {
		return nil, err
	}

	if err := db.Table("messages").Select("*").Where("message_id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	slog.Info(m.Title)
	// TiDB上からLIKE検索
	return &m, nil
}

func CreateDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(getDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDSN() string {
	tidbHost := getEnvWithDefault("TIDB_HOST", "127.0.0.1")
	tidbPort := getEnvWithDefault("TIDB_PORT", "4000")
	tidbUser := getEnvWithDefault("TIDB_USER", "root")
	tidbPassword := getEnvWithDefault("TIDB_PASSWORD", "")
	tidbDBName := getEnvWithDefault("TIDB_DB_NAME", "test")
	useSSL := getEnvWithDefault("USE_SSL", "true")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&tls=%s",
		tidbUser, tidbPassword, tidbHost, tidbPort, tidbDBName, useSSL)
}

func getEnvWithDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
