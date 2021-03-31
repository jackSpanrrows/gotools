package main

import (
	"fmt"

	"github.com/jackSpanrrows/go-share-library/src/jlogger"

	"github.com/gin-gonic/gin"
)

const ServerPort = 8081

var ServerVersion = "v1.0.0"
var SadminUrl = ""

func main() {
	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong", "status": 200, "data": "[]"})
	})

	jlogger.Logger.LogInfo("日志测试", "aaaaa")

	r.Run(fmt.Sprintf(":%v", ServerPort))
}

// 载入配置文件
func LoadConfig() {

}
