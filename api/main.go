package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
	"log"
	. "teamup/api/filter"
	"teamup/api/server"
)


func main() {

	service := web.NewService(
		web.Name("go.micro.teamup.api"),
		web.Address("0.0.0.0:8888"),
	)

	service.Init()
	router := gin.Default()

	userRouter := router.Group("/user")
	userRouter.GET("/info", server.UserInfoHandler, LoginRequired)
	userRouter.POST("/login", server.UserLoginHandler)

	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}