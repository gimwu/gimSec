package utils

import (
	"fmt"
	"gimSec/basic/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	GormMysqlDatabase *gorm.DB
)

func init() {
	var err error
	GormMysqlDatabase, err = GormMysqlConnection("localhost", 3306, "root", "Gbdv470365234", "gimmick")
	if err != nil {
		logging.Error(err)
	}
}

func GormMysqlConnection(hostname string, port int, username string, password string, dbname string) (*gorm.DB, error) {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, port, dbname)
	GormMysqlDatabase, err = gorm.Open(mysql.Open(dns))
	if err != nil {
		return nil, err
	}
	return GormMysqlDatabase, nil
}
