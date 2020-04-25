package main

import (
	"context"
	"github.com/emirpasic/gods/utils"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"log"
	"teamup/config"
	proto "teamup/user/proto"
	util "teamup/user/util"
)

type userServiceImpl struct {}

func (us *userServiceImpl) Login(ctx context.Context, req *proto.LoginReq, rsp *proto.LoginRes) error {
	wechatSession, err := util.Code2Session(req.Code)
	if err != nil {
		return err
	}
	if wechatSession.Errcode != 0 {
		rsp.Code = wechatSession.Errcode
		rsp.Msg = wechatSession.Errmsg
		rsp.Token = ""
		return nil
	}

	// 查询数据库，无此人需要创建
	userId := int32(1)


	rsp.Msg = "登录成功"
	rsp.Code = 0
	rsp.Token, err = util.GenerateToken(userId, config.Cfg.Get("auth", "secret").String(""))
	if err != nil {
		return err
	}
	return nil
}

func (us *userServiceImpl) GetUserInfo(ctx context.Context, req *proto.GetUserInfoReq, rsp *proto.UserInfo) error {
	rsp.Username = "user" + utils.ToString(req.UserId)
	rsp.Avatar = "https://micro.mu/logo.png"
	return nil
}

func main() {

	service := micro.NewService(
		micro.Name("go.micro.teamup.svc.user"),
		micro.Flags(
			&cli.StringFlag{
				Name:  "app-secret",
				Required: true,
			},
			&cli.StringFlag{
				Name: "app-id",
				Required: true,
			},
		),
	)

	service.Init(micro.Action(func(c *cli.Context) error {
		config.Init()
		config.Cfg.Set(c.String("app-secret"), "wechat", "app_secret")
		config.Cfg.Set(c.String("app-id"), "wechat", "app_id")
		return nil
	}))

	config.Cfg.Set("", "app_secret")

	proto.RegisterUserServiceHandler(service.Server(), &userServiceImpl{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
