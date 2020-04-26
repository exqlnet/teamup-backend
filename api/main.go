package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
	. "teamup/api/filter"
	"teamup/api/server"
	_ "teamup/api/swagger/statik"
)


func main() {

	service := web.NewService(
		web.Name("go.micro.teamup.api"),
		web.Address("0.0.0.0:8888"),
	)

	service.Init()
	router := gin.Default()

	userRouter := router.Group("/v1")
	userRouter.GET("/info", server.UserInfoHandler, LoginRequired)
	userRouter.POST("/login", server.UserLoginHandler)

	// swagger 接口文档
	staticFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	http.FileServer(staticFS)
	router.StaticFS("/docs", staticFS)
	service.Handle("/v1", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
