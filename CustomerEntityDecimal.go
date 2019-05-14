package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CustomerEntityDecimal :
type CustomerEntityDecimal struct {
	ValueId      uint16  `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16  `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16  `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	EntityId     uint16  `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        float64 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CustomerEntityDecimal) TableName() string {
	return "customer_entity_decimal"
}

// handler create
func initRoutersCustomerEntityDecimal(r *gin.Engine, customerentitydecimal string) {
	route := r.Group("customerentitydecimal")
	route.GET("/", getCustomerEntityDecimals)
	route.GET("/:id", getCustomerEntityDecimal)
	route.POST("/", createCustomerEntityDecimal)
	route.PUT("/:id", updateCustomerEntityDecimal)
	route.DELETE("/:id", deleteCustomerEntityDecimal)
}

func getCustomerEntityDecimals(c *gin.Context) {
	var customerEntityDecimals []CustomerEntityDecimal
	if err := g.Find(&customerEntityDecimals).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEntityDecimals)
	}
}

func getCustomerEntityDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEntityDecimal CustomerEntityDecimal
	if err := g.Where("id = ?", id).First(&customerEntityDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEntityDecimal)
	}
}

func createCustomerEntityDecimal(c *gin.Context) {
	var customerEntityDecimal CustomerEntityDecimal
	c.BindJSON(&customerEntityDecimal)
	g.Create(&customerEntityDecimal)
	c.JSON(200, customerEntityDecimal)
}

func updateCustomerEntityDecimal(c *gin.Context) {
	var customerEntityDecimal CustomerEntityDecimal
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerEntityDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerEntityDecimal)
	g.Save(&customerEntityDecimal)
	c.JSON(200, customerEntityDecimal)
}
func deleteCustomerEntityDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEntityDecimal CustomerEntityDecimal
	d := g.Where("id = ?", id).Delete(&customerEntityDecimal)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
