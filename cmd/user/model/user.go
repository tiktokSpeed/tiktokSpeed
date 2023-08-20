package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string `gorm:"type:varchar(32);not null;default:''"`
	Password      string `gorm:"type:varchar(32);not null;default:''"`
	FollowCount   int32  `gorm:"type:int(11);not null;default:0"`
	FollowerCount int32  `gorm:"type:int(11);not null;default:0"`
}
