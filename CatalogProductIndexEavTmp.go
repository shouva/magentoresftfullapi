package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexEavTmp :
type CatalogProductIndexEavTmp struct {
	EntityId    uint16 `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	AttributeId uint16 `gorm:"column:attribute_id;primary_key" form:"attribute_id;primary_key" json:"attribute_id;primary_key"`
	StoreId     uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Value       uint16 `gorm:"column:value;primary_key" form:"value;primary_key" json:"value;primary_key"`
}

// TableName :
func (*CatalogProductIndexEavTmp) TableName() string {
	return "catalog_product_index_eav_tmp"
}

// handler create
func initRoutersCatalogProductIndexEavTmp(r *gin.Engine, catalogproductindexeavtmp string) {
	route := r.Group("catalogproductindexeavtmp")
	route.GET("/", getCatalogProductIndexEavTmps)
	route.GET("/:id", getCatalogProductIndexEavTmp)
	route.POST("/", createCatalogProductIndexEavTmp)
	route.PUT("/:id", updateCatalogProductIndexEavTmp)
	route.DELETE("/:id", deleteCatalogProductIndexEavTmp)
}

func getCatalogProductIndexEavTmps(c *gin.Context) {
	var catalogProductIndexEavTmps []CatalogProductIndexEavTmp
	if err := g.Find(&catalogProductIndexEavTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexEavTmps)
	}
}

func getCatalogProductIndexEavTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexEavTmp CatalogProductIndexEavTmp
	if err := g.Where("id = ?", id).First(&catalogProductIndexEavTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexEavTmp)
	}
}

func createCatalogProductIndexEavTmp(c *gin.Context) {
	var catalogProductIndexEavTmp CatalogProductIndexEavTmp
	c.BindJSON(&catalogProductIndexEavTmp)
	g.Create(&catalogProductIndexEavTmp)
	c.JSON(200, catalogProductIndexEavTmp)
}

func updateCatalogProductIndexEavTmp(c *gin.Context) {
	var catalogProductIndexEavTmp CatalogProductIndexEavTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexEavTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexEavTmp)
	g.Save(&catalogProductIndexEavTmp)
	c.JSON(200, catalogProductIndexEavTmp)
}
func deleteCatalogProductIndexEavTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexEavTmp CatalogProductIndexEavTmp
	d := g.Where("id = ?", id).Delete(&catalogProductIndexEavTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
