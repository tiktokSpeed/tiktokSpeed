package main

import (
	"fmt"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/user/initialize"
	"github.com/tiktokSpeed/tiktokSpeed/conf"
	"github.com/tiktokSpeed/tiktokSpeed/pkg/utils"
	apiService "github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api/apiservice"
)

func main() {
	initialize.InitMySql()
	opts := kitexInit()
	svr := apiService.NewServer(new(UserServiceImpl), opts...)
	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	freePort, err := utils.GetFreePort()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", freePort))
	if err != nil {
		panic(err)
	}
	r, err := etcd.NewEtcdRegistry(conf.GetConf().Registry.RegistryAddress)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "user",
	}))
	opts = append(opts, server.WithRegistry(r))
	opts = append(opts, server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
	return
}
