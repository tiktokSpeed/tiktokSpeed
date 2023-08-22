package rpc

import (
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/tiktokSpeed/tiktokSpeed/conf"
	apiService "github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api/apiservice"
)

var UserClient apiService.Client

// Initialize the user client
func InitUser() {
	r, err := etcd.NewEtcdResolver(conf.GetConf().Registry.RegistryAddress)
	if err != nil {
		log.Fatal(err)
	}
	c, err := apiService.NewClient("user", client.WithResolver(r))
	if err != nil {
		panic(err)
	}

	log.Println(err)
	UserClient = c
	log.Println(UserClient)
}
