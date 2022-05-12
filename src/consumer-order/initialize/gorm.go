package initialize

import (
	"gimSec/basic/global"
	"gimSec/basic/logging"
	"gimSec/basic/utils"
	"gimSec/src/consumer-order/model"
)

func Gorm() {
	var err error
	global.DB, err = utils.GormMysqlConnection("124.221.132.236", 3306, "gimmick", "Gbdv470365234,.", "gimmick")
	if err != nil {
		logging.Error(err)
	}

	global.DB.AutoMigrate(&model.Order{})
}
