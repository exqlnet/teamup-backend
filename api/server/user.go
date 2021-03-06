package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"teamup/api/svc"
	userPB "teamup/user/proto"
)


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
//   "404":
//     "$ref": "#/responses/notFound"
func UserInfoHandler (c *gin.Context) {
	userId, _ := c.Get("userID")
	info, err := svc.UserServiceClient.GetUserInfo(context.Background(), &userPB.GetUserInfoReq{UserId:userId.(int32)})
	if err != nil {
		log.Printf("%v", err)
		c.JSON(401, gin.H{
			"code": 401,
			"msg": "unauthorized",
		})
		return
	}
	c.JSON(200, info)
}