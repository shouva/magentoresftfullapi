package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CustomerEavAttribute :
type CustomerEavAttribute struct {
	AttributeId    uint16 `gorm:"column:attribute_id;primary_key" form:"attribute_id;primary_key" json:"attribute_id;primary_key"`
	IsVisible      uint16 `gorm:"column:is_visible" form:"is_visible" json:"is_visible"`
	InputFilter    string `gorm:"column:input_filter" form:"input_filter" json:"input_filter"`
	MultilineCount uint16 `gorm:"column:multiline_count" form:"multiline_count" json:"multiline_count"`
	ValidateRules  string `gorm:"column:validate_rules" form:"validate_rules" json:"validate_rules"`
	IsSystem       uint16 `gorm:"column:is_system" form:"is_system" json:"is_system"`
	SortOrder      uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
	DataModel      string `gorm:"column:data_model" form:"data_model" json:"data_model"`
}

// TableName :
func (*CustomerEavAttribute) TableName() string {
	return "customer_eav_attribute"
}

// handler create
func initRoutersCustomerEavAttribute(r *gin.Engine, customereavattribute string) {
	route := r.Group("customereavattribute")
	route.GET("/", getCustomerEavAttributes)
	route.GET("/:id", getCustomerEavAttribute)
	route.POST("/", createCustomerEavAttribute)
	route.PUT("/:id", updateCustomerEavAttribute)
	route.DELETE("/:id", deleteCustomerEavAttribute)
}

func getCustomerEavAttributes(c *gin.Context) {
	var customerEavAttributes []CustomerEavAttribute
	if err := g.Find(&customerEavAttributes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEavAttributes)
	}
}

func getCustomerEavAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEavAttribute CustomerEavAttribute
	if err := g.Where("id = ?", id).First(&customerEavAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEavAttribute)
	}
}

func createCustomerEavAttribute(c *gin.Context) {
	var customerEavAttribute CustomerEavAttribute
	c.BindJSON(&customerEavAttribute)
	g.Create(&customerEavAttribute)
	c.JSON(200, customerEavAttribute)
}

func updateCustomerEavAttribute(c *gin.Context) {
	var customerEavAttribute CustomerEavAttribute
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerEavAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerEavAttribute)
	g.Save(&customerEavAttribute)
	c.JSON(200, customerEavAttribute)
}
func deleteCustomerEavAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEavAttribute CustomerEavAttribute
	d := g.Where("id = ?", id).Delete(&customerEavAttribute)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
