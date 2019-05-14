package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ReportComparedProductIndex :
type ReportComparedProductIndex struct {
	IndexId    uint64     `gorm:"column:index_id;primary_key" form:"index_id;primary_key" json:"index_id;primary_key"`
	VisitorId  uint16     `gorm:"column:visitor_id" form:"visitor_id" json:"visitor_id"`
	CustomerId uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	ProductId  uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	StoreId    uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	AddedAt    *time.Time `gorm:"column:added_at" form:"added_at" json:"added_at"`
}

// TableName :
func (*ReportComparedProductIndex) TableName() string {
	return "report_compared_product_index"
}

// handler create
func initRoutersReportComparedProductIndex(r *gin.Engine, reportcomparedproductindex string) {
	route := r.Group("reportcomparedproductindex")
	route.GET("/", getReportComparedProductIndexs)
	route.GET("/:id", getReportComparedProductIndex)
	route.POST("/", createReportComparedProductIndex)
	route.PUT("/:id", updateReportComparedProductIndex)
	route.DELETE("/:id", deleteReportComparedProductIndex)
}

func getReportComparedProductIndexs(c *gin.Context) {
	var reportComparedProductIndexs []ReportComparedProductIndex
	if err := g.Find(&reportComparedProductIndexs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportComparedProductIndexs)
	}
}

func getReportComparedProductIndex(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportComparedProductIndex ReportComparedProductIndex
	if err := g.Where("id = ?", id).First(&reportComparedProductIndex).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reportComparedProductIndex)
	}
}

func createReportComparedProductIndex(c *gin.Context) {
	var reportComparedProductIndex ReportComparedProductIndex
	c.BindJSON(&reportComparedProductIndex)
	g.Create(&reportComparedProductIndex)
	c.JSON(200, reportComparedProductIndex)
}

func updateReportComparedProductIndex(c *gin.Context) {
	var reportComparedProductIndex ReportComparedProductIndex
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&reportComparedProductIndex).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&reportComparedProductIndex)
	g.Save(&reportComparedProductIndex)
	c.JSON(200, reportComparedProductIndex)
}
func deleteReportComparedProductIndex(c *gin.Context) {
	id := c.Params.ByName("id")
	var reportComparedProductIndex ReportComparedProductIndex
	d := g.Where("id = ?", id).Delete(&reportComparedProductIndex)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
