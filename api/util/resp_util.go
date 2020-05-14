package util

import "github.com/gin-gonic/gin"

func InternalError(c *gin.Context)  {
	c.AbortWithStatusJSON(500, gin.H{
		"code": 500,
		"msg": "Internal Server Error",
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "success",
		"data": data,
	})
}