package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USER     = "guest_user"
	PASSWORD = "guest_password"
	DBNAME   = "exam"
	HOST     = "127.0.0.1"
	PORT     = "13306"
)

func main() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/exam?charset=utf8mb4&parseTime=True&loc=Asia%%2FTokyo", USER, PASSWORD, HOST, PORT)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB接続に失敗しました。接続情報、MySQLコンテナの起動確認を行ってください。: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("DBオブジェクト取得失敗しました。: %v", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("DB接続のクローズ失敗: %v", err)
		}
	}()

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Ping失敗: %v", err)
	}

	fmt.Println("DB接続成功 DBNAME:", db.Migrator().CurrentDatabase())
}
