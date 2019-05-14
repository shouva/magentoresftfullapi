package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ReportViewedProductAggregatedYearly :
type ReportViewedProductAggregatedYearly struct {
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
func (*ReportViewedProductAggregatedYearly) TableName() string {
	return "report_viewed_product_aggregated_yearly"
}

// handler create
func initRoutersReportViewedProductAggregatedYearly(r *gin.Engine, reportviewedproductaggregatedyearly string) {
	route := r.Group("reportviewedproductaggregatedyearly")
	route.GET("/", getReportViewedProductAggregatedYearlys)
	route.GET("/:id", getReportViewedProductAggregatedYearly)
	route.POST("/", createReportViewedProductAggregatedYearly)
	route.PUT("/:id", updateReportViewedProductAggregatedYearly)
	route.DELETE("/:id", deleteReportViewedProductAggregatedYearly)
}

func getReportViewedProductAggregatedYearlys(c *gin.Context) {
	var reportViewedProductAggregatedYearlys []ReportViewedProductAggregatedYearly
	if err := g.Find(&reportViewedProductAggregatedYearlys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportViewedProductAggregatedYearlys)
	}
}

func getReportViewedProductAggregatedYearly(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportViewedProductAggregatedYearly ReportViewedProductAggregatedYearly
	if err := g.Where("id = ?", id).First(&reportViewedProductAggregatedYearly).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportViewedProductAggregatedYearly)
	}
}

func createReportViewedProductAggregatedYearly(c *gin.Context) {
	var reportViewedProductAggregatedYearly ReportViewedProductAggregatedYearly
	c.BindJSON(&reportViewedProductAggregatedYearly)
	g.Create(&reportViewedProductAggregatedYearly)
	c.JSON(200, reportViewedProductAggregatedYearly)
}

func updateReportViewedProductAggregatedYearly(c *gin.Context) {
	var reportViewedProductAggregatedYearly ReportViewedProductAggregatedYearly
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&reportViewedProductAggregatedYearly).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&reportViewedProductAggregatedYearly)
	g.Save(&reportViewedProductAggregatedYearly)
	c.JSON(200, reportViewedProductAggregatedYearly)
}
func deleteReportViewedProductAggregatedYearly(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportViewedProductAggregatedYearly ReportViewedProductAggregatedYearly
	d := g.Where("id = ?", id).Delete(&reportViewedProductAggregatedYearly)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
