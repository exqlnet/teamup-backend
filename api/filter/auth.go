package filter

import (
	"github.com/gin-gonic/gin"
	"teamup/config"
	"teamup/user/util"
)

func LoginRequired(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" || len(tokenString) < 8 {
		unAuth(c)
		return
	}
	token := c.GetHeader("Authorization")[7:]
	secret := config.Cfg.Get("auth", "secret").String("")
	if secret == "" {
		unAuth(c)
		return
	}

	userId, err := util.ParseToken(token, secret)
	if err != nil {
		unAuth(c)
		return
	}

	c.Set("userId", userId)
}


func unAuth(c *gin.Context) {
	c.JSON(401, gin.H{
		"code": 401,
		"msg": "unauthorized",
	})
	c.AbortWithStatus(401)
}