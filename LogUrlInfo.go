package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// LogUrlInfo :
type LogUrlInfo struct {
	UrlId   uint64 `gorm:"column:url_id;primary_key" form:"url_id;primary_key" json:"url_id;primary_key"`
	Url     string `gorm:"column:url" form:"url" json:"url"`
	Referer string `gorm:"column:referer" form:"referer" json:"referer"`
}

// TableName :
func (*LogUrlInfo) TableName() string {
	return "log_url_info"
}

// handler create
func initRoutersLogUrlInfo(r *gin.Engine, logurlinfo string) {
	route := r.Group("logurlinfo")
	route.GET("/", getLogUrlInfos)
	route.GET("/:id", getLogUrlInfo)
	route.POST("/", createLogUrlInfo)
	route.PUT("/:id", updateLogUrlInfo)
	route.DELETE("/:id", deleteLogUrlInfo)
}

func getLogUrlInfos(c *gin.Context) {
	var logUrlInfos []LogUrlInfo
	if err := g.Find(&logUrlInfos).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logUrlInfos)
	}
}

func getLogUrlInfo(c *gin.Context) {
	id := c.Params.ByName("id")
	var logUrlInfo LogUrlInfo
	if err := g.Where("id = ?", id).First(&logUrlInfo).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logUrlInfo)
	}
}

func createLogUrlInfo(c *gin.Context) {
	var logUrlInfo LogUrlInfo
	c.BindJSON(&logUrlInfo)
	g.Create(&logUrlInfo)
	c.JSON(200, logUrlInfo)
}

func updateLogUrlInfo(c *gin.Context) {
	var logUrlInfo LogUrlInfo
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&logUrlInfo).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&logUrlInfo)
	g.Save(&logUrlInfo)
	c.JSON(200, logUrlInfo)
}
func deleteLogUrlInfo(c *gin.Context) {
	id := c.Params.ByName("id")
	var logUrlInfo LogUrlInfo
	d := g.Where("id = ?", id).Delete(&logUrlInfo)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
