package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/api/rpc"
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
		Id:            v.Id,
		//User先置为空
		Author:        nil,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v. FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}
}
