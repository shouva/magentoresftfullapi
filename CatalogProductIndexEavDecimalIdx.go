package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexEavDecimalIdx :
type CatalogProductIndexEavDecimalIdx struct {
	EntityId    uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	AttributeId uint16  `gorm:"column:attribute_id;primary_key" form:"attribute_id;primary_key" json:"attribute_id;primary_key"`
	StoreId     uint16  `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Value       float64 `gorm:"column:value;primary_key" form:"value;primary_key" json:"value;primary_key"`
}

// TableName :
func (*CatalogProductIndexEavDecimalIdx) TableName() string {
	return "catalog_product_index_eav_decimal_idx"
}

// handler create
func initRoutersCatalogProductIndexEavDecimalIdx(r *gin.Engine, catalogproductindexeavdecimalidx string) {
	route := r.Group("catalogproductindexeavdecimalidx")
	route.GET("/", getCatalogProductIndexEavDecimalIdxs)
	route.GET("/:id", getCatalogProductIndexEavDecimalIdx)
	route.POST("/", createCatalogProductIndexEavDecimalIdx)
	route.PUT("/:id", updateCatalogProductIndexEavDecimalIdx)
	route.DELETE("/:id", deleteCatalogProductIndexEavDecimalIdx)
}

func getCatalogProductIndexEavDecimalIdxs(c *gin.Context) {
	var catalogProductIndexEavDecimalIdxs []CatalogProductIndexEavDecimalIdx
	if err := g.Find(&catalogProductIndexEavDecimalIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexEavDecimalIdxs)
	}
}

func getCatalogProductIndexEavDecimalIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexEavDecimalIdx CatalogProductIndexEavDecimalIdx
	if err := g.Where("id = ?", id).First(&catalogProductIndexEavDecimalIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexEavDecimalIdx)
	}
}

func createCatalogProductIndexEavDecimalIdx(c *gin.Context) {
	var catalogProductIndexEavDecimalIdx CatalogProductIndexEavDecimalIdx
	c.BindJSON(&catalogProductIndexEavDecimalIdx)
	g.Create(&catalogProductIndexEavDecimalIdx)
	c.JSON(200, catalogProductIndexEavDecimalIdx)
}

func updateCatalogProductIndexEavDecimalIdx(c *gin.Context) {
	var catalogProductIndexEavDecimalIdx CatalogProductIndexEavDecimalIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexEavDecimalIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexEavDecimalIdx)
	g.Save(&catalogProductIndexEavDecimalIdx)
	c.JSON(200, catalogProductIndexEavDecimalIdx)
}
func deleteCatalogProductIndexEavDecimalIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexEavDecimalIdx CatalogProductIndexEavDecimalIdx
	d := g.Where("id = ?", id).Delete(&catalogProductIndexEavDecimalIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
