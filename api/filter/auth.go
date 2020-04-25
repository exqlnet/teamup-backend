package filter

import (
	"github.com/gin-gonic/gin"
	"teamup/config"
	"teamup/user/util"
)

func LoginRequired(c *gin.Context) {
	token := c.GetHeader("Authorization")[7:]
	secret := config.Cfg.Get("auth", "secret").String("")
	if secret == "" {
		c.AbortWithStatus(401)
	}

	userId, err := util.ParseToken(token, secret)
	if err != nil {
		c.AbortWithStatus(401)
	}

	c.Set("userId", userId)
}