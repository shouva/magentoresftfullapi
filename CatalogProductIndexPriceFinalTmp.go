package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceFinalTmp :
type CatalogProductIndexPriceFinalTmp struct {
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
func (*CatalogProductIndexPriceFinalTmp) TableName() string {
	return "catalog_product_index_price_final_tmp"
}

// handler create
func initRoutersCatalogProductIndexPriceFinalTmp(r *gin.Engine, catalogproductindexpricefinaltmp string) {
	route := r.Group("catalogproductindexpricefinaltmp")
	route.GET("/", getCatalogProductIndexPriceFinalTmps)
	route.GET("/:id", getCatalogProductIndexPriceFinalTmp)
	route.POST("/", createCatalogProductIndexPriceFinalTmp)
	route.PUT("/:id", updateCatalogProductIndexPriceFinalTmp)
	route.DELETE("/:id", deleteCatalogProductIndexPriceFinalTmp)
}

func getCatalogProductIndexPriceFinalTmps(c *gin.Context) {
	var catalogProductIndexPriceFinalTmps []CatalogProductIndexPriceFinalTmp
	if err := g.Find(&catalogProductIndexPriceFinalTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceFinalTmps)
	}
}

func getCatalogProductIndexPriceFinalTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceFinalTmp CatalogProductIndexPriceFinalTmp
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceFinalTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceFinalTmp)
	}
}

func createCatalogProductIndexPriceFinalTmp(c *gin.Context) {
	var catalogProductIndexPriceFinalTmp CatalogProductIndexPriceFinalTmp
	c.BindJSON(&catalogProductIndexPriceFinalTmp)
	g.Create(&catalogProductIndexPriceFinalTmp)
	c.JSON(200, catalogProductIndexPriceFinalTmp)
}

func updateCatalogProductIndexPriceFinalTmp(c *gin.Context) {
	var catalogProductIndexPriceFinalTmp CatalogProductIndexPriceFinalTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceFinalTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceFinalTmp)
	g.Save(&catalogProductIndexPriceFinalTmp)
	c.JSON(200, catalogProductIndexPriceFinalTmp)
}
func deleteCatalogProductIndexPriceFinalTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceFinalTmp CatalogProductIndexPriceFinalTmp
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceFinalTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
