package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceOptIdx :
type CatalogProductIndexPriceOptIdx struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	MaxPrice        float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
	TierPrice       float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	GroupPrice      float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
}

// TableName :
func (*CatalogProductIndexPriceOptIdx) TableName() string {
	return "catalog_product_index_price_opt_idx"
}

// handler create
func initRoutersCatalogProductIndexPriceOptIdx(r *gin.Engine, catalogproductindexpriceoptidx string) {
	route := r.Group("catalogproductindexpriceoptidx")
	route.GET("/", getCatalogProductIndexPriceOptIdxs)
	route.GET("/:id", getCatalogProductIndexPriceOptIdx)
	route.POST("/", createCatalogProductIndexPriceOptIdx)
	route.PUT("/:id", updateCatalogProductIndexPriceOptIdx)
	route.DELETE("/:id", deleteCatalogProductIndexPriceOptIdx)
}

func getCatalogProductIndexPriceOptIdxs(c *gin.Context) {
	var catalogProductIndexPriceOptIdxs []CatalogProductIndexPriceOptIdx
	if err := g.Find(&catalogProductIndexPriceOptIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceOptIdxs)
	}
}

func getCatalogProductIndexPriceOptIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceOptIdx CatalogProductIndexPriceOptIdx
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceOptIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceOptIdx)
	}
}

func createCatalogProductIndexPriceOptIdx(c *gin.Context) {
	var catalogProductIndexPriceOptIdx CatalogProductIndexPriceOptIdx
	c.BindJSON(&catalogProductIndexPriceOptIdx)
	g.Create(&catalogProductIndexPriceOptIdx)
	c.JSON(200, catalogProductIndexPriceOptIdx)
}

func updateCatalogProductIndexPriceOptIdx(c *gin.Context) {
	var catalogProductIndexPriceOptIdx CatalogProductIndexPriceOptIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceOptIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceOptIdx)
	g.Save(&catalogProductIndexPriceOptIdx)
	c.JSON(200, catalogProductIndexPriceOptIdx)
}
func deleteCatalogProductIndexPriceOptIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceOptIdx CatalogProductIndexPriceOptIdx
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceOptIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
