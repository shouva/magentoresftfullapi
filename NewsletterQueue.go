package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// NewsletterQueue :
type NewsletterQueue struct {
	QueueId               uint16     `gorm:"column:queue_id;primary_key" form:"queue_id;primary_key" json:"queue_id;primary_key"`
	TemplateId            uint16     `gorm:"column:template_id" form:"template_id" json:"template_id"`
	NewsletterType        uint16     `gorm:"column:newsletter_type" form:"newsletter_type" json:"newsletter_type"`
	NewsletterText        string     `gorm:"column:newsletter_text" form:"newsletter_text" json:"newsletter_text"`
	NewsletterStyles      string     `gorm:"column:newsletter_styles" form:"newsletter_styles" json:"newsletter_styles"`
	NewsletterSubject     string     `gorm:"column:newsletter_subject" form:"newsletter_subject" json:"newsletter_subject"`
	NewsletterSenderName  string     `gorm:"column:newsletter_sender_name" form:"newsletter_sender_name" json:"newsletter_sender_name"`
	NewsletterSenderEmail string     `gorm:"column:newsletter_sender_email" form:"newsletter_sender_email" json:"newsletter_sender_email"`
	QueueStatus           uint16     `gorm:"column:queue_status" form:"queue_status" json:"queue_status"`
	QueueStartAt          *time.Time `gorm:"column:queue_start_at" form:"queue_start_at" json:"queue_start_at"`
	QueueFinishAt         *time.Time `gorm:"column:queue_finish_at" form:"queue_finish_at" json:"queue_finish_at"`
}

// TableName :
func (*NewsletterQueue) TableName() string {
	return "newsletter_queue"
}

// handler create
func initRoutersNewsletterQueue(r *gin.Engine, newsletterqueue string) {
	route := r.Group("newsletterqueue")
	route.GET("/", getNewsletterQueues)
	route.GET("/:id", getNewsletterQueue)
	route.POST("/", createNewsletterQueue)
	route.PUT("/:id", updateNewsletterQueue)
	route.DELETE("/:id", deleteNewsletterQueue)
}

func getNewsletterQueues(c *gin.Context) {
	var newsletterQueues []NewsletterQueue
	if err := g.Find(&newsletterQueues).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, newsletterQueues)
	}
}

func getNewsletterQueue(c *gin.Context) {
	id := c.Params.ByName("id")
	var newsletterQueue NewsletterQueue
	if err := g.Where("id = ?", id).First(&newsletterQueue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, newsletterQueue)
	}
}

func createNewsletterQueue(c *gin.Context) {
	var newsletterQueue NewsletterQueue
	c.BindJSON(&newsletterQueue)
	g.Create(&newsletterQueue)
	c.JSON(200, newsletterQueue)
}

func updateNewsletterQueue(c *gin.Context) {
	var newsletterQueue NewsletterQueue
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&newsletterQueue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&newsletterQueue)
	g.Save(&newsletterQueue)
	c.JSON(200, newsletterQueue)
}
func deleteNewsletterQueue(c *gin.Context) {
	id := c.Params.ByName("id")
	var newsletterQueue NewsletterQueue
	d := g.Where("id = ?", id).Delete(&newsletterQueue)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
