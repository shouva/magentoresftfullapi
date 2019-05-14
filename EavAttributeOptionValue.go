package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavAttributeOptionValue :
type EavAttributeOptionValue struct {
	ValueId  uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	OptionId uint16 `gorm:"column:option_id" form:"option_id" json:"option_id"`
	StoreId  uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Value    string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*EavAttributeOptionValue) TableName() string {
	return "eav_attribute_option_value"
}

// handler create
func initRoutersEavAttributeOptionValue(r *gin.Engine, eavattributeoptionvalue string) {
	route := r.Group("eavattributeoptionvalue")
	route.GET("/", getEavAttributeOptionValues)
	route.GET("/:id", getEavAttributeOptionValue)
	route.POST("/", createEavAttributeOptionValue)
	route.PUT("/:id", updateEavAttributeOptionValue)
	route.DELETE("/:id", deleteEavAttributeOptionValue)
}

func getEavAttributeOptionValues(c *gin.Context) {
	var eavAttributeOptionValues []EavAttributeOptionValue
	if err := g.Find(&eavAttributeOptionValues).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavAttributeOptionValues)
	}
}

func getEavAttributeOptionValue(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavAttributeOptionValue EavAttributeOptionValue
	if err := g.Where("id = ?", id).First(&eavAttributeOptionValue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavAttributeOptionValue)
	}
}

func createEavAttributeOptionValue(c *gin.Context) {
	var eavAttributeOptionValue EavAttributeOptionValue
	c.BindJSON(&eavAttributeOptionValue)
	g.Create(&eavAttributeOptionValue)
	c.JSON(200, eavAttributeOptionValue)
}

func updateEavAttributeOptionValue(c *gin.Context) {
	var eavAttributeOptionValue EavAttributeOptionValue
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavAttributeOptionValue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavAttributeOptionValue)
	g.Save(&eavAttributeOptionValue)
	c.JSON(200, eavAttributeOptionValue)
}
func deleteEavAttributeOptionValue(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavAttributeOptionValue EavAttributeOptionValue
	d := g.Where("id = ?", id).Delete(&eavAttributeOptionValue)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
