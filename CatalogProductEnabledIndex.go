package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductEnabledIndex :
type CatalogProductEnabledIndex struct {
	ProductId  uint16 `gorm:"column:product_id;primary_key" form:"product_id;primary_key" json:"product_id;primary_key"`
	StoreId    uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Visibility uint16 `gorm:"column:visibility" form:"visibility" json:"visibility"`
}

// TableName :
func (*CatalogProductEnabledIndex) TableName() string {
	return "catalog_product_enabled_index"
}

// handler create
func initRoutersCatalogProductEnabledIndex(r *gin.Engine, catalogproductenabledindex string) {
	route := r.Group("catalogproductenabledindex")
	route.GET("/", getCatalogProductEnabledIndexs)
	route.GET("/:id", getCatalogProductEnabledIndex)
	route.POST("/", createCatalogProductEnabledIndex)
	route.PUT("/:id", updateCatalogProductEnabledIndex)
	route.DELETE("/:id", deleteCatalogProductEnabledIndex)
}

func getCatalogProductEnabledIndexs(c *gin.Context) {
	var catalogProductEnabledIndexs []CatalogProductEnabledIndex
	if err := g.Find(&catalogProductEnabledIndexs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEnabledIndexs)
	}
}

func getCatalogProductEnabledIndex(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEnabledIndex CatalogProductEnabledIndex
	if err := g.Where("id = ?", id).First(&catalogProductEnabledIndex).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductEnabledIndex)
	}
}

func createCatalogProductEnabledIndex(c *gin.Context) {
	var catalogProductEnabledIndex CatalogProductEnabledIndex
	c.BindJSON(&catalogProductEnabledIndex)
	g.Create(&catalogProductEnabledIndex)
	c.JSON(200, catalogProductEnabledIndex)
}

func updateCatalogProductEnabledIndex(c *gin.Context) {
	var catalogProductEnabledIndex CatalogProductEnabledIndex
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductEnabledIndex).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductEnabledIndex)
	g.Save(&catalogProductEnabledIndex)
	c.JSON(200, catalogProductEnabledIndex)
}
func deleteCatalogProductEnabledIndex(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductEnabledIndex CatalogProductEnabledIndex
	d := g.Where("id = ?", id).Delete(&catalogProductEnabledIndex)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
