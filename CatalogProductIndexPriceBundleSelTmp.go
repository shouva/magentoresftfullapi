package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceBundleSelTmp :
type CatalogProductIndexPriceBundleSelTmp struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	OptionId        uint16  `gorm:"column:option_id;primary_key" form:"option_id;primary_key" json:"option_id;primary_key"`
	SelectionId     uint16  `gorm:"column:selection_id;primary_key" form:"selection_id;primary_key" json:"selection_id;primary_key"`
	GroupType       uint16  `gorm:"column:group_type" form:"group_type" json:"group_type"`
	IsRequired      uint16  `gorm:"column:is_required" form:"is_required" json:"is_required"`
	Price           float64 `gorm:"column:price" form:"price" json:"price"`
	TierPrice       float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	GroupPrice      float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
}

// TableName :
func (*CatalogProductIndexPriceBundleSelTmp) TableName() string {
	return "catalog_product_index_price_bundle_sel_tmp"
}

// handler create
func initRoutersCatalogProductIndexPriceBundleSelTmp(r *gin.Engine, catalogproductindexpricebundleseltmp string) {
	route := r.Group("catalogproductindexpricebundleseltmp")
	route.GET("/", getCatalogProductIndexPriceBundleSelTmps)
	route.GET("/:id", getCatalogProductIndexPriceBundleSelTmp)
	route.POST("/", createCatalogProductIndexPriceBundleSelTmp)
	route.PUT("/:id", updateCatalogProductIndexPriceBundleSelTmp)
	route.DELETE("/:id", deleteCatalogProductIndexPriceBundleSelTmp)
}

func getCatalogProductIndexPriceBundleSelTmps(c *gin.Context) {
	var catalogProductIndexPriceBundleSelTmps []CatalogProductIndexPriceBundleSelTmp
	if err := g.Find(&catalogProductIndexPriceBundleSelTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceBundleSelTmps)
	}
}

func getCatalogProductIndexPriceBundleSelTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceBundleSelTmp CatalogProductIndexPriceBundleSelTmp
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceBundleSelTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceBundleSelTmp)
	}
}

func createCatalogProductIndexPriceBundleSelTmp(c *gin.Context) {
	var catalogProductIndexPriceBundleSelTmp CatalogProductIndexPriceBundleSelTmp
	c.BindJSON(&catalogProductIndexPriceBundleSelTmp)
	g.Create(&catalogProductIndexPriceBundleSelTmp)
	c.JSON(200, catalogProductIndexPriceBundleSelTmp)
}

func updateCatalogProductIndexPriceBundleSelTmp(c *gin.Context) {
	var catalogProductIndexPriceBundleSelTmp CatalogProductIndexPriceBundleSelTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceBundleSelTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceBundleSelTmp)
	g.Save(&catalogProductIndexPriceBundleSelTmp)
	c.JSON(200, catalogProductIndexPriceBundleSelTmp)
}
func deleteCatalogProductIndexPriceBundleSelTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceBundleSelTmp CatalogProductIndexPriceBundleSelTmp
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceBundleSelTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
