package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CoreEmailQueue :
type CoreEmailQueue struct {
	MessageId         uint16     `gorm:"column:message_id;primary_key" form:"message_id;primary_key" json:"message_id;primary_key"`
	EntityId          uint16     `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	EntityType        string     `gorm:"column:entity_type" form:"entity_type" json:"entity_type"`
	EventType         string     `gorm:"column:event_type" form:"event_type" json:"event_type"`
	MessageBodyHash   string     `gorm:"column:message_body_hash" form:"message_body_hash" json:"message_body_hash"`
	MessageBody       string     `gorm:"column:message_body" form:"message_body" json:"message_body"`
	MessageParameters string     `gorm:"column:message_parameters" form:"message_parameters" json:"message_parameters"`
	CreatedAt         *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	ProcessedAt       *time.Time `gorm:"column:processed_at" form:"processed_at" json:"processed_at"`
}

// TableName :
func (*CoreEmailQueue) TableName() string {
	return "core_email_queue"
}

// handler create
func initRoutersCoreEmailQueue(r *gin.Engine, coreemailqueue string) {
	route := r.Group("coreemailqueue")
	route.GET("/", getCoreEmailQueues)
	route.GET("/:id", getCoreEmailQueue)
	route.POST("/", createCoreEmailQueue)
	route.PUT("/:id", updateCoreEmailQueue)
	route.DELETE("/:id", deleteCoreEmailQueue)
}

func getCoreEmailQueues(c *gin.Context) {
	var coreEmailQueues []CoreEmailQueue
	if err := g.Find(&coreEmailQueues).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreEmailQueues)
	}
}

func getCoreEmailQueue(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreEmailQueue CoreEmailQueue
	if err := g.Where("id = ?", id).First(&coreEmailQueue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreEmailQueue)
	}
}

func createCoreEmailQueue(c *gin.Context) {
	var coreEmailQueue CoreEmailQueue
	c.BindJSON(&coreEmailQueue)
	g.Create(&coreEmailQueue)
	c.JSON(200, coreEmailQueue)
}

func updateCoreEmailQueue(c *gin.Context) {
	var coreEmailQueue CoreEmailQueue
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreEmailQueue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreEmailQueue)
	g.Save(&coreEmailQueue)
	c.JSON(200, coreEmailQueue)
}
func deleteCoreEmailQueue(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreEmailQueue CoreEmailQueue
	d := g.Where("id = ?", id).Delete(&coreEmailQueue)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
