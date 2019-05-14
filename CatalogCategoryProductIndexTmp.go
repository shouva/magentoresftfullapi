package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryProductIndexTmp :
type CatalogCategoryProductIndexTmp struct {
	CategoryId uint16 `gorm:"column:category_id" form:"category_id" json:"category_id"`
	ProductId  uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	Position   uint16 `gorm:"column:position" form:"position" json:"position"`
	IsParent   uint16 `gorm:"column:is_parent" form:"is_parent" json:"is_parent"`
	StoreId    uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Visibility uint16 `gorm:"column:visibility" form:"visibility" json:"visibility"`
}

// TableName :
func (*CatalogCategoryProductIndexTmp) TableName() string {
	return "catalog_category_product_index_tmp"
}

// handler create
func initRoutersCatalogCategoryProductIndexTmp(r *gin.Engine, catalogcategoryproductindextmp string) {
	route := r.Group("catalogcategoryproductindextmp")
	route.GET("/", getCatalogCategoryProductIndexTmps)
	route.GET("/:id", getCatalogCategoryProductIndexTmp)
	route.POST("/", createCatalogCategoryProductIndexTmp)
	route.PUT("/:id", updateCatalogCategoryProductIndexTmp)
	route.DELETE("/:id", deleteCatalogCategoryProductIndexTmp)
}

func getCatalogCategoryProductIndexTmps(c *gin.Context) {
	var catalogCategoryProductIndexTmps []CatalogCategoryProductIndexTmp
	if err := g.Find(&catalogCategoryProductIndexTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryProductIndexTmps)
	}
}

func getCatalogCategoryProductIndexTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryProductIndexTmp CatalogCategoryProductIndexTmp
	if err := g.Where("id = ?", id).First(&catalogCategoryProductIndexTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryProductIndexTmp)
	}
}

func createCatalogCategoryProductIndexTmp(c *gin.Context) {
	var catalogCategoryProductIndexTmp CatalogCategoryProductIndexTmp
	c.BindJSON(&catalogCategoryProductIndexTmp)
	g.Create(&catalogCategoryProductIndexTmp)
	c.JSON(200, catalogCategoryProductIndexTmp)
}

func updateCatalogCategoryProductIndexTmp(c *gin.Context) {
	var catalogCategoryProductIndexTmp CatalogCategoryProductIndexTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryProductIndexTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryProductIndexTmp)
	g.Save(&catalogCategoryProductIndexTmp)
	c.JSON(200, catalogCategoryProductIndexTmp)
}
func deleteCatalogCategoryProductIndexTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryProductIndexTmp CatalogCategoryProductIndexTmp
	d := g.Where("id = ?", id).Delete(&catalogCategoryProductIndexTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
