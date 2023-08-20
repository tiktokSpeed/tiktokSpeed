package dao

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/user/initialize"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/user/model"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/base"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// create a new user and save it into database
func NewUser(username string, password string) (*base.User, error) {
	// use bcrypt to hash password for security
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if username == "" {
		return nil, gorm.ErrInvalidValue
	}
	err = initialize.DB.Table("user").Where(&model.User{Username: username}).First(&model.User{}).Error
	if err != gorm.ErrRecordNotFound || err == nil {
		klog.Info("Have record before")
		return nil, err
	}

	usr := &model.User{
		Username: username,
		Password: string(hashedPW),
	}

	res := initialize.DB.Table("user").Create(&usr)
	if res.Error != nil {
		return nil, res.Error
	}

	return &base.User{
		Id:   usr.ID,
		Name: username,
	}, nil
}
