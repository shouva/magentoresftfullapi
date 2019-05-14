package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexEavIdx :
type CatalogProductIndexEavIdx struct {
	EntityId    uint16 `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	AttributeId uint16 `gorm:"column:attribute_id;primary_key" form:"attribute_id;primary_key" json:"attribute_id;primary_key"`
	StoreId     uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Value       uint16 `gorm:"column:value;primary_key" form:"value;primary_key" json:"value;primary_key"`
}

// TableName :
func (*CatalogProductIndexEavIdx) TableName() string {
	return "catalog_product_index_eav_idx"
}

// handler create
func initRoutersCatalogProductIndexEavIdx(r *gin.Engine, catalogproductindexeavidx string) {
	route := r.Group("catalogproductindexeavidx")
	route.GET("/", getCatalogProductIndexEavIdxs)
	route.GET("/:id", getCatalogProductIndexEavIdx)
	route.POST("/", createCatalogProductIndexEavIdx)
	route.PUT("/:id", updateCatalogProductIndexEavIdx)
	route.DELETE("/:id", deleteCatalogProductIndexEavIdx)
}

func getCatalogProductIndexEavIdxs(c *gin.Context) {
	var catalogProductIndexEavIdxs []CatalogProductIndexEavIdx
	if err := g.Find(&catalogProductIndexEavIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexEavIdxs)
	}
}

func getCatalogProductIndexEavIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexEavIdx CatalogProductIndexEavIdx
	if err := g.Where("id = ?", id).First(&catalogProductIndexEavIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexEavIdx)
	}
}

func createCatalogProductIndexEavIdx(c *gin.Context) {
	var catalogProductIndexEavIdx CatalogProductIndexEavIdx
	c.BindJSON(&catalogProductIndexEavIdx)
	g.Create(&catalogProductIndexEavIdx)
	c.JSON(200, catalogProductIndexEavIdx)
}

func updateCatalogProductIndexEavIdx(c *gin.Context) {
	var catalogProductIndexEavIdx CatalogProductIndexEavIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexEavIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexEavIdx)
	g.Save(&catalogProductIndexEavIdx)
	c.JSON(200, catalogProductIndexEavIdx)
}
func deleteCatalogProductIndexEavIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexEavIdx CatalogProductIndexEavIdx
	d := g.Where("id = ?", id).Delete(&catalogProductIndexEavIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
