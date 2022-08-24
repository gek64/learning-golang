package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Article json 用类
type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	// 创建默认路由引擎
	router := gin.Default()
	// 加载多层文件夹html模板文件
	router.LoadHTMLGlob("templates/**/*")

	// 前台首页
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
		})
	})

	// 前台新闻
	router.GET("/news", func(c *gin.Context) {

		news := &Article{
			Title:   "标题",
			Content: "新闻详情",
		}

		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title":  "新闻",
			"news":   news,
			"number": 99,
		})
	})

	// 后台首页
	router.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "首页",
		})
	})

	// 后台新闻
	router.GET("/admin_news", func(c *gin.Context) {

		news := &Article{
			Title:   "标题",
			Content: "新闻详情",
		}

		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "新闻",
			"news":  news,
		})
	})

	// 在127.0.0.1:80上启动
	err := router.Run("127.0.0.1:80")
	if err != nil {
		log.Panicln(err)
	}

}
