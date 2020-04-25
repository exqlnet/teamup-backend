package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/web"
	"log"
	. "teamup/api/filter"
	proto "teamup/user/proto"
)

var cl proto.UserService

func main() {

	service := web.NewService(
		web.Name("go.micro.teamup.api"),
	)

	service.Init()

	cl = proto.NewUserService("go.micro.teamup.svc.user", micro.NewService().Client())

	_, err := cl.GetUserInfo(context.TODO(), &proto.GetUserInfoReq{UserId: 1})
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	userRouter := router.Group("/user")
	userRouter.GET("/info", UserInfoHandler, LoginRequired)
	userRouter.POST("/login", UserLoginHandler)

	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}