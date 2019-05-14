package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexEavDecimal :
type CatalogProductIndexEavDecimal struct {
	EntityId    uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	AttributeId uint16  `gorm:"column:attribute_id;primary_key" form:"attribute_id;primary_key" json:"attribute_id;primary_key"`
	StoreId     uint16  `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Value       float64 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*CatalogProductIndexEavDecimal) TableName() string {
	return "catalog_product_index_eav_decimal"
}

// handler create
func initRoutersCatalogProductIndexEavDecimal(r *gin.Engine, catalogproductindexeavdecimal string) {
	route := r.Group("catalogproductindexeavdecimal")
	route.GET("/", getCatalogProductIndexEavDecimals)
	route.GET("/:id", getCatalogProductIndexEavDecimal)
	route.POST("/", createCatalogProductIndexEavDecimal)
	route.PUT("/:id", updateCatalogProductIndexEavDecimal)
	route.DELETE("/:id", deleteCatalogProductIndexEavDecimal)
}

func getCatalogProductIndexEavDecimals(c *gin.Context) {
	var catalogProductIndexEavDecimals []CatalogProductIndexEavDecimal
	if err := g.Find(&catalogProductIndexEavDecimals).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexEavDecimals)
	}
}

func getCatalogProductIndexEavDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexEavDecimal CatalogProductIndexEavDecimal
	if err := g.Where("id = ?", id).First(&catalogProductIndexEavDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexEavDecimal)
	}
}

func createCatalogProductIndexEavDecimal(c *gin.Context) {
	var catalogProductIndexEavDecimal CatalogProductIndexEavDecimal
	c.BindJSON(&catalogProductIndexEavDecimal)
	g.Create(&catalogProductIndexEavDecimal)
	c.JSON(200, catalogProductIndexEavDecimal)
}

func updateCatalogProductIndexEavDecimal(c *gin.Context) {
	var catalogProductIndexEavDecimal CatalogProductIndexEavDecimal
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexEavDecimal).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexEavDecimal)
	g.Save(&catalogProductIndexEavDecimal)
	c.JSON(200, catalogProductIndexEavDecimal)
}
func deleteCatalogProductIndexEavDecimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexEavDecimal CatalogProductIndexEavDecimal
	d := g.Where("id = ?", id).Delete(&catalogProductIndexEavDecimal)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
