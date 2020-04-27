package test

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	"log"
	proto "teamup/user/proto"
	"testing"
)

const name  = "go.micro.teamup.svc.user"

func TestGetUserInfo(t *testing.T)  {
	cli := proto.NewUserService(name, client.DefaultClient)
	_, err := cli.GetUserInfo(context.Background(), &proto.GetUserInfoReq{UserId:int32(1)})
	if err != nil {
		log.Println(err)
		t.Fatal(err)
	}

}

func TestLogin(t *testing.T) {
	cli := proto.NewUserService(name, client.DefaultClient)
	_, err := cli.Login(context.Background(), &proto.LoginReq{
		Code:                 "ncuhomelaile",
	})

	if err != nil {
		t.Fatal(err)
	}

}