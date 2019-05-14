package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceFinalIdx :
type CatalogProductIndexPriceFinalIdx struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	TaxClassId      uint16  `gorm:"column:tax_class_id" form:"tax_class_id" json:"tax_class_id"`
	OrigPrice       float64 `gorm:"column:orig_price" form:"orig_price" json:"orig_price"`
	Price           float64 `gorm:"column:price" form:"price" json:"price"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	MaxPrice        float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
	TierPrice       float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	BaseTier        float64 `gorm:"column:base_tier" form:"base_tier" json:"base_tier"`
	GroupPrice      float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
	BaseGroupPrice  float64 `gorm:"column:base_group_price" form:"base_group_price" json:"base_group_price"`
}

// TableName :
func (*CatalogProductIndexPriceFinalIdx) TableName() string {
	return "catalog_product_index_price_final_idx"
}

// handler create
func initRoutersCatalogProductIndexPriceFinalIdx(r *gin.Engine, catalogproductindexpricefinalidx string) {
	route := r.Group("catalogproductindexpricefinalidx")
	route.GET("/", getCatalogProductIndexPriceFinalIdxs)
	route.GET("/:id", getCatalogProductIndexPriceFinalIdx)
	route.POST("/", createCatalogProductIndexPriceFinalIdx)
	route.PUT("/:id", updateCatalogProductIndexPriceFinalIdx)
	route.DELETE("/:id", deleteCatalogProductIndexPriceFinalIdx)
}

func getCatalogProductIndexPriceFinalIdxs(c *gin.Context) {
	var catalogProductIndexPriceFinalIdxs []CatalogProductIndexPriceFinalIdx
	if err := g.Find(&catalogProductIndexPriceFinalIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceFinalIdxs)
	}
}

func getCatalogProductIndexPriceFinalIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceFinalIdx CatalogProductIndexPriceFinalIdx
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceFinalIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceFinalIdx)
	}
}

func createCatalogProductIndexPriceFinalIdx(c *gin.Context) {
	var catalogProductIndexPriceFinalIdx CatalogProductIndexPriceFinalIdx
	c.BindJSON(&catalogProductIndexPriceFinalIdx)
	g.Create(&catalogProductIndexPriceFinalIdx)
	c.JSON(200, catalogProductIndexPriceFinalIdx)
}

func updateCatalogProductIndexPriceFinalIdx(c *gin.Context) {
	var catalogProductIndexPriceFinalIdx CatalogProductIndexPriceFinalIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceFinalIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceFinalIdx)
	g.Save(&catalogProductIndexPriceFinalIdx)
	c.JSON(200, catalogProductIndexPriceFinalIdx)
}
func deleteCatalogProductIndexPriceFinalIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceFinalIdx CatalogProductIndexPriceFinalIdx
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceFinalIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
