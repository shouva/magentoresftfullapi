package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatOrderGrid :
type SalesFlatOrderGrid struct {
	EntityId          uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	Status            string     `gorm:"column:status" form:"status" json:"status"`
	StoreId           uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	StoreName         string     `gorm:"column:store_name" form:"store_name" json:"store_name"`
	CustomerId        uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	BaseGrandTotal    float64    `gorm:"column:base_grand_total" form:"base_grand_total" json:"base_grand_total"`
	BaseTotalPaid     float64    `gorm:"column:base_total_paid" form:"base_total_paid" json:"base_total_paid"`
	GrandTotal        float64    `gorm:"column:grand_total" form:"grand_total" json:"grand_total"`
	TotalPaid         float64    `gorm:"column:total_paid" form:"total_paid" json:"total_paid"`
	IncrementId       string     `gorm:"column:increment_id" form:"increment_id" json:"increment_id"`
	BaseCurrencyCode  string     `gorm:"column:base_currency_code" form:"base_currency_code" json:"base_currency_code"`
	OrderCurrencyCode string     `gorm:"column:order_currency_code" form:"order_currency_code" json:"order_currency_code"`
	ShippingName      string     `gorm:"column:shipping_name" form:"shipping_name" json:"shipping_name"`
	BillingName       string     `gorm:"column:billing_name" form:"billing_name" json:"billing_name"`
	CreatedAt         *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
}

// TableName :
func (*SalesFlatOrderGrid) TableName() string {
	return "sales_flat_order_grid"
}

// handler create
func initRoutersSalesFlatOrderGrid(r *gin.Engine, salesflatordergrid string) {
	route := r.Group("salesflatordergrid")
	route.GET("/", getSalesFlatOrderGrids)
	route.GET("/:id", getSalesFlatOrderGrid)
	route.POST("/", createSalesFlatOrderGrid)
	route.PUT("/:id", updateSalesFlatOrderGrid)
	route.DELETE("/:id", deleteSalesFlatOrderGrid)
}

func getSalesFlatOrderGrids(c *gin.Context) {
	var salesFlatOrderGrids []SalesFlatOrderGrid
	if err := g.Find(&salesFlatOrderGrids).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatOrderGrids)
	}
}

func getSalesFlatOrderGrid(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatOrderGrid SalesFlatOrderGrid
	if err := g.Where("id = ?", id).First(&salesFlatOrderGrid).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatOrderGrid)
	}
}

func createSalesFlatOrderGrid(c *gin.Context) {
	var salesFlatOrderGrid SalesFlatOrderGrid
	c.BindJSON(&salesFlatOrderGrid)
	g.Create(&salesFlatOrderGrid)
	c.JSON(200, salesFlatOrderGrid)
}

func updateSalesFlatOrderGrid(c *gin.Context) {
	var salesFlatOrderGrid SalesFlatOrderGrid
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatOrderGrid).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatOrderGrid)
	g.Save(&salesFlatOrderGrid)
	c.JSON(200, salesFlatOrderGrid)
}
func deleteSalesFlatOrderGrid(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatOrderGrid SalesFlatOrderGrid
	d := g.Where("id = ?", id).Delete(&salesFlatOrderGrid)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
