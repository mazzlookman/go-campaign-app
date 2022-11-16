package app

import (
	"database/sql"
	"go-campaign-app/helper"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

func DBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bwastartup")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

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
