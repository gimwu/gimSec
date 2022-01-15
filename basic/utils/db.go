package utils

import (
	"database/sql"
	"fmt"
)

var Database *sql.DB

func mysql(hostname string, port int, username string, password string, dbname string) (*sql.DB, error) {
	var err error
	Database, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			username, password, hostname, port, dbname))
	if err != nil {
		panic(err.Error())
	}
	return Database, nil
}
