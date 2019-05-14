package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogsearchFulltext :
type CatalogsearchFulltext struct {
	FulltextId uint16 `gorm:"column:fulltext_id;primary_key" form:"fulltext_id;primary_key" json:"fulltext_id;primary_key"`
	ProductId  uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
	StoreId    uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	DataIndex  string `gorm:"column:data_index" form:"data_index" json:"data_index"`
}

// TableName :
func (*CatalogsearchFulltext) TableName() string {
	return "catalogsearch_fulltext"
}

// handler create
func initRoutersCatalogsearchFulltext(r *gin.Engine, catalogsearchfulltext string) {
	route := r.Group("catalogsearchfulltext")
	route.GET("/", getCatalogsearchFulltexts)
	route.GET("/:id", getCatalogsearchFulltext)
	route.POST("/", createCatalogsearchFulltext)
	route.PUT("/:id", updateCatalogsearchFulltext)
	route.DELETE("/:id", deleteCatalogsearchFulltext)
}

func getCatalogsearchFulltexts(c *gin.Context) {
	var catalogsearchFulltexts []CatalogsearchFulltext
	if err := g.Find(&catalogsearchFulltexts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogsearchFulltexts)
	}
}

func getCatalogsearchFulltext(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogsearchFulltext CatalogsearchFulltext
	if err := g.Where("id = ?", id).First(&catalogsearchFulltext).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogsearchFulltext)
	}
}

func createCatalogsearchFulltext(c *gin.Context) {
	var catalogsearchFulltext CatalogsearchFulltext
	c.BindJSON(&catalogsearchFulltext)
	g.Create(&catalogsearchFulltext)
	c.JSON(200, catalogsearchFulltext)
}

func updateCatalogsearchFulltext(c *gin.Context) {
	var catalogsearchFulltext CatalogsearchFulltext
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogsearchFulltext).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogsearchFulltext)
	g.Save(&catalogsearchFulltext)
	c.JSON(200, catalogsearchFulltext)
}
func deleteCatalogsearchFulltext(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogsearchFulltext CatalogsearchFulltext
	d := g.Where("id = ?", id).Delete(&catalogsearchFulltext)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
