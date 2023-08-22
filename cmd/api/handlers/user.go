package handlers

import (
	"fmt"
	"strconv"

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

	userName, ok := c.GetQuery("username")
	password, ok := c.GetQuery("password")
	if !ok {
		handleError(fmt.Errorf("Invalid input"), "Invalid input", c, apiResp)
		return
	}

	resp, err := rpc.UserClient.Register(ctx, &api.DouyinUserRegisterRequest{
		Username: userName,
		Password: password,
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
	apiResp := new(api.DouyinUserResponse)

	var req api.DouyinUserRequest
	if condition := c.BindAndValidate(&req); condition != nil {
		handleError(condition, "Request validation failed", c, apiResp)
		return
	}

	userid, ok := c.GetQuery("user_id")
	token, ok := c.GetQuery("token")
	if !ok {
		handleError(fmt.Errorf("Invalid input"), "Invalid input", c, apiResp)
		return
	}

	id, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		handleError(fmt.Errorf("Invalid input"), "Invalid input", c, apiResp)
		return
	}

	apiResp, err = rpc.UserClient.GetUserInfo(context.Background(), &api.DouyinUserRequest{
		UserId: id,
		Token:  token,
	})

	if err != nil {
		handleError(err, "Failed to get user info", c, apiResp)
		return
	}

	consts.SendResponse(c, apiResp)

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

	username, ok := c.GetQuery("username")
	password, ok := c.GetQuery("password")
	if !ok {
		handleError(fmt.Errorf("Invalid input"), "Invalid input", c, apiResp)
		return
	}

	apiResp, err := rpc.UserClient.Login(context.Background(), &api.DouyinUserLoginRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		handleError(err, "Failed to login", c, apiResp)
		return
	}

	consts.SendResponse(c, apiResp)

}

func handleError(err error, message string, c *app.RequestContext, apiResp interface{}) {
	hlog.Info(message, err)
	consts.SendResponse(c, struct {
		StatusCode int    `json:"status_code"`
		StatusMsg  string `json:"status_msg"`
	}{consts.ErrCode, message})
}
