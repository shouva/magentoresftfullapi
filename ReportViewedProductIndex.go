package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ReportViewedProductIndex :
type ReportViewedProductIndex struct {
	IndexId    uint64     `gorm:"column:index_id;primary_key" form:"index_id;primary_key" json:"index_id;primary_key"`
	VisitorId  uint16     `gorm:"column:visitor_id" form:"visitor_id" json:"visitor_id"`
	CustomerId uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	ProductId  uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	StoreId    uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	AddedAt    *time.Time `gorm:"column:added_at" form:"added_at" json:"added_at"`
}

// TableName :
func (*ReportViewedProductIndex) TableName() string {
	return "report_viewed_product_index"
}

// handler create
func initRoutersReportViewedProductIndex(r *gin.Engine, reportviewedproductindex string) {
	route := r.Group("reportviewedproductindex")
	route.GET("/", getReportViewedProductIndexs)
	route.GET("/:id", getReportViewedProductIndex)
	route.POST("/", createReportViewedProductIndex)
	route.PUT("/:id", updateReportViewedProductIndex)
	route.DELETE("/:id", deleteReportViewedProductIndex)
}

func getReportViewedProductIndexs(c *gin.Context) {
	var reportViewedProductIndexs []ReportViewedProductIndex
	if err := g.Find(&reportViewedProductIndexs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportViewedProductIndexs)
	}
}

func getReportViewedProductIndex(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportViewedProductIndex ReportViewedProductIndex
	if err := g.Where("id = ?", id).First(&reportViewedProductIndex).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportViewedProductIndex)
	}
}

func createReportViewedProductIndex(c *gin.Context) {
	var reportViewedProductIndex ReportViewedProductIndex
	c.BindJSON(&reportViewedProductIndex)
	g.Create(&reportViewedProductIndex)
	c.JSON(200, reportViewedProductIndex)
}

func updateReportViewedProductIndex(c *gin.Context) {
	var reportViewedProductIndex ReportViewedProductIndex
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&reportViewedProductIndex).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&reportViewedProductIndex)
	g.Save(&reportViewedProductIndex)
	c.JSON(200, reportViewedProductIndex)
}
func deleteReportViewedProductIndex(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportViewedProductIndex ReportViewedProductIndex
	d := g.Where("id = ?", id).Delete(&reportViewedProductIndex)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
