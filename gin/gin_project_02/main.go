package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Article json 用类
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	// 创建默认路由引擎
	router := gin.Default()
	// 加载html模板文件
	router.LoadHTMLGlob("templates/*")

	// 返回响应类型
	// 返回响应-字符串
	router.GET("/string", func(c *gin.Context) {
		c.String(http.StatusOK, "返回响应-字符串")
	})

	// 返回响应-json
	// gin.H = map[string]any
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "返回响应-json",
		})
	})

	// 返回响应-json
	// 可以采用结构体
	router.GET("/json_struct", func(c *gin.Context) {
		article := Article{
			Title:   "标题",
			Desc:    "描述",
			Content: "内容",
		}

		c.JSON(http.StatusOK, article)
	})

	// 返回响应-jsonp
	// 用于跨域请求,http://localhost/json_p?callback=xxxx
	router.GET("/json_p", func(c *gin.Context) {
		article := Article{
			Title:   "标题",
			Desc:    "描述",
			Content: "内容",
		}

		c.JSONP(http.StatusOK, article)
	})

	// 返回响应-html
	router.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据",
		})
	})

	// 在127.0.0.1:80上启动
	err := router.Run("127.0.0.1:80")
	if err != nil {
		log.Panicln(err)
	}

}
