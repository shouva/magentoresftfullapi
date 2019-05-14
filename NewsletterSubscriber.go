package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// NewsletterSubscriber :
type NewsletterSubscriber struct {
	SubscriberId          uint16     `gorm:"column:subscriber_id;primary_key" form:"subscriber_id;primary_key" json:"subscriber_id;primary_key"`
	StoreId               uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	ChangeStatusAt        *time.Time `gorm:"column:change_status_at" form:"change_status_at" json:"change_status_at"`
	CustomerId            uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	SubscriberEmail       string     `gorm:"column:subscriber_email" form:"subscriber_email" json:"subscriber_email"`
	SubscriberStatus      uint16     `gorm:"column:subscriber_status" form:"subscriber_status" json:"subscriber_status"`
	SubscriberConfirmCode string     `gorm:"column:subscriber_confirm_code" form:"subscriber_confirm_code" json:"subscriber_confirm_code"`
}

// TableName :
func (*NewsletterSubscriber) TableName() string {
	return "newsletter_subscriber"
}

// handler create
func initRoutersNewsletterSubscriber(r *gin.Engine, newslettersubscriber string) {
	route := r.Group("newslettersubscriber")
	route.GET("/", getNewsletterSubscribers)
	route.GET("/:id", getNewsletterSubscriber)
	route.POST("/", createNewsletterSubscriber)
	route.PUT("/:id", updateNewsletterSubscriber)
	route.DELETE("/:id", deleteNewsletterSubscriber)
}

func getNewsletterSubscribers(c *gin.Context) {
	var newsletterSubscribers []NewsletterSubscriber
	if err := g.Find(&newsletterSubscribers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, newsletterSubscribers)
	}
}

func getNewsletterSubscriber(c *gin.Context) {
	id := c.Params.ByName("id")
	var newsletterSubscriber NewsletterSubscriber
	if err := g.Where("id = ?", id).First(&newsletterSubscriber).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, newsletterSubscriber)
	}
}

func createNewsletterSubscriber(c *gin.Context) {
	var newsletterSubscriber NewsletterSubscriber
	c.BindJSON(&newsletterSubscriber)
	g.Create(&newsletterSubscriber)
	c.JSON(200, newsletterSubscriber)
}

func updateNewsletterSubscriber(c *gin.Context) {
	var newsletterSubscriber NewsletterSubscriber
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&newsletterSubscriber).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&newsletterSubscriber)
	g.Save(&newsletterSubscriber)
	c.JSON(200, newsletterSubscriber)
}
func deleteNewsletterSubscriber(c *gin.Context) {
	id := c.Params.ByName("id")
	var newsletterSubscriber NewsletterSubscriber
	d := g.Where("id = ?", id).Delete(&newsletterSubscriber)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
