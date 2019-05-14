package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceDownlodTmp :
type CatalogProductIndexPriceDownlodTmp struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	MaxPrice        float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
}

// TableName :
func (*CatalogProductIndexPriceDownlodTmp) TableName() string {
	return "catalog_product_index_price_downlod_tmp"
}

// handler create
func initRoutersCatalogProductIndexPriceDownlodTmp(r *gin.Engine, catalogproductindexpricedownlodtmp string) {
	route := r.Group("catalogproductindexpricedownlodtmp")
	route.GET("/", getCatalogProductIndexPriceDownlodTmps)
	route.GET("/:id", getCatalogProductIndexPriceDownlodTmp)
	route.POST("/", createCatalogProductIndexPriceDownlodTmp)
	route.PUT("/:id", updateCatalogProductIndexPriceDownlodTmp)
	route.DELETE("/:id", deleteCatalogProductIndexPriceDownlodTmp)
}

func getCatalogProductIndexPriceDownlodTmps(c *gin.Context) {
	var catalogProductIndexPriceDownlodTmps []CatalogProductIndexPriceDownlodTmp
	if err := g.Find(&catalogProductIndexPriceDownlodTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceDownlodTmps)
	}
}

func getCatalogProductIndexPriceDownlodTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceDownlodTmp CatalogProductIndexPriceDownlodTmp
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceDownlodTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceDownlodTmp)
	}
}

func createCatalogProductIndexPriceDownlodTmp(c *gin.Context) {
	var catalogProductIndexPriceDownlodTmp CatalogProductIndexPriceDownlodTmp
	c.BindJSON(&catalogProductIndexPriceDownlodTmp)
	g.Create(&catalogProductIndexPriceDownlodTmp)
	c.JSON(200, catalogProductIndexPriceDownlodTmp)
}

func updateCatalogProductIndexPriceDownlodTmp(c *gin.Context) {
	var catalogProductIndexPriceDownlodTmp CatalogProductIndexPriceDownlodTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceDownlodTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceDownlodTmp)
	g.Save(&catalogProductIndexPriceDownlodTmp)
	c.JSON(200, catalogProductIndexPriceDownlodTmp)
}
func deleteCatalogProductIndexPriceDownlodTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceDownlodTmp CatalogProductIndexPriceDownlodTmp
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceDownlodTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
