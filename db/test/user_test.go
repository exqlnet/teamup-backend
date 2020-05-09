package test

import (
	"teamup/db"
	"teamup/db/model"
	"testing"
)

func TestUserQuery(t *testing.T)  {
	user := &model.User{}
	db.Conn.Where("openid = ?", 1).First(user)
}

func TestPrimaryKeyReturn(t *testing.T)  {
	user := model.User{
		Openid:   "testopenidid",
		Username: "test_user",
		Avatar:   "-",
	}
	t.Log(db.Conn.Debug().NewRecord(&user))
	err := db.Conn.Debug().Save(&user).Error
	t.Log(db.Conn.Debug().NewRecord(&user))
	t.Log(err)
	t.Log("user id", user.UserID)
}

func TestUserCount(t *testing.T) {
	var count int
	db.Conn.Table("user").Count(&count)
	t.Log(count)
}