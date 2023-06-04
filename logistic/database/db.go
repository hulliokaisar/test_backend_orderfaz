package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/logistic")
	if err != nil {
		panic(err.Error())
	}
}

func GetDB() *sql.DB {
	return db
}
