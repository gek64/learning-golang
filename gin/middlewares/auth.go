package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m Middle) AuthMiddleWare(c *gin.Context) {
	cookie, err := c.Cookie("cookie_key")
	// cookie 验证出错
	if err != nil {
		c.Abort()
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"status": "cookie 不存在,验证中间件停止后续运行",
		})
		return
	}

	// cookie 验证成功
	if cookie == "test_cookie_value" {
		c.Next()
		return
	}

	// cookie 验证失败
	c.Abort()
	c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
		"status": "cookie 不存在,验证中间件停止后续运行",
	})
	return

}
