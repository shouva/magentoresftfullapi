package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// NewsletterQueueStoreLink :
type NewsletterQueueStoreLink struct {
	QueueId uint16 `gorm:"column:queue_id;primary_key" form:"queue_id;primary_key" json:"queue_id;primary_key"`
	StoreId uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
}

// TableName :
func (*NewsletterQueueStoreLink) TableName() string {
	return "newsletter_queue_store_link"
}

// handler create
func initRoutersNewsletterQueueStoreLink(r *gin.Engine, newsletterqueuestorelink string) {
	route := r.Group("newsletterqueuestorelink")
	route.GET("/", getNewsletterQueueStoreLinks)
	route.GET("/:id", getNewsletterQueueStoreLink)
	route.POST("/", createNewsletterQueueStoreLink)
	route.PUT("/:id", updateNewsletterQueueStoreLink)
	route.DELETE("/:id", deleteNewsletterQueueStoreLink)
}

func getNewsletterQueueStoreLinks(c *gin.Context) {
	var newsletterQueueStoreLinks []NewsletterQueueStoreLink
	if err := g.Find(&newsletterQueueStoreLinks).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, newsletterQueueStoreLinks)
	}
}

func getNewsletterQueueStoreLink(c *gin.Context) {
	id := c.Params.ByName("id")
	var newsletterQueueStoreLink NewsletterQueueStoreLink
	if err := g.Where("id = ?", id).First(&newsletterQueueStoreLink).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, newsletterQueueStoreLink)
	}
}

func createNewsletterQueueStoreLink(c *gin.Context) {
	var newsletterQueueStoreLink NewsletterQueueStoreLink
	c.BindJSON(&newsletterQueueStoreLink)
	g.Create(&newsletterQueueStoreLink)
	c.JSON(200, newsletterQueueStoreLink)
}

func updateNewsletterQueueStoreLink(c *gin.Context) {
	var newsletterQueueStoreLink NewsletterQueueStoreLink
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&newsletterQueueStoreLink).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&newsletterQueueStoreLink)
	g.Save(&newsletterQueueStoreLink)
	c.JSON(200, newsletterQueueStoreLink)
}
func deleteNewsletterQueueStoreLink(c *gin.Context) {
	id := c.Params.ByName("id")
	var newsletterQueueStoreLink NewsletterQueueStoreLink
	d := g.Where("id = ?", id).Delete(&newsletterQueueStoreLink)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
