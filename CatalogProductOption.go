package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductOption :
type CatalogProductOption struct {
	OptionId      uint16 `gorm:"column:option_id;primary_key" form:"option_id;primary_key" json:"option_id;primary_key"`
	ProductId     uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	Type          string `gorm:"column:type" form:"type" json:"type"`
	IsRequire     uint16 `gorm:"column:is_require" form:"is_require" json:"is_require"`
	Sku           string `gorm:"column:sku" form:"sku" json:"sku"`
	MaxCharacters uint16 `gorm:"column:max_characters" form:"max_characters" json:"max_characters"`
	FileExtension string `gorm:"column:file_extension" form:"file_extension" json:"file_extension"`
	ImageSizeX    uint16 `gorm:"column:image_size_x" form:"image_size_x" json:"image_size_x"`
	ImageSizeY    uint16 `gorm:"column:image_size_y" form:"image_size_y" json:"image_size_y"`
	SortOrder     uint16 `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
}

// TableName :
func (*CatalogProductOption) TableName() string {
	return "catalog_product_option"
}

// handler create
func initRoutersCatalogProductOption(r *gin.Engine, catalogproductoption string) {
	route := r.Group("catalogproductoption")
	route.GET("/", getCatalogProductOptions)
	route.GET("/:id", getCatalogProductOption)
	route.POST("/", createCatalogProductOption)
	route.PUT("/:id", updateCatalogProductOption)
	route.DELETE("/:id", deleteCatalogProductOption)
}

func getCatalogProductOptions(c *gin.Context) {
	var catalogProductOptions []CatalogProductOption
	if err := g.Find(&catalogProductOptions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductOptions)
	}
}

func getCatalogProductOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductOption CatalogProductOption
	if err := g.Where("id = ?", id).First(&catalogProductOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductOption)
	}
}

func createCatalogProductOption(c *gin.Context) {
	var catalogProductOption CatalogProductOption
	c.BindJSON(&catalogProductOption)
	g.Create(&catalogProductOption)
	c.JSON(200, catalogProductOption)
}

func updateCatalogProductOption(c *gin.Context) {
	var catalogProductOption CatalogProductOption
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductOption).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductOption)
	g.Save(&catalogProductOption)
	c.JSON(200, catalogProductOption)
}
func deleteCatalogProductOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductOption CatalogProductOption
	d := g.Where("id = ?", id).Delete(&catalogProductOption)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
