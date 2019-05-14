package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavEntityText :
type EavEntityText struct {
	ValueId      uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeId  uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId      uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	EntityId     uint16 `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Value        string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*EavEntityText) TableName() string {
	return "eav_entity_text"
}

// handler create
func initRoutersEavEntityText(r *gin.Engine, eaventitytext string) {
	route := r.Group("eaventitytext")
	route.GET("/", getEavEntityTexts)
	route.GET("/:id", getEavEntityText)
	route.POST("/", createEavEntityText)
	route.PUT("/:id", updateEavEntityText)
	route.DELETE("/:id", deleteEavEntityText)
}

func getEavEntityTexts(c *gin.Context) {
	var eavEntityTexts []EavEntityText
	if err := g.Find(&eavEntityTexts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityTexts)
	}
}

func getEavEntityText(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityText EavEntityText
	if err := g.Where("id = ?", id).First(&eavEntityText).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityText)
	}
}

func createEavEntityText(c *gin.Context) {
	var eavEntityText EavEntityText
	c.BindJSON(&eavEntityText)
	g.Create(&eavEntityText)
	c.JSON(200, eavEntityText)
}

func updateEavEntityText(c *gin.Context) {
	var eavEntityText EavEntityText
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavEntityText).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavEntityText)
	g.Save(&eavEntityText)
	c.JSON(200, eavEntityText)
}
func deleteEavEntityText(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityText EavEntityText
	d := g.Where("id = ?", id).Delete(&eavEntityText)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
