package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogProductIndexEav :
type CatalogProductIndexEav struct {
	EntityId    uint16 `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	AttributeId uint16 `gorm:"column:attribute_id;primary_key" form:"attribute_id;primary_key" json:"attribute_id;primary_key"`
	StoreId     uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Value       uint16 `gorm:"column:value;primary_key" form:"value;primary_key" json:"value;primary_key"`
}

// TableName :
func (*CatalogProductIndexEav) TableName() string {
	return "catalog_product_index_eav"
}

// handler create
func initRoutersCatalogProductIndexEav(r *gin.Engine, catalogproductindexeav string) {
	route := r.Group("catalogproductindexeav")
	route.GET("/", getCatalogProductIndexEavs)
	route.GET("/:id", getCatalogProductIndexEav)
	route.POST("/", createCatalogProductIndexEav)
	route.PUT("/:id", updateCatalogProductIndexEav)
	route.DELETE("/:id", deleteCatalogProductIndexEav)
}

func getCatalogProductIndexEavs(c *gin.Context) {
	var catalogProductIndexEavs []CatalogProductIndexEav
	if err := g.Find(&catalogProductIndexEavs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexEavs)
	}
}

func getCatalogProductIndexEav(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexEav CatalogProductIndexEav
	if err := g.Where("id = ?", id).First(&catalogProductIndexEav).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogProductIndexEav)
	}
}

func createCatalogProductIndexEav(c *gin.Context) {
	var catalogProductIndexEav CatalogProductIndexEav
	c.BindJSON(&catalogProductIndexEav)
	g.Create(&catalogProductIndexEav)
	c.JSON(200, catalogProductIndexEav)
}

func updateCatalogProductIndexEav(c *gin.Context) {
	var catalogProductIndexEav CatalogProductIndexEav
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogProductIndexEav).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogProductIndexEav)
	g.Save(&catalogProductIndexEav)
	c.JSON(200, catalogProductIndexEav)
}
func deleteCatalogProductIndexEav(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogProductIndexEav CatalogProductIndexEav
	d := g.Where("id = ?", id).Delete(&catalogProductIndexEav)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
