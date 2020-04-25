package svc

import (
	"github.com/micro/go-micro/v2/client"
	userPB "teamup/user/proto"
)

var UserServiceClient userPB.UserService

func init() {
	UserServiceClient = userPB.NewUserService("go.micro.teamup.svc.user", client.DefaultClient)
}

