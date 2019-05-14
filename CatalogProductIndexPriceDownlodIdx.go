package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceDownlodIdx :
type CatalogProductIndexPriceDownlodIdx struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	MaxPrice        float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
}

// TableName :
func (*CatalogProductIndexPriceDownlodIdx) TableName() string {
	return "catalog_product_index_price_downlod_idx"
}

// handler create
func initRoutersCatalogProductIndexPriceDownlodIdx(r *gin.Engine, catalogproductindexpricedownlodidx string) {
	route := r.Group("catalogproductindexpricedownlodidx")
	route.GET("/", getCatalogProductIndexPriceDownlodIdxs)
	route.GET("/:id", getCatalogProductIndexPriceDownlodIdx)
	route.POST("/", createCatalogProductIndexPriceDownlodIdx)
	route.PUT("/:id", updateCatalogProductIndexPriceDownlodIdx)
	route.DELETE("/:id", deleteCatalogProductIndexPriceDownlodIdx)
}

func getCatalogProductIndexPriceDownlodIdxs(c *gin.Context) {
	var catalogProductIndexPriceDownlodIdxs []CatalogProductIndexPriceDownlodIdx
	if err := g.Find(&catalogProductIndexPriceDownlodIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceDownlodIdxs)
	}
}

func getCatalogProductIndexPriceDownlodIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceDownlodIdx CatalogProductIndexPriceDownlodIdx
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceDownlodIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceDownlodIdx)
	}
}

func createCatalogProductIndexPriceDownlodIdx(c *gin.Context) {
	var catalogProductIndexPriceDownlodIdx CatalogProductIndexPriceDownlodIdx
	c.BindJSON(&catalogProductIndexPriceDownlodIdx)
	g.Create(&catalogProductIndexPriceDownlodIdx)
	c.JSON(200, catalogProductIndexPriceDownlodIdx)
}

func updateCatalogProductIndexPriceDownlodIdx(c *gin.Context) {
	var catalogProductIndexPriceDownlodIdx CatalogProductIndexPriceDownlodIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceDownlodIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceDownlodIdx)
	g.Save(&catalogProductIndexPriceDownlodIdx)
	c.JSON(200, catalogProductIndexPriceDownlodIdx)
}
func deleteCatalogProductIndexPriceDownlodIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceDownlodIdx CatalogProductIndexPriceDownlodIdx
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceDownlodIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
