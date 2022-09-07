package main

import (
	"github.com/gin-gonic/gin"
	"learning_gin/middlewares"
	"learning_gin/routers"
	"log"
)

func main() {
	// 创建默认路由引擎
	router := gin.Default()

	// 加载路由中的全局中间件
	router.Use(middlewares.Middle{}.GlobeMiddleware)

	// 加载路由
	routers.MainRouter(router)
	routers.MiddleRouter(router)
	routers.UploadRouter(router)
	routers.DataRouter(router)
	routers.CookieRouter(router)

	// 在127.0.0.1:80上启动
	err := router.Run("127.0.0.1:80")
	if err != nil {
		log.Panicln(err)
	}
}
