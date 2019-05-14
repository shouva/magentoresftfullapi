package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// IndexEvent :
type IndexEvent struct {
	EventId   uint64     `gorm:"column:event_id;primary_key" form:"event_id;primary_key" json:"event_id;primary_key"`
	Type      string     `gorm:"column:type" form:"type" json:"type"`
	Entity    string     `gorm:"column:entity" form:"entity" json:"entity"`
	EntityPk  uint64     `gorm:"column:entity_pk" form:"entity_pk" json:"entity_pk"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	OldData   string     `gorm:"column:old_data" form:"old_data" json:"old_data"`
	NewData   string     `gorm:"column:new_data" form:"new_data" json:"new_data"`
}

// TableName :
func (*IndexEvent) TableName() string {
	return "index_event"
}

// handler create
func initRoutersIndexEvent(r *gin.Engine, indexevent string) {
	route := r.Group("indexevent")
	route.GET("/", getIndexEvents)
	route.GET("/:id", getIndexEvent)
	route.POST("/", createIndexEvent)
	route.PUT("/:id", updateIndexEvent)
	route.DELETE("/:id", deleteIndexEvent)
}

func getIndexEvents(c *gin.Context) {
	var indexEvents []IndexEvent
	if err := g.Find(&indexEvents).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, indexEvents)
	}
}

func getIndexEvent(c *gin.Context) {
	id := c.Params.ByName("id")
	var indexEvent IndexEvent
	if err := g.Where("id = ?", id).First(&indexEvent).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, indexEvent)
	}
}

func createIndexEvent(c *gin.Context) {
	var indexEvent IndexEvent
	c.BindJSON(&indexEvent)
	g.Create(&indexEvent)
	c.JSON(200, indexEvent)
}

func updateIndexEvent(c *gin.Context) {
	var indexEvent IndexEvent
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&indexEvent).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&indexEvent)
	g.Save(&indexEvent)
	c.JSON(200, indexEvent)
}
func deleteIndexEvent(c *gin.Context) {
	id := c.Params.ByName("id")
	var indexEvent IndexEvent
	d := g.Where("id = ?", id).Delete(&indexEvent)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
