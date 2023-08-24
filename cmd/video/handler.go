package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/video/dao"
	"github.com/tiktokSpeed/tiktokSpeed/shared/consts"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/base"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/video"
)

type VideoServiceImpl struct {
}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	//klog.Info("调用video服务端 Feed方法 ")
	resp = new(video.DouyinFeedResponse)
	if req.LatestTime <= 0 {
		req.LatestTime = time.Now().UnixNano()
	}
	//req.LatestTime=1692441970550877010
	//               1692337473000000000
	videos, err := dao.GetVideoListByLatestTime(req.LatestTime)
	nextTime := videos[len(videos)-1].CreatedAt
	resp.NextTime = nextTime.UnixNano()

	if err != nil {
		klog.Info("video服务端 Feed方法 错误")
		resp.BaseResp = &base.DouyinBaseResponse{
			StatusCode: int32(consts.ErrCode),
			StatusMsg:  "video服务端 Feed方法 错误",
		}
		return resp, err
	}
	resp.VideoList, err = fillVideoList(videos)
	resp.BaseResp = &base.DouyinBaseResponse{
		StatusCode: int32(consts.CorrectCode),
		StatusMsg:  "video服务端 Feed方法 成功",
	}

	return resp, nil
}

func fillVideoList(videoList []*dao.VideoPo) ([]*base.Video, error) {
	videolistVo := make([]*base.Video, len(videoList))
	ids := make([]int64, len(videoList))

	for i := 0; i < len(videoList); i++ {
		ids[i] = videoList[i].UserId
		//查找用户
		//users, err := dao.GetUserById(ids[i])
		// if err != nil {
		// 	klog.Info("video  服务端 查找用户详细信息 错误")
		// }
		//userVo := feed.User{Id: users.ID, Name: users.Name, FollowCount: &users.FollowerCount, FollowerCount: &users.FollowerCount}
		var flag bool
		if i%2 == 0 {
			flag = false
		} else {
			flag = true
		}
		videolistVo[i] = &base.Video{
			Id:            videoList[i].ID,
			Author:        &base.User{},
			PlayUrl:       videoList[i].PlayUrl,
			CoverUrl:      videoList[i].CoverUrl,
			FavoriteCount: videoList[i].FavoriteCount,
			CommentCount:  videoList[i].CommentCount,
			IsFavorite:    flag,
			Title:         videoList[i].Title,
		}
	}
	return videolistVo, nil
	//return []*base.Video{}, nil
}

// GetFavoriteVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFavoriteVideoList(ctx context.Context, req *video.DouyinGetFavoriteListRequest) (resp *video.DouyinGetFavoriteListResponse, err error) {
	return resp, err
}

// GetPublishedVideoIdList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishedVideoIdList(ctx context.Context, req *video.DouyinGetPublishedVideoIdListRequest) (resp *video.DouyinGetPublishedVideoIdListResponse, err error) {
	return resp, err
}

// GetPublishedVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishedVideoList(ctx context.Context, req *video.DouyinGetPublishedListRequest) (resp *video.DouyinGetPublishedListResponse, err error) {
	resp = new(video.DouyinGetPublishedListResponse)
	videos, err := dao.GetVideoListByUserID(req.OwnerId)
	if err != nil {
		resp.BaseResp = &base.DouyinBaseResponse{
			StatusCode: int32(consts.ErrCode),
			StatusMsg:  fmt.Sprintf("failed to get video list by user ID: %v", err),
		}
		return resp, err
	}
	resp.VideoList, err = fillVideoList(videos)
	resp.BaseResp = &base.DouyinBaseResponse{
		StatusCode: int32(consts.CorrectCode),
		StatusMsg:  "OK",
	}
	return resp, nil
}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	resp = new(video.DouyinPublishActionResponse)
	err = dao.SavePublishVideo(req)
	if err != nil {
		resp.BaseResp = &base.DouyinBaseResponse{
			StatusCode: int32(consts.ErrCode),
			StatusMsg:  fmt.Sprintf("failed to save new video %v", err),
		}
	} else {
		resp.BaseResp = &base.DouyinBaseResponse{
			StatusCode: int32(consts.CorrectCode),
			StatusMsg:  "OK",
		}
	}
	return resp, err
}
