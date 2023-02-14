package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/golang-basic1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil

}
