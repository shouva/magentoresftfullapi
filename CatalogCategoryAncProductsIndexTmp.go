package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryAncProductsIndexTmp :
type CatalogCategoryAncProductsIndexTmp struct {
	CategoryId uint16 `gorm:"column:category_id" form:"category_id" json:"category_id"`
	ProductId  uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	Position   uint16 `gorm:"column:position" form:"position" json:"position"`
}

// TableName :
func (*CatalogCategoryAncProductsIndexTmp) TableName() string {
	return "catalog_category_anc_products_index_tmp"
}

// handler create
func initRoutersCatalogCategoryAncProductsIndexTmp(r *gin.Engine, catalogcategoryancproductsindextmp string) {
	route := r.Group("catalogcategoryancproductsindextmp")
	route.GET("/", getCatalogCategoryAncProductsIndexTmps)
	route.GET("/:id", getCatalogCategoryAncProductsIndexTmp)
	route.POST("/", createCatalogCategoryAncProductsIndexTmp)
	route.PUT("/:id", updateCatalogCategoryAncProductsIndexTmp)
	route.DELETE("/:id", deleteCatalogCategoryAncProductsIndexTmp)
}

func getCatalogCategoryAncProductsIndexTmps(c *gin.Context) {
	var catalogCategoryAncProductsIndexTmps []CatalogCategoryAncProductsIndexTmp
	if err := g.Find(&catalogCategoryAncProductsIndexTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryAncProductsIndexTmps)
	}
}

func getCatalogCategoryAncProductsIndexTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryAncProductsIndexTmp CatalogCategoryAncProductsIndexTmp
	if err := g.Where("id = ?", id).First(&catalogCategoryAncProductsIndexTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryAncProductsIndexTmp)
	}
}

func createCatalogCategoryAncProductsIndexTmp(c *gin.Context) {
	var catalogCategoryAncProductsIndexTmp CatalogCategoryAncProductsIndexTmp
	c.BindJSON(&catalogCategoryAncProductsIndexTmp)
	g.Create(&catalogCategoryAncProductsIndexTmp)
	c.JSON(200, catalogCategoryAncProductsIndexTmp)
}

func updateCatalogCategoryAncProductsIndexTmp(c *gin.Context) {
	var catalogCategoryAncProductsIndexTmp CatalogCategoryAncProductsIndexTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryAncProductsIndexTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryAncProductsIndexTmp)
	g.Save(&catalogCategoryAncProductsIndexTmp)
	c.JSON(200, catalogCategoryAncProductsIndexTmp)
}
func deleteCatalogCategoryAncProductsIndexTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryAncProductsIndexTmp CatalogCategoryAncProductsIndexTmp
	d := g.Where("id = ?", id).Delete(&catalogCategoryAncProductsIndexTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
