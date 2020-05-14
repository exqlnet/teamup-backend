package main

import (
	"github.com/micro/go-micro/v2"
	"log"
	proto "teamup/user/proto"
	"teamup/user/svc"
)


const SvcName = "go.micro.teamup.svc.user"


func main() {

	service := micro.NewService(
		micro.Name(SvcName),
	)

	service.Init()
	proto.RegisterUserServiceHandler(service.Server(), &svc.UserServiceImpl{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
