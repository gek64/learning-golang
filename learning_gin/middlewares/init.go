package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GlobeMiddleware(c *gin.Context) {
	// 设置中间件与其他控制器共享的数据
	c.Set("dst", "upload")

	fmt.Println("start")

	c.Next()

	fmt.Println("done")
}
