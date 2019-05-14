package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryProductIndexIdx :
type CatalogCategoryProductIndexIdx struct {
	CategoryId uint16 `gorm:"column:category_id" form:"category_id" json:"category_id"`
	ProductId  uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	Position   uint16 `gorm:"column:position" form:"position" json:"position"`
	IsParent   uint16 `gorm:"column:is_parent" form:"is_parent" json:"is_parent"`
	StoreId    uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Visibility uint16 `gorm:"column:visibility" form:"visibility" json:"visibility"`
}

// TableName :
func (*CatalogCategoryProductIndexIdx) TableName() string {
	return "catalog_category_product_index_idx"
}

// handler create
func initRoutersCatalogCategoryProductIndexIdx(r *gin.Engine, catalogcategoryproductindexidx string) {
	route := r.Group("catalogcategoryproductindexidx")
	route.GET("/", getCatalogCategoryProductIndexIdxs)
	route.GET("/:id", getCatalogCategoryProductIndexIdx)
	route.POST("/", createCatalogCategoryProductIndexIdx)
	route.PUT("/:id", updateCatalogCategoryProductIndexIdx)
	route.DELETE("/:id", deleteCatalogCategoryProductIndexIdx)
}

func getCatalogCategoryProductIndexIdxs(c *gin.Context) {
	var catalogCategoryProductIndexIdxs []CatalogCategoryProductIndexIdx
	if err := g.Find(&catalogCategoryProductIndexIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryProductIndexIdxs)
	}
}

func getCatalogCategoryProductIndexIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryProductIndexIdx CatalogCategoryProductIndexIdx
	if err := g.Where("id = ?", id).First(&catalogCategoryProductIndexIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryProductIndexIdx)
	}
}

func createCatalogCategoryProductIndexIdx(c *gin.Context) {
	var catalogCategoryProductIndexIdx CatalogCategoryProductIndexIdx
	c.BindJSON(&catalogCategoryProductIndexIdx)
	g.Create(&catalogCategoryProductIndexIdx)
	c.JSON(200, catalogCategoryProductIndexIdx)
}

func updateCatalogCategoryProductIndexIdx(c *gin.Context) {
	var catalogCategoryProductIndexIdx CatalogCategoryProductIndexIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryProductIndexIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryProductIndexIdx)
	g.Save(&catalogCategoryProductIndexIdx)
	c.JSON(200, catalogCategoryProductIndexIdx)
}
func deleteCatalogCategoryProductIndexIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryProductIndexIdx CatalogCategoryProductIndexIdx
	d := g.Where("id = ?", id).Delete(&catalogCategoryProductIndexIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
