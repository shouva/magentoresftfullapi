package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavEntityInt :
type EavEntityInt struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        uint16 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*EavEntityInt) TableName() string {
	return "eav_entity_int"
}

// handler create
func initRoutersEavEntityInt(r *gin.Engine, eaventityint string) {
	route := r.Group("eaventityint")
	route.GET("/", getEavEntityInts)
	route.GET("/:id", getEavEntityInt)
	route.POST("/", createEavEntityInt)
	route.PUT("/:id", updateEavEntityInt)
	route.DELETE("/:id", deleteEavEntityInt)
}

func getEavEntityInts(c *gin.Context) {
	var eavEntityInts []EavEntityInt
	if err := g.Find(&eavEntityInts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityInts)
	}
}

func getEavEntityInt(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityInt EavEntityInt
	if err := g.Where("id = ?", id).First(&eavEntityInt).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityInt)
	}
}

func createEavEntityInt(c *gin.Context) {
	var eavEntityInt EavEntityInt
	c.BindJSON(&eavEntityInt)
	g.Create(&eavEntityInt)
	c.JSON(200, eavEntityInt)
}

func updateEavEntityInt(c *gin.Context) {
	var eavEntityInt EavEntityInt
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavEntityInt).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavEntityInt)
	g.Save(&eavEntityInt)
	c.JSON(200, eavEntityInt)
}
func deleteEavEntityInt(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityInt EavEntityInt
	d := g.Where("id = ?", id).Delete(&eavEntityInt)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
