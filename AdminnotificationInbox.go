package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// AdminnotificationInbox :
type AdminnotificationInbox struct {
	NotificationId uint16     `gorm:"column:notification_id;primary_key" form:"notification_id;primary_key" json:"notification_id;primary_key"`
	Severity       uint16     `gorm:"column:severity" form:"severity" json:"severity"`
	DateAdded      *time.Time `gorm:"column:date_added" form:"date_added" json:"date_added"`
	Title          string     `gorm:"column:title" form:"title" json:"title"`
	Description    string     `gorm:"column:description" form:"description" json:"description"`
	Url            string     `gorm:"column:url" form:"url" json:"url"`
	IsRead         uint16     `gorm:"column:is_read" form:"is_read" json:"is_read"`
	IsRemove       uint16     `gorm:"column:is_remove" form:"is_remove" json:"is_remove"`
}

// TableName :
func (*AdminnotificationInbox) TableName() string {
	return "adminnotification_inbox"
}

// handler create
func initRoutersAdminnotificationInbox(r *gin.Engine, adminnotificationinbox string) {
	route := r.Group("adminnotificationinbox")
	route.GET("/", getAdminnotificationInboxs)
	route.GET("/:id", getAdminnotificationInbox)
	route.POST("/", createAdminnotificationInbox)
	route.PUT("/:id", updateAdminnotificationInbox)
	route.DELETE("/:id", deleteAdminnotificationInbox)
}

func getAdminnotificationInboxs(c *gin.Context) {
	var adminnotificationInboxs []AdminnotificationInbox
	if err := g.Find(&adminnotificationInboxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, adminnotificationInboxs)
	}
}

func getAdminnotificationInbox(c *gin.Context) {
	id := c.Params.ByName("id")
	var adminnotificationInbox AdminnotificationInbox
	if err := g.Where("id = ?", id).First(&adminnotificationInbox).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, adminnotificationInbox)
	}
}

func createAdminnotificationInbox(c *gin.Context) {
	var adminnotificationInbox AdminnotificationInbox
	c.BindJSON(&adminnotificationInbox)
	g.Create(&adminnotificationInbox)
	c.JSON(200, adminnotificationInbox)
}

func updateAdminnotificationInbox(c *gin.Context) {
	var adminnotificationInbox AdminnotificationInbox
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&adminnotificationInbox).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&adminnotificationInbox)
	g.Save(&adminnotificationInbox)
	c.JSON(200, adminnotificationInbox)
}
func deleteAdminnotificationInbox(c *gin.Context) {
	id := c.Params.ByName("id")
	var adminnotificationInbox AdminnotificationInbox
	d := g.Where("id = ?", id).Delete(&adminnotificationInbox)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
