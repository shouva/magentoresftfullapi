package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesOrderStatusLabel :
type SalesOrderStatusLabel struct {
	Status  string `gorm:"column:status;primary_key" form:"status;primary_key" json:"status;primary_key"`
	StoreId uint16 `gorm:"column:store_id;primary_key" form:"store_id;primary_key" json:"store_id;primary_key"`
	Label   string `gorm:"column:label" form:"label" json:"label"`
}

// TableName :
func (*SalesOrderStatusLabel) TableName() string {
	return "sales_order_status_label"
}

// handler create
func initRoutersSalesOrderStatusLabel(r *gin.Engine, salesorderstatuslabel string) {
	route := r.Group("salesorderstatuslabel")
	route.GET("/", getSalesOrderStatusLabels)
	route.GET("/:id", getSalesOrderStatusLabel)
	route.POST("/", createSalesOrderStatusLabel)
	route.PUT("/:id", updateSalesOrderStatusLabel)
	route.DELETE("/:id", deleteSalesOrderStatusLabel)
}

func getSalesOrderStatusLabels(c *gin.Context) {
	var salesOrderStatusLabels []SalesOrderStatusLabel
	if err := g.Find(&salesOrderStatusLabels).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderStatusLabels)
	}
}

func getSalesOrderStatusLabel(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderStatusLabel SalesOrderStatusLabel
	if err := g.Where("id = ?", id).First(&salesOrderStatusLabel).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderStatusLabel)
	}
}

func createSalesOrderStatusLabel(c *gin.Context) {
	var salesOrderStatusLabel SalesOrderStatusLabel
	c.BindJSON(&salesOrderStatusLabel)
	g.Create(&salesOrderStatusLabel)
	c.JSON(200, salesOrderStatusLabel)
}

func updateSalesOrderStatusLabel(c *gin.Context) {
	var salesOrderStatusLabel SalesOrderStatusLabel
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesOrderStatusLabel).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesOrderStatusLabel)
	g.Save(&salesOrderStatusLabel)
	c.JSON(200, salesOrderStatusLabel)
}
func deleteSalesOrderStatusLabel(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderStatusLabel SalesOrderStatusLabel
	d := g.Where("id = ?", id).Delete(&salesOrderStatusLabel)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
