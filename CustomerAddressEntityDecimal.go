package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CustomerAddressEntityDecimal :
type CustomerAddressEntityDecimal struct {
	ValueId      uint16  `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16  `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16  `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	EntityId     uint16  `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        float64 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CustomerAddressEntityDecimal) TableName() string {
	return "customer_address_entity_decimal"
}

// handler create
func initRoutersCustomerAddressEntityDecimal(r *gin.Engine, customeraddressentitydecimal string) {
	route := r.Group("customeraddressentitydecimal")
	route.GET("/", getCustomerAddressEntityDecimals)
	route.GET("/:id", getCustomerAddressEntityDecimal)
	route.POST("/", createCustomerAddressEntityDecimal)
	route.PUT("/:id", updateCustomerAddressEntityDecimal)
	route.DELETE("/:id", deleteCustomerAddressEntityDecimal)
}

func getCustomerAddressEntityDecimals(c *gin.Context) {
	var customerAddressEntityDecimals []CustomerAddressEntityDecimal
	if err := g.Find(&customerAddressEntityDecimals).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerAddressEntityDecimals)
	}
}

func getCustomerAddressEntityDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerAddressEntityDecimal CustomerAddressEntityDecimal
	if err := g.Where("id = ?", id).First(&customerAddressEntityDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerAddressEntityDecimal)
	}
}

func createCustomerAddressEntityDecimal(c *gin.Context) {
	var customerAddressEntityDecimal CustomerAddressEntityDecimal
	c.BindJSON(&customerAddressEntityDecimal)
	g.Create(&customerAddressEntityDecimal)
	c.JSON(200, customerAddressEntityDecimal)
}

func updateCustomerAddressEntityDecimal(c *gin.Context) {
	var customerAddressEntityDecimal CustomerAddressEntityDecimal
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerAddressEntityDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerAddressEntityDecimal)
	g.Save(&customerAddressEntityDecimal)
	c.JSON(200, customerAddressEntityDecimal)
}
func deleteCustomerAddressEntityDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerAddressEntityDecimal CustomerAddressEntityDecimal
	d := g.Where("id = ?", id).Delete(&customerAddressEntityDecimal)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
