package toStruct

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Login 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段,若接收为空值,则报错,是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
	Info     string `form:"info" json:"info" uri:"info" xml:"info"`
}

// JsonToStruct 解析json
func JsonToStruct(c *gin.Context) {
	var login Login

	// body中的json转换到结构体
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Panicln(err)
	}

	// 返回输入
	c.JSON(http.StatusOK, gin.H{
		"user":     login.User,
		"password": login.Password,
		"info":     login.Info,
	})
}

// FormToStruct 解析multipart form
func FormToStruct(c *gin.Context) {
	var login Login

	// body中的json转换到结构体
	err := c.Bind(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Panicln(err)
	}

	// 返回输入
	c.JSON(http.StatusOK, gin.H{
		"user":     login.User,
		"password": login.Password,
		"info":     login.Info,
	})
}

// URLToStruct 解析URL
func URLToStruct(c *gin.Context) {
	var login Login

	// body中的json转换到结构体
	err := c.ShouldBindUri(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Panicln(err)
	}

	// 返回输入
	c.JSON(http.StatusOK, gin.H{
		"user":     login.User,
		"password": login.Password,
	})
}
