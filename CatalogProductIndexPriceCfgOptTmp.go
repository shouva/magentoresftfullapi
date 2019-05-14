package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceCfgOptTmp :
type CatalogProductIndexPriceCfgOptTmp struct {
	EntityId        uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	MinPrice        float64 `gorm:"column:min_price" form:"min_price" json:"min_price"`
	MaxPrice        float64 `gorm:"column:max_price" form:"max_price" json:"max_price"`
	TierPrice       float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	GroupPrice      float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
}

// TableName :
func (*CatalogProductIndexPriceCfgOptTmp) TableName() string {
	return "catalog_product_index_price_cfg_opt_tmp"
}

// handler create
func initRoutersCatalogProductIndexPriceCfgOptTmp(r *gin.Engine, catalogproductindexpricecfgopttmp string) {
	route := r.Group("catalogproductindexpricecfgopttmp")
	route.GET("/", getCatalogProductIndexPriceCfgOptTmps)
	route.GET("/:id", getCatalogProductIndexPriceCfgOptTmp)
	route.POST("/", createCatalogProductIndexPriceCfgOptTmp)
	route.PUT("/:id", updateCatalogProductIndexPriceCfgOptTmp)
	route.DELETE("/:id", deleteCatalogProductIndexPriceCfgOptTmp)
}

func getCatalogProductIndexPriceCfgOptTmps(c *gin.Context) {
	var catalogProductIndexPriceCfgOptTmps []CatalogProductIndexPriceCfgOptTmp
	if err := g.Find(&catalogProductIndexPriceCfgOptTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceCfgOptTmps)
	}
}

func getCatalogProductIndexPriceCfgOptTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceCfgOptTmp CatalogProductIndexPriceCfgOptTmp
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceCfgOptTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceCfgOptTmp)
	}
}

func createCatalogProductIndexPriceCfgOptTmp(c *gin.Context) {
	var catalogProductIndexPriceCfgOptTmp CatalogProductIndexPriceCfgOptTmp
	c.BindJSON(&catalogProductIndexPriceCfgOptTmp)
	g.Create(&catalogProductIndexPriceCfgOptTmp)
	c.JSON(200, catalogProductIndexPriceCfgOptTmp)
}

func updateCatalogProductIndexPriceCfgOptTmp(c *gin.Context) {
	var catalogProductIndexPriceCfgOptTmp CatalogProductIndexPriceCfgOptTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceCfgOptTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceCfgOptTmp)
	g.Save(&catalogProductIndexPriceCfgOptTmp)
	c.JSON(200, catalogProductIndexPriceCfgOptTmp)
}
func deleteCatalogProductIndexPriceCfgOptTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceCfgOptTmp CatalogProductIndexPriceCfgOptTmp
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceCfgOptTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
