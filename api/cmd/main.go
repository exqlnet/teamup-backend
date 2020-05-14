package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
	"teamup/api/filter"
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

	userRouter := router.Group("/api/v1/user")
	userRouter.GET("/info", filter.LoginRequired, server.UserInfoHandler)
	userRouter.POST("/login", server.UserLoginHandler)

	actRouter := router.Group("/api/v1/activity", filter.LoginRequired)
	actRouter.POST("", server.ActivityCreateHandler)
	actRouter.DELETE("", server.DeleteActivityHandler)
	actRouter.GET("", server.GetActivityDetailByIDHandler)
	actRouter.PUT("", server.UpdateActivityHandler)
	actRouter.GET("/hot", server.GetHotActivityListHandler)
	actRouter.GET("/joined", server.GetJoinedActivityHandler)
	actRouter.GET("/created", server.GetCreatedActivityHandler)
	actRouter.GET("/teams", server.GetTeamListByActivityIDHandler)
	actRouter.POST("/team", server.CreateTeamHandler)
	actRouter.GET("/team", server.GetTeamByIDHandler)


	// swagger 接口文档
	staticFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	http.FileServer(staticFS)
	router.StaticFS("/docs", staticFS)
	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
