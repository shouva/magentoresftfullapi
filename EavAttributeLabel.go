package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavAttributeLabel :
type EavAttributeLabel struct {
	AttributeLabelId uint16 `gorm:"column:attribute_label_id;primary_key" form:"attribute_label_id;primary_key" json:"attribute_label_id;primary_key"`
	AttributeId      uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	StoreId          uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Value            string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*EavAttributeLabel) TableName() string {
	return "eav_attribute_label"
}

// handler create
func initRoutersEavAttributeLabel(r *gin.Engine, eavattributelabel string) {
	route := r.Group("eavattributelabel")
	route.GET("/", getEavAttributeLabels)
	route.GET("/:id", getEavAttributeLabel)
	route.POST("/", createEavAttributeLabel)
	route.PUT("/:id", updateEavAttributeLabel)
	route.DELETE("/:id", deleteEavAttributeLabel)
}

func getEavAttributeLabels(c *gin.Context) {
	var eavAttributeLabels []EavAttributeLabel
	if err := g.Find(&eavAttributeLabels).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavAttributeLabels)
	}
}

func getEavAttributeLabel(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavAttributeLabel EavAttributeLabel
	if err := g.Where("id = ?", id).First(&eavAttributeLabel).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavAttributeLabel)
	}
}

func createEavAttributeLabel(c *gin.Context) {
	var eavAttributeLabel EavAttributeLabel
	c.BindJSON(&eavAttributeLabel)
	g.Create(&eavAttributeLabel)
	c.JSON(200, eavAttributeLabel)
}

func updateEavAttributeLabel(c *gin.Context) {
	var eavAttributeLabel EavAttributeLabel
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavAttributeLabel).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavAttributeLabel)
	g.Save(&eavAttributeLabel)
	c.JSON(200, eavAttributeLabel)
}
func deleteEavAttributeLabel(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavAttributeLabel EavAttributeLabel
	d := g.Where("id = ?", id).Delete(&eavAttributeLabel)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
