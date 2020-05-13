package svc

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	"teamup/user/proto"
	"testing"
)

const name  = "go.micro.teamup.svc.user"

func TestGetUserInfo(t *testing.T)  {
	handler := &UserServiceImpl{}
	resp := &teamup_svc_user.UserInfo{}
	err := handler.GetUserInfo(context.Background(), &teamup_svc_user.GetUserInfoReq{UserId: int32(1)}, resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", resp)
}

func TestLogin(t *testing.T) {
	cli := teamup_svc_user.NewUserService(name, client.DefaultClient)
	_, err := cli.Login(context.Background(), &teamup_svc_user.LoginReq{
		Code:                 "ncuhomelaile",
	})

	if err != nil {
		t.Fatal(err)
	}

}