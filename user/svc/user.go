package svc

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"teamup/config"
	"teamup/db"
	"teamup/db/model"
	proto "teamup/user/proto"
	"teamup/user/util"
)


type UserServiceImpl struct {}

func (us *UserServiceImpl) Login(ctx context.Context, req *proto.LoginReq, rsp *proto.LoginRes) error {
	var wechatSession *util.WechatSession
	var err error

	// 获取wechatSession 区分测试环境和生产环境
	if config.Cfg.Get("env").String("") != "product" {
		wechatSession = &util.WechatSession{
			Openid:     req.Code,
			SessionKey: "-",
			Unionid:    "-",
			Errcode:    0,
			Errmsg:     "",
		}
	} else {
		wechatSession, err = util.Code2Session(req.Code)
		if err != nil {
			log.Println("err: ", err)
			return err
		}
	}

	if wechatSession.Errcode != 0 {
		rsp.Code = wechatSession.Errcode
		rsp.Msg = wechatSession.Errmsg
		rsp.Token = ""
		return nil
	}

	// 查询数据库，无此人需要创建
	user := &model.User{}
	if err := db.Conn.Where("openid = ?", wechatSession.Openid).First(user).Error; gorm.IsRecordNotFoundError(err) {
		user.Openid = wechatSession.Openid
		user.Avatar = "https://golang.org/lib/godoc/images/footer-gopher.jpg"
		user.Username = "wx_" + util.Hash(wechatSession.Openid)[:8]
		db.Conn.Save(user)
		db.Conn.Where("openid = ?", wechatSession.Openid).First(user)
	}

	rsp.Msg = "登录成功"
	rsp.Code = 0
	rsp.Token, err = util.GenerateToken(int32(user.UserID), config.Cfg.Get("auth", "secret").String(""))
	if err != nil {
		return err
	}
	return nil
}

func (us *UserServiceImpl) GetUserInfo(ctx context.Context, req *proto.GetUserInfoReq, rsp *proto.UserInfo) error {
	log.Println("getting user info of userid : ", req.UserId)
	user := &model.User{}
	db.Conn.Where("user_id = ?", req.UserId).First(user)
	if user.UserID == 0 {
		return errors.New("user not found")
	}
	rsp.Username = user.Username
	rsp.Avatar = user.Avatar
	return nil
}
