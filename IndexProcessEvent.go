package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// IndexProcessEvent :
type IndexProcessEvent struct {
	ProcessId uint16 `gorm:"column:process_id;primary_key" form:"process_id;primary_key" json:"process_id;primary_key"`
	EventId   uint64 `gorm:"column:event_id;primary_key" form:"event_id;primary_key" json:"event_id;primary_key"`
	Status    string `gorm:"column:status" form:"status" json:"status"`
}

// TableName :
func (*IndexProcessEvent) TableName() string {
	return "index_process_event"
}

// handler create
func initRoutersIndexProcessEvent(r *gin.Engine, indexprocessevent string) {
	route := r.Group("indexprocessevent")
	route.GET("/", getIndexProcessEvents)
	route.GET("/:id", getIndexProcessEvent)
	route.POST("/", createIndexProcessEvent)
	route.PUT("/:id", updateIndexProcessEvent)
	route.DELETE("/:id", deleteIndexProcessEvent)
}

func getIndexProcessEvents(c *gin.Context) {
	var indexProcessEvents []IndexProcessEvent
	if err := g.Find(&indexProcessEvents).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, indexProcessEvents)
	}
}

func getIndexProcessEvent(c *gin.Context) {
	id := c.Params.ByName("id")
	var indexProcessEvent IndexProcessEvent
	if err := g.Where("id = ?", id).First(&indexProcessEvent).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, indexProcessEvent)
	}
}

func createIndexProcessEvent(c *gin.Context) {
	var indexProcessEvent IndexProcessEvent
	c.BindJSON(&indexProcessEvent)
	g.Create(&indexProcessEvent)
	c.JSON(200, indexProcessEvent)
}

func updateIndexProcessEvent(c *gin.Context) {
	var indexProcessEvent IndexProcessEvent
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&indexProcessEvent).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&indexProcessEvent)
	g.Save(&indexProcessEvent)
	c.JSON(200, indexProcessEvent)
}
func deleteIndexProcessEvent(c *gin.Context) {
	id := c.Params.ByName("id")
	var indexProcessEvent IndexProcessEvent
	d := g.Where("id = ?", id).Delete(&indexProcessEvent)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
