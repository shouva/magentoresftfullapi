package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// EavEntity :
type EavEntity struct {
	EntityId       uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	EntityTypeId   uint16     `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeSetId uint16     `gorm:"column:attribute_set_id" form:"attribute_set_id" json:"attribute_set_id"`
	IncrementId    string     `gorm:"column:increment_id" form:"increment_id" json:"increment_id"`
	ParentId       uint16     `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	StoreId        uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	CreatedAt      *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	IsActive       uint16     `gorm:"column:is_active" form:"is_active" json:"is_active"`
}

// TableName :
func (*EavEntity) TableName() string {
	return "eav_entity"
}

// handler create
func initRoutersEavEntity(r *gin.Engine, eaventity string) {
	route := r.Group("eaventity")
	route.GET("/", getEavEntitys)
	route.GET("/:id", getEavEntity)
	route.POST("/", createEavEntity)
	route.PUT("/:id", updateEavEntity)
	route.DELETE("/:id", deleteEavEntity)
}

func getEavEntitys(c *gin.Context) {
	var eavEntitys []EavEntity
	if err := g.Find(&eavEntitys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntitys)
	}
}

func getEavEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntity EavEntity
	if err := g.Where("id = ?", id).First(&eavEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntity)
	}
}

func createEavEntity(c *gin.Context) {
	var eavEntity EavEntity
	c.BindJSON(&eavEntity)
	g.Create(&eavEntity)
	c.JSON(200, eavEntity)
}

func updateEavEntity(c *gin.Context) {
	var eavEntity EavEntity
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavEntity)
	g.Save(&eavEntity)
	c.JSON(200, eavEntity)
}
func deleteEavEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntity EavEntity
	d := g.Where("id = ?", id).Delete(&eavEntity)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
