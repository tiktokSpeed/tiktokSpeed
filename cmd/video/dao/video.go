package dao

import (
	"github.com/tiktokSpeed/tiktokSpeed/cmd/video/initialize"
	"github.com/tiktokSpeed/tiktokSpeed/shared/consts"
)

// type Video struct {
// 	db *gorm.DB
// }

//后面的注释不影响映射
type VideoPo struct {
	ID            int64  `gorm:"primarykey"`
	FavoriteCount int64  `gorm:"column:favorite_count;`
	CommentCount  int64  `gorm:"column:comment_count;"`
	UserId        int64  `gorm:"column:user_id;"`
	PlayUrl       string `gorm:"not null; type: varchar(255)"`
	CoverUrl      string `gorm:"not null; type: varchar(255)"`
	Title         string `gorm:"not null; type: varchar(255)"`
}

// GetVideoListByLatestTime gets videos for feed.
func GetVideoListByLatestTime(latestTime int64) ([]*VideoPo, error) {
	//每次返回10个
	videos := make([]*VideoPo, 0)
	if err := initialize.DB.
		Table("video").
		Where("created_at <= ?", latestTime).
		Order("created_at desc").
		Limit(consts.VideosLimit).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}
