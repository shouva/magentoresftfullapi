package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceBundleIdx :
type CatalogProductIndexPriceBundleIdx struct {
	EntityId          uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId   uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId         uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	TaxClassId        uint16  `gorm:"column:tax_class_id" form:"tax_class_id" json:"tax_class_id"`
	PriceType         uint16  `gorm:"column:price_type" form:"price_type" json:"price_type"`
	SpecialPrice      float64 `gorm:"column:special_price" form:"special_price" json:"special_price"`
	TierPercent       float64 `gorm:"column:tier_percent" form:"tier_percent" json:"tier_percent"`
	OrigPrice         float64 `gorm:"column:orig_price" form:"orig_price" json:"orig_price"`
	Price             float64 `gorm:"column:price" form:"price" json:"price"`
	MinPrice          float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	MaxPrice          float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
	TierPrice         float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	BaseTier          float64 `gorm:"column:base_tier" form:"base_tier" json:"base_tier"`
	GroupPrice        float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
	BaseGroupPrice    float64 `gorm:"column:base_group_price" form:"base_group_price" json:"base_group_price"`
	GroupPricePercent float64 `gorm:"column:group_price_percent" form:"group_price_percent" json:"group_price_percent"`
}

// TableName :
func (*CatalogProductIndexPriceBundleIdx) TableName() string {
	return "catalog_product_index_price_bundle_idx"
}

// handler create
func initRoutersCatalogProductIndexPriceBundleIdx(r *gin.Engine, catalogproductindexpricebundleidx string) {
	route := r.Group("catalogproductindexpricebundleidx")
	route.GET("/", getCatalogProductIndexPriceBundleIdxs)
	route.GET("/:id", getCatalogProductIndexPriceBundleIdx)
	route.POST("/", createCatalogProductIndexPriceBundleIdx)
	route.PUT("/:id", updateCatalogProductIndexPriceBundleIdx)
	route.DELETE("/:id", deleteCatalogProductIndexPriceBundleIdx)
}

func getCatalogProductIndexPriceBundleIdxs(c *gin.Context) {
	var catalogProductIndexPriceBundleIdxs []CatalogProductIndexPriceBundleIdx
	if err := g.Find(&catalogProductIndexPriceBundleIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceBundleIdxs)
	}
}

func getCatalogProductIndexPriceBundleIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceBundleIdx CatalogProductIndexPriceBundleIdx
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceBundleIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceBundleIdx)
	}
}

func createCatalogProductIndexPriceBundleIdx(c *gin.Context) {
	var catalogProductIndexPriceBundleIdx CatalogProductIndexPriceBundleIdx
	c.BindJSON(&catalogProductIndexPriceBundleIdx)
	g.Create(&catalogProductIndexPriceBundleIdx)
	c.JSON(200, catalogProductIndexPriceBundleIdx)
}

func updateCatalogProductIndexPriceBundleIdx(c *gin.Context) {
	var catalogProductIndexPriceBundleIdx CatalogProductIndexPriceBundleIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceBundleIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceBundleIdx)
	g.Save(&catalogProductIndexPriceBundleIdx)
	c.JSON(200, catalogProductIndexPriceBundleIdx)
}
func deleteCatalogProductIndexPriceBundleIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceBundleIdx CatalogProductIndexPriceBundleIdx
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceBundleIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
