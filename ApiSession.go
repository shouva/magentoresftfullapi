package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ApiSession :
type ApiSession struct {
	UserId  uint16     `gorm:"column:user_id" form:"user_id" json:"user_id"`
	Logdate *time.Time `gorm:"column:logdate" form:"logdate" json:"logdate"`
	Sessid  string     `gorm:"column:sessid" form:"sessid" json:"sessid"`
}

// TableName :
func (*ApiSession) TableName() string {
	return "api_session"
}

// handler create
func initRoutersApiSession(r *gin.Engine, apisession string) {
	route := r.Group("apisession")
	route.GET("/", getApiSessions)
	route.GET("/:id", getApiSession)
	route.POST("/", createApiSession)
	route.PUT("/:id", updateApiSession)
	route.DELETE("/:id", deleteApiSession)
}

func getApiSessions(c *gin.Context) {
	var apiSessions []ApiSession
	if err := g.Find(&apiSessions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, apiSessions)
	}
}

func getApiSession(c *gin.Context) {
	id := c.Params.ByName("id")
	var apiSession ApiSession
	if err := g.Where("id = ?", id).First(&apiSession).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, apiSession)
	}
}

func createApiSession(c *gin.Context) {
	var apiSession ApiSession
	c.BindJSON(&apiSession)
	g.Create(&apiSession)
	c.JSON(200, apiSession)
}

func updateApiSession(c *gin.Context) {
	var apiSession ApiSession
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&apiSession).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&apiSession)
	g.Save(&apiSession)
	c.JSON(200, apiSession)
}
func deleteApiSession(c *gin.Context) {
	id := c.Params.ByName("id")
	var apiSession ApiSession
	d := g.Where("id = ?", id).Delete(&apiSession)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
