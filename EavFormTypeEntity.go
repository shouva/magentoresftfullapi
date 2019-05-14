package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavFormTypeEntity :
type EavFormTypeEntity struct {
	TypeId       uint16 `gorm:"column:type_id;primary_key" form:"type_id;primary_key" json:"type_id;primary_key"`
	EntityTypeId uint16 `gorm:"column:entity_type_id;primary_key" form:"entity_type_id;primary_key" json:"entity_type_id;primary_key"`
}

// TableName :
func (*EavFormTypeEntity) TableName() string {
	return "eav_form_type_entity"
}

// handler create
func initRoutersEavFormTypeEntity(r *gin.Engine, eavformtypeentity string) {
	route := r.Group("eavformtypeentity")
	route.GET("/", getEavFormTypeEntitys)
	route.GET("/:id", getEavFormTypeEntity)
	route.POST("/", createEavFormTypeEntity)
	route.PUT("/:id", updateEavFormTypeEntity)
	route.DELETE("/:id", deleteEavFormTypeEntity)
}

func getEavFormTypeEntitys(c *gin.Context) {
	var eavFormTypeEntitys []EavFormTypeEntity
	if err := g.Find(&eavFormTypeEntitys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavFormTypeEntitys)
	}
}

func getEavFormTypeEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavFormTypeEntity EavFormTypeEntity
	if err := g.Where("id = ?", id).First(&eavFormTypeEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavFormTypeEntity)
	}
}

func createEavFormTypeEntity(c *gin.Context) {
	var eavFormTypeEntity EavFormTypeEntity
	c.BindJSON(&eavFormTypeEntity)
	g.Create(&eavFormTypeEntity)
	c.JSON(200, eavFormTypeEntity)
}

func updateEavFormTypeEntity(c *gin.Context) {
	var eavFormTypeEntity EavFormTypeEntity
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavFormTypeEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavFormTypeEntity)
	g.Save(&eavFormTypeEntity)
	c.JSON(200, eavFormTypeEntity)
}
func deleteEavFormTypeEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavFormTypeEntity EavFormTypeEntity
	d := g.Where("id = ?", id).Delete(&eavFormTypeEntity)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
