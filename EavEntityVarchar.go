package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavEntityVarchar :
type EavEntityVarchar struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*EavEntityVarchar) TableName() string {
	return "eav_entity_varchar"
}

// handler create
func initRoutersEavEntityVarchar(r *gin.Engine, eaventityvarchar string) {
	route := r.Group("eaventityvarchar")
	route.GET("/", getEavEntityVarchars)
	route.GET("/:id", getEavEntityVarchar)
	route.POST("/", createEavEntityVarchar)
	route.PUT("/:id", updateEavEntityVarchar)
	route.DELETE("/:id", deleteEavEntityVarchar)
}

func getEavEntityVarchars(c *gin.Context) {
	var eavEntityVarchars []EavEntityVarchar
	if err := g.Find(&eavEntityVarchars).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityVarchars)
	}
}

func getEavEntityVarchar(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityVarchar EavEntityVarchar
	if err := g.Where("id = ?", id).First(&eavEntityVarchar).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityVarchar)
	}
}

func createEavEntityVarchar(c *gin.Context) {
	var eavEntityVarchar EavEntityVarchar
	c.BindJSON(&eavEntityVarchar)
	g.Create(&eavEntityVarchar)
	c.JSON(200, eavEntityVarchar)
}

func updateEavEntityVarchar(c *gin.Context) {
	var eavEntityVarchar EavEntityVarchar
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavEntityVarchar).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavEntityVarchar)
	g.Save(&eavEntityVarchar)
	c.JSON(200, eavEntityVarchar)
}
func deleteEavEntityVarchar(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityVarchar EavEntityVarchar
	d := g.Where("id = ?", id).Delete(&eavEntityVarchar)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
