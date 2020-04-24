package test

import (
	"teamup/user/util"
	"testing"
)

func TestJwt(t *testing.T)  {
	userIdList := []int32{1, 2, 3, 4,}
	secret := "ncuhomeok~"

	for _, userId := range userIdList {
		tokenString, err := util.GenerateToken(userId, secret)
		if err != nil {
			t.Fatalf("error: %s", err)
		}

		user, err := util.ParseToken(tokenString, secret)
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
	ws, err := util.Code2Session("123456")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ws.Errcode, ws.Errmsg)
}