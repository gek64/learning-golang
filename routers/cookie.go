package routers

import (
	"github.com/gin-gonic/gin"
	"learning-golang-gin/controllers/cookie"
	"learning-golang-gin/middlewares"
)

func CookieRouter(router *gin.Engine) {
	routerGroup := router.Group("/cookie")
	{
		routerGroup.GET("/get", cookie.SetCookie)
		routerGroup.GET("/show", middlewares.Middle{}.AuthMiddleWare, cookie.ShowCookie)
	}
}
