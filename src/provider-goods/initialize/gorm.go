package initialize

import (
	"gimSec/basic/global"
	"gimSec/basic/logging"
	"gimSec/basic/utils"
)

func Gorm() {
	var err error
	global.GOODS_DB, err = utils.GormMysqlConnection("124.221.132.236", 3306, "gimmick", "Gbdv470365234,.", "gimmick")
	if err != nil {
		logging.Error(err)
	}
}
