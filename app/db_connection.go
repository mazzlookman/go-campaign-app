package app

import (
	"database/sql"
	"go-campaign-app/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

func DBConnection() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	return db
}

func DBConnectionTest() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bwastartup_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
