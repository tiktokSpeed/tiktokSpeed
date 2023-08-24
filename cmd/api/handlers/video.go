package handlers

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/api/rpc"
	"github.com/tiktokSpeed/tiktokSpeed/pkg/jwt"
	"github.com/tiktokSpeed/tiktokSpeed/pkg/minio"
	"github.com/tiktokSpeed/tiktokSpeed/shared/consts"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/base"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/video"
)

// 视频流接口  /douyin/feed/
func GetUserFeed(ctx context.Context, c *app.RequestContext) {
	apiResp := new(api.DouyinFeedResponse)
	var err error
	var req api.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	//hlog.Info("调用视频流接口")
	videoResp, err := rpc.VideoClient.Feed(ctx, &video.DouyinFeedRequest{
		LatestTime: req.LatestTime,
		ViewerId:   817,
	})
	if err != nil {
		handleError(err, "Failed to get video", c)
		return
	}
	apiResp.VideoList = Videos(videoResp.VideoList)
	apiResp.StatusCode = int32(consts.CorrectCode)
	apiResp.StatusMsg = "调用视频流接口 成功"
	apiResp.NextTime = videoResp.NextTime
	if err != nil {
		hlog.Info("调用视频流接口 出错", err)
		apiResp.StatusCode = int32(consts.ErrCode)
		apiResp.StatusMsg = "调用视频流接口 成功"
		consts.SendResponse(c, apiResp)

	}

	consts.SendResponse(c, apiResp)

}

func Videos(videos []*base.Video) []*base.Video {
	vs := make([]*base.Video, 0)
	for _, video := range videos {
		if v := Video(video); v != nil {
			vs = append(vs, v)
		}
	}
	return vs
}

func Video(v *base.Video) *base.Video {
	if v == nil {
		return nil
	}
	return &base.Video{
		Id: v.Id,
		//User先置为空
		Author:        nil,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}
}

// /douyin/publish/action/
func PublishVideo(ctx context.Context, c *app.RequestContext) {
	hlog.Info("-----App calls Publish Video-----")
	token := c.PostForm("token")
	userClaims, err := jwt.ParseToken(token)
	if err != nil {
		handleError(err, "Failed to parse token", c)
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		handleError(err, "Failed to get video", c)
		return
	}
	body, err := data.Open()
	if err != nil {
		handleError(err, "Failed to open video", c)
		return
	}
	defer body.Close()
	fileKey, err := minio.Upload(ctx, userClaims.ID, &minio.File{
		Size: data.Size,
		Body: body,
		Name: data.Filename,
	})
	req := new(video.DouyinPublishActionRequest)
	req.UserId = userClaims.ID
	title := c.PostForm("title")
	req.Title = title
	req.PlayUrl = fileKey
	resp, err := rpc.VideoClient.PublishVideo(ctx, req)
	if err != nil {
		handleError(err, "Failed to publish video", c)
		return
	}
	consts.SendResponse(c, resp.BaseResp)
}

// /douyin/publish/list/
func GetPublishList(ctx context.Context, c *app.RequestContext) {
	hlog.Info("-----App calls Publish List-----")
	userID := c.Query("user_id")
	ownerID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		handleError(err, "Failed to parse user_id", c)
		return
	}
	req := new(video.DouyinGetPublishedListRequest)
	req.OwnerId = ownerID
	resp, err := rpc.VideoClient.GetPublishedVideoList(ctx, req)
	if err != nil {
		handleError(err, "Failed to get publish list", c)
		return
	}
	consts.SendResponse(c, resp)
}
