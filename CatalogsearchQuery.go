package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CatalogsearchQuery :
type CatalogsearchQuery struct {
	QueryId        uint16     `gorm:"column:query_id;primary_key" form:"query_id;primary_key" json:"query_id;primary_key"`
	QueryText      string     `gorm:"column:query_text" form:"query_text" json:"query_text"`
	NumResults     uint16     `gorm:"column:num_results" form:"num_results" json:"num_results"`
	Popularity     uint16     `gorm:"column:popularity" form:"popularity" json:"popularity"`
	Redirect       string     `gorm:"column:redirect" form:"redirect" json:"redirect"`
	SynonymFor     string     `gorm:"column:synonym_for" form:"synonym_for" json:"synonym_for"`
	StoreId        uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	DisplayInTerms uint16     `gorm:"column:display_in_terms" form:"display_in_terms" json:"display_in_terms"`
	IsActive       uint16     `gorm:"column:is_active" form:"is_active" json:"is_active"`
	IsProcessed    uint16     `gorm:"column:is_processed" form:"is_processed" json:"is_processed"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
}

// TableName :
func (*CatalogsearchQuery) TableName() string {
	return "catalogsearch_query"
}

// handler create
func initRoutersCatalogsearchQuery(r *gin.Engine, catalogsearchquery string) {
	route := r.Group("catalogsearchquery")
	route.GET("/", getCatalogsearchQuerys)
	route.GET("/:id", getCatalogsearchQuery)
	route.POST("/", createCatalogsearchQuery)
	route.PUT("/:id", updateCatalogsearchQuery)
	route.DELETE("/:id", deleteCatalogsearchQuery)
}

func getCatalogsearchQuerys(c *gin.Context) {
	var catalogsearchQuerys []CatalogsearchQuery
	if err := g.Find(&catalogsearchQuerys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogsearchQuerys)
	}
}

func getCatalogsearchQuery(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogsearchQuery CatalogsearchQuery
	if err := g.Where("id = ?", id).First(&catalogsearchQuery).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogsearchQuery)
	}
}

func createCatalogsearchQuery(c *gin.Context) {
	var catalogsearchQuery CatalogsearchQuery
	c.BindJSON(&catalogsearchQuery)
	g.Create(&catalogsearchQuery)
	c.JSON(200, catalogsearchQuery)
}

func updateCatalogsearchQuery(c *gin.Context) {
	var catalogsearchQuery CatalogsearchQuery
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogsearchQuery).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogsearchQuery)
	g.Save(&catalogsearchQuery)
	c.JSON(200, catalogsearchQuery)
}
func deleteCatalogsearchQuery(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogsearchQuery CatalogsearchQuery
	d := g.Where("id = ?", id).Delete(&catalogsearchQuery)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
