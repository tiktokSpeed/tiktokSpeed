package initialize

import (
	"context"

	server "github.com/cloudwego/kitex/server"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api"
	apiservice "github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api/apiservice"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler UserService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(apiservice.NewServiceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

type UserService interface {
	Register(ctx context.Context, req *api.DouyinUserRegisterRequest) (resp *api.DouyinUserRegisterResponse, err error)
	Login(ctx context.Context, req *api.DouyinUserLoginRequest) (r *api.DouyinUserLoginResponse, err error)
	GetUserInfo(ctx context.Context, req *api.DouyinUserRequest) (r *api.DouyinUserResponse, err error)
}
