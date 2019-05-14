package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CustomerAddressEntityVarchar :
type CustomerAddressEntityVarchar struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CustomerAddressEntityVarchar) TableName() string {
	return "customer_address_entity_varchar"
}

// handler create
func initRoutersCustomerAddressEntityVarchar(r *gin.Engine, customeraddressentityvarchar string) {
	route := r.Group("customeraddressentityvarchar")
	route.GET("/", getCustomerAddressEntityVarchars)
	route.GET("/:id", getCustomerAddressEntityVarchar)
	route.POST("/", createCustomerAddressEntityVarchar)
	route.PUT("/:id", updateCustomerAddressEntityVarchar)
	route.DELETE("/:id", deleteCustomerAddressEntityVarchar)
}

func getCustomerAddressEntityVarchars(c *gin.Context) {
	var customerAddressEntityVarchars []CustomerAddressEntityVarchar
	if err := g.Find(&customerAddressEntityVarchars).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerAddressEntityVarchars)
	}
}

func getCustomerAddressEntityVarchar(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerAddressEntityVarchar CustomerAddressEntityVarchar
	if err := g.Where("id = ?", id).First(&customerAddressEntityVarchar).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerAddressEntityVarchar)
	}
}

func createCustomerAddressEntityVarchar(c *gin.Context) {
	var customerAddressEntityVarchar CustomerAddressEntityVarchar
	c.BindJSON(&customerAddressEntityVarchar)
	g.Create(&customerAddressEntityVarchar)
	c.JSON(200, customerAddressEntityVarchar)
}

func updateCustomerAddressEntityVarchar(c *gin.Context) {
	var customerAddressEntityVarchar CustomerAddressEntityVarchar
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerAddressEntityVarchar).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerAddressEntityVarchar)
	g.Save(&customerAddressEntityVarchar)
	c.JSON(200, customerAddressEntityVarchar)
}
func deleteCustomerAddressEntityVarchar(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerAddressEntityVarchar CustomerAddressEntityVarchar
	d := g.Where("id = ?", id).Delete(&customerAddressEntityVarchar)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
