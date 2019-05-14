package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CustomerEntityVarchar :
type CustomerEntityVarchar struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CustomerEntityVarchar) TableName() string {
	return "customer_entity_varchar"
}

// handler create
func initRoutersCustomerEntityVarchar(r *gin.Engine, customerentityvarchar string) {
	route := r.Group("customerentityvarchar")
	route.GET("/", getCustomerEntityVarchars)
	route.GET("/:id", getCustomerEntityVarchar)
	route.POST("/", createCustomerEntityVarchar)
	route.PUT("/:id", updateCustomerEntityVarchar)
	route.DELETE("/:id", deleteCustomerEntityVarchar)
}

func getCustomerEntityVarchars(c *gin.Context) {
	var customerEntityVarchars []CustomerEntityVarchar
	if err := g.Find(&customerEntityVarchars).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEntityVarchars)
	}
}

func getCustomerEntityVarchar(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEntityVarchar CustomerEntityVarchar
	if err := g.Where("id = ?", id).First(&customerEntityVarchar).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEntityVarchar)
	}
}

func createCustomerEntityVarchar(c *gin.Context) {
	var customerEntityVarchar CustomerEntityVarchar
	c.BindJSON(&customerEntityVarchar)
	g.Create(&customerEntityVarchar)
	c.JSON(200, customerEntityVarchar)
}

func updateCustomerEntityVarchar(c *gin.Context) {
	var customerEntityVarchar CustomerEntityVarchar
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerEntityVarchar).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerEntityVarchar)
	g.Save(&customerEntityVarchar)
	c.JSON(200, customerEntityVarchar)
}
func deleteCustomerEntityVarchar(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEntityVarchar CustomerEntityVarchar
	d := g.Where("id = ?", id).Delete(&customerEntityVarchar)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
