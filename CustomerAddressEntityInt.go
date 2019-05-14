package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CustomerAddressEntityInt :
type CustomerAddressEntityInt struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        uint16 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CustomerAddressEntityInt) TableName() string {
	return "customer_address_entity_int"
}

// handler create
func initRoutersCustomerAddressEntityInt(r *gin.Engine, customeraddressentityint string) {
	route := r.Group("customeraddressentityint")
	route.GET("/", getCustomerAddressEntityInts)
	route.GET("/:id", getCustomerAddressEntityInt)
	route.POST("/", createCustomerAddressEntityInt)
	route.PUT("/:id", updateCustomerAddressEntityInt)
	route.DELETE("/:id", deleteCustomerAddressEntityInt)
}

func getCustomerAddressEntityInts(c *gin.Context) {
	var customerAddressEntityInts []CustomerAddressEntityInt
	if err := g.Find(&customerAddressEntityInts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerAddressEntityInts)
	}
}

func getCustomerAddressEntityInt(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerAddressEntityInt CustomerAddressEntityInt
	if err := g.Where("id = ?", id).First(&customerAddressEntityInt).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerAddressEntityInt)
	}
}

func createCustomerAddressEntityInt(c *gin.Context) {
	var customerAddressEntityInt CustomerAddressEntityInt
	c.BindJSON(&customerAddressEntityInt)
	g.Create(&customerAddressEntityInt)
	c.JSON(200, customerAddressEntityInt)
}

func updateCustomerAddressEntityInt(c *gin.Context) {
	var customerAddressEntityInt CustomerAddressEntityInt
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerAddressEntityInt).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerAddressEntityInt)
	g.Save(&customerAddressEntityInt)
	c.JSON(200, customerAddressEntityInt)
}
func deleteCustomerAddressEntityInt(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerAddressEntityInt CustomerAddressEntityInt
	d := g.Where("id = ?", id).Delete(&customerAddressEntityInt)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
