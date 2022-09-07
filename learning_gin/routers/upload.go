package routers

import (
	"github.com/gin-gonic/gin"
	"learning_gin/controllers/upload"
)

func UploadRouter(router *gin.Engine) {
	routerGroup := router.Group("/upload")
	{
		routerGroup.POST("/raw", upload.UpdateRawData)
		routerGroup.POST("/", upload.UpdateFiles)
	}
}
