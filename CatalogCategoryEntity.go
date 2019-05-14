package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryEntity :
type CatalogCategoryEntity struct {
	EntityId       uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	EntityTypeId   uint16     `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeSetId uint16     `gorm:"column:attribute_set_id" form:"attribute_set_id" json:"attribute_set_id"`
	ParentId       uint16     `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	CreatedAt      *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	Path           string     `gorm:"column:path" form:"path" json:"path"`
	Position       uint16     `gorm:"column:position" form:"position" json:"position"`
	Level          uint16     `gorm:"column:level" form:"level" json:"level"`
	ChildrenCount  uint16     `gorm:"column:children_count" form:"children_count" json:"children_count"`
}

// TableName :
func (*CatalogCategoryEntity) TableName() string {
	return "catalog_category_entity"
}

// handler create
func initRoutersCatalogCategoryEntity(r *gin.Engine, catalogcategoryentity string) {
	route := r.Group("catalogcategoryentity")
	route.GET("/", getCatalogCategoryEntitys)
	route.GET("/:id", getCatalogCategoryEntity)
	route.POST("/", createCatalogCategoryEntity)
	route.PUT("/:id", updateCatalogCategoryEntity)
	route.DELETE("/:id", deleteCatalogCategoryEntity)
}

func getCatalogCategoryEntitys(c *gin.Context) {
	var catalogCategoryEntitys []CatalogCategoryEntity
	if err := g.Find(&catalogCategoryEntitys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryEntitys)
	}
}

func getCatalogCategoryEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryEntity CatalogCategoryEntity
	if err := g.Where("id = ?", id).First(&catalogCategoryEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryEntity)
	}
}

func createCatalogCategoryEntity(c *gin.Context) {
	var catalogCategoryEntity CatalogCategoryEntity
	c.BindJSON(&catalogCategoryEntity)
	g.Create(&catalogCategoryEntity)
	c.JSON(200, catalogCategoryEntity)
}

func updateCatalogCategoryEntity(c *gin.Context) {
	var catalogCategoryEntity CatalogCategoryEntity
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryEntity).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryEntity)
	g.Save(&catalogCategoryEntity)
	c.JSON(200, catalogCategoryEntity)
}
func deleteCatalogCategoryEntity(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryEntity CatalogCategoryEntity
	d := g.Where("id = ?", id).Delete(&catalogCategoryEntity)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
