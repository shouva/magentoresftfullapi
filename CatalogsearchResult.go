package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogsearchResult :
type CatalogsearchResult struct {
	QueryId   uint16  `gorm:"column:query_id;primary_key" form:"query_id;primary_key" json:"query_id;primary_key"`
	ProductId uint16  `gorm:"column:product_id;primary_key" form:"product_id;primary_key" json:"product_id;primary_key"`
	Relevance float64 `gorm:"column:relevance" form:"relevance" json:"relevance"`
}

// TableName :
func (*CatalogsearchResult) TableName() string {
	return "catalogsearch_result"
}

// handler create
func initRoutersCatalogsearchResult(r *gin.Engine, catalogsearchresult string) {
	route := r.Group("catalogsearchresult")
	route.GET("/", getCatalogsearchResults)
	route.GET("/:id", getCatalogsearchResult)
	route.POST("/", createCatalogsearchResult)
	route.PUT("/:id", updateCatalogsearchResult)
	route.DELETE("/:id", deleteCatalogsearchResult)
}

func getCatalogsearchResults(c *gin.Context) {
	var catalogsearchResults []CatalogsearchResult
	if err := g.Find(&catalogsearchResults).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogsearchResults)
	}
}

func getCatalogsearchResult(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogsearchResult CatalogsearchResult
	if err := g.Where("id = ?", id).First(&catalogsearchResult).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogsearchResult)
	}
}

func createCatalogsearchResult(c *gin.Context) {
	var catalogsearchResult CatalogsearchResult
	c.BindJSON(&catalogsearchResult)
	g.Create(&catalogsearchResult)
	c.JSON(200, catalogsearchResult)
}

func updateCatalogsearchResult(c *gin.Context) {
	var catalogsearchResult CatalogsearchResult
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogsearchResult).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogsearchResult)
	g.Save(&catalogsearchResult)
	c.JSON(200, catalogsearchResult)
}
func deleteCatalogsearchResult(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogsearchResult CatalogsearchResult
	d := g.Where("id = ?", id).Delete(&catalogsearchResult)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
