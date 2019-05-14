package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogCategoryAncCategsIndexIdx :
type CatalogCategoryAncCategsIndexIdx struct {
	CategoryId uint16 `gorm:"column:category_id" form:"category_id" json:"category_id"`
	Path       string `gorm:"column:path" form:"path" json:"path"`
}

// TableName :
func (*CatalogCategoryAncCategsIndexIdx) TableName() string {
	return "catalog_category_anc_categs_index_idx"
}

// handler create
func initRoutersCatalogCategoryAncCategsIndexIdx(r *gin.Engine, catalogcategoryanccategsindexidx string) {
	route := r.Group("catalogcategoryanccategsindexidx")
	route.GET("/", getCatalogCategoryAncCategsIndexIdxs)
	route.GET("/:id", getCatalogCategoryAncCategsIndexIdx)
	route.POST("/", createCatalogCategoryAncCategsIndexIdx)
	route.PUT("/:id", updateCatalogCategoryAncCategsIndexIdx)
	route.DELETE("/:id", deleteCatalogCategoryAncCategsIndexIdx)
}

func getCatalogCategoryAncCategsIndexIdxs(c *gin.Context) {
	var catalogCategoryAncCategsIndexIdxs []CatalogCategoryAncCategsIndexIdx
	if err := g.Find(&catalogCategoryAncCategsIndexIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryAncCategsIndexIdxs)
	}
}

func getCatalogCategoryAncCategsIndexIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryAncCategsIndexIdx CatalogCategoryAncCategsIndexIdx
	if err := g.Where("id = ?", id).First(&catalogCategoryAncCategsIndexIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogCategoryAncCategsIndexIdx)
	}
}

func createCatalogCategoryAncCategsIndexIdx(c *gin.Context) {
	var catalogCategoryAncCategsIndexIdx CatalogCategoryAncCategsIndexIdx
	c.BindJSON(&catalogCategoryAncCategsIndexIdx)
	g.Create(&catalogCategoryAncCategsIndexIdx)
	c.JSON(200, catalogCategoryAncCategsIndexIdx)
}

func updateCatalogCategoryAncCategsIndexIdx(c *gin.Context) {
	var catalogCategoryAncCategsIndexIdx CatalogCategoryAncCategsIndexIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogCategoryAncCategsIndexIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogCategoryAncCategsIndexIdx)
	g.Save(&catalogCategoryAncCategsIndexIdx)
	c.JSON(200, catalogCategoryAncCategsIndexIdx)
}
func deleteCatalogCategoryAncCategsIndexIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogCategoryAncCategsIndexIdx CatalogCategoryAncCategsIndexIdx
	d := g.Where("id = ?", id).Delete(&catalogCategoryAncCategsIndexIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
