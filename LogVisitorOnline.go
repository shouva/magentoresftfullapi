package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LogVisitorOnline :
type LogVisitorOnline struct {
	VisitorId    uint64     `gorm:"column:visitor_id;primary_key" form:"visitor_id;primary_key" json:"visitor_id;primary_key"`
	VisitorType  string     `gorm:"column:visitor_type" form:"visitor_type" json:"visitor_type"`
	RemoteAddr   []byte     `gorm:"column:remote_addr" form:"remote_addr" json:"remote_addr"`
	FirstVisitAt *time.Time `gorm:"column:first_visit_at" form:"first_visit_at" json:"first_visit_at"`
	LastVisitAt  *time.Time `gorm:"column:last_visit_at" form:"last_visit_at" json:"last_visit_at"`
	CustomerId   uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	LastUrl      string     `gorm:"column:last_url" form:"last_url" json:"last_url"`
}

// TableName :
func (*LogVisitorOnline) TableName() string {
	return "log_visitor_online"
}

// handler create
func initRoutersLogVisitorOnline(r *gin.Engine, logvisitoronline string) {
	route := r.Group("logvisitoronline")
	route.GET("/", getLogVisitorOnlines)
	route.GET("/:id", getLogVisitorOnline)
	route.POST("/", createLogVisitorOnline)
	route.PUT("/:id", updateLogVisitorOnline)
	route.DELETE("/:id", deleteLogVisitorOnline)
}

func getLogVisitorOnlines(c *gin.Context) {
	var logVisitorOnlines []LogVisitorOnline
	if err := g.Find(&logVisitorOnlines).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logVisitorOnlines)
	}
}

func getLogVisitorOnline(c *gin.Context) {
	id := c.Params.ByName("id")
	var logVisitorOnline LogVisitorOnline
	if err := g.Where("id = ?", id).First(&logVisitorOnline).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logVisitorOnline)
	}
}

func createLogVisitorOnline(c *gin.Context) {
	var logVisitorOnline LogVisitorOnline
	c.BindJSON(&logVisitorOnline)
	g.Create(&logVisitorOnline)
	c.JSON(200, logVisitorOnline)
}

func updateLogVisitorOnline(c *gin.Context) {
	var logVisitorOnline LogVisitorOnline
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&logVisitorOnline).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&logVisitorOnline)
	g.Save(&logVisitorOnline)
	c.JSON(200, logVisitorOnline)
}
func deleteLogVisitorOnline(c *gin.Context) {
	id := c.Params.ByName("id")
	var logVisitorOnline LogVisitorOnline
	d := g.Where("id = ?", id).Delete(&logVisitorOnline)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
