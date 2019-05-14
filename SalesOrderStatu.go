package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesOrderStatu :
type SalesOrderStatu struct {
	Status string `gorm:"column:status;primary_key" form:"status;primary_key" json:"status;primary_key"`
	Label  string `gorm:"column:label" form:"label" json:"label"`
}

// TableName :
func (*SalesOrderStatu) TableName() string {
	return "sales_order_status"
}

// handler create
func initRoutersSalesOrderStatu(r *gin.Engine, salesorderstatu string) {
	route := r.Group("salesorderstatu")
	route.GET("/", getSalesOrderStatus)
	route.GET("/:id", getSalesOrderStatu)
	route.POST("/", createSalesOrderStatu)
	route.PUT("/:id", updateSalesOrderStatu)
	route.DELETE("/:id", deleteSalesOrderStatu)
}

func getSalesOrderStatus(c *gin.Context) {
	var salesOrderStatus []SalesOrderStatu
	if err := g.Find(&salesOrderStatus).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderStatus)
	}
}

func getSalesOrderStatu(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderStatu SalesOrderStatu
	if err := g.Where("id = ?", id).First(&salesOrderStatu).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderStatu)
	}
}

func createSalesOrderStatu(c *gin.Context) {
	var salesOrderStatu SalesOrderStatu
	c.BindJSON(&salesOrderStatu)
	g.Create(&salesOrderStatu)
	c.JSON(200, salesOrderStatu)
}

func updateSalesOrderStatu(c *gin.Context) {
	var salesOrderStatu SalesOrderStatu
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesOrderStatu).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesOrderStatu)
	g.Save(&salesOrderStatu)
	c.JSON(200, salesOrderStatu)
}
func deleteSalesOrderStatu(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderStatu SalesOrderStatu
	d := g.Where("id = ?", id).Delete(&salesOrderStatu)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
