package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TagSummary :
type TagSummary struct {
	TagId          uint16 `gorm:"column:tag_id;primary_key" form:"tag_id;primary_key" json:"tag_id;primary_key"`
	StoreId        uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Customers      uint16 `gorm:"column:customers" form:"customers" json:"customers"`
	Products       uint16 `gorm:"column:products" form:"products" json:"products"`
	Uses           uint16 `gorm:"column:uses" form:"uses" json:"uses"`
	HistoricalUses uint16 `gorm:"column:historical_uses" form:"historical_uses" json:"historical_uses"`
	Popularity     uint16 `gorm:"column:popularity" form:"popularity" json:"popularity"`
	BasePopularity uint16 `gorm:"column:base_popularity" form:"base_popularity" json:"base_popularity"`
}

// TableName :
func (*TagSummary) TableName() string {
	return "tag_summary"
}

// handler create
func initRoutersTagSummary(r *gin.Engine, tagsummary string) {
	route := r.Group("tagsummary")
	route.GET("/", getTagSummarys)
	route.GET("/:id", getTagSummary)
	route.POST("/", createTagSummary)
	route.PUT("/:id", updateTagSummary)
	route.DELETE("/:id", deleteTagSummary)
}

func getTagSummarys(c *gin.Context) {
	var tagSummarys []TagSummary
	if err := g.Find(&tagSummarys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, tagSummarys)
	}
}

func getTagSummary(c *gin.Context) {
	id := c.Params.ByName("id")
	var tagSummary TagSummary
	if err := g.Where("id = ?", id).First(&tagSummary).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, tagSummary)
	}
}

func createTagSummary(c *gin.Context) {
	var tagSummary TagSummary
	c.BindJSON(&tagSummary)
	g.Create(&tagSummary)
	c.JSON(200, tagSummary)
}

func updateTagSummary(c *gin.Context) {
	var tagSummary TagSummary
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&tagSummary).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&tagSummary)
	g.Save(&tagSummary)
	c.JSON(200, tagSummary)
}
func deleteTagSummary(c *gin.Context) {
	id := c.Params.ByName("id")
	var tagSummary TagSummary
	d := g.Where("id = ?", id).Delete(&tagSummary)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
