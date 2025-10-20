package app

import (
	"database/sql"
	"learn/api/helpers"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang?parseTime=true")
	helpers.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(time.Minute * 1)
	db.SetConnMaxLifetime(time.Hour * 1)

	return db
}
