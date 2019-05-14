package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryProductIndexEnblTmp :
type CatalogCategoryProductIndexEnblTmp struct {
	ProductId  uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	Visibility uint16 `gorm:"column:visibility" form:"visibility" json:"visibility"`
}

// TableName :
func (*CatalogCategoryProductIndexEnblTmp) TableName() string {
	return "catalog_category_product_index_enbl_tmp"
}

// handler create
func initRoutersCatalogCategoryProductIndexEnblTmp(r *gin.Engine, catalogcategoryproductindexenbltmp string) {
	route := r.Group("catalogcategoryproductindexenbltmp")
	route.GET("/", getCatalogCategoryProductIndexEnblTmps)
	route.GET("/:id", getCatalogCategoryProductIndexEnblTmp)
	route.POST("/", createCatalogCategoryProductIndexEnblTmp)
	route.PUT("/:id", updateCatalogCategoryProductIndexEnblTmp)
	route.DELETE("/:id", deleteCatalogCategoryProductIndexEnblTmp)
}

func getCatalogCategoryProductIndexEnblTmps(c *gin.Context) {
	var catalogCategoryProductIndexEnblTmps []CatalogCategoryProductIndexEnblTmp
	if err := g.Find(&catalogCategoryProductIndexEnblTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryProductIndexEnblTmps)
	}
}

func getCatalogCategoryProductIndexEnblTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryProductIndexEnblTmp CatalogCategoryProductIndexEnblTmp
	if err := g.Where("id = ?", id).First(&catalogCategoryProductIndexEnblTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryProductIndexEnblTmp)
	}
}

func createCatalogCategoryProductIndexEnblTmp(c *gin.Context) {
	var catalogCategoryProductIndexEnblTmp CatalogCategoryProductIndexEnblTmp
	c.BindJSON(&catalogCategoryProductIndexEnblTmp)
	g.Create(&catalogCategoryProductIndexEnblTmp)
	c.JSON(200, catalogCategoryProductIndexEnblTmp)
}

func updateCatalogCategoryProductIndexEnblTmp(c *gin.Context) {
	var catalogCategoryProductIndexEnblTmp CatalogCategoryProductIndexEnblTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryProductIndexEnblTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryProductIndexEnblTmp)
	g.Save(&catalogCategoryProductIndexEnblTmp)
	c.JSON(200, catalogCategoryProductIndexEnblTmp)
}
func deleteCatalogCategoryProductIndexEnblTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryProductIndexEnblTmp CatalogCategoryProductIndexEnblTmp
	d := g.Where("id = ?", id).Delete(&catalogCategoryProductIndexEnblTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
