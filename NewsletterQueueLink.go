package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// NewsletterQueueLink :
type NewsletterQueueLink struct {
	QueueLinkId  uint16     `gorm:"column:queue_link_id;primary_key" form:"queue_link_id;primary_key" json:"queue_link_id;primary_key"`
	QueueId      uint16     `gorm:"column:queue_id" form:"queue_id" json:"queue_id"`
	SubscriberId uint16     `gorm:"column:subscriber_id" form:"subscriber_id" json:"subscriber_id"`
	LetterSentAt *time.Time `gorm:"column:letter_sent_at" form:"letter_sent_at" json:"letter_sent_at"`
}

// TableName :
func (*NewsletterQueueLink) TableName() string {
	return "newsletter_queue_link"
}

// handler create
func initRoutersNewsletterQueueLink(r *gin.Engine, newsletterqueuelink string) {
	route := r.Group("newsletterqueuelink")
	route.GET("/", getNewsletterQueueLinks)
	route.GET("/:id", getNewsletterQueueLink)
	route.POST("/", createNewsletterQueueLink)
	route.PUT("/:id", updateNewsletterQueueLink)
	route.DELETE("/:id", deleteNewsletterQueueLink)
}

func getNewsletterQueueLinks(c *gin.Context) {
	var newsletterQueueLinks []NewsletterQueueLink
	if err := g.Find(&newsletterQueueLinks).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, newsletterQueueLinks)
	}
}

func getNewsletterQueueLink(c *gin.Context) {
	id := c.Params.ByName("id")
	var newsletterQueueLink NewsletterQueueLink
	if err := g.Where("id = ?", id).First(&newsletterQueueLink).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, newsletterQueueLink)
	}
}

func createNewsletterQueueLink(c *gin.Context) {
	var newsletterQueueLink NewsletterQueueLink
	c.BindJSON(&newsletterQueueLink)
	g.Create(&newsletterQueueLink)
	c.JSON(200, newsletterQueueLink)
}

func updateNewsletterQueueLink(c *gin.Context) {
	var newsletterQueueLink NewsletterQueueLink
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&newsletterQueueLink).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&newsletterQueueLink)
	g.Save(&newsletterQueueLink)
	c.JSON(200, newsletterQueueLink)
}
func deleteNewsletterQueueLink(c *gin.Context) {
	id := c.Params.ByName("id")
	var newsletterQueueLink NewsletterQueueLink
	d := g.Where("id = ?", id).Delete(&newsletterQueueLink)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
