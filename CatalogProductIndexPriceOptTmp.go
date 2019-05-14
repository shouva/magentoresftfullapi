package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceOptTmp :
type CatalogProductIndexPriceOptTmp struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	MaxPrice        float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
	TierPrice       float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	GroupPrice      float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
}

// TableName :
func (*CatalogProductIndexPriceOptTmp) TableName() string {
	return "catalog_product_index_price_opt_tmp"
}

// handler create
func initRoutersCatalogProductIndexPriceOptTmp(r *gin.Engine, catalogproductindexpriceopttmp string) {
	route := r.Group("catalogproductindexpriceopttmp")
	route.GET("/", getCatalogProductIndexPriceOptTmps)
	route.GET("/:id", getCatalogProductIndexPriceOptTmp)
	route.POST("/", createCatalogProductIndexPriceOptTmp)
	route.PUT("/:id", updateCatalogProductIndexPriceOptTmp)
	route.DELETE("/:id", deleteCatalogProductIndexPriceOptTmp)
}

func getCatalogProductIndexPriceOptTmps(c *gin.Context) {
	var catalogProductIndexPriceOptTmps []CatalogProductIndexPriceOptTmp
	if err := g.Find(&catalogProductIndexPriceOptTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceOptTmps)
	}
}

func getCatalogProductIndexPriceOptTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceOptTmp CatalogProductIndexPriceOptTmp
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceOptTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceOptTmp)
	}
}

func createCatalogProductIndexPriceOptTmp(c *gin.Context) {
	var catalogProductIndexPriceOptTmp CatalogProductIndexPriceOptTmp
	c.BindJSON(&catalogProductIndexPriceOptTmp)
	g.Create(&catalogProductIndexPriceOptTmp)
	c.JSON(200, catalogProductIndexPriceOptTmp)
}

func updateCatalogProductIndexPriceOptTmp(c *gin.Context) {
	var catalogProductIndexPriceOptTmp CatalogProductIndexPriceOptTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceOptTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceOptTmp)
	g.Save(&catalogProductIndexPriceOptTmp)
	c.JSON(200, catalogProductIndexPriceOptTmp)
}
func deleteCatalogProductIndexPriceOptTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceOptTmp CatalogProductIndexPriceOptTmp
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceOptTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
