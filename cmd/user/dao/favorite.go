package dao

import (
	"errors"

	"github.com/tiktokSpeed/tiktokSpeed/cmd/user/initialize"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/user/model"
	"gorm.io/gorm"
)

func FavoriteAction(userID, videoID int64, actionType int8) error {
	exist, err := QueryFavorite(userID, videoID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if exist != nil {
		if exist.ActionType == actionType {
			return nil
		}
		exist.ActionType = actionType
		return initialize.DB.Table("favorite").Save(&exist).Error
	}
	favorite := &model.Favorite{
		UserID:     userID,
		VideoID:    videoID,
		ActionType: actionType,
	}
	return initialize.DB.Table("favorite").Create(&favorite).Error
}

func QueryFavorite(userID, videoID int64) (*model.Favorite, error) {
	var favorite model.Favorite
	err := initialize.DB.
		Table("favorite").
		Where(&model.Favorite{UserID: userID, VideoID: videoID}).
		First(&favorite).Error
	if err != nil {
		return nil, err
	}
	return &favorite, nil
}
