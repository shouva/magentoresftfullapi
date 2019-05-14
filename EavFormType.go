package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavFormType :
type EavFormType struct {
	TypeId   uint16 `gorm:"column:type_id;primary_key" form:"type_id;primary_key" json:"type_id;primary_key"`
	Code     string `gorm:"column:code" form:"code" json:"code"`
	Label    string `gorm:"column:label" form:"label" json:"label"`
	IsSystem uint16 `gorm:"column:is_system" form:"is_system" json:"is_system"`
	Theme    string `gorm:"column:theme" form:"theme" json:"theme"`
	StoreId  uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
}

// TableName :
func (*EavFormType) TableName() string {
	return "eav_form_type"
}

// handler create
func initRoutersEavFormType(r *gin.Engine, eavformtype string) {
	route := r.Group("eavformtype")
	route.GET("/", getEavFormTypes)
	route.GET("/:id", getEavFormType)
	route.POST("/", createEavFormType)
	route.PUT("/:id", updateEavFormType)
	route.DELETE("/:id", deleteEavFormType)
}

func getEavFormTypes(c *gin.Context) {
	var eavFormTypes []EavFormType
	if err := g.Find(&eavFormTypes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavFormTypes)
	}
}

func getEavFormType(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavFormType EavFormType
	if err := g.Where("id = ?", id).First(&eavFormType).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavFormType)
	}
}

func createEavFormType(c *gin.Context) {
	var eavFormType EavFormType
	c.BindJSON(&eavFormType)
	g.Create(&eavFormType)
	c.JSON(200, eavFormType)
}

func updateEavFormType(c *gin.Context) {
	var eavFormType EavFormType
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavFormType).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavFormType)
	g.Save(&eavFormType)
	c.JSON(200, eavFormType)
}
func deleteEavFormType(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavFormType EavFormType
	d := g.Where("id = ?", id).Delete(&eavFormType)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
