package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SendfriendLog :
type SendfriendLog struct {
	LogId     uint16 `gorm:"column:log_id;primary_key" form:"log_id;primary_key" json:"log_id;primary_key"`
	Ip        []byte `gorm:"column:ip" form:"ip" json:"ip"`
	Time      uint16 `gorm:"column:time" form:"time" json:"time"`
	WebsiteId uint16 `gorm:"column:website_id" form:"website_id" json:"website_id"`
}

// TableName :
func (*SendfriendLog) TableName() string {
	return "sendfriend_log"
}

// handler create
func initRoutersSendfriendLog(r *gin.Engine, sendfriendlog string) {
	route := r.Group("sendfriendlog")
	route.GET("/", getSendfriendLogs)
	route.GET("/:id", getSendfriendLog)
	route.POST("/", createSendfriendLog)
	route.PUT("/:id", updateSendfriendLog)
	route.DELETE("/:id", deleteSendfriendLog)
}

func getSendfriendLogs(c *gin.Context) {
	var sendfriendLogs []SendfriendLog
	if err := g.Find(&sendfriendLogs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, sendfriendLogs)
	}
}

func getSendfriendLog(c *gin.Context) {
	id := c.Params.ByName("id")
	var sendfriendLog SendfriendLog
	if err := g.Where("id = ?", id).First(&sendfriendLog).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, sendfriendLog)
	}
}

func createSendfriendLog(c *gin.Context) {
	var sendfriendLog SendfriendLog
	c.BindJSON(&sendfriendLog)
	g.Create(&sendfriendLog)
	c.JSON(200, sendfriendLog)
}

func updateSendfriendLog(c *gin.Context) {
	var sendfriendLog SendfriendLog
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&sendfriendLog).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&sendfriendLog)
	g.Save(&sendfriendLog)
	c.JSON(200, sendfriendLog)
}
func deleteSendfriendLog(c *gin.Context) {
	id := c.Params.ByName("id")
	var sendfriendLog SendfriendLog
	d := g.Where("id = ?", id).Delete(&sendfriendLog)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
