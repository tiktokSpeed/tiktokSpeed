package model

type Favorite struct {
	ID         int64 `gorm:"column:id;primary_key"`
	UserID     int64 `gorm:"column:user_id"`
	VideoID    int64 `gorm:"column:video_id"`
	ActionType int8  `gorm:"column:action_type"`
}

func (m *Favorite) TableName() string {
	return "favorite"
}
