package svc

import (
	"github.com/micro/go-micro/v2/client"
	userPB "teamup/user/proto"
	activityPB "teamup/activity/proto"
)

var UserServiceClient userPB.UserService
var ActivitySvc activityPB.ActivityService

func init() {
	UserServiceClient = userPB.NewUserService("go.micro.teamup.svc.user", client.DefaultClient)
	ActivitySvc = activityPB.NewActivityService("go.micro.teamup.svc.activity", client.DefaultClient)
}

