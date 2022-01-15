package utils

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Database  *sql.DB
	msyqlGorm *gorm.DB
)

func MysqlConnection(hostname string, port int, username string, password string, dbname string) (*sql.DB, error) {
	var err error
	Database, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			username, password, hostname, port, dbname))
	if err != nil {
		panic(err.Error())
	}
	return Database, nil
}

func gormMysqlConnection(hostname string, port int, username string, password string, dbname string) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, port, dbname)
	msyqlGorm, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		return nil, err
	}
	return msyqlGorm, nil
}
