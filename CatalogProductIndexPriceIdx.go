package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceIdx :
type CatalogProductIndexPriceIdx struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	TaxClassId      uint16  `gorm:"column:tax_class_id" form:"tax_class_id" json:"tax_class_id"`
	Price           float64 `gorm:"column:price" form:"price" json:"price"`
	FinalPrice      float64 `gorm:"column:final_price" form:"final_price" json:"final_price"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	MaxPrice        float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
	TierPrice       float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	GroupPrice      float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
}

// TableName :
func (*CatalogProductIndexPriceIdx) TableName() string {
	return "catalog_product_index_price_idx"
}

// handler create
func initRoutersCatalogProductIndexPriceIdx(r *gin.Engine, catalogproductindexpriceidx string) {
	route := r.Group("catalogproductindexpriceidx")
	route.GET("/", getCatalogProductIndexPriceIdxs)
	route.GET("/:id", getCatalogProductIndexPriceIdx)
	route.POST("/", createCatalogProductIndexPriceIdx)
	route.PUT("/:id", updateCatalogProductIndexPriceIdx)
	route.DELETE("/:id", deleteCatalogProductIndexPriceIdx)
}

func getCatalogProductIndexPriceIdxs(c *gin.Context) {
	var catalogProductIndexPriceIdxs []CatalogProductIndexPriceIdx
	if err := g.Find(&catalogProductIndexPriceIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceIdxs)
	}
}

func getCatalogProductIndexPriceIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceIdx CatalogProductIndexPriceIdx
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceIdx)
	}
}

func createCatalogProductIndexPriceIdx(c *gin.Context) {
	var catalogProductIndexPriceIdx CatalogProductIndexPriceIdx
	c.BindJSON(&catalogProductIndexPriceIdx)
	g.Create(&catalogProductIndexPriceIdx)
	c.JSON(200, catalogProductIndexPriceIdx)
}

func updateCatalogProductIndexPriceIdx(c *gin.Context) {
	var catalogProductIndexPriceIdx CatalogProductIndexPriceIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceIdx)
	g.Save(&catalogProductIndexPriceIdx)
	c.JSON(200, catalogProductIndexPriceIdx)
}
func deleteCatalogProductIndexPriceIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceIdx CatalogProductIndexPriceIdx
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
