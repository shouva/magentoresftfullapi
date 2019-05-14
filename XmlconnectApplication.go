package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// XmlconnectApplication :
type XmlconnectApplication struct {
	ApplicationId uint16     `gorm:"column:application_id;primary_key" form:"application_id;primary_key" json:"application_id;primary_key"`
	Name          string     `gorm:"column:name" form:"name" json:"name"`
	Code          string     `gorm:"column:code" form:"code" json:"code"`
	Type          string     `gorm:"column:type" form:"type" json:"type"`
	StoreId       uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	ActiveFrom    *time.Time `gorm:"column:active_from" form:"active_from" json:"active_from"`
	ActiveTo      *time.Time `gorm:"column:active_to" form:"active_to" json:"active_to"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	Status        uint16     `gorm:"column:status" form:"status" json:"status"`
	BrowsingMode  uint16     `gorm:"column:browsing_mode" form:"browsing_mode" json:"browsing_mode"`
}

// TableName :
func (*XmlconnectApplication) TableName() string {
	return "xmlconnect_application"
}

// handler create
func initRoutersXmlconnectApplication(r *gin.Engine, xmlconnectapplication string) {
	route := r.Group("xmlconnectapplication")
	route.GET("/", getXmlconnectApplications)
	route.GET("/:id", getXmlconnectApplication)
	route.POST("/", createXmlconnectApplication)
	route.PUT("/:id", updateXmlconnectApplication)
	route.DELETE("/:id", deleteXmlconnectApplication)
}

func getXmlconnectApplications(c *gin.Context) {
	var xmlconnectApplications []XmlconnectApplication
	if err := g.Find(&xmlconnectApplications).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, xmlconnectApplications)
	}
}

func getXmlconnectApplication(c *gin.Context) {
	id := c.Params.ByName("id")
	var xmlconnectApplication XmlconnectApplication
	if err := g.Where("id = ?", id).First(&xmlconnectApplication).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, xmlconnectApplication)
	}
}

func createXmlconnectApplication(c *gin.Context) {
	var xmlconnectApplication XmlconnectApplication
	c.BindJSON(&xmlconnectApplication)
	g.Create(&xmlconnectApplication)
	c.JSON(200, xmlconnectApplication)
}

func updateXmlconnectApplication(c *gin.Context) {
	var xmlconnectApplication XmlconnectApplication
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&xmlconnectApplication).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&xmlconnectApplication)
	g.Save(&xmlconnectApplication)
	c.JSON(200, xmlconnectApplication)
}
func deleteXmlconnectApplication(c *gin.Context) {
	id := c.Params.ByName("id")
	var xmlconnectApplication XmlconnectApplication
	d := g.Where("id = ?", id).Delete(&xmlconnectApplication)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
