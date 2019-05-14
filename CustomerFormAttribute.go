package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CustomerFormAttribute :
type CustomerFormAttribute struct {
	FormCode    string `gorm:"column:form_code;primary_key" form:"form_code;primary_key" json:"form_code;primary_key"`
	AttributeId uint16 `gorm:"column:attribute_id;primary_key" form:"attribute_id;primary_key" json:"attribute_id;primary_key"`
}

// TableName :
func (*CustomerFormAttribute) TableName() string {
	return "customer_form_attribute"
}

// handler create
func initRoutersCustomerFormAttribute(r *gin.Engine, customerformattribute string) {
	route := r.Group("customerformattribute")
	route.GET("/", getCustomerFormAttributes)
	route.GET("/:id", getCustomerFormAttribute)
	route.POST("/", createCustomerFormAttribute)
	route.PUT("/:id", updateCustomerFormAttribute)
	route.DELETE("/:id", deleteCustomerFormAttribute)
}

func getCustomerFormAttributes(c *gin.Context) {
	var customerFormAttributes []CustomerFormAttribute
	if err := g.Find(&customerFormAttributes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerFormAttributes)
	}
}

func getCustomerFormAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerFormAttribute CustomerFormAttribute
	if err := g.Where("id = ?", id).First(&customerFormAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerFormAttribute)
	}
}

func createCustomerFormAttribute(c *gin.Context) {
	var customerFormAttribute CustomerFormAttribute
	c.BindJSON(&customerFormAttribute)
	g.Create(&customerFormAttribute)
	c.JSON(200, customerFormAttribute)
}

func updateCustomerFormAttribute(c *gin.Context) {
	var customerFormAttribute CustomerFormAttribute
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerFormAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerFormAttribute)
	g.Save(&customerFormAttribute)
	c.JSON(200, customerFormAttribute)
}
func deleteCustomerFormAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerFormAttribute CustomerFormAttribute
	d := g.Where("id = ?", id).Delete(&customerFormAttribute)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
