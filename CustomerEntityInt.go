package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CustomerEntityInt :
type CustomerEntityInt struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        uint16 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CustomerEntityInt) TableName() string {
	return "customer_entity_int"
}

// handler create
func initRoutersCustomerEntityInt(r *gin.Engine, customerentityint string) {
	route := r.Group("customerentityint")
	route.GET("/", getCustomerEntityInts)
	route.GET("/:id", getCustomerEntityInt)
	route.POST("/", createCustomerEntityInt)
	route.PUT("/:id", updateCustomerEntityInt)
	route.DELETE("/:id", deleteCustomerEntityInt)
}

func getCustomerEntityInts(c *gin.Context) {
	var customerEntityInts []CustomerEntityInt
	if err := g.Find(&customerEntityInts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEntityInts)
	}
}

func getCustomerEntityInt(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEntityInt CustomerEntityInt
	if err := g.Where("id = ?", id).First(&customerEntityInt).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEntityInt)
	}
}

func createCustomerEntityInt(c *gin.Context) {
	var customerEntityInt CustomerEntityInt
	c.BindJSON(&customerEntityInt)
	g.Create(&customerEntityInt)
	c.JSON(200, customerEntityInt)
}

func updateCustomerEntityInt(c *gin.Context) {
	var customerEntityInt CustomerEntityInt
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerEntityInt).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerEntityInt)
	g.Save(&customerEntityInt)
	c.JSON(200, customerEntityInt)
}
func deleteCustomerEntityInt(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEntityInt CustomerEntityInt
	d := g.Where("id = ?", id).Delete(&customerEntityInt)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
