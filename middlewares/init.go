package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Middle 全局中间件
type Middle struct {
}

func (m Middle) Hello(c *gin.Context) {
	fmt.Println("全局中间件开始")

	// 设置中间件与其他控制器共享的数据
	c.Set("userId", "123456")

	c.Next()
	fmt.Println("全局中间件结束")
}
