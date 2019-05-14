package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CustomerEntityDatetime :
type CustomerEntityDatetime struct {
	ValueId      uint16     `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16     `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16     `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	EntityId     uint16     `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        *time.Time `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CustomerEntityDatetime) TableName() string {
	return "customer_entity_datetime"
}

// handler create
func initRoutersCustomerEntityDatetime(r *gin.Engine, customerentitydatetime string) {
	route := r.Group("customerentitydatetime")
	route.GET("/", getCustomerEntityDatetimes)
	route.GET("/:id", getCustomerEntityDatetime)
	route.POST("/", createCustomerEntityDatetime)
	route.PUT("/:id", updateCustomerEntityDatetime)
	route.DELETE("/:id", deleteCustomerEntityDatetime)
}

func getCustomerEntityDatetimes(c *gin.Context) {
	var customerEntityDatetimes []CustomerEntityDatetime
	if err := g.Find(&customerEntityDatetimes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEntityDatetimes)
	}
}

func getCustomerEntityDatetime(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEntityDatetime CustomerEntityDatetime
	if err := g.Where("id = ?", id).First(&customerEntityDatetime).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEntityDatetime)
	}
}

func createCustomerEntityDatetime(c *gin.Context) {
	var customerEntityDatetime CustomerEntityDatetime
	c.BindJSON(&customerEntityDatetime)
	g.Create(&customerEntityDatetime)
	c.JSON(200, customerEntityDatetime)
}

func updateCustomerEntityDatetime(c *gin.Context) {
	var customerEntityDatetime CustomerEntityDatetime
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerEntityDatetime).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerEntityDatetime)
	g.Save(&customerEntityDatetime)
	c.JSON(200, customerEntityDatetime)
}
func deleteCustomerEntityDatetime(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEntityDatetime CustomerEntityDatetime
	d := g.Where("id = ?", id).Delete(&customerEntityDatetime)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
