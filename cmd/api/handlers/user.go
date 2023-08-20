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

// API:  /douyin/user/ [GET]
// GetUser implements getting user info
func GetUser(ctx context.Context, c *app.RequestContext) {
	hlog.Info("-----App calles Get User Info-----")
	
}

// API:  /douyin/user/login [POST]
// Login implements user login
func Login(ctx context.Context, c *app.RequestContext) {
	hlog.Info("-----App calles Login-----")
	apiResp := new(api.DouyinUserLoginResponse)

	var req api.DouyinUserLoginRequest
	if condition := c.BindAndValidate(&req); condition != nil {
		handleError(condition, "Request validation failed", c, apiResp)
		return
	}

	apiResp, err := rpc.UserClient.Login(context.Background(), &api.DouyinUserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		handleError(err, "Failed to login", c, apiResp)
		return
	}

	consts.SendResponse(c, apiResp)

}

func handleError(err error, message string, c *app.RequestContext, apiResp interface{}) {
	hlog.Info(message, err)
	resp, ok := apiResp.(interface {
		SetStatusCode(int32)
		SetStatusMsg(string)
	})
	if !ok {
		hlog.Info("apiResp is not a pointer")
		return
	}

	resp.SetStatusCode(consts.ErrCode)
	resp.SetStatusMsg(message)
	consts.SendResponse(c, resp)
}
