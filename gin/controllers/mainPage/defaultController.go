package mainPage

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct {
}

func (d DefaultController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get请求,从服务器取出资源(一项或多项)",
	})
}

func (d DefaultController) Post(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Post请求,在服务器新建一个资源",
	})
}

func (d DefaultController) Put(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Put请求,在服务器更新资源(客户端提供改变后的完整资源)",
	})
}

func (d DefaultController) Patch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Patch请求,在服务器更新资源(客户端提供改变的属性)",
	})
}

func (d DefaultController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete请求,从服务器删除资源",
	})
}

func (d DefaultController) Options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Options请求,获取信息,关于资源的哪些属性是客户端可以改变的",
	})
}

func (d DefaultController) Head(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Head请求,获取资源的元数据,无返回值",
	})
}
