package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"teamup/api/svc"
	userPB "teamup/user/proto"
)

// swagger:operation POST /user/login user login
// ---
// summary: 用微信授权码登陆
// parameters:
// - name: code
//   in: path
//   description: 微信授权码
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/userLoginResp"
//   "400":
//     "$ref": "#/responses/badReq"
func UserLoginHandler(c *gin.Context) {
	params := make(map[string]string)
	if err := c.BindJSON(&params); err != nil {
		log.Println(err)
		c.AbortWithStatus(400)
		return
	}

	res, err := svc.UserServiceClient.Login(context.Background(), &userPB.LoginReq{Code: params["code"]})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, res)
}

// swagger:operation GET /user/info user getUserInfo
// ---
// summary: 获取用户信息
// responses:
//   "200":
//     "$ref": "#/responses/userInfoResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "401":
//     "$ref": "#/responses/unauthorized"
func UserInfoHandler (c *gin.Context) {
	userId, _ := c.Get("userId")
	info, _ := svc.UserServiceClient.GetUserInfo(context.Background(), &userPB.GetUserInfoReq{UserId:userId.(int32)})

	c.JSON(200, info)
}