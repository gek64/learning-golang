package main

import (
	"github.com/gin-gonic/gin"
	"learning-golang-gin/routers"
	"log"
)

func main() {
	// 创建默认路由引擎
	router := gin.Default()

	// 加载mainPage路由
	routers.MainPageRouter(router)

	// 在127.0.0.1:80上启动
	err := router.Run("127.0.0.1:80")
	if err != nil {
		log.Panicln(err)
	}

}
