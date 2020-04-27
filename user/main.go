package main

import (
	"context"
	"github.com/emirpasic/gods/utils"
	"github.com/micro/go-micro/v2"
	"log"
	"os"
	"teamup/config"
	"teamup/db"
	"teamup/db/model"
	proto "teamup/user/proto"
	util "teamup/user/util"
)

type userServiceImpl struct {}

func (us *userServiceImpl) Login(ctx context.Context, req *proto.LoginReq, rsp *proto.LoginRes) error {
	wechatSession, err := util.Code2Session(req.Code)
	if err != nil {
		log.Println("err: ", err)
		return err
	}
	if wechatSession.Errcode != 0 {
		rsp.Code = wechatSession.Errcode
		rsp.Msg = wechatSession.Errmsg
		rsp.Token = ""
		return nil
	}

	// 查询数据库，无此人需要创建
	user := &model.User{}
	db.Conn.Where("open_id = ?", wechatSession.Openid).First(user)
	if user.UserID == 0 {
		user.Openid = wechatSession.Openid
		user.Avatar = "https://golang.org/lib/godoc/images/footer-gopher.jpg"
		user.Username = "wx_" + util.Hash(wechatSession.Openid)[:8]
	}


	rsp.Msg = "登录成功"
	rsp.Code = 0
	rsp.Token, err = util.GenerateToken(int32(user.UserID), config.Cfg.Get("auth", "secret").String(""))
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
	)

	appSecret := os.Getenv("APP_SECRET")
	appID := os.Getenv("APP_ID")
	if appSecret == "" || appID == "" {
		log.Fatal("APP_SECRET or APP_ID is not set")
	}
	service.Init()
	config.Cfg.Set(appID, "wechat", "app_id")
	config.Cfg.Set(appSecret, "wechat", "app_secret")

	proto.RegisterUserServiceHandler(service.Server(), &userServiceImpl{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
