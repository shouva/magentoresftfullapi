package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// XmlconnectQueue :
type XmlconnectQueue struct {
	QueueId      uint16     `gorm:"column:queue_id;primary_key" form:"queue_id;primary_key" json:"queue_id;primary_key"`
	CreateTime   *time.Time `gorm:"column:create_time" form:"create_time" json:"create_time"`
	ExecTime     *time.Time `gorm:"column:exec_time" form:"exec_time" json:"exec_time"`
	TemplateId   uint16     `gorm:"column:template_id" form:"template_id" json:"template_id"`
	PushTitle    string     `gorm:"column:push_title" form:"push_title" json:"push_title"`
	MessageTitle string     `gorm:"column:message_title" form:"message_title" json:"message_title"`
	Content      string     `gorm:"column:content" form:"content" json:"content"`
	Status       uint16     `gorm:"column:status" form:"status" json:"status"`
	Type         string     `gorm:"column:type" form:"type" json:"type"`
}

// TableName :
func (*XmlconnectQueue) TableName() string {
	return "xmlconnect_queue"
}

// handler create
func initRoutersXmlconnectQueue(r *gin.Engine, xmlconnectqueue string) {
	route := r.Group("xmlconnectqueue")
	route.GET("/", getXmlconnectQueues)
	route.GET("/:id", getXmlconnectQueue)
	route.POST("/", createXmlconnectQueue)
	route.PUT("/:id", updateXmlconnectQueue)
	route.DELETE("/:id", deleteXmlconnectQueue)
}

func getXmlconnectQueues(c *gin.Context) {
	var xmlconnectQueues []XmlconnectQueue
	if err := g.Find(&xmlconnectQueues).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, xmlconnectQueues)
	}
}

func getXmlconnectQueue(c *gin.Context) {
	id := c.Params.ByName("id")
	var xmlconnectQueue XmlconnectQueue
	if err := g.Where("id = ?", id).First(&xmlconnectQueue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, xmlconnectQueue)
	}
}

func createXmlconnectQueue(c *gin.Context) {
	var xmlconnectQueue XmlconnectQueue
	c.BindJSON(&xmlconnectQueue)
	g.Create(&xmlconnectQueue)
	c.JSON(200, xmlconnectQueue)
}

func updateXmlconnectQueue(c *gin.Context) {
	var xmlconnectQueue XmlconnectQueue
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&xmlconnectQueue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&xmlconnectQueue)
	g.Save(&xmlconnectQueue)
	c.JSON(200, xmlconnectQueue)
}
func deleteXmlconnectQueue(c *gin.Context) {
	id := c.Params.ByName("id")
	var xmlconnectQueue XmlconnectQueue
	d := g.Where("id = ?", id).Delete(&xmlconnectQueue)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
