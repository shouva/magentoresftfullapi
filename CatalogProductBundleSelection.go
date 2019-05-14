package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductBundleSelection :
type CatalogProductBundleSelection struct {
	SelectionId           uint16  `gorm:"column:selection_id;primary_key" form:"selection_id;primary_key" json:"selection_id;primary_key"`
	OptionId              uint16  `gorm:"column:option_id" form:"option_id" json:"option_id"`
	ParentProductId       uint16  `gorm:"column:parent_product_id" form:"parent_product_id" json:"parent_product_id"`
	ProductId             uint16  `gorm:"column:product_id" form:"product_id" json:"product_id"`
	Position              uint16  `gorm:"column:position" form:"position" json:"position"`
	IsDefault             uint16  `gorm:"column:is_default" form:"is_default" json:"is_default"`
	SelectionPriceType    uint16  `gorm:"column:selection_price_type" form:"selection_price_type" json:"selection_price_type"`
	SelectionPriceValue   float64 `gorm:"column:selection_price_value" form:"selection_price_value" json:"selection_price_value"`
	SelectionQty          float64 `gorm:"column:selection_qty" form:"selection_qty" json:"selection_qty"`
	SelectionCanChangeQty uint16  `gorm:"column:selection_can_change_qty" form:"selection_can_change_qty" json:"selection_can_change_qty"`
}

// TableName :
func (*CatalogProductBundleSelection) TableName() string {
	return "catalog_product_bundle_selection"
}

// handler create
func initRoutersCatalogProductBundleSelection(r *gin.Engine, catalogproductbundleselection string) {
	route := r.Group("catalogproductbundleselection")
	route.GET("/", getCatalogProductBundleSelections)
	route.GET("/:id", getCatalogProductBundleSelection)
	route.POST("/", createCatalogProductBundleSelection)
	route.PUT("/:id", updateCatalogProductBundleSelection)
	route.DELETE("/:id", deleteCatalogProductBundleSelection)
}

func getCatalogProductBundleSelections(c *gin.Context) {
	var catalogProductBundleSelections []CatalogProductBundleSelection
	if err := g.Find(&catalogProductBundleSelections).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductBundleSelections)
	}
}

func getCatalogProductBundleSelection(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductBundleSelection CatalogProductBundleSelection
	if err := g.Where("id = ?", id).First(&catalogProductBundleSelection).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductBundleSelection)
	}
}

func createCatalogProductBundleSelection(c *gin.Context) {
	var catalogProductBundleSelection CatalogProductBundleSelection
	c.BindJSON(&catalogProductBundleSelection)
	g.Create(&catalogProductBundleSelection)
	c.JSON(200, catalogProductBundleSelection)
}

func updateCatalogProductBundleSelection(c *gin.Context) {
	var catalogProductBundleSelection CatalogProductBundleSelection
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductBundleSelection).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductBundleSelection)
	g.Save(&catalogProductBundleSelection)
	c.JSON(200, catalogProductBundleSelection)
}
func deleteCatalogProductBundleSelection(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductBundleSelection CatalogProductBundleSelection
	d := g.Where("id = ?", id).Delete(&catalogProductBundleSelection)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
