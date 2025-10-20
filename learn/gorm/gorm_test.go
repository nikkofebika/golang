package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dialector := mysql.Open("root@tcp(localhost:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func TestConnection(t *testing.T) {
	db := OpenConnection()
	assert.NotNil(t, db)
}

type User struct {
	Id      int
	Name    string
	Email   string
	IsAdmin bool
}

func TestRawSql(t *testing.T) {
	db := OpenConnection()

	var user User

	db.Raw("SELECT id, name, email, is_admin FROM users limit 1").Scan(&user)
	fmt.Println(user)

	var users []User
	db.Raw("SELECT id, name, email, is_admin FROM users").Scan(&users)
	fmt.Println(users)
}

func TestExec(t *testing.T) {
	db := OpenConnection()

	err := db.Exec("INSERT INTO users (name, email, password, is_admin) VALUES (?, ?, ?, ?)", "Nikko", "nikko@gmail.com", "password", true).Error
	assert.Nil(t, err)
}

func TestScanRow(t *testing.T) {
	db := OpenConnection()

	var user User
	row := db.Raw("SELECT id, name, email, is_admin FROM users limit 1").Row()
	row.Scan(user)
	fmt.Println(user)

	var users []User
	rows, _ := db.Raw("SELECT id, name, email, is_admin FROM users").Rows()
	defer rows.Close()

	for rows.Next() {
		db.ScanRows(rows, &users)
	}
	fmt.Println(users)
}
