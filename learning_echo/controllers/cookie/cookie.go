package cookie

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Resp struct {
	Status string      `json:"status" xml:"status"`
	Cookie http.Cookie `json:"cookie" xml:"cookie"`
}

func SetCookie(c echo.Context) error {
	cookie, err := c.Cookie("cookie_key")
	// 如果cookie不存在就设置cookie
	if err != nil {
		cookie = new(http.Cookie)
		cookie.Name = "cookie_key"
		cookie.Value = "test_cookie_value"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.Secure = false
		cookie.HttpOnly = true
		c.SetCookie(cookie)
	}

	err = ShowCookie(c)
	if err != nil {
		return err
	}

	return nil
}

func ShowCookie(c echo.Context) error {
	cookieCurrent, err := c.Cookie("cookie_key")
	if err != nil {
		err := c.JSON(http.StatusOK, Resp{
			Status: "cookie 不存在",
		})
		if err != nil {
			return err
		}
	} else {
		err := c.JSON(http.StatusOK, Resp{
			Status: "cookie 存在",
			Cookie: *cookieCurrent,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
