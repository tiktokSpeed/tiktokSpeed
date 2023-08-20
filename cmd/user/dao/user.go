package dao

import (
	"github.com/tiktokSpeed/tiktokSpeed/cmd/user/model"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/video/initialize"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/base"
	"golang.org/x/crypto/bcrypt"
)

// create a new user and save it into database
func NewUser(username string, password string) (*base.User, error) {
	m := initialize.DB
	if !m.Migrator().HasTable(&model.User{}) {
		if err := m.AutoMigrate(&model.User{}); err != nil {
			return nil, err
		}
	}

	// use bcrypt to hash password for security
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	usr := &model.User{
		Username: username,
		Password: string(hashedPW),
	}

	res := m.Create(&usr)
	if res.Error != nil {
		return nil, res.Error
	}

	return &base.User{
		Id:   int64(usr.ID),
		Name: username,
	}, nil
}
