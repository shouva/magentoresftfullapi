package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceBundleSelIdx :
type CatalogProductIndexPriceBundleSelIdx struct {
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
func (*CatalogProductIndexPriceBundleSelIdx) TableName() string {
	return "catalog_product_index_price_bundle_sel_idx"
}

// handler create
func initRoutersCatalogProductIndexPriceBundleSelIdx(r *gin.Engine, catalogproductindexpricebundleselidx string) {
	route := r.Group("catalogproductindexpricebundleselidx")
	route.GET("/", getCatalogProductIndexPriceBundleSelIdxs)
	route.GET("/:id", getCatalogProductIndexPriceBundleSelIdx)
	route.POST("/", createCatalogProductIndexPriceBundleSelIdx)
	route.PUT("/:id", updateCatalogProductIndexPriceBundleSelIdx)
	route.DELETE("/:id", deleteCatalogProductIndexPriceBundleSelIdx)
}

func getCatalogProductIndexPriceBundleSelIdxs(c *gin.Context) {
	var catalogProductIndexPriceBundleSelIdxs []CatalogProductIndexPriceBundleSelIdx
	if err := g.Find(&catalogProductIndexPriceBundleSelIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceBundleSelIdxs)
	}
}

func getCatalogProductIndexPriceBundleSelIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceBundleSelIdx CatalogProductIndexPriceBundleSelIdx
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceBundleSelIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceBundleSelIdx)
	}
}

func createCatalogProductIndexPriceBundleSelIdx(c *gin.Context) {
	var catalogProductIndexPriceBundleSelIdx CatalogProductIndexPriceBundleSelIdx
	c.BindJSON(&catalogProductIndexPriceBundleSelIdx)
	g.Create(&catalogProductIndexPriceBundleSelIdx)
	c.JSON(200, catalogProductIndexPriceBundleSelIdx)
}

func updateCatalogProductIndexPriceBundleSelIdx(c *gin.Context) {
	var catalogProductIndexPriceBundleSelIdx CatalogProductIndexPriceBundleSelIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceBundleSelIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceBundleSelIdx)
	g.Save(&catalogProductIndexPriceBundleSelIdx)
	c.JSON(200, catalogProductIndexPriceBundleSelIdx)
}
func deleteCatalogProductIndexPriceBundleSelIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceBundleSelIdx CatalogProductIndexPriceBundleSelIdx
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceBundleSelIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
