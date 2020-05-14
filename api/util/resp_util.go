package util

import "github.com/gin-gonic/gin"


type RspData struct {
	Code int	`json:"code"`
	Msg string	`json:"msg"`
	Data interface{}	`json:"data,omitempty"`
	Total int	`json:"total,omitempty"`
}
func InternalError(c *gin.Context)  {
	c.JSON(500, gin.H{
		"code": 500,
		"msg": "Internal Server Error",
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, RspData{
		Code:  0,
		Msg:   "success",
		Data:  data,
		Total: 0,
	})
}

func SuccessWithTotal(c *gin.Context, data interface{}, total int)  {
	c.JSON(200, RspData{
		Code:  0,
		Msg:   "success",
		Data:  data,
		Total: total,
	})
}

func BadRequest(c *gin.Context)  {
	c.JSON(400, gin.H{
		"code": 400,
		"msg": "bad request",
	})
}

func Unauthorize(c *gin.Context)  {
	c.JSON(401, gin.H{
		"code": 401,
		"msg": "unauthorized",
	})
}

func Failed(c *gin.Context, code int, msg string)  {
	// 在body中返回具体的错误码和错误信息
	c.JSON(200, gin.H{
		"code": code,
		"msg": msg,
	})
}