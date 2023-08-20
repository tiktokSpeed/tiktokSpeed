package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// TODO: configrate mysql address
const address = "root:Css187456@tcp(123.249.68.61:3306)/tiktok_speed?charset=utf8&parseTime=True&loc=Local"

// initialize mysql
func InitMySql() {
	var err error
	DB, err = gorm.Open(mysql.Open(address))
	if err != nil {
		klog.Info(err)
	}
}
