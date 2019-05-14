package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceCfgOptIdx :
type CatalogProductIndexPriceCfgOptIdx struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	MaxPrice        float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
	TierPrice       float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	GroupPrice      float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
}

// TableName :
func (*CatalogProductIndexPriceCfgOptIdx) TableName() string {
	return "catalog_product_index_price_cfg_opt_idx"
}

// handler create
func initRoutersCatalogProductIndexPriceCfgOptIdx(r *gin.Engine, catalogproductindexpricecfgoptidx string) {
	route := r.Group("catalogproductindexpricecfgoptidx")
	route.GET("/", getCatalogProductIndexPriceCfgOptIdxs)
	route.GET("/:id", getCatalogProductIndexPriceCfgOptIdx)
	route.POST("/", createCatalogProductIndexPriceCfgOptIdx)
	route.PUT("/:id", updateCatalogProductIndexPriceCfgOptIdx)
	route.DELETE("/:id", deleteCatalogProductIndexPriceCfgOptIdx)
}

func getCatalogProductIndexPriceCfgOptIdxs(c *gin.Context) {
	var catalogProductIndexPriceCfgOptIdxs []CatalogProductIndexPriceCfgOptIdx
	if err := g.Find(&catalogProductIndexPriceCfgOptIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceCfgOptIdxs)
	}
}

func getCatalogProductIndexPriceCfgOptIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceCfgOptIdx CatalogProductIndexPriceCfgOptIdx
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceCfgOptIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceCfgOptIdx)
	}
}

func createCatalogProductIndexPriceCfgOptIdx(c *gin.Context) {
	var catalogProductIndexPriceCfgOptIdx CatalogProductIndexPriceCfgOptIdx
	c.BindJSON(&catalogProductIndexPriceCfgOptIdx)
	g.Create(&catalogProductIndexPriceCfgOptIdx)
	c.JSON(200, catalogProductIndexPriceCfgOptIdx)
}

func updateCatalogProductIndexPriceCfgOptIdx(c *gin.Context) {
	var catalogProductIndexPriceCfgOptIdx CatalogProductIndexPriceCfgOptIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceCfgOptIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceCfgOptIdx)
	g.Save(&catalogProductIndexPriceCfgOptIdx)
	c.JSON(200, catalogProductIndexPriceCfgOptIdx)
}
func deleteCatalogProductIndexPriceCfgOptIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceCfgOptIdx CatalogProductIndexPriceCfgOptIdx
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceCfgOptIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
