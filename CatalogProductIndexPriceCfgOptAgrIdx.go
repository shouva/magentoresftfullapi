package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceCfgOptAgrIdx :
type CatalogProductIndexPriceCfgOptAgrIdx struct {
	ParentId        uint16  `gorm:"column:parent_id;primary_key" form:"parent_id;primary_key" json:"parent_id;primary_key"`
	ChildId         uint16  `gorm:"column:child_id;primary_key" form:"child_id;primary_key" json:"child_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	Price           float64 `gorm:"column:price" form:"price" json:"price"`
	TierPrice       float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	GroupPrice      float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
}

// TableName :
func (*CatalogProductIndexPriceCfgOptAgrIdx) TableName() string {
	return "catalog_product_index_price_cfg_opt_agr_idx"
}

// handler create
func initRoutersCatalogProductIndexPriceCfgOptAgrIdx(r *gin.Engine, catalogproductindexpricecfgoptagridx string) {
	route := r.Group("catalogproductindexpricecfgoptagridx")
	route.GET("/", getCatalogProductIndexPriceCfgOptAgrIdxs)
	route.GET("/:id", getCatalogProductIndexPriceCfgOptAgrIdx)
	route.POST("/", createCatalogProductIndexPriceCfgOptAgrIdx)
	route.PUT("/:id", updateCatalogProductIndexPriceCfgOptAgrIdx)
	route.DELETE("/:id", deleteCatalogProductIndexPriceCfgOptAgrIdx)
}

func getCatalogProductIndexPriceCfgOptAgrIdxs(c *gin.Context) {
	var catalogProductIndexPriceCfgOptAgrIdxs []CatalogProductIndexPriceCfgOptAgrIdx
	if err := g.Find(&catalogProductIndexPriceCfgOptAgrIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceCfgOptAgrIdxs)
	}
}

func getCatalogProductIndexPriceCfgOptAgrIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceCfgOptAgrIdx CatalogProductIndexPriceCfgOptAgrIdx
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceCfgOptAgrIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceCfgOptAgrIdx)
	}
}

func createCatalogProductIndexPriceCfgOptAgrIdx(c *gin.Context) {
	var catalogProductIndexPriceCfgOptAgrIdx CatalogProductIndexPriceCfgOptAgrIdx
	c.BindJSON(&catalogProductIndexPriceCfgOptAgrIdx)
	g.Create(&catalogProductIndexPriceCfgOptAgrIdx)
	c.JSON(200, catalogProductIndexPriceCfgOptAgrIdx)
}

func updateCatalogProductIndexPriceCfgOptAgrIdx(c *gin.Context) {
	var catalogProductIndexPriceCfgOptAgrIdx CatalogProductIndexPriceCfgOptAgrIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceCfgOptAgrIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceCfgOptAgrIdx)
	g.Save(&catalogProductIndexPriceCfgOptAgrIdx)
	c.JSON(200, catalogProductIndexPriceCfgOptAgrIdx)
}
func deleteCatalogProductIndexPriceCfgOptAgrIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceCfgOptAgrIdx CatalogProductIndexPriceCfgOptAgrIdx
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceCfgOptAgrIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
