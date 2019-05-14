package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryProductIndexEnblIdx :
type CatalogCategoryProductIndexEnblIdx struct {
	ProductId  uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	Visibility uint16 `gorm:"column:visibility" form:"visibility" json:"visibility"`
}

// TableName :
func (*CatalogCategoryProductIndexEnblIdx) TableName() string {
	return "catalog_category_product_index_enbl_idx"
}

// handler create
func initRoutersCatalogCategoryProductIndexEnblIdx(r *gin.Engine, catalogcategoryproductindexenblidx string) {
	route := r.Group("catalogcategoryproductindexenblidx")
	route.GET("/", getCatalogCategoryProductIndexEnblIdxs)
	route.GET("/:id", getCatalogCategoryProductIndexEnblIdx)
	route.POST("/", createCatalogCategoryProductIndexEnblIdx)
	route.PUT("/:id", updateCatalogCategoryProductIndexEnblIdx)
	route.DELETE("/:id", deleteCatalogCategoryProductIndexEnblIdx)
}

func getCatalogCategoryProductIndexEnblIdxs(c *gin.Context) {
	var catalogCategoryProductIndexEnblIdxs []CatalogCategoryProductIndexEnblIdx
	if err := g.Find(&catalogCategoryProductIndexEnblIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryProductIndexEnblIdxs)
	}
}

func getCatalogCategoryProductIndexEnblIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryProductIndexEnblIdx CatalogCategoryProductIndexEnblIdx
	if err := g.Where("id = ?", id).First(&catalogCategoryProductIndexEnblIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryProductIndexEnblIdx)
	}
}

func createCatalogCategoryProductIndexEnblIdx(c *gin.Context) {
	var catalogCategoryProductIndexEnblIdx CatalogCategoryProductIndexEnblIdx
	c.BindJSON(&catalogCategoryProductIndexEnblIdx)
	g.Create(&catalogCategoryProductIndexEnblIdx)
	c.JSON(200, catalogCategoryProductIndexEnblIdx)
}

func updateCatalogCategoryProductIndexEnblIdx(c *gin.Context) {
	var catalogCategoryProductIndexEnblIdx CatalogCategoryProductIndexEnblIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryProductIndexEnblIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryProductIndexEnblIdx)
	g.Save(&catalogCategoryProductIndexEnblIdx)
	c.JSON(200, catalogCategoryProductIndexEnblIdx)
}
func deleteCatalogCategoryProductIndexEnblIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryProductIndexEnblIdx CatalogCategoryProductIndexEnblIdx
	d := g.Where("id = ?", id).Delete(&catalogCategoryProductIndexEnblIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
