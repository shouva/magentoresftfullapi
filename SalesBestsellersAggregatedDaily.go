package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesBestsellersAggregatedDaily :
type SalesBestsellersAggregatedDaily struct {
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
func (*SalesBestsellersAggregatedDaily) TableName() string {
	return "sales_bestsellers_aggregated_daily"
}

// handler create
func initRoutersSalesBestsellersAggregatedDaily(r *gin.Engine, salesbestsellersaggregateddaily string) {
	route := r.Group("salesbestsellersaggregateddaily")
	route.GET("/", getSalesBestsellersAggregatedDailys)
	route.GET("/:id", getSalesBestsellersAggregatedDaily)
	route.POST("/", createSalesBestsellersAggregatedDaily)
	route.PUT("/:id", updateSalesBestsellersAggregatedDaily)
	route.DELETE("/:id", deleteSalesBestsellersAggregatedDaily)
}

func getSalesBestsellersAggregatedDailys(c *gin.Context) {
	var salesBestsellersAggregatedDailys []SalesBestsellersAggregatedDaily
	if err := g.Find(&salesBestsellersAggregatedDailys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesBestsellersAggregatedDailys)
	}
}

func getSalesBestsellersAggregatedDaily(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesBestsellersAggregatedDaily SalesBestsellersAggregatedDaily
	if err := g.Where("id = ?", id).First(&salesBestsellersAggregatedDaily).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesBestsellersAggregatedDaily)
	}
}

func createSalesBestsellersAggregatedDaily(c *gin.Context) {
	var salesBestsellersAggregatedDaily SalesBestsellersAggregatedDaily
	c.BindJSON(&salesBestsellersAggregatedDaily)
	g.Create(&salesBestsellersAggregatedDaily)
	c.JSON(200, salesBestsellersAggregatedDaily)
}

func updateSalesBestsellersAggregatedDaily(c *gin.Context) {
	var salesBestsellersAggregatedDaily SalesBestsellersAggregatedDaily
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesBestsellersAggregatedDaily).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesBestsellersAggregatedDaily)
	g.Save(&salesBestsellersAggregatedDaily)
	c.JSON(200, salesBestsellersAggregatedDaily)
}
func deleteSalesBestsellersAggregatedDaily(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesBestsellersAggregatedDaily SalesBestsellersAggregatedDaily
	d := g.Where("id = ?", id).Delete(&salesBestsellersAggregatedDaily)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
