package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductOptionTitle :
type CatalogProductOptionTitle struct {
	OptionTitleId uint16 `gorm:"column:option_title_id;primary_key" form:"option_title_id;primary_key" json:"option_title_id;primary_key"`
	OptionId      uint16 `gorm:"column:option_id" form:"option_id" json:"option_id"`
	StoreId       uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Title         string `gorm:"column:title" form:"title" json:"title"`
}

// TableName :
func (*CatalogProductOptionTitle) TableName() string {
	return "catalog_product_option_title"
}

// handler create
func initRoutersCatalogProductOptionTitle(r *gin.Engine, catalogproductoptiontitle string) {
	route := r.Group("catalogproductoptiontitle")
	route.GET("/", getCatalogProductOptionTitles)
	route.GET("/:id", getCatalogProductOptionTitle)
	route.POST("/", createCatalogProductOptionTitle)
	route.PUT("/:id", updateCatalogProductOptionTitle)
	route.DELETE("/:id", deleteCatalogProductOptionTitle)
}

func getCatalogProductOptionTitles(c *gin.Context) {
	var catalogProductOptionTitles []CatalogProductOptionTitle
	if err := g.Find(&catalogProductOptionTitles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductOptionTitles)
	}
}

func getCatalogProductOptionTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductOptionTitle CatalogProductOptionTitle
	if err := g.Where("id = ?", id).First(&catalogProductOptionTitle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductOptionTitle)
	}
}

func createCatalogProductOptionTitle(c *gin.Context) {
	var catalogProductOptionTitle CatalogProductOptionTitle
	c.BindJSON(&catalogProductOptionTitle)
	g.Create(&catalogProductOptionTitle)
	c.JSON(200, catalogProductOptionTitle)
}

func updateCatalogProductOptionTitle(c *gin.Context) {
	var catalogProductOptionTitle CatalogProductOptionTitle
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductOptionTitle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductOptionTitle)
	g.Save(&catalogProductOptionTitle)
	c.JSON(200, catalogProductOptionTitle)
}
func deleteCatalogProductOptionTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductOptionTitle CatalogProductOptionTitle
	d := g.Where("id = ?", id).Delete(&catalogProductOptionTitle)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
