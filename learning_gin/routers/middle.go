package routers

import (
	"github.com/gin-gonic/gin"
	"learning_gin/controllers/middlePage"
)

func MiddleRouter(router *gin.Engine) {
	routerGroup := router.Group("/middle")
	{
		routerGroup.GET("/", middlePage.DefaultController{}.Timer, middlePage.DefaultController{}.Resp)
		routerGroup.POST("/", middlePage.DefaultController{}.Error, middlePage.DefaultController{}.Resp)
	}
}
