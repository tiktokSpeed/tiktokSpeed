package dao

import (
	"context"
	"path"
	"time"

	"github.com/tiktokSpeed/tiktokSpeed/cmd/video/initialize"
	"github.com/tiktokSpeed/tiktokSpeed/pkg/minio"
	"github.com/tiktokSpeed/tiktokSpeed/shared/consts"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/video"
	"gorm.io/gorm"
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
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func getPlayAndCoverUrl(videoKey string) (string, string, error) {
	coverKey := path.Dir(videoKey) + "/cover.jpg"
	playURL, err := minio.PresignedURL(context.Background(), videoKey)
	if err != nil {
		return "", "", err
	}
	coverURL, err := minio.PresignedURL(context.Background(), coverKey)
	if err != nil {
		return "", "", err
	}
	return playURL, coverURL, nil
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
	for i := range videos {
		videoKey := videos[i].PlayUrl
		playURL, coverURL, err := getPlayAndCoverUrl(videoKey)
		if err != nil {
			return nil, err
		}
		videos[i].PlayUrl = playURL
		videos[i].CoverUrl = coverURL
	}
	return videos, nil
}

// SavePublishVideo saves video to db.
func SavePublishVideo(req *video.DouyinPublishActionRequest) error {
	videoPo := VideoPo{
		FavoriteCount: 0,
		CommentCount:  0,
		UserId:        req.UserId,
		PlayUrl:       req.PlayUrl,
		CoverUrl:      req.CoverUrl,
		Title:         req.Title,
		CreatedAt:     time.Now(),
	}
	return initialize.DB.Table("video").Create(&videoPo).Error
}

// GetVideoListByUserID gets videos by user id.
func GetVideoListByUserID(userID int64) ([]*VideoPo, error) {
	var videos []*VideoPo
	if err := initialize.DB.
		Table("video").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&videos).Error; err != nil {
		return nil, err
	}
	for i := range videos {
		videoKey := videos[i].PlayUrl
		playURL, coverURL, err := getPlayAndCoverUrl(videoKey)
		if err != nil {
			return nil, err
		}
		videos[i].PlayUrl = playURL
		videos[i].CoverUrl = coverURL
	}
	return videos, nil
}
