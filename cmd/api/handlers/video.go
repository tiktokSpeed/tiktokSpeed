package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/api/rpc"
	"github.com/tiktokSpeed/tiktokSpeed/shared/consts"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/video"
)

// 视频流接口  /douyin/feed/
func GetUserFeed(ctx context.Context, c *app.RequestContext) {
	hlog.Info("hlog 日志")
	apiResp := new(api.DouyinFeedResponse)
	var err error
	var req api.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	hlog.Info("调用视频流接口")
	videoResp, err := rpc.VideoClient.Feed(ctx, &video.DouyinFeedRequest{
		LatestTime: req.LatestTime,
		ViewerId:   817,
	})
	if err != nil {
		hlog.Info("调用视频流接口 出错", err)

		//不需要
		apiResp.StatusCode = int32(consts.ErrCode)
		apiResp.StatusMsg = "调用视频流接口 出错"

		consts.SendResponse(c, videoResp)

	}

}
