package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ReportViewedProductAggregatedDaily :
type ReportViewedProductAggregatedDaily struct {
	ID           uint16     `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	Period       *time.Time `gorm:"column:period" form:"period" json:"period"`
	StoreId      uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	ProductId    uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	ProductName  string     `gorm:"column:product_name" form:"product_name" json:"product_name"`
	ProductPrice float64    `gorm:"column:product_price" form:"product_price" json:"product_price"`
	ViewsNum     uint16     `gorm:"column:views_num" form:"views_num" json:"views_num"`
	RatingPos    uint16     `gorm:"column:rating_pos" form:"rating_pos" json:"rating_pos"`
}

// TableName :
func (*ReportViewedProductAggregatedDaily) TableName() string {
	return "report_viewed_product_aggregated_daily"
}

// handler create
func initRoutersReportViewedProductAggregatedDaily(r *gin.Engine, reportviewedproductaggregateddaily string) {
	route := r.Group("reportviewedproductaggregateddaily")
	route.GET("/", getReportViewedProductAggregatedDailys)
	route.GET("/:id", getReportViewedProductAggregatedDaily)
	route.POST("/", createReportViewedProductAggregatedDaily)
	route.PUT("/:id", updateReportViewedProductAggregatedDaily)
	route.DELETE("/:id", deleteReportViewedProductAggregatedDaily)
}

func getReportViewedProductAggregatedDailys(c *gin.Context) {
	var reportViewedProductAggregatedDailys []ReportViewedProductAggregatedDaily
	if err := g.Find(&reportViewedProductAggregatedDailys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportViewedProductAggregatedDailys)
	}
}

func getReportViewedProductAggregatedDaily(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportViewedProductAggregatedDaily ReportViewedProductAggregatedDaily
	if err := g.Where("id = ?", id).First(&reportViewedProductAggregatedDaily).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportViewedProductAggregatedDaily)
	}
}

func createReportViewedProductAggregatedDaily(c *gin.Context) {
	var reportViewedProductAggregatedDaily ReportViewedProductAggregatedDaily
	c.BindJSON(&reportViewedProductAggregatedDaily)
	g.Create(&reportViewedProductAggregatedDaily)
	c.JSON(200, reportViewedProductAggregatedDaily)
}

func updateReportViewedProductAggregatedDaily(c *gin.Context) {
	var reportViewedProductAggregatedDaily ReportViewedProductAggregatedDaily
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&reportViewedProductAggregatedDaily).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&reportViewedProductAggregatedDaily)
	g.Save(&reportViewedProductAggregatedDaily)
	c.JSON(200, reportViewedProductAggregatedDaily)
}
func deleteReportViewedProductAggregatedDaily(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportViewedProductAggregatedDaily ReportViewedProductAggregatedDaily
	d := g.Where("id = ?", id).Delete(&reportViewedProductAggregatedDaily)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
