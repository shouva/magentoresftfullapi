package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavAttributeGroup :
type EavAttributeGroup struct {
	AttributeGroupId   uint16 `gorm:"column:attribute_group_id;primary_key" form:"attribute_group_id;primary_key" json:"attribute_group_id;primary_key"`
	AttributeSetId     uint16 `gorm:"column:attribute_set_id" form:"attribute_set_id" json:"attribute_set_id"`
	AttributeGroupName string `gorm:"column:attribute_group_name" form:"attribute_group_name" json:"attribute_group_name"`
	SortOrder          uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
	DefaultId          uint16 `gorm:"column:default_id" form:"default_id" json:"default_id"`
}

// TableName :
func (*EavAttributeGroup) TableName() string {
	return "eav_attribute_group"
}

// handler create
func initRoutersEavAttributeGroup(r *gin.Engine, eavattributegroup string) {
	route := r.Group("eavattributegroup")
	route.GET("/", getEavAttributeGroups)
	route.GET("/:id", getEavAttributeGroup)
	route.POST("/", createEavAttributeGroup)
	route.PUT("/:id", updateEavAttributeGroup)
	route.DELETE("/:id", deleteEavAttributeGroup)
}

func getEavAttributeGroups(c *gin.Context) {
	var eavAttributeGroups []EavAttributeGroup
	if err := g.Find(&eavAttributeGroups).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavAttributeGroups)
	}
}

func getEavAttributeGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavAttributeGroup EavAttributeGroup
	if err := g.Where("id = ?", id).First(&eavAttributeGroup).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavAttributeGroup)
	}
}

func createEavAttributeGroup(c *gin.Context) {
	var eavAttributeGroup EavAttributeGroup
	c.BindJSON(&eavAttributeGroup)
	g.Create(&eavAttributeGroup)
	c.JSON(200, eavAttributeGroup)
}

func updateEavAttributeGroup(c *gin.Context) {
	var eavAttributeGroup EavAttributeGroup
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavAttributeGroup).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavAttributeGroup)
	g.Save(&eavAttributeGroup)
	c.JSON(200, eavAttributeGroup)
}
func deleteEavAttributeGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavAttributeGroup EavAttributeGroup
	d := g.Where("id = ?", id).Delete(&eavAttributeGroup)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
