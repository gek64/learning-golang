package routers

import (
	"github.com/gin-gonic/gin"
	"learning_gin/controllers/toStruct"
)

func DataRouter(router *gin.Engine) {
	routerGroup := router.Group("/data")
	{
		routerGroup.POST("/json", toStruct.JsonToStruct)
		routerGroup.POST("/form", toStruct.FormToStruct)
		routerGroup.POST("/:user/:password", toStruct.URLToStruct)
	}
}
