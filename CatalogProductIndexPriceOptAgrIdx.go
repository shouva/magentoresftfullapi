package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceOptAgrIdx :
type CatalogProductIndexPriceOptAgrIdx struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	OptionId        uint16  `gorm:"column:option_id;primary_key" form:"option_id;primary_key" json:"option_id;primary_key"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	MaxPrice        float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
	TierPrice       float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	GroupPrice      float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
}

// TableName :
func (*CatalogProductIndexPriceOptAgrIdx) TableName() string {
	return "catalog_product_index_price_opt_agr_idx"
}

// handler create
func initRoutersCatalogProductIndexPriceOptAgrIdx(r *gin.Engine, catalogproductindexpriceoptagridx string) {
	route := r.Group("catalogproductindexpriceoptagridx")
	route.GET("/", getCatalogProductIndexPriceOptAgrIdxs)
	route.GET("/:id", getCatalogProductIndexPriceOptAgrIdx)
	route.POST("/", createCatalogProductIndexPriceOptAgrIdx)
	route.PUT("/:id", updateCatalogProductIndexPriceOptAgrIdx)
	route.DELETE("/:id", deleteCatalogProductIndexPriceOptAgrIdx)
}

func getCatalogProductIndexPriceOptAgrIdxs(c *gin.Context) {
	var catalogProductIndexPriceOptAgrIdxs []CatalogProductIndexPriceOptAgrIdx
	if err := g.Find(&catalogProductIndexPriceOptAgrIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceOptAgrIdxs)
	}
}

func getCatalogProductIndexPriceOptAgrIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceOptAgrIdx CatalogProductIndexPriceOptAgrIdx
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceOptAgrIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceOptAgrIdx)
	}
}

func createCatalogProductIndexPriceOptAgrIdx(c *gin.Context) {
	var catalogProductIndexPriceOptAgrIdx CatalogProductIndexPriceOptAgrIdx
	c.BindJSON(&catalogProductIndexPriceOptAgrIdx)
	g.Create(&catalogProductIndexPriceOptAgrIdx)
	c.JSON(200, catalogProductIndexPriceOptAgrIdx)
}

func updateCatalogProductIndexPriceOptAgrIdx(c *gin.Context) {
	var catalogProductIndexPriceOptAgrIdx CatalogProductIndexPriceOptAgrIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceOptAgrIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceOptAgrIdx)
	g.Save(&catalogProductIndexPriceOptAgrIdx)
	c.JSON(200, catalogProductIndexPriceOptAgrIdx)
}
func deleteCatalogProductIndexPriceOptAgrIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceOptAgrIdx CatalogProductIndexPriceOptAgrIdx
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceOptAgrIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
