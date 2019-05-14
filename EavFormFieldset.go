package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavFormFieldset :
type EavFormFieldset struct {
	FieldsetId uint16 `gorm:"column:fieldset_id;primary_key" form:"fieldset_id;primary_key" json:"fieldset_id;primary_key"`
	TypeId     uint16 `gorm:"column:type_id" form:"type_id" json:"type_id"`
	Code       string `gorm:"column:code" form:"code" json:"code"`
	SortOrder  uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
}

// TableName :
func (*EavFormFieldset) TableName() string {
	return "eav_form_fieldset"
}

// handler create
func initRoutersEavFormFieldset(r *gin.Engine, eavformfieldset string) {
	route := r.Group("eavformfieldset")
	route.GET("/", getEavFormFieldsets)
	route.GET("/:id", getEavFormFieldset)
	route.POST("/", createEavFormFieldset)
	route.PUT("/:id", updateEavFormFieldset)
	route.DELETE("/:id", deleteEavFormFieldset)
}

func getEavFormFieldsets(c *gin.Context) {
	var eavFormFieldsets []EavFormFieldset
	if err := g.Find(&eavFormFieldsets).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavFormFieldsets)
	}
}

func getEavFormFieldset(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavFormFieldset EavFormFieldset
	if err := g.Where("id = ?", id).First(&eavFormFieldset).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavFormFieldset)
	}
}

func createEavFormFieldset(c *gin.Context) {
	var eavFormFieldset EavFormFieldset
	c.BindJSON(&eavFormFieldset)
	g.Create(&eavFormFieldset)
	c.JSON(200, eavFormFieldset)
}

func updateEavFormFieldset(c *gin.Context) {
	var eavFormFieldset EavFormFieldset
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavFormFieldset).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavFormFieldset)
	g.Save(&eavFormFieldset)
	c.JSON(200, eavFormFieldset)
}
func deleteEavFormFieldset(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavFormFieldset EavFormFieldset
	d := g.Where("id = ?", id).Delete(&eavFormFieldset)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
