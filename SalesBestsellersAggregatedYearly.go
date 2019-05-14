package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesBestsellersAggregatedYearly :
type SalesBestsellersAggregatedYearly struct {
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
func (*SalesBestsellersAggregatedYearly) TableName() string {
	return "sales_bestsellers_aggregated_yearly"
}

// handler create
func initRoutersSalesBestsellersAggregatedYearly(r *gin.Engine, salesbestsellersaggregatedyearly string) {
	route := r.Group("salesbestsellersaggregatedyearly")
	route.GET("/", getSalesBestsellersAggregatedYearlys)
	route.GET("/:id", getSalesBestsellersAggregatedYearly)
	route.POST("/", createSalesBestsellersAggregatedYearly)
	route.PUT("/:id", updateSalesBestsellersAggregatedYearly)
	route.DELETE("/:id", deleteSalesBestsellersAggregatedYearly)
}

func getSalesBestsellersAggregatedYearlys(c *gin.Context) {
	var salesBestsellersAggregatedYearlys []SalesBestsellersAggregatedYearly
	if err := g.Find(&salesBestsellersAggregatedYearlys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesBestsellersAggregatedYearlys)
	}
}

func getSalesBestsellersAggregatedYearly(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesBestsellersAggregatedYearly SalesBestsellersAggregatedYearly
	if err := g.Where("id = ?", id).First(&salesBestsellersAggregatedYearly).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesBestsellersAggregatedYearly)
	}
}

func createSalesBestsellersAggregatedYearly(c *gin.Context) {
	var salesBestsellersAggregatedYearly SalesBestsellersAggregatedYearly
	c.BindJSON(&salesBestsellersAggregatedYearly)
	g.Create(&salesBestsellersAggregatedYearly)
	c.JSON(200, salesBestsellersAggregatedYearly)
}

func updateSalesBestsellersAggregatedYearly(c *gin.Context) {
	var salesBestsellersAggregatedYearly SalesBestsellersAggregatedYearly
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesBestsellersAggregatedYearly).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesBestsellersAggregatedYearly)
	g.Save(&salesBestsellersAggregatedYearly)
	c.JSON(200, salesBestsellersAggregatedYearly)
}
func deleteSalesBestsellersAggregatedYearly(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesBestsellersAggregatedYearly SalesBestsellersAggregatedYearly
	d := g.Where("id = ?", id).Delete(&salesBestsellersAggregatedYearly)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
