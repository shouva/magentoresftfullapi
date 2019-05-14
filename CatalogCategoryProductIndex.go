package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryProductIndex :
type CatalogCategoryProductIndex struct {
	CategoryId uint16 `gorm:"column:category_id;primary_key" form:"category_id;primary_key" json:"category_id;primary_key"`
	ProductId  uint16 `gorm:"column:product_id;primary_key" form:"product_id;primary_key" json:"product_id;primary_key"`
	Position   uint16 `gorm:"column:position" form:"position" json:"position"`
	IsParent   uint16 `gorm:"column:is_parent" form:"is_parent" json:"is_parent"`
	StoreId    uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Visibility uint16 `gorm:"column:visibility" form:"visibility" json:"visibility"`
}

// TableName :
func (*CatalogCategoryProductIndex) TableName() string {
	return "catalog_category_product_index"
}

// handler create
func initRoutersCatalogCategoryProductIndex(r *gin.Engine, catalogcategoryproductindex string) {
	route := r.Group("catalogcategoryproductindex")
	route.GET("/", getCatalogCategoryProductIndexs)
	route.GET("/:id", getCatalogCategoryProductIndex)
	route.POST("/", createCatalogCategoryProductIndex)
	route.PUT("/:id", updateCatalogCategoryProductIndex)
	route.DELETE("/:id", deleteCatalogCategoryProductIndex)
}

func getCatalogCategoryProductIndexs(c *gin.Context) {
	var catalogCategoryProductIndexs []CatalogCategoryProductIndex
	if err := g.Find(&catalogCategoryProductIndexs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryProductIndexs)
	}
}

func getCatalogCategoryProductIndex(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryProductIndex CatalogCategoryProductIndex
	if err := g.Where("id = ?", id).First(&catalogCategoryProductIndex).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryProductIndex)
	}
}

func createCatalogCategoryProductIndex(c *gin.Context) {
	var catalogCategoryProductIndex CatalogCategoryProductIndex
	c.BindJSON(&catalogCategoryProductIndex)
	g.Create(&catalogCategoryProductIndex)
	c.JSON(200, catalogCategoryProductIndex)
}

func updateCatalogCategoryProductIndex(c *gin.Context) {
	var catalogCategoryProductIndex CatalogCategoryProductIndex
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryProductIndex).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryProductIndex)
	g.Save(&catalogCategoryProductIndex)
	c.JSON(200, catalogCategoryProductIndex)
}
func deleteCatalogCategoryProductIndex(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryProductIndex CatalogCategoryProductIndex
	d := g.Where("id = ?", id).Delete(&catalogCategoryProductIndex)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
