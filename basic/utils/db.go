package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	GormMysqlDatabase *gorm.DB
)

func GormMysqlConnection(hostname string, port int, username string, password string, dbname string) (*gorm.DB, error) {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, port, dbname)
	GormMysqlDatabase, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return GormMysqlDatabase, nil
}
