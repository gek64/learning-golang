package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 创建默认路由引擎
	router := gin.Default()

	// 4中restful请求路由配置
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get请求,主要用于获取数据",
		})
	})

	router.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Post请求,主要用于创建、增加数据",
		})
	})

	router.PUT("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Put请求,主要用于更新、编辑数据",
		})
	})

	router.DELETE("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete请求,主要用于删除数据",
		})
	})

	// 在127.0.0.1:80上启动
	err := router.Run("127.0.0.1:80")
	if err != nil {
		log.Panicln(err)
	}

}
