package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexPriceTmp :
type CatalogProductIndexPriceTmp struct {
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
func (*CatalogProductIndexPriceTmp) TableName() string {
	return "catalog_product_index_price_tmp"
}

// handler create
func initRoutersCatalogProductIndexPriceTmp(r *gin.Engine, catalogproductindexpricetmp string) {
	route := r.Group("catalogproductindexpricetmp")
	route.GET("/", getCatalogProductIndexPriceTmps)
	route.GET("/:id", getCatalogProductIndexPriceTmp)
	route.POST("/", createCatalogProductIndexPriceTmp)
	route.PUT("/:id", updateCatalogProductIndexPriceTmp)
	route.DELETE("/:id", deleteCatalogProductIndexPriceTmp)
}

func getCatalogProductIndexPriceTmps(c *gin.Context) {
	var catalogProductIndexPriceTmps []CatalogProductIndexPriceTmp
	if err := g.Find(&catalogProductIndexPriceTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceTmps)
	}
}

func getCatalogProductIndexPriceTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceTmp CatalogProductIndexPriceTmp
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexPriceTmp)
	}
}

func createCatalogProductIndexPriceTmp(c *gin.Context) {
	var catalogProductIndexPriceTmp CatalogProductIndexPriceTmp
	c.BindJSON(&catalogProductIndexPriceTmp)
	g.Create(&catalogProductIndexPriceTmp)
	c.JSON(200, catalogProductIndexPriceTmp)
}

func updateCatalogProductIndexPriceTmp(c *gin.Context) {
	var catalogProductIndexPriceTmp CatalogProductIndexPriceTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexPriceTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexPriceTmp)
	g.Save(&catalogProductIndexPriceTmp)
	c.JSON(200, catalogProductIndexPriceTmp)
}
func deleteCatalogProductIndexPriceTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexPriceTmp CatalogProductIndexPriceTmp
	d := g.Where("id = ?", id).Delete(&catalogProductIndexPriceTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
