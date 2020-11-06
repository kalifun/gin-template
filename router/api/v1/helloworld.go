package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kalifun/gin-template/global"
	"net/http"
)

func Hello(c *gin.Context) {
	global.Log.Warning("log test")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Hello",
	})
}
