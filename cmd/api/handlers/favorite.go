package handlers

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/api/rpc"
	"github.com/tiktokSpeed/tiktokSpeed/shared/consts"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api"
)

func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		handleError(err, "Invalid input", c)
		return
	}
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 8)
	if err != nil {
		handleError(err, "Invalid input", c)
		return
	}
	resp, err := rpc.UserClient.Favorite(ctx, &api.DouyinFavoriteActionRequest{
		Token:      c.Query("token"),
		VideoId:    videoId,
		ActionType: int8(actionType),
	})
	if err != nil {
		handleError(err, "Failed to favorite", c)
		return
	}
	consts.SendResponse(c, resp)
}
