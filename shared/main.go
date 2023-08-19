package main

import (
	api "github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api/apiservice"
	"log"
)

func main() {
	svr := api.NewServer(new(ApiServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
