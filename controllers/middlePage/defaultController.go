package middlePage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type DefaultController struct {
}

func (d DefaultController) Resp(c *gin.Context) {
	// 在协程中使用的c需要复制一份
	cc := c.Copy()
	go func() {
		cc.Set("time", time.Now())
	}()

	// 获取中间件传递而来的用户id数据
	if userId, exist := c.Get("userId"); exist {
		c.JSON(http.StatusOK, gin.H{
			"userId":  userId,
			"message": "请求返回",
			"time":    time.Now(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"userId":  "not exist",
			"message": "请求返回",
			"time":    time.Now(),
		})
	}

}

// Timer 中间件,用于计算程序执行时间
func (d DefaultController) Timer(c *gin.Context) {
	t1 := time.Now()

	// 先执行后面的其他函数,执行完成再跳回到当前函数
	c.Next()

	t2 := time.Now()

	fmt.Printf("程序启动时间:%v,程序结束时间时间:%v\n", t1, t2)
}

// Error 中间件,终止后续函数的执行
func (d DefaultController) Error(c *gin.Context) {
	// 终止后续函数的执行,并执行下面的函数
	c.Abort()
	fmt.Printf("后续程序的运行已被终止\n")
}
