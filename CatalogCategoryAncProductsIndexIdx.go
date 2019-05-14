package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryAncProductsIndexIdx :
type CatalogCategoryAncProductsIndexIdx struct {
	CategoryId uint16 `gorm:"column:category_id" form:"category_id" json:"category_id"`
	ProductId  uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	Position   uint16 `gorm:"column:position" form:"position" json:"position"`
}

// TableName :
func (*CatalogCategoryAncProductsIndexIdx) TableName() string {
	return "catalog_category_anc_products_index_idx"
}

// handler create
func initRoutersCatalogCategoryAncProductsIndexIdx(r *gin.Engine, catalogcategoryancproductsindexidx string) {
	route := r.Group("catalogcategoryancproductsindexidx")
	route.GET("/", getCatalogCategoryAncProductsIndexIdxs)
	route.GET("/:id", getCatalogCategoryAncProductsIndexIdx)
	route.POST("/", createCatalogCategoryAncProductsIndexIdx)
	route.PUT("/:id", updateCatalogCategoryAncProductsIndexIdx)
	route.DELETE("/:id", deleteCatalogCategoryAncProductsIndexIdx)
}

func getCatalogCategoryAncProductsIndexIdxs(c *gin.Context) {
	var catalogCategoryAncProductsIndexIdxs []CatalogCategoryAncProductsIndexIdx
	if err := g.Find(&catalogCategoryAncProductsIndexIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryAncProductsIndexIdxs)
	}
}

func getCatalogCategoryAncProductsIndexIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryAncProductsIndexIdx CatalogCategoryAncProductsIndexIdx
	if err := g.Where("id = ?", id).First(&catalogCategoryAncProductsIndexIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryAncProductsIndexIdx)
	}
}

func createCatalogCategoryAncProductsIndexIdx(c *gin.Context) {
	var catalogCategoryAncProductsIndexIdx CatalogCategoryAncProductsIndexIdx
	c.BindJSON(&catalogCategoryAncProductsIndexIdx)
	g.Create(&catalogCategoryAncProductsIndexIdx)
	c.JSON(200, catalogCategoryAncProductsIndexIdx)
}

func updateCatalogCategoryAncProductsIndexIdx(c *gin.Context) {
	var catalogCategoryAncProductsIndexIdx CatalogCategoryAncProductsIndexIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryAncProductsIndexIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryAncProductsIndexIdx)
	g.Save(&catalogCategoryAncProductsIndexIdx)
	c.JSON(200, catalogCategoryAncProductsIndexIdx)
}
func deleteCatalogCategoryAncProductsIndexIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryAncProductsIndexIdx CatalogCategoryAncProductsIndexIdx
	d := g.Where("id = ?", id).Delete(&catalogCategoryAncProductsIndexIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
