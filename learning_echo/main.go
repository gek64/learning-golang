package main

import (
	"github.com/labstack/echo/v4"
	"learning_echo/middlewares"
	"learning_echo/routers"
)

func main() {
	// 创建默认echo实例
	e := echo.New()

	e.Use(middlewares.Process, middlewares.ProcessNext, middlewares.ProcessNextNext)

	// 加载路由
	routers.MainRouter(e)
	routers.CookieRouter(e)

	// 在127.0.0.1:80上启动
	e.Logger.Panic(e.Start("localhost:80"))
}
