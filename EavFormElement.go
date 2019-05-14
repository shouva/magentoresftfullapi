package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavFormElement :
type EavFormElement struct {
	ElementId   uint16 `gorm:"column:element_id;primary_key" form:"element_id;primary_key" json:"element_id;primary_key"`
	TypeId      uint16 `gorm:"column:type_id" form:"type_id" json:"type_id"`
	FieldsetId  uint16 `gorm:"column:fieldset_id" form:"fieldset_id" json:"fieldset_id"`
	AttributeId uint16 `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	SortOrder   uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
}

// TableName :
func (*EavFormElement) TableName() string {
	return "eav_form_element"
}

// handler create
func initRoutersEavFormElement(r *gin.Engine, eavformelement string) {
	route := r.Group("eavformelement")
	route.GET("/", getEavFormElements)
	route.GET("/:id", getEavFormElement)
	route.POST("/", createEavFormElement)
	route.PUT("/:id", updateEavFormElement)
	route.DELETE("/:id", deleteEavFormElement)
}

func getEavFormElements(c *gin.Context) {
	var eavFormElements []EavFormElement
	if err := g.Find(&eavFormElements).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavFormElements)
	}
}

func getEavFormElement(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavFormElement EavFormElement
	if err := g.Where("id = ?", id).First(&eavFormElement).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavFormElement)
	}
}

func createEavFormElement(c *gin.Context) {
	var eavFormElement EavFormElement
	c.BindJSON(&eavFormElement)
	g.Create(&eavFormElement)
	c.JSON(200, eavFormElement)
}

func updateEavFormElement(c *gin.Context) {
	var eavFormElement EavFormElement
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavFormElement).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavFormElement)
	g.Save(&eavFormElement)
	c.JSON(200, eavFormElement)
}
func deleteEavFormElement(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavFormElement EavFormElement
	d := g.Where("id = ?", id).Delete(&eavFormElement)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
