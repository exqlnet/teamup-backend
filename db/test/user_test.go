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