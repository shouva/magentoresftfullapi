package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductRelation :
type CatalogProductRelation struct {
	ParentId uint16 `gorm:"column:parent_id;primary_key" form:"parent_id;primary_key" json:"parent_id;primary_key"`
	ChildId  uint16 `gorm:"column:child_id;primary_key" form:"child_id;primary_key" json:"child_id;primary_key"`
}

// TableName :
func (*CatalogProductRelation) TableName() string {
	return "catalog_product_relation"
}

// handler create
func initRoutersCatalogProductRelation(r *gin.Engine, catalogproductrelation string) {
	route := r.Group("catalogproductrelation")
	route.GET("/", getCatalogProductRelations)
	route.GET("/:id", getCatalogProductRelation)
	route.POST("/", createCatalogProductRelation)
	route.PUT("/:id", updateCatalogProductRelation)
	route.DELETE("/:id", deleteCatalogProductRelation)
}

func getCatalogProductRelations(c *gin.Context) {
	var catalogProductRelations []CatalogProductRelation
	if err := g.Find(&catalogProductRelations).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductRelations)
	}
}

func getCatalogProductRelation(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductRelation CatalogProductRelation
	if err := g.Where("id = ?", id).First(&catalogProductRelation).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductRelation)
	}
}

func createCatalogProductRelation(c *gin.Context) {
	var catalogProductRelation CatalogProductRelation
	c.BindJSON(&catalogProductRelation)
	g.Create(&catalogProductRelation)
	c.JSON(200, catalogProductRelation)
}

func updateCatalogProductRelation(c *gin.Context) {
	var catalogProductRelation CatalogProductRelation
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductRelation).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductRelation)
	g.Save(&catalogProductRelation)
	c.JSON(200, catalogProductRelation)
}
func deleteCatalogProductRelation(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductRelation CatalogProductRelation
	d := g.Where("id = ?", id).Delete(&catalogProductRelation)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
