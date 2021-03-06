package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"log"
	teamup_svc_user "teamup/user/proto"
)

func main() {

	service := micro.NewService()

	userService := teamup_svc_user.NewUserService("go.micro.teamup.svc.user", service.Client())
	u, err := userService.GetUserInfo(context.Background(), &teamup_svc_user.GetUserInfoReq{UserId: 1})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Found user: ", u)
}
