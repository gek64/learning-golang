package routers

import (
	"github.com/gin-gonic/gin"
	"learning-golang-gin/controllers/mainPage"
)

func MainPageRouter(router *gin.Engine) {
	routerGroup := router.Group("/")
	{
		routerGroup.GET("/", mainPage.DefaultController{}.Get)
		routerGroup.POST("/", mainPage.DefaultController{}.Post)
		routerGroup.PUT("/", mainPage.DefaultController{}.Put)
		routerGroup.PATCH("/", mainPage.DefaultController{}.Patch)
		routerGroup.DELETE("/", mainPage.DefaultController{}.Delete)
		routerGroup.OPTIONS("/", mainPage.DefaultController{}.Options)
		routerGroup.HEAD("/", mainPage.DefaultController{}.Head)
	}
}
