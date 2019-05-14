package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceOptAgrTmp :
type CatalogProductIndexPriceOptAgrTmp struct {
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
func (*CatalogProductIndexPriceOptAgrTmp) TableName() string {
	return "catalog_product_index_price_opt_agr_tmp"
}

// handler create
func initRoutersCatalogProductIndexPriceOptAgrTmp(r *gin.Engine, catalogproductindexpriceoptagrtmp string) {
	route := r.Group("catalogproductindexpriceoptagrtmp")
	route.GET("/", getCatalogProductIndexPriceOptAgrTmps)
	route.GET("/:id", getCatalogProductIndexPriceOptAgrTmp)
	route.POST("/", createCatalogProductIndexPriceOptAgrTmp)
	route.PUT("/:id", updateCatalogProductIndexPriceOptAgrTmp)
	route.DELETE("/:id", deleteCatalogProductIndexPriceOptAgrTmp)
}

func getCatalogProductIndexPriceOptAgrTmps(c *gin.Context) {
	var catalogProductIndexPriceOptAgrTmps []CatalogProductIndexPriceOptAgrTmp
	if err := g.Find(&catalogProductIndexPriceOptAgrTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceOptAgrTmps)
	}
}

func getCatalogProductIndexPriceOptAgrTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceOptAgrTmp CatalogProductIndexPriceOptAgrTmp
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceOptAgrTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceOptAgrTmp)
	}
}

func createCatalogProductIndexPriceOptAgrTmp(c *gin.Context) {
	var catalogProductIndexPriceOptAgrTmp CatalogProductIndexPriceOptAgrTmp
	c.BindJSON(&catalogProductIndexPriceOptAgrTmp)
	g.Create(&catalogProductIndexPriceOptAgrTmp)
	c.JSON(200, catalogProductIndexPriceOptAgrTmp)
}

func updateCatalogProductIndexPriceOptAgrTmp(c *gin.Context) {
	var catalogProductIndexPriceOptAgrTmp CatalogProductIndexPriceOptAgrTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceOptAgrTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceOptAgrTmp)
	g.Save(&catalogProductIndexPriceOptAgrTmp)
	c.JSON(200, catalogProductIndexPriceOptAgrTmp)
}
func deleteCatalogProductIndexPriceOptAgrTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceOptAgrTmp CatalogProductIndexPriceOptAgrTmp
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceOptAgrTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
