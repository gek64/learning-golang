package main

import (
	"github.com/gin-gonic/gin"
	"learning-golang-gin/middlewares"
	"learning-golang-gin/routers"
	"log"
)

func main() {
	// 创建默认路由引擎
	router := gin.Default()

	// 加载路由中的全局中间件
	router.Use(middlewares.Middle{}.SetUploadLocation)

	// 加载路由
	routers.MainRouter(router)
	routers.MiddleRouter(router)
	routers.UploadRouter(router)

	// 在127.0.0.1:80上启动
	err := router.Run("127.0.0.1:80")
	if err != nil {
		log.Panicln(err)
	}

}
