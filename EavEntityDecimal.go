package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavEntityDecimal :
type EavEntityDecimal struct {
	ValueId      uint16  `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16  `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16  `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16  `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16  `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        float64 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*EavEntityDecimal) TableName() string {
	return "eav_entity_decimal"
}

// handler create
func initRoutersEavEntityDecimal(r *gin.Engine, eaventitydecimal string) {
	route := r.Group("eaventitydecimal")
	route.GET("/", getEavEntityDecimals)
	route.GET("/:id", getEavEntityDecimal)
	route.POST("/", createEavEntityDecimal)
	route.PUT("/:id", updateEavEntityDecimal)
	route.DELETE("/:id", deleteEavEntityDecimal)
}

func getEavEntityDecimals(c *gin.Context) {
	var eavEntityDecimals []EavEntityDecimal
	if err := g.Find(&eavEntityDecimals).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityDecimals)
	}
}

func getEavEntityDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityDecimal EavEntityDecimal
	if err := g.Where("id = ?", id).First(&eavEntityDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityDecimal)
	}
}

func createEavEntityDecimal(c *gin.Context) {
	var eavEntityDecimal EavEntityDecimal
	c.BindJSON(&eavEntityDecimal)
	g.Create(&eavEntityDecimal)
	c.JSON(200, eavEntityDecimal)
}

func updateEavEntityDecimal(c *gin.Context) {
	var eavEntityDecimal EavEntityDecimal
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavEntityDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavEntityDecimal)
	g.Save(&eavEntityDecimal)
	c.JSON(200, eavEntityDecimal)
}
func deleteEavEntityDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityDecimal EavEntityDecimal
	d := g.Where("id = ?", id).Delete(&eavEntityDecimal)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
