package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// XmlconnectNotificationTemplate :
type XmlconnectNotificationTemplate struct {
	TemplateId    uint16     `gorm:"column:template_id;primary_key" form:"template_id;primary_key" json:"template_id;primary_key"`
	Name          string     `gorm:"column:name" form:"name" json:"name"`
	PushTitle     string     `gorm:"column:push_title" form:"push_title" json:"push_title"`
	MessageTitle  string     `gorm:"column:message_title" form:"message_title" json:"message_title"`
	Content       string     `gorm:"column:content" form:"content" json:"content"`
	CreatedAt     *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	ModifiedAt    *time.Time `gorm:"column:modified_at" form:"modified_at" json:"modified_at"`
	ApplicationId uint16     `gorm:"column:application_id" form:"application_id" json:"application_id"`
}

// TableName :
func (*XmlconnectNotificationTemplate) TableName() string {
	return "xmlconnect_notification_template"
}

// handler create
func initRoutersXmlconnectNotificationTemplate(r *gin.Engine, xmlconnectnotificationtemplate string) {
	route := r.Group("xmlconnectnotificationtemplate")
	route.GET("/", getXmlconnectNotificationTemplates)
	route.GET("/:id", getXmlconnectNotificationTemplate)
	route.POST("/", createXmlconnectNotificationTemplate)
	route.PUT("/:id", updateXmlconnectNotificationTemplate)
	route.DELETE("/:id", deleteXmlconnectNotificationTemplate)
}

func getXmlconnectNotificationTemplates(c *gin.Context) {
	var xmlconnectNotificationTemplates []XmlconnectNotificationTemplate
	if err := g.Find(&xmlconnectNotificationTemplates).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, xmlconnectNotificationTemplates)
	}
}

func getXmlconnectNotificationTemplate(c *gin.Context) {
	id := c.Params.ByName("id")
	var xmlconnectNotificationTemplate XmlconnectNotificationTemplate
	if err := g.Where("id = ?", id).First(&xmlconnectNotificationTemplate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, xmlconnectNotificationTemplate)
	}
}

func createXmlconnectNotificationTemplate(c *gin.Context) {
	var xmlconnectNotificationTemplate XmlconnectNotificationTemplate
	c.BindJSON(&xmlconnectNotificationTemplate)
	g.Create(&xmlconnectNotificationTemplate)
	c.JSON(200, xmlconnectNotificationTemplate)
}

func updateXmlconnectNotificationTemplate(c *gin.Context) {
	var xmlconnectNotificationTemplate XmlconnectNotificationTemplate
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&xmlconnectNotificationTemplate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&xmlconnectNotificationTemplate)
	g.Save(&xmlconnectNotificationTemplate)
	c.JSON(200, xmlconnectNotificationTemplate)
}
func deleteXmlconnectNotificationTemplate(c *gin.Context) {
	id := c.Params.ByName("id")
	var xmlconnectNotificationTemplate XmlconnectNotificationTemplate
	d := g.Where("id = ?", id).Delete(&xmlconnectNotificationTemplate)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
