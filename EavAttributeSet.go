package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavAttributeSet :
type EavAttributeSet struct {
	AttributeSetId   uint16 `gorm:"column:attribute_set_id;primary_key" form:"attribute_set_id;primary_key" json:"attribute_set_id;primary_key"`
	EntityTypeId     uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeSetName string `gorm:"column:attribute_set_name" form:"attribute_set_name" json:"attribute_set_name"`
	SortOrder        uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
}

// TableName :
func (*EavAttributeSet) TableName() string {
	return "eav_attribute_set"
}

// handler create
func initRoutersEavAttributeSet(r *gin.Engine, eavattributeset string) {
	route := r.Group("eavattributeset")
	route.GET("/", getEavAttributeSets)
	route.GET("/:id", getEavAttributeSet)
	route.POST("/", createEavAttributeSet)
	route.PUT("/:id", updateEavAttributeSet)
	route.DELETE("/:id", deleteEavAttributeSet)
}

func getEavAttributeSets(c *gin.Context) {
	var eavAttributeSets []EavAttributeSet
	if err := g.Find(&eavAttributeSets).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavAttributeSets)
	}
}

func getEavAttributeSet(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavAttributeSet EavAttributeSet
	if err := g.Where("id = ?", id).First(&eavAttributeSet).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavAttributeSet)
	}
}

func createEavAttributeSet(c *gin.Context) {
	var eavAttributeSet EavAttributeSet
	c.BindJSON(&eavAttributeSet)
	g.Create(&eavAttributeSet)
	c.JSON(200, eavAttributeSet)
}

func updateEavAttributeSet(c *gin.Context) {
	var eavAttributeSet EavAttributeSet
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavAttributeSet).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavAttributeSet)
	g.Save(&eavAttributeSet)
	c.JSON(200, eavAttributeSet)
}
func deleteEavAttributeSet(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavAttributeSet EavAttributeSet
	d := g.Where("id = ?", id).Delete(&eavAttributeSet)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
