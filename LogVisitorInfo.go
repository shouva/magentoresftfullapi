package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// LogVisitorInfo :
type LogVisitorInfo struct {
	VisitorId          uint64 `gorm:"column:visitor_id;primary_key" form:"visitor_id;primary_key" json:"visitor_id;primary_key"`
	HttpReferer        string `gorm:"column:http_referer" form:"http_referer" json:"http_referer"`
	HttpUserAgent      string `gorm:"column:http_user_agent" form:"http_user_agent" json:"http_user_agent"`
	HttpAcceptCharset  string `gorm:"column:http_accept_charset" form:"http_accept_charset" json:"http_accept_charset"`
	HttpAcceptLanguage string `gorm:"column:http_accept_language" form:"http_accept_language" json:"http_accept_language"`
	ServerAddr         []byte `gorm:"column:server_addr" form:"server_addr" json:"server_addr"`
	RemoteAddr         []byte `gorm:"column:remote_addr" form:"remote_addr" json:"remote_addr"`
}

// TableName :
func (*LogVisitorInfo) TableName() string {
	return "log_visitor_info"
}

// handler create
func initRoutersLogVisitorInfo(r *gin.Engine, logvisitorinfo string) {
	route := r.Group("logvisitorinfo")
	route.GET("/", getLogVisitorInfos)
	route.GET("/:id", getLogVisitorInfo)
	route.POST("/", createLogVisitorInfo)
	route.PUT("/:id", updateLogVisitorInfo)
	route.DELETE("/:id", deleteLogVisitorInfo)
}

func getLogVisitorInfos(c *gin.Context) {
	var logVisitorInfos []LogVisitorInfo
	if err := g.Find(&logVisitorInfos).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logVisitorInfos)
	}
}

func getLogVisitorInfo(c *gin.Context) {
	id := c.Params.ByName("id")
	var logVisitorInfo LogVisitorInfo
	if err := g.Where("id = ?", id).First(&logVisitorInfo).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logVisitorInfo)
	}
}

func createLogVisitorInfo(c *gin.Context) {
	var logVisitorInfo LogVisitorInfo
	c.BindJSON(&logVisitorInfo)
	g.Create(&logVisitorInfo)
	c.JSON(200, logVisitorInfo)
}

func updateLogVisitorInfo(c *gin.Context) {
	var logVisitorInfo LogVisitorInfo
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&logVisitorInfo).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&logVisitorInfo)
	g.Save(&logVisitorInfo)
	c.JSON(200, logVisitorInfo)
}
func deleteLogVisitorInfo(c *gin.Context) {
	id := c.Params.ByName("id")
	var logVisitorInfo LogVisitorInfo
	d := g.Where("id = ?", id).Delete(&logVisitorInfo)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
