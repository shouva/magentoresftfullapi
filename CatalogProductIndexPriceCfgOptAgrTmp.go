package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceCfgOptAgrTmp :
type CatalogProductIndexPriceCfgOptAgrTmp struct {
	ParentId        uint16  `gorm:"column:parent_id;primary_key" form:"parent_id;primary_key" json:"parent_id;primary_key"`
	ChildId         uint16  `gorm:"column:child_id;primary_key" form:"child_id;primary_key" json:"child_id;primary_key"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	WebsiteId       uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	Price           float64 `gorm:"column:price" form:"price" json:"price"`
	TierPrice       float64 `gorm:"column:tier_price" form:"tier_price" json:"tier_price"`
	GroupPrice      float64 `gorm:"column:group_price" form:"group_price" json:"group_price"`
}

// TableName :
func (*CatalogProductIndexPriceCfgOptAgrTmp) TableName() string {
	return "catalog_product_index_price_cfg_opt_agr_tmp"
}

// handler create
func initRoutersCatalogProductIndexPriceCfgOptAgrTmp(r *gin.Engine, catalogproductindexpricecfgoptagrtmp string) {
	route := r.Group("catalogproductindexpricecfgoptagrtmp")
	route.GET("/", getCatalogProductIndexPriceCfgOptAgrTmps)
	route.GET("/:id", getCatalogProductIndexPriceCfgOptAgrTmp)
	route.POST("/", createCatalogProductIndexPriceCfgOptAgrTmp)
	route.PUT("/:id", updateCatalogProductIndexPriceCfgOptAgrTmp)
	route.DELETE("/:id", deleteCatalogProductIndexPriceCfgOptAgrTmp)
}

func getCatalogProductIndexPriceCfgOptAgrTmps(c *gin.Context) {
	var catalogProductIndexPriceCfgOptAgrTmps []CatalogProductIndexPriceCfgOptAgrTmp
	if err := g.Find(&catalogProductIndexPriceCfgOptAgrTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceCfgOptAgrTmps)
	}
}

func getCatalogProductIndexPriceCfgOptAgrTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceCfgOptAgrTmp CatalogProductIndexPriceCfgOptAgrTmp
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceCfgOptAgrTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceCfgOptAgrTmp)
	}
}

func createCatalogProductIndexPriceCfgOptAgrTmp(c *gin.Context) {
	var catalogProductIndexPriceCfgOptAgrTmp CatalogProductIndexPriceCfgOptAgrTmp
	c.BindJSON(&catalogProductIndexPriceCfgOptAgrTmp)
	g.Create(&catalogProductIndexPriceCfgOptAgrTmp)
	c.JSON(200, catalogProductIndexPriceCfgOptAgrTmp)
}

func updateCatalogProductIndexPriceCfgOptAgrTmp(c *gin.Context) {
	var catalogProductIndexPriceCfgOptAgrTmp CatalogProductIndexPriceCfgOptAgrTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceCfgOptAgrTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceCfgOptAgrTmp)
	g.Save(&catalogProductIndexPriceCfgOptAgrTmp)
	c.JSON(200, catalogProductIndexPriceCfgOptAgrTmp)
}
func deleteCatalogProductIndexPriceCfgOptAgrTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceCfgOptAgrTmp CatalogProductIndexPriceCfgOptAgrTmp
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceCfgOptAgrTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
