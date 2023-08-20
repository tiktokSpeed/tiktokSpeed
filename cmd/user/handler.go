package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/user/dao"
	"github.com/tiktokSpeed/tiktokSpeed/pkg/jwt"
	"github.com/tiktokSpeed/tiktokSpeed/shared/consts"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api"
)

type UserServiceImpl struct {
}

// TODO: configurable secret key
const jwtSecretKey = "asecretkey"

// Feed implements the UserService interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *api.DouyinUserRegisterRequest) (resp *api.DouyinUserRegisterResponse, err error) {

	resp = new(api.DouyinUserRegisterResponse)

	user, err := dao.NewUser(req.Username, req.Password)
	if err != nil {
		return handleError(err, "Failed to create user", resp)
	}

	// generate token by Jason Web Token
	token, err := jwt.NewJWT([]byte(jwtSecretKey)).CreateToken(jwt.CustomClaims{ID: user.Id})

	if err != nil {
		handleError(err, "Failed to generate token", resp)
		return
	}

	// generate response
	resp.Token = token
	resp.UserId = user.Id
	resp.StatusCode = int32(consts.CorrectCode)
	resp.StatusMsg = "Register successfully"

	return resp, nil
}

func handleError(err error, message string, apiResp *api.DouyinUserRegisterResponse) (*api.DouyinUserRegisterResponse, error) {
	klog.Info(message, err)
	apiResp.StatusCode = int32(consts.ErrCode)
	apiResp.StatusMsg = message
	return apiResp, err
}

func (s *UserServiceImpl) Login(ctx context.Context, req *api.DouyinUserLoginRequest) (r *api.DouyinUserLoginResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *api.DouyinUserRequest) (r *api.DouyinUserResponse, err error) {
	// TODO: implement this method
	return nil, nil
}
