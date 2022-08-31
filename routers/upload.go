package routers

import (
	"github.com/gin-gonic/gin"
	"learning-golang-gin/controllers/upload"
)

func UploadRouter(router *gin.Engine) {
	routerGroup := router.Group("/upload")
	{
		routerGroup.POST("/raw", upload.UpdateRawData)
		routerGroup.POST("/", upload.UpdateFiles)
	}
}
