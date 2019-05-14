package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LogUrl :
type LogUrl struct {
	UrlId     uint64     `gorm:"column:url_id" form:"url_id" json:"url_id"`
	VisitorId uint64     `gorm:"column:visitor_id" form:"visitor_id" json:"visitor_id"`
	VisitTime *time.Time `gorm:"column:visit_time" form:"visit_time" json:"visit_time"`
}

// TableName :
func (*LogUrl) TableName() string {
	return "log_url"
}

// handler create
func initRoutersLogUrl(r *gin.Engine, logurl string) {
	route := r.Group("logurl")
	route.GET("/", getLogUrls)
	route.GET("/:id", getLogUrl)
	route.POST("/", createLogUrl)
	route.PUT("/:id", updateLogUrl)
	route.DELETE("/:id", deleteLogUrl)
}

func getLogUrls(c *gin.Context) {
	var logUrls []LogUrl
	if err := g.Find(&logUrls).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logUrls)
	}
}

func getLogUrl(c *gin.Context) {
	id := c.Params.ByName("id")
	var logUrl LogUrl
	if err := g.Where("id = ?", id).First(&logUrl).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logUrl)
	}
}

func createLogUrl(c *gin.Context) {
	var logUrl LogUrl
	c.BindJSON(&logUrl)
	g.Create(&logUrl)
	c.JSON(200, logUrl)
}

func updateLogUrl(c *gin.Context) {
	var logUrl LogUrl
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&logUrl).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&logUrl)
	g.Save(&logUrl)
	c.JSON(200, logUrl)
}
func deleteLogUrl(c *gin.Context) {
	id := c.Params.ByName("id")
	var logUrl LogUrl
	d := g.Where("id = ?", id).Delete(&logUrl)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
