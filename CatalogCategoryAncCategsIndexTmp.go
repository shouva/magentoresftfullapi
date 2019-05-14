package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryAncCategsIndexTmp :
type CatalogCategoryAncCategsIndexTmp struct {
	CategoryId uint16 `gorm:"column:category_id" form:"category_id" json:"category_id"`
	Path       string `gorm:"column:path" form:"path" json:"path"`
}

// TableName :
func (*CatalogCategoryAncCategsIndexTmp) TableName() string {
	return "catalog_category_anc_categs_index_tmp"
}

// handler create
func initRoutersCatalogCategoryAncCategsIndexTmp(r *gin.Engine, catalogcategoryanccategsindextmp string) {
	route := r.Group("catalogcategoryanccategsindextmp")
	route.GET("/", getCatalogCategoryAncCategsIndexTmps)
	route.GET("/:id", getCatalogCategoryAncCategsIndexTmp)
	route.POST("/", createCatalogCategoryAncCategsIndexTmp)
	route.PUT("/:id", updateCatalogCategoryAncCategsIndexTmp)
	route.DELETE("/:id", deleteCatalogCategoryAncCategsIndexTmp)
}

func getCatalogCategoryAncCategsIndexTmps(c *gin.Context) {
	var catalogCategoryAncCategsIndexTmps []CatalogCategoryAncCategsIndexTmp
	if err := g.Find(&catalogCategoryAncCategsIndexTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryAncCategsIndexTmps)
	}
}

func getCatalogCategoryAncCategsIndexTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryAncCategsIndexTmp CatalogCategoryAncCategsIndexTmp
	if err := g.Where("id = ?", id).First(&catalogCategoryAncCategsIndexTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryAncCategsIndexTmp)
	}
}

func createCatalogCategoryAncCategsIndexTmp(c *gin.Context) {
	var catalogCategoryAncCategsIndexTmp CatalogCategoryAncCategsIndexTmp
	c.BindJSON(&catalogCategoryAncCategsIndexTmp)
	g.Create(&catalogCategoryAncCategsIndexTmp)
	c.JSON(200, catalogCategoryAncCategsIndexTmp)
}

func updateCatalogCategoryAncCategsIndexTmp(c *gin.Context) {
	var catalogCategoryAncCategsIndexTmp CatalogCategoryAncCategsIndexTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryAncCategsIndexTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryAncCategsIndexTmp)
	g.Save(&catalogCategoryAncCategsIndexTmp)
	c.JSON(200, catalogCategoryAncCategsIndexTmp)
}
func deleteCatalogCategoryAncCategsIndexTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryAncCategsIndexTmp CatalogCategoryAncCategsIndexTmp
	d := g.Where("id = ?", id).Delete(&catalogCategoryAncCategsIndexTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
