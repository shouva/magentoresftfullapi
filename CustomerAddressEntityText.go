package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CustomerAddressEntityText :
type CustomerAddressEntityText struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CustomerAddressEntityText) TableName() string {
	return "customer_address_entity_text"
}

// handler create
func initRoutersCustomerAddressEntityText(r *gin.Engine, customeraddressentitytext string) {
	route := r.Group("customeraddressentitytext")
	route.GET("/", getCustomerAddressEntityTexts)
	route.GET("/:id", getCustomerAddressEntityText)
	route.POST("/", createCustomerAddressEntityText)
	route.PUT("/:id", updateCustomerAddressEntityText)
	route.DELETE("/:id", deleteCustomerAddressEntityText)
}

func getCustomerAddressEntityTexts(c *gin.Context) {
	var customerAddressEntityTexts []CustomerAddressEntityText
	if err := g.Find(&customerAddressEntityTexts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerAddressEntityTexts)
	}
}

func getCustomerAddressEntityText(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerAddressEntityText CustomerAddressEntityText
	if err := g.Where("id = ?", id).First(&customerAddressEntityText).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerAddressEntityText)
	}
}

func createCustomerAddressEntityText(c *gin.Context) {
	var customerAddressEntityText CustomerAddressEntityText
	c.BindJSON(&customerAddressEntityText)
	g.Create(&customerAddressEntityText)
	c.JSON(200, customerAddressEntityText)
}

func updateCustomerAddressEntityText(c *gin.Context) {
	var customerAddressEntityText CustomerAddressEntityText
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerAddressEntityText).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerAddressEntityText)
	g.Save(&customerAddressEntityText)
	c.JSON(200, customerAddressEntityText)
}
func deleteCustomerAddressEntityText(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerAddressEntityText CustomerAddressEntityText
	d := g.Where("id = ?", id).Delete(&customerAddressEntityText)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
