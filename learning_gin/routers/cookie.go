package routers

import (
	"github.com/gin-gonic/gin"
	"learning_gin/controllers/cookie"
	"learning_gin/middlewares"
)

func CookieRouter(router *gin.Engine) {
	routerGroup := router.Group("/cookie")
	{
		routerGroup.GET("/get", cookie.SetCookie)
		routerGroup.GET("/show", middlewares.Middle{}.AuthMiddleWare, cookie.ShowCookie)
	}
}
