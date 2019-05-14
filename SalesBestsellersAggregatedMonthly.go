package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesBestsellersAggregatedMonthly :
type SalesBestsellersAggregatedMonthly struct {
	ID           uint16     `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	Period       *time.Time `gorm:"column:period" form:"period" json:"period"`
	StoreId      uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	ProductId    uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	ProductName  string     `gorm:"column:product_name" form:"product_name" json:"product_name"`
	ProductPrice float64    `gorm:"column:product_price" form:"product_price" json:"product_price"`
	QtyOrdered   float64    `gorm:"column:qty_ordered" form:"qty_ordered" json:"qty_ordered"`
	RatingPos    uint16     `gorm:"column:rating_pos" form:"rating_pos" json:"rating_pos"`
}

// TableName :
func (*SalesBestsellersAggregatedMonthly) TableName() string {
	return "sales_bestsellers_aggregated_monthly"
}

// handler create
func initRoutersSalesBestsellersAggregatedMonthly(r *gin.Engine, salesbestsellersaggregatedmonthly string) {
	route := r.Group("salesbestsellersaggregatedmonthly")
	route.GET("/", getSalesBestsellersAggregatedMonthlys)
	route.GET("/:id", getSalesBestsellersAggregatedMonthly)
	route.POST("/", createSalesBestsellersAggregatedMonthly)
	route.PUT("/:id", updateSalesBestsellersAggregatedMonthly)
	route.DELETE("/:id", deleteSalesBestsellersAggregatedMonthly)
}

func getSalesBestsellersAggregatedMonthlys(c *gin.Context) {
	var salesBestsellersAggregatedMonthlys []SalesBestsellersAggregatedMonthly
	if err := g.Find(&salesBestsellersAggregatedMonthlys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesBestsellersAggregatedMonthlys)
	}
}

func getSalesBestsellersAggregatedMonthly(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesBestsellersAggregatedMonthly SalesBestsellersAggregatedMonthly
	if err := g.Where("id = ?", id).First(&salesBestsellersAggregatedMonthly).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesBestsellersAggregatedMonthly)
	}
}

func createSalesBestsellersAggregatedMonthly(c *gin.Context) {
	var salesBestsellersAggregatedMonthly SalesBestsellersAggregatedMonthly
	c.BindJSON(&salesBestsellersAggregatedMonthly)
	g.Create(&salesBestsellersAggregatedMonthly)
	c.JSON(200, salesBestsellersAggregatedMonthly)
}

func updateSalesBestsellersAggregatedMonthly(c *gin.Context) {
	var salesBestsellersAggregatedMonthly SalesBestsellersAggregatedMonthly
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesBestsellersAggregatedMonthly).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesBestsellersAggregatedMonthly)
	g.Save(&salesBestsellersAggregatedMonthly)
	c.JSON(200, salesBestsellersAggregatedMonthly)
}
func deleteSalesBestsellersAggregatedMonthly(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesBestsellersAggregatedMonthly SalesBestsellersAggregatedMonthly
	d := g.Where("id = ?", id).Delete(&salesBestsellersAggregatedMonthly)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
