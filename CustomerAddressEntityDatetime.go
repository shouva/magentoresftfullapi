package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CustomerAddressEntityDatetime :
type CustomerAddressEntityDatetime struct {
	ValueId      uint16     `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16     `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16     `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	EntityId     uint16     `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        *time.Time `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CustomerAddressEntityDatetime) TableName() string {
	return "customer_address_entity_datetime"
}

// handler create
func initRoutersCustomerAddressEntityDatetime(r *gin.Engine, customeraddressentitydatetime string) {
	route := r.Group("customeraddressentitydatetime")
	route.GET("/", getCustomerAddressEntityDatetimes)
	route.GET("/:id", getCustomerAddressEntityDatetime)
	route.POST("/", createCustomerAddressEntityDatetime)
	route.PUT("/:id", updateCustomerAddressEntityDatetime)
	route.DELETE("/:id", deleteCustomerAddressEntityDatetime)
}

func getCustomerAddressEntityDatetimes(c *gin.Context) {
	var customerAddressEntityDatetimes []CustomerAddressEntityDatetime
	if err := g.Find(&customerAddressEntityDatetimes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerAddressEntityDatetimes)
	}
}

func getCustomerAddressEntityDatetime(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerAddressEntityDatetime CustomerAddressEntityDatetime
	if err := g.Where("id = ?", id).First(&customerAddressEntityDatetime).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerAddressEntityDatetime)
	}
}

func createCustomerAddressEntityDatetime(c *gin.Context) {
	var customerAddressEntityDatetime CustomerAddressEntityDatetime
	c.BindJSON(&customerAddressEntityDatetime)
	g.Create(&customerAddressEntityDatetime)
	c.JSON(200, customerAddressEntityDatetime)
}

func updateCustomerAddressEntityDatetime(c *gin.Context) {
	var customerAddressEntityDatetime CustomerAddressEntityDatetime
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerAddressEntityDatetime).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerAddressEntityDatetime)
	g.Save(&customerAddressEntityDatetime)
	c.JSON(200, customerAddressEntityDatetime)
}
func deleteCustomerAddressEntityDatetime(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerAddressEntityDatetime CustomerAddressEntityDatetime
	d := g.Where("id = ?", id).Delete(&customerAddressEntityDatetime)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
