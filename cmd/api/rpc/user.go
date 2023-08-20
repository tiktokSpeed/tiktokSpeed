package rpc

import (
	"log"

	"github.com/cloudwego/kitex/client"
	apiService "github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api/apiservice"
)

var UserClient apiService.Client

// Initialize the user client
func InitUser() {
	c, err := apiService.NewClient("User", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}

	log.Println(err)
	UserClient = c
	log.Println(UserClient)
}
