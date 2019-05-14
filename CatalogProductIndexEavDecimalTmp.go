package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexEavDecimalTmp :
type CatalogProductIndexEavDecimalTmp struct {
	EntityId    uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	AttributeId uint16  `gorm:"column:attribute_id;primary_key" form:"attribute_id;primary_key" json:"attribute_id;primary_key"`
	StoreId     uint16  `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Value       float64 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductIndexEavDecimalTmp) TableName() string {
	return "catalog_product_index_eav_decimal_tmp"
}

// handler create
func initRoutersCatalogProductIndexEavDecimalTmp(r *gin.Engine, catalogproductindexeavdecimaltmp string) {
	route := r.Group("catalogproductindexeavdecimaltmp")
	route.GET("/", getCatalogProductIndexEavDecimalTmps)
	route.GET("/:id", getCatalogProductIndexEavDecimalTmp)
	route.POST("/", createCatalogProductIndexEavDecimalTmp)
	route.PUT("/:id", updateCatalogProductIndexEavDecimalTmp)
	route.DELETE("/:id", deleteCatalogProductIndexEavDecimalTmp)
}

func getCatalogProductIndexEavDecimalTmps(c *gin.Context) {
	var catalogProductIndexEavDecimalTmps []CatalogProductIndexEavDecimalTmp
	if err := g.Find(&catalogProductIndexEavDecimalTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexEavDecimalTmps)
	}
}

func getCatalogProductIndexEavDecimalTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexEavDecimalTmp CatalogProductIndexEavDecimalTmp
	if err := g.Where("id = ?", id).First(&catalogProductIndexEavDecimalTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexEavDecimalTmp)
	}
}

func createCatalogProductIndexEavDecimalTmp(c *gin.Context) {
	var catalogProductIndexEavDecimalTmp CatalogProductIndexEavDecimalTmp
	c.BindJSON(&catalogProductIndexEavDecimalTmp)
	g.Create(&catalogProductIndexEavDecimalTmp)
	c.JSON(200, catalogProductIndexEavDecimalTmp)
}

func updateCatalogProductIndexEavDecimalTmp(c *gin.Context) {
	var catalogProductIndexEavDecimalTmp CatalogProductIndexEavDecimalTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexEavDecimalTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexEavDecimalTmp)
	g.Save(&catalogProductIndexEavDecimalTmp)
	c.JSON(200, catalogProductIndexEavDecimalTmp)
}
func deleteCatalogProductIndexEavDecimalTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexEavDecimalTmp CatalogProductIndexEavDecimalTmp
	d := g.Where("id = ?", id).Delete(&catalogProductIndexEavDecimalTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
