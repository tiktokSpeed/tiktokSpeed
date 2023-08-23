package main

import (
	// "context"
	// "errors"
	// "fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/tiktokSpeed/tiktokSpeed/cmd/api/handlers"
	myRpc "github.com/tiktokSpeed/tiktokSpeed/cmd/api/rpc"
)

// 初始化 Hertz
func InitHertz() *server.Hertz {
	opts := []config.Option{server.WithMaxRequestBodySize(104857600),
		server.WithHostPorts("0.0.0.0:9999"),
	}
	h := server.Default(opts...)
	return h

}

// 注册 Router组
func registerGroup(h *server.Hertz) {
	douyin := h.Group("/douyin")
	feed := douyin.Group("/feed")
	feed.GET("/", handlers.GetUserFeed)

	user := douyin.Group("/user")
	user.GET("/", handlers.GetUser)
	{
		register := user.Group("/register")
		register.POST("/", handlers.Register)
	}
	{
		login := user.Group("/login")
		login.POST("/", handlers.Login)
	}
	publish := douyin.Group("/publish")
	publish.POST("/action/", handlers.PublishVideo)
}

// 运行API模块
func main() {

	myRpc.InitVideo()
	myRpc.InitUser()
	h := InitHertz()
	registerGroup(h)

	h.Spin()
}
