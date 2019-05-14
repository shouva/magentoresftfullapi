package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductOptionTypeTitle :
type CatalogProductOptionTypeTitle struct {
	OptionTypeTitleId uint16 `gorm:"column:option_type_title_id;primary_key" form:"option_type_title_id;primary_key" json:"option_type_title_id;primary_key"`
	OptionTypeId      uint16 `gorm:"column:option_type_id" form:"option_type_id" json:"option_type_id"`
	StoreId           uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Title             string `gorm:"column:title" form:"title" json:"title"`
}

// TableName :
func (*CatalogProductOptionTypeTitle) TableName() string {
	return "catalog_product_option_type_title"
}

// handler create
func initRoutersCatalogProductOptionTypeTitle(r *gin.Engine, catalogproductoptiontypetitle string) {
	route := r.Group("catalogproductoptiontypetitle")
	route.GET("/", getCatalogProductOptionTypeTitles)
	route.GET("/:id", getCatalogProductOptionTypeTitle)
	route.POST("/", createCatalogProductOptionTypeTitle)
	route.PUT("/:id", updateCatalogProductOptionTypeTitle)
	route.DELETE("/:id", deleteCatalogProductOptionTypeTitle)
}

func getCatalogProductOptionTypeTitles(c *gin.Context) {
	var catalogProductOptionTypeTitles []CatalogProductOptionTypeTitle
	if err := g.Find(&catalogProductOptionTypeTitles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductOptionTypeTitles)
	}
}

func getCatalogProductOptionTypeTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductOptionTypeTitle CatalogProductOptionTypeTitle
	if err := g.Where("id = ?", id).First(&catalogProductOptionTypeTitle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductOptionTypeTitle)
	}
}

func createCatalogProductOptionTypeTitle(c *gin.Context) {
	var catalogProductOptionTypeTitle CatalogProductOptionTypeTitle
	c.BindJSON(&catalogProductOptionTypeTitle)
	g.Create(&catalogProductOptionTypeTitle)
	c.JSON(200, catalogProductOptionTypeTitle)
}

func updateCatalogProductOptionTypeTitle(c *gin.Context) {
	var catalogProductOptionTypeTitle CatalogProductOptionTypeTitle
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductOptionTypeTitle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductOptionTypeTitle)
	g.Save(&catalogProductOptionTypeTitle)
	c.JSON(200, catalogProductOptionTypeTitle)
}
func deleteCatalogProductOptionTypeTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductOptionTypeTitle CatalogProductOptionTypeTitle
	d := g.Where("id = ?", id).Delete(&catalogProductOptionTypeTitle)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
