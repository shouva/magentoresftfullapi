package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CustomerEavAttributeWebsite :
type CustomerEavAttributeWebsite struct {
	AttributeId    uint16 `gorm:"column:attribute_id;primary_key" form:"attribute_id;primary_key" json:"attribute_id;primary_key"`
	WebsiteId      uint16 `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	IsVisible      uint16 `gorm:"column:is_visible" form:"is_visible" json:"is_visible"`
	IsRequired     uint16 `gorm:"column:is_required" form:"is_required" json:"is_required"`
	DefaultValue   string `gorm:"column:default_value" form:"default_value" json:"default_value"`
	MultilineCount uint16 `gorm:"column:multiline_count" form:"multiline_count" json:"multiline_count"`
}

// TableName :
func (*CustomerEavAttributeWebsite) TableName() string {
	return "customer_eav_attribute_website"
}

// handler create
func initRoutersCustomerEavAttributeWebsite(r *gin.Engine, customereavattributewebsite string) {
	route := r.Group("customereavattributewebsite")
	route.GET("/", getCustomerEavAttributeWebsites)
	route.GET("/:id", getCustomerEavAttributeWebsite)
	route.POST("/", createCustomerEavAttributeWebsite)
	route.PUT("/:id", updateCustomerEavAttributeWebsite)
	route.DELETE("/:id", deleteCustomerEavAttributeWebsite)
}

func getCustomerEavAttributeWebsites(c *gin.Context) {
	var customerEavAttributeWebsites []CustomerEavAttributeWebsite
	if err := g.Find(&customerEavAttributeWebsites).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEavAttributeWebsites)
	}
}

func getCustomerEavAttributeWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEavAttributeWebsite CustomerEavAttributeWebsite
	if err := g.Where("id = ?", id).First(&customerEavAttributeWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEavAttributeWebsite)
	}
}

func createCustomerEavAttributeWebsite(c *gin.Context) {
	var customerEavAttributeWebsite CustomerEavAttributeWebsite
	c.BindJSON(&customerEavAttributeWebsite)
	g.Create(&customerEavAttributeWebsite)
	c.JSON(200, customerEavAttributeWebsite)
}

func updateCustomerEavAttributeWebsite(c *gin.Context) {
	var customerEavAttributeWebsite CustomerEavAttributeWebsite
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerEavAttributeWebsite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerEavAttributeWebsite)
	g.Save(&customerEavAttributeWebsite)
	c.JSON(200, customerEavAttributeWebsite)
}
func deleteCustomerEavAttributeWebsite(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEavAttributeWebsite CustomerEavAttributeWebsite
	d := g.Where("id = ?", id).Delete(&customerEavAttributeWebsite)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
