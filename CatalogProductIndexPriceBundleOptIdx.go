package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceBundleOptIdx :
type CatalogProductIndexPriceBundleOptIdx struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	OptionId        uint16  `gorm:"column:option_id;primary_key" form:"option_id;primary_key" json:"option_id;primary_key"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	AltPrice        float64 `gorm:"column:alt_price" form:"alt_price" json:"alt_price"`
	MaxPrice        float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
	TierPrice       float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	AltTierPrice    float64 `gorm:"column:alt_tier_price" form:"alt_tier_price" json:"alt_tier_price"`
	GroupPrice      float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
	AltGroupPrice   float64 `gorm:"column:alt_group_price" form:"alt_group_price" json:"alt_group_price"`
}

// TableName :
func (*CatalogProductIndexPriceBundleOptIdx) TableName() string {
	return "catalog_product_index_price_bundle_opt_idx"
}

// handler create
func initRoutersCatalogProductIndexPriceBundleOptIdx(r *gin.Engine, catalogproductindexpricebundleoptidx string) {
	route := r.Group("catalogproductindexpricebundleoptidx")
	route.GET("/", getCatalogProductIndexPriceBundleOptIdxs)
	route.GET("/:id", getCatalogProductIndexPriceBundleOptIdx)
	route.POST("/", createCatalogProductIndexPriceBundleOptIdx)
	route.PUT("/:id", updateCatalogProductIndexPriceBundleOptIdx)
	route.DELETE("/:id", deleteCatalogProductIndexPriceBundleOptIdx)
}

func getCatalogProductIndexPriceBundleOptIdxs(c *gin.Context) {
	var catalogProductIndexPriceBundleOptIdxs []CatalogProductIndexPriceBundleOptIdx
	if err := g.Find(&catalogProductIndexPriceBundleOptIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceBundleOptIdxs)
	}
}

func getCatalogProductIndexPriceBundleOptIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceBundleOptIdx CatalogProductIndexPriceBundleOptIdx
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceBundleOptIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceBundleOptIdx)
	}
}

func createCatalogProductIndexPriceBundleOptIdx(c *gin.Context) {
	var catalogProductIndexPriceBundleOptIdx CatalogProductIndexPriceBundleOptIdx
	c.BindJSON(&catalogProductIndexPriceBundleOptIdx)
	g.Create(&catalogProductIndexPriceBundleOptIdx)
	c.JSON(200, catalogProductIndexPriceBundleOptIdx)
}

func updateCatalogProductIndexPriceBundleOptIdx(c *gin.Context) {
	var catalogProductIndexPriceBundleOptIdx CatalogProductIndexPriceBundleOptIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceBundleOptIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceBundleOptIdx)
	g.Save(&catalogProductIndexPriceBundleOptIdx)
	c.JSON(200, catalogProductIndexPriceBundleOptIdx)
}
func deleteCatalogProductIndexPriceBundleOptIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceBundleOptIdx CatalogProductIndexPriceBundleOptIdx
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceBundleOptIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
