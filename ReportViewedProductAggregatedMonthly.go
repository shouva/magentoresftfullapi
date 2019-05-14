package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ReportViewedProductAggregatedMonthly :
type ReportViewedProductAggregatedMonthly struct {
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
func (*ReportViewedProductAggregatedMonthly) TableName() string {
	return "report_viewed_product_aggregated_monthly"
}

// handler create
func initRoutersReportViewedProductAggregatedMonthly(r *gin.Engine, reportviewedproductaggregatedmonthly string) {
	route := r.Group("reportviewedproductaggregatedmonthly")
	route.GET("/", getReportViewedProductAggregatedMonthlys)
	route.GET("/:id", getReportViewedProductAggregatedMonthly)
	route.POST("/", createReportViewedProductAggregatedMonthly)
	route.PUT("/:id", updateReportViewedProductAggregatedMonthly)
	route.DELETE("/:id", deleteReportViewedProductAggregatedMonthly)
}

func getReportViewedProductAggregatedMonthlys(c *gin.Context) {
	var reportViewedProductAggregatedMonthlys []ReportViewedProductAggregatedMonthly
	if err := g.Find(&reportViewedProductAggregatedMonthlys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportViewedProductAggregatedMonthlys)
	}
}

func getReportViewedProductAggregatedMonthly(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportViewedProductAggregatedMonthly ReportViewedProductAggregatedMonthly
	if err := g.Where("id = ?", id).First(&reportViewedProductAggregatedMonthly).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportViewedProductAggregatedMonthly)
	}
}

func createReportViewedProductAggregatedMonthly(c *gin.Context) {
	var reportViewedProductAggregatedMonthly ReportViewedProductAggregatedMonthly
	c.BindJSON(&reportViewedProductAggregatedMonthly)
	g.Create(&reportViewedProductAggregatedMonthly)
	c.JSON(200, reportViewedProductAggregatedMonthly)
}

func updateReportViewedProductAggregatedMonthly(c *gin.Context) {
	var reportViewedProductAggregatedMonthly ReportViewedProductAggregatedMonthly
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&reportViewedProductAggregatedMonthly).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&reportViewedProductAggregatedMonthly)
	g.Save(&reportViewedProductAggregatedMonthly)
	c.JSON(200, reportViewedProductAggregatedMonthly)
}
func deleteReportViewedProductAggregatedMonthly(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportViewedProductAggregatedMonthly ReportViewedProductAggregatedMonthly
	d := g.Where("id = ?", id).Delete(&reportViewedProductAggregatedMonthly)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
