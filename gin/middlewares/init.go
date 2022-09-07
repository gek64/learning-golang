package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Middle 全局中间件
type Middle struct {
}

func (m Middle) GlobeMiddleware(c *gin.Context) {
	// 设置中间件与其他控制器共享的数据
	c.Set("dst", "upload")

	c.Next()
}
