package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CustomerEntityText :
type CustomerEntityText struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CustomerEntityText) TableName() string {
	return "customer_entity_text"
}

// handler create
func initRoutersCustomerEntityText(r *gin.Engine, customerentitytext string) {
	route := r.Group("customerentitytext")
	route.GET("/", getCustomerEntityTexts)
	route.GET("/:id", getCustomerEntityText)
	route.POST("/", createCustomerEntityText)
	route.PUT("/:id", updateCustomerEntityText)
	route.DELETE("/:id", deleteCustomerEntityText)
}

func getCustomerEntityTexts(c *gin.Context) {
	var customerEntityTexts []CustomerEntityText
	if err := g.Find(&customerEntityTexts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEntityTexts)
	}
}

func getCustomerEntityText(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEntityText CustomerEntityText
	if err := g.Where("id = ?", id).First(&customerEntityText).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customerEntityText)
	}
}

func createCustomerEntityText(c *gin.Context) {
	var customerEntityText CustomerEntityText
	c.BindJSON(&customerEntityText)
	g.Create(&customerEntityText)
	c.JSON(200, customerEntityText)
}

func updateCustomerEntityText(c *gin.Context) {
	var customerEntityText CustomerEntityText
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&customerEntityText).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customerEntityText)
	g.Save(&customerEntityText)
	c.JSON(200, customerEntityText)
}
func deleteCustomerEntityText(c *gin.Context) {
	id := c.Params.ByName("id")
	var customerEntityText CustomerEntityText
	d := g.Where("id = ?", id).Delete(&customerEntityText)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
