package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// GiftMessage :
type GiftMessage struct {
	GiftMessageId uint16 `gorm:"column:gift_message_id;primary_key" form:"gift_message_id;primary_key" json:"gift_message_id;primary_key"`
	CustomerId    uint16 `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	Sender        string `gorm:"column:sender" form:"sender" json:"sender"`
	Recipient     string `gorm:"column:recipient" form:"recipient" json:"recipient"`
	Message       string `gorm:"column:message" form:"message" json:"message"`
}

// TableName :
func (*GiftMessage) TableName() string {
	return "gift_message"
}

// handler create
func initRoutersGiftMessage(r *gin.Engine, giftmessage string) {
	route := r.Group("giftmessage")
	route.GET("/", getGiftMessages)
	route.GET("/:id", getGiftMessage)
	route.POST("/", createGiftMessage)
	route.PUT("/:id", updateGiftMessage)
	route.DELETE("/:id", deleteGiftMessage)
}

func getGiftMessages(c *gin.Context) {
	var giftMessages []GiftMessage
	if err := g.Find(&giftMessages).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, giftMessages)
	}
}

func getGiftMessage(c *gin.Context) {
	id := c.Params.ByName("id")
	var giftMessage GiftMessage
	if err := g.Where("id = ?", id).First(&giftMessage).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, giftMessage)
	}
}

func createGiftMessage(c *gin.Context) {
	var giftMessage GiftMessage
	c.BindJSON(&giftMessage)
	g.Create(&giftMessage)
	c.JSON(200, giftMessage)
}

func updateGiftMessage(c *gin.Context) {
	var giftMessage GiftMessage
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&giftMessage).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&giftMessage)
	g.Save(&giftMessage)
	c.JSON(200, giftMessage)
}
func deleteGiftMessage(c *gin.Context) {
	id := c.Params.ByName("id")
	var giftMessage GiftMessage
	d := g.Where("id = ?", id).Delete(&giftMessage)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
