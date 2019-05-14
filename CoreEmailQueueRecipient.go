package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreEmailQueueRecipient :
type CoreEmailQueueRecipient struct {
	RecipientId    uint16 `gorm:"column:recipient_id;primary_key" form:"recipient_id;primary_key" json:"recipient_id;primary_key"`
	MessageId      uint16 `gorm:"column:message_id" form:"message_id" json:"message_id"`
	RecipientEmail string `gorm:"column:recipient_email" form:"recipient_email" json:"recipient_email"`
	RecipientName  string `gorm:"column:recipient_name" form:"recipient_name" json:"recipient_name"`
	EmailType      uint16 `gorm:"column:email_type" form:"email_type" json:"email_type"`
}

// TableName :
func (*CoreEmailQueueRecipient) TableName() string {
	return "core_email_queue_recipients"
}

// handler create
func initRoutersCoreEmailQueueRecipient(r *gin.Engine, coreemailqueuerecipient string) {
	route := r.Group("coreemailqueuerecipient")
	route.GET("/", getCoreEmailQueueRecipients)
	route.GET("/:id", getCoreEmailQueueRecipient)
	route.POST("/", createCoreEmailQueueRecipient)
	route.PUT("/:id", updateCoreEmailQueueRecipient)
	route.DELETE("/:id", deleteCoreEmailQueueRecipient)
}

func getCoreEmailQueueRecipients(c *gin.Context) {
	var coreEmailQueueRecipients []CoreEmailQueueRecipient
	if err := g.Find(&coreEmailQueueRecipients).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreEmailQueueRecipients)
	}
}

func getCoreEmailQueueRecipient(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreEmailQueueRecipient CoreEmailQueueRecipient
	if err := g.Where("id = ?", id).First(&coreEmailQueueRecipient).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreEmailQueueRecipient)
	}
}

func createCoreEmailQueueRecipient(c *gin.Context) {
	var coreEmailQueueRecipient CoreEmailQueueRecipient
	c.BindJSON(&coreEmailQueueRecipient)
	g.Create(&coreEmailQueueRecipient)
	c.JSON(200, coreEmailQueueRecipient)
}

func updateCoreEmailQueueRecipient(c *gin.Context) {
	var coreEmailQueueRecipient CoreEmailQueueRecipient
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreEmailQueueRecipient).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreEmailQueueRecipient)
	g.Save(&coreEmailQueueRecipient)
	c.JSON(200, coreEmailQueueRecipient)
}
func deleteCoreEmailQueueRecipient(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreEmailQueueRecipient CoreEmailQueueRecipient
	d := g.Where("id = ?", id).Delete(&coreEmailQueueRecipient)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
