package routers

import (
	"github.com/labstack/echo/v4"
	"learning_echo/controllers/mainPage"
)

func MainRouter(e *echo.Echo) {
	routerGroup := e.Group("")
	{
		routerGroup.GET("/", mainPage.Hello)
	}
}
