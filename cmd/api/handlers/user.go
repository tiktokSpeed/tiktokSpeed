package handlers

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/api/rpc"
	"github.com/tiktokSpeed/tiktokSpeed/shared/consts"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api"
	"golang.org/x/net/context"
)

// API:  /douyin/user/register
// Register implements creating a user
func Register(ctx context.Context, c *app.RequestContext) {
	hlog.Info("-----App calles Register-----")
	apiResp := new(api.DouyinUserRegisterResponse)

	var req api.DouyinUserRegisterRequest

	if err := c.BindAndValidate(&req); err != nil {
		handleError(err, "Request validation failed", c, apiResp)
		return
	}

	resp, err := rpc.UserClient.Register(ctx, &api.DouyinUserRegisterRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})

	if err != nil {
		handleError(err, "Failed to register", c, apiResp)
		return
	}

	consts.SendResponse(c, resp)
}

func handleError(err error, message string, c *app.RequestContext, apiResp *api.DouyinUserRegisterResponse) {
	hlog.Info(message, err)
	apiResp.StatusCode = int32(consts.ErrCode)
	apiResp.StatusMsg = message
	consts.SendResponse(c, apiResp)
}
