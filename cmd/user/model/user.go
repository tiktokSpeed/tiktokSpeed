package model

type User struct {
	ID            int64  `gorm:"primarykey;autoIncrement"`
	Username      string `gorm:"type:varchar(32);not null;default:''"`
	Password      string `gorm:"type:varchar(32);not null;default:''"`
	FollowCount   int32  `gorm:"type:int(11);not null;default:0"`
	FollowerCount int32  `gorm:"type:int(11);not null;default:0"`
}
