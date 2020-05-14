package main

import (
	"github.com/micro/go-micro/v2"
	"log"
	"os"
	"teamup/config"
	proto "teamup/user/proto"
	"teamup/user/svc"
)


const SvcName = "go.micro.teamup.svc.user"

func init() {
	if config.Cfg.Get("env").String("") != "product" {
		config.Cfg.Set("test_app_id", "wechat", "app_id")
		config.Cfg.Set("test_app_secret", "wechat", "app_secret")
	} else {
		appSecret := os.Getenv("APP_SECRET")
		appID := os.Getenv("APP_ID")
		if appSecret == "" || appID == "" {
			log.Fatal("APP_SECRET or APP_ID is not set")
		}
		config.Cfg.Set(appID, "wechat", "app_id")
		config.Cfg.Set(appSecret, "wechat", "app_secret")
	}
}

func main() {

	service := micro.NewService(
		micro.Name(SvcName),
		micro.Address("127.0.0.1:8999"),
	)

	service.Init()
	proto.RegisterUserServiceHandler(service.Server(), &svc.UserServiceImpl{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
