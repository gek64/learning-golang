package routers

import (
	"github.com/labstack/echo/v4"
	"learning_echo/controllers/cookie"
)

func CookieRouter(e *echo.Echo) {
	routerGroup := e.Group("cookie")
	{
		routerGroup.GET("/set", cookie.SetCookie)
		routerGroup.GET("/show", cookie.ShowCookie)
	}
}
