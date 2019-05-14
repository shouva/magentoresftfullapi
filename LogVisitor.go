package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LogVisitor :
type LogVisitor struct {
	VisitorId    uint64     `gorm:"column:visitor_id;primary_key" form:"visitor_id;primary_key" json:"visitor_id;primary_key"`
	SessionId    string     `gorm:"column:session_id" form:"session_id" json:"session_id"`
	FirstVisitAt *time.Time `gorm:"column:first_visit_at" form:"first_visit_at" json:"first_visit_at"`
	LastVisitAt  *time.Time `gorm:"column:last_visit_at" form:"last_visit_at" json:"last_visit_at"`
	LastUrlId    uint64     `gorm:"column:last_url_id" form:"last_url_id" json:"last_url_id"`
	StoreId      uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
}

// TableName :
func (*LogVisitor) TableName() string {
	return "log_visitor"
}

// handler create
func initRoutersLogVisitor(r *gin.Engine, logvisitor string) {
	route := r.Group("logvisitor")
	route.GET("/", getLogVisitors)
	route.GET("/:id", getLogVisitor)
	route.POST("/", createLogVisitor)
	route.PUT("/:id", updateLogVisitor)
	route.DELETE("/:id", deleteLogVisitor)
}

func getLogVisitors(c *gin.Context) {
	var logVisitors []LogVisitor
	if err := g.Find(&logVisitors).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logVisitors)
	}
}

func getLogVisitor(c *gin.Context) {
	id := c.Params.ByName("id")
	var logVisitor LogVisitor
	if err := g.Where("id = ?", id).First(&logVisitor).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logVisitor)
	}
}

func createLogVisitor(c *gin.Context) {
	var logVisitor LogVisitor
	c.BindJSON(&logVisitor)
	g.Create(&logVisitor)
	c.JSON(200, logVisitor)
}

func updateLogVisitor(c *gin.Context) {
	var logVisitor LogVisitor
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&logVisitor).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&logVisitor)
	g.Save(&logVisitor)
	c.JSON(200, logVisitor)
}
func deleteLogVisitor(c *gin.Context) {
	id := c.Params.ByName("id")
	var logVisitor LogVisitor
	d := g.Where("id = ?", id).Delete(&logVisitor)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
