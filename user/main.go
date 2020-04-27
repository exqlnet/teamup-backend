package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	"log"
	"os"
	"teamup/config"
	"teamup/db"
	"teamup/db/model"
	proto "teamup/user/proto"
	util "teamup/user/util"
)

const SVC_NAME  = "go.micro.teamup.svc.user"

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
	db.Conn.Where("openid = ?", wechatSession.Openid).First(user)
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
	log.Println("getting user info of userid : ", req.UserId)
	user := &model.User{}
	db.Conn.Where("user_id = ?", req.UserId).First(user)
	if user.UserID == 0 {
		return errors.New(SVC_NAME, "user not found", 404)
	}
	rsp.Username = user.Username
	rsp.Avatar = user.Avatar
	return nil
}

func main() {

	service := micro.NewService(
		micro.Name(SVC_NAME),
	)

	appSecret := os.Getenv("APP_SECRET")
	appID := os.Getenv("APP_ID")
	if appSecret == "" || appID == "" {
		log.Fatal("APP_SECRET or APP_ID is not set")
	}
	service.Init()
	config.Cfg.Set(appID, "wechat", "app_id")
	config.Cfg.Set(appSecret, "wechat", "app_secret")

	println("-----", config.Cfg.Get("env").String(""))
	println("-----", config.Cfg.Get("db", "host").String("host"))
	proto.RegisterUserServiceHandler(service.Server(), &userServiceImpl{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
