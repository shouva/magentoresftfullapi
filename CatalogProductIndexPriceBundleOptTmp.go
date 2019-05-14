package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceBundleOptTmp :
type CatalogProductIndexPriceBundleOptTmp struct {
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
func (*CatalogProductIndexPriceBundleOptTmp) TableName() string {
	return "catalog_product_index_price_bundle_opt_tmp"
}

// handler create
func initRoutersCatalogProductIndexPriceBundleOptTmp(r *gin.Engine, catalogproductindexpricebundleopttmp string) {
	route := r.Group("catalogproductindexpricebundleopttmp")
	route.GET("/", getCatalogProductIndexPriceBundleOptTmps)
	route.GET("/:id", getCatalogProductIndexPriceBundleOptTmp)
	route.POST("/", createCatalogProductIndexPriceBundleOptTmp)
	route.PUT("/:id", updateCatalogProductIndexPriceBundleOptTmp)
	route.DELETE("/:id", deleteCatalogProductIndexPriceBundleOptTmp)
}

func getCatalogProductIndexPriceBundleOptTmps(c *gin.Context) {
	var catalogProductIndexPriceBundleOptTmps []CatalogProductIndexPriceBundleOptTmp
	if err := g.Find(&catalogProductIndexPriceBundleOptTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceBundleOptTmps)
	}
}

func getCatalogProductIndexPriceBundleOptTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceBundleOptTmp CatalogProductIndexPriceBundleOptTmp
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceBundleOptTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceBundleOptTmp)
	}
}

func createCatalogProductIndexPriceBundleOptTmp(c *gin.Context) {
	var catalogProductIndexPriceBundleOptTmp CatalogProductIndexPriceBundleOptTmp
	c.BindJSON(&catalogProductIndexPriceBundleOptTmp)
	g.Create(&catalogProductIndexPriceBundleOptTmp)
	c.JSON(200, catalogProductIndexPriceBundleOptTmp)
}

func updateCatalogProductIndexPriceBundleOptTmp(c *gin.Context) {
	var catalogProductIndexPriceBundleOptTmp CatalogProductIndexPriceBundleOptTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceBundleOptTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceBundleOptTmp)
	g.Save(&catalogProductIndexPriceBundleOptTmp)
	c.JSON(200, catalogProductIndexPriceBundleOptTmp)
}
func deleteCatalogProductIndexPriceBundleOptTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceBundleOptTmp CatalogProductIndexPriceBundleOptTmp
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceBundleOptTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
