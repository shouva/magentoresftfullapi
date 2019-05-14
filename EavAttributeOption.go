package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavAttributeOption :
type EavAttributeOption struct {
	OptionId    uint16 `gorm:"column:option_id;primary_key" form:"option_id;primary_key" json:"option_id;primary_key"`
	AttributeId uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	SortOrder   uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
}

// TableName :
func (*EavAttributeOption) TableName() string {
	return "eav_attribute_option"
}

// handler create
func initRoutersEavAttributeOption(r *gin.Engine, eavattributeoption string) {
	route := r.Group("eavattributeoption")
	route.GET("/", getEavAttributeOptions)
	route.GET("/:id", getEavAttributeOption)
	route.POST("/", createEavAttributeOption)
	route.PUT("/:id", updateEavAttributeOption)
	route.DELETE("/:id", deleteEavAttributeOption)
}

func getEavAttributeOptions(c *gin.Context) {
	var eavAttributeOptions []EavAttributeOption
	if err := g.Find(&eavAttributeOptions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavAttributeOptions)
	}
}

func getEavAttributeOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavAttributeOption EavAttributeOption
	if err := g.Where("id = ?", id).First(&eavAttributeOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavAttributeOption)
	}
}

func createEavAttributeOption(c *gin.Context) {
	var eavAttributeOption EavAttributeOption
	c.BindJSON(&eavAttributeOption)
	g.Create(&eavAttributeOption)
	c.JSON(200, eavAttributeOption)
}

func updateEavAttributeOption(c *gin.Context) {
	var eavAttributeOption EavAttributeOption
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavAttributeOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavAttributeOption)
	g.Save(&eavAttributeOption)
	c.JSON(200, eavAttributeOption)
}
func deleteEavAttributeOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavAttributeOption EavAttributeOption
	d := g.Where("id = ?", id).Delete(&eavAttributeOption)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
