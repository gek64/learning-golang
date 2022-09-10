package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("process")
		return next(c)
	}
}

func ProcessNext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("process Next")
		return next(c)
	}
}

func ProcessNextNext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("process Next Next")
		// 使用next(c)来启动下一个中间件
		return next(c)
	}
}

func Check(adminName string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.QueryParam("username") == adminName {
				return next(c)
			}
			return echo.NewHTTPError(http.StatusUnauthorized, "")
		}
	}
}
