package mainPage

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// Msg 主页返回消息
type Msg struct {
	Message string    `json:"name" xml:"name"`
	Time    time.Time `json:"time" xml:"time"`
}

func Hello(c echo.Context) error {
	msg := &Msg{
		Message: "用户你好,欢迎来到主页",
		Time:    time.Now().UTC(),
	}

	return c.JSON(http.StatusOK, msg)
}
