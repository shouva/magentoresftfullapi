package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavEntityAttribute :
type EavEntityAttribute struct {
	EntityAttributeId uint16 `gorm:"column:entity_attribute_id;primary_key" form:"entity_attribute_id;primary_key" json:"entity_attribute_id;primary_key"`
	EntityTypeId      uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeSetId    uint16 `gorm:"column:attribute_set_id" form:"attribute_set_id" json:"attribute_set_id"`
	AttributeGroupId  uint16 `gorm:"column:attribute_group_id" form:"attribute_group_id" json:"attribute_group_id"`
	AttributeId       uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	SortOrder         uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
}

// TableName :
func (*EavEntityAttribute) TableName() string {
	return "eav_entity_attribute"
}

// handler create
func initRoutersEavEntityAttribute(r *gin.Engine, eaventityattribute string) {
	route := r.Group("eaventityattribute")
	route.GET("/", getEavEntityAttributes)
	route.GET("/:id", getEavEntityAttribute)
	route.POST("/", createEavEntityAttribute)
	route.PUT("/:id", updateEavEntityAttribute)
	route.DELETE("/:id", deleteEavEntityAttribute)
}

func getEavEntityAttributes(c *gin.Context) {
	var eavEntityAttributes []EavEntityAttribute
	if err := g.Find(&eavEntityAttributes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityAttributes)
	}
}

func getEavEntityAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityAttribute EavEntityAttribute
	if err := g.Where("id = ?", id).First(&eavEntityAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityAttribute)
	}
}

func createEavEntityAttribute(c *gin.Context) {
	var eavEntityAttribute EavEntityAttribute
	c.BindJSON(&eavEntityAttribute)
	g.Create(&eavEntityAttribute)
	c.JSON(200, eavEntityAttribute)
}

func updateEavEntityAttribute(c *gin.Context) {
	var eavEntityAttribute EavEntityAttribute
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavEntityAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavEntityAttribute)
	g.Save(&eavEntityAttribute)
	c.JSON(200, eavEntityAttribute)
}
func deleteEavEntityAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityAttribute EavEntityAttribute
	d := g.Where("id = ?", id).Delete(&eavEntityAttribute)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
