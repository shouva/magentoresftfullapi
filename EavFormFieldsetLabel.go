package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavFormFieldsetLabel :
type EavFormFieldsetLabel struct {
	FieldsetId uint16 `gorm:"column:fieldset_id;primary_key" form:"fieldset_id;primary_key" json:"fieldset_id;primary_key"`
	StoreId    uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Label      string `gorm:"column:label" form:"label" json:"label"`
}

// TableName :
func (*EavFormFieldsetLabel) TableName() string {
	return "eav_form_fieldset_label"
}

// handler create
func initRoutersEavFormFieldsetLabel(r *gin.Engine, eavformfieldsetlabel string) {
	route := r.Group("eavformfieldsetlabel")
	route.GET("/", getEavFormFieldsetLabels)
	route.GET("/:id", getEavFormFieldsetLabel)
	route.POST("/", createEavFormFieldsetLabel)
	route.PUT("/:id", updateEavFormFieldsetLabel)
	route.DELETE("/:id", deleteEavFormFieldsetLabel)
}

func getEavFormFieldsetLabels(c *gin.Context) {
	var eavFormFieldsetLabels []EavFormFieldsetLabel
	if err := g.Find(&eavFormFieldsetLabels).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavFormFieldsetLabels)
	}
}

func getEavFormFieldsetLabel(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavFormFieldsetLabel EavFormFieldsetLabel
	if err := g.Where("id = ?", id).First(&eavFormFieldsetLabel).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavFormFieldsetLabel)
	}
}

func createEavFormFieldsetLabel(c *gin.Context) {
	var eavFormFieldsetLabel EavFormFieldsetLabel
	c.BindJSON(&eavFormFieldsetLabel)
	g.Create(&eavFormFieldsetLabel)
	c.JSON(200, eavFormFieldsetLabel)
}

func updateEavFormFieldsetLabel(c *gin.Context) {
	var eavFormFieldsetLabel EavFormFieldsetLabel
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavFormFieldsetLabel).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavFormFieldsetLabel)
	g.Save(&eavFormFieldsetLabel)
	c.JSON(200, eavFormFieldsetLabel)
}
func deleteEavFormFieldsetLabel(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavFormFieldsetLabel EavFormFieldsetLabel
	d := g.Where("id = ?", id).Delete(&eavFormFieldsetLabel)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
