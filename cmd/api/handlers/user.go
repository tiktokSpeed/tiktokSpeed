package handlers

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/tiktokSpeed/tiktokSpeed/pkg/jwt"
	"github.com/tiktokSpeed/tiktokSpeed/shared/consts"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/base"
	"golang.org/x/net/context"
)

// TODO: configurable secret key
const jwtSecretKey = "asecretkey"

// API:  /douyin/user/register
// Register implements creating a user
func Register(_ context.Context, c *app.RequestContext) {
	hlog.Info("-----App calles Register-----")
	apiResp := new(api.DouyinUserRegisterResponse)
	var err error
	var req api.DouyinUserRegisterRequest

	if err := c.BindAndValidate(&req); err != nil {
		handleError(err, "Request validation failed", c, apiResp)
		return
	}

	usr, err := createUser(req.GetUsername())
	if err != nil {
		handleError(err, "Failed to create user", c, apiResp)
		return
	}

	// generate token by Jason Web Token
	token, err := jwt.NewJWT([]byte(jwtSecretKey)).CreateToken(jwt.CustomClaims{ID: usr.Id})

	if err != nil {
		handleError(err, "Failed to generate token", c, apiResp)
		return
	}

	// generate response
	apiResp.Token = token
	apiResp.UserId = usr.Id
	apiResp.StatusCode = int32(consts.CorrectCode)
	apiResp.StatusMsg = "Register successfully"

	consts.SendResponse(c, apiResp)
}

func createUser(username string) (*base.User, error) {
	// TODO: check free node and create a user
	// Temporary solution to generate ID using timestamps
	usr := &base.User{
		Id:   time.Since(time.Now()).Nanoseconds() / 1000000,
		Name: username,
	}

	// TODO: Create a user in database
	return usr, nil
}

func handleError(err error, message string, c *app.RequestContext, apiResp *api.DouyinUserRegisterResponse) {
	hlog.Info(message, err)
	apiResp.StatusCode = int32(consts.ErrCode)
	apiResp.StatusMsg = message
	consts.SendResponse(c, apiResp)
}
