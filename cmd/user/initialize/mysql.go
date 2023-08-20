package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// initialize mysql
func InitMySql() {
	var err error
	// TODO: configrate mysql address
	DB, err = gorm.Open(mysql.Open("root:Css187456@tcp(123.249.68.61:3306)/tiktok_speed?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		klog.Info(err)
	}
}
