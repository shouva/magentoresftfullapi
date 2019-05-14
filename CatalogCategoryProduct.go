package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryProduct :
type CatalogCategoryProduct struct {
	CategoryId uint16 `gorm:"column:category_id;primary_key" form:"category_id;primary_key" json:"category_id;primary_key"`
	ProductId  uint16 `gorm:"column:product_id;primary_key" form:"product_id;primary_key" json:"product_id;primary_key"`
	Position   uint16 `gorm:"column:position" form:"position" json:"position"`
}

// TableName :
func (*CatalogCategoryProduct) TableName() string {
	return "catalog_category_product"
}

// handler create
func initRoutersCatalogCategoryProduct(r *gin.Engine, catalogcategoryproduct string) {
	route := r.Group("catalogcategoryproduct")
	route.GET("/", getCatalogCategoryProducts)
	route.GET("/:id", getCatalogCategoryProduct)
	route.POST("/", createCatalogCategoryProduct)
	route.PUT("/:id", updateCatalogCategoryProduct)
	route.DELETE("/:id", deleteCatalogCategoryProduct)
}

func getCatalogCategoryProducts(c *gin.Context) {
	var catalogCategoryProducts []CatalogCategoryProduct
	if err := g.Find(&catalogCategoryProducts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryProducts)
	}
}

func getCatalogCategoryProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryProduct CatalogCategoryProduct
	if err := g.Where("id = ?", id).First(&catalogCategoryProduct).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryProduct)
	}
}

func createCatalogCategoryProduct(c *gin.Context) {
	var catalogCategoryProduct CatalogCategoryProduct
	c.BindJSON(&catalogCategoryProduct)
	g.Create(&catalogCategoryProduct)
	c.JSON(200, catalogCategoryProduct)
}

func updateCatalogCategoryProduct(c *gin.Context) {
	var catalogCategoryProduct CatalogCategoryProduct
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryProduct).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryProduct)
	g.Save(&catalogCategoryProduct)
	c.JSON(200, catalogCategoryProduct)
}
func deleteCatalogCategoryProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryProduct CatalogCategoryProduct
	d := g.Where("id = ?", id).Delete(&catalogCategoryProduct)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
