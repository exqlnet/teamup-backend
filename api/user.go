package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	teamup_svc_user "teamup/user/proto"
)

// Login User
// params:
// - code string: 微信授权码
// response:
// -
func UserLoginHandler(c *gin.Context) {
	params := make(map[string]string)
	log.Println("user login...")
	if err := c.BindJSON(&params); err != nil {
		log.Println(err)
		c.AbortWithStatus(400)
		return
	}
	
	res, err := cl.Login(context.Background(), &teamup_svc_user.LoginReq{Code: params["code"]})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, res)
}

// Get User Info
//
// params:
// -
func UserInfoHandler (c *gin.Context) {

}