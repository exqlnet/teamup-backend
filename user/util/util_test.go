package util

import (
	"log"
	"os"
	"teamup/config"
	"testing"
)

func TestJwt(t *testing.T)  {
	userIdList := []int32{1, 2, 3, 4,}
	secret := "ncuhomeok~"

	for _, userId := range userIdList {
		tokenString, err := GenerateToken(userId, secret)
		if err != nil {
			t.Fatalf("error: %s", err)
		}

		user, err := ParseToken(tokenString, secret)
		if err != nil {
			t.Fatalf("error: %s", err)
		}
		if user != userId {
			t.Fatalf("error: %d not equals to %d", user, userId)
		}
		t.Logf("Success: %d => %s", userId, tokenString)
	}
}

func TestCode2Session(t *testing.T) {
	appId := os.Getenv("APP_ID")
	appSecret := os.Getenv("APP_SECRET")
	config.Cfg.Set(appId, "wechat", "app_id")
	config.Cfg.Set(appSecret, "wechat", "app_secret")
	if appId == "" || appSecret == "" {
		log.Fatal("err: unable to read config from env")
	}
	ws, err := Code2Session("123456")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ws.Errcode, ws.Errmsg)
}