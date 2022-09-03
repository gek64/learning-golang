package cookie

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Cookies

func SetCookie(c *gin.Context) {
	cookie, err := c.Cookie("cookie_key")
	// 如果cookie不存在就设置cookie
	if err != nil {
		// name cookie名称
		// name cookie值
		// maxAge cookie有效时间,单位为秒
		// path cookie所在目录
		// domain 域名 可以使用 .example.com 这种来匹配多个二级域名
		// secure 是否仅允许通过https传递而来的cookie(避免中间人攻击)
		// httpOnly 是否允许通过js获取cookie(true 为不允许,避免跨站脚本攻击)
		c.SetCookie("cookie_key", "test_cookie_value", 86400, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"status": "cookie 已生成",
			"cookie": "test_cookie_value",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "cookie 已经存在",
			"cookie": cookie,
		})
	}
}

func ShowCookie(c *gin.Context) {
	cookie, err := c.Cookie("cookie_key")
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"status": "cookie 不存在",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "cookie 存在",
			"cookie": cookie,
		})
	}

}
