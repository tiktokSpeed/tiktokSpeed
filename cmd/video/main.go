package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/video/initialize"
	"github.com/tiktokSpeed/tiktokSpeed/conf"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/video/videoservice"
)

func main() {
	initialize.InitMySql()
	opts := kitexInit()
	svr := videoservice.NewServer(new(VideoServiceImpl), opts...)
	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))
	opts = append(opts, server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
	return
}
