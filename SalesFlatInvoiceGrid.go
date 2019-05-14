package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatInvoiceGrid :
type SalesFlatInvoiceGrid struct {
	EntityId           uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	StoreId            uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	BaseGrandTotal     float64    `gorm:"column:base_grand_total" form:"base_grand_total" json:"base_grand_total"`
	GrandTotal         float64    `gorm:"column:grand_total" form:"grand_total" json:"grand_total"`
	OrderId            uint16     `gorm:"column:order_id" form:"order_id" json:"order_id"`
	State              uint16     `gorm:"column:state" form:"state" json:"state"`
	StoreCurrencyCode  string     `gorm:"column:store_currency_code" form:"store_currency_code" json:"store_currency_code"`
	OrderCurrencyCode  string     `gorm:"column:order_currency_code" form:"order_currency_code" json:"order_currency_code"`
	BaseCurrencyCode   string     `gorm:"column:base_currency_code" form:"base_currency_code" json:"base_currency_code"`
	GlobalCurrencyCode string     `gorm:"column:global_currency_code" form:"global_currency_code" json:"global_currency_code"`
	IncrementId        string     `gorm:"column:increment_id" form:"increment_id" json:"increment_id"`
	OrderIncrementId   string     `gorm:"column:order_increment_id" form:"order_increment_id" json:"order_increment_id"`
	CreatedAt          *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	OrderCreatedAt     *time.Time `gorm:"column:order_created_at" form:"order_created_at" json:"order_created_at"`
	BillingName        string     `gorm:"column:billing_name" form:"billing_name" json:"billing_name"`
}

// TableName :
func (*SalesFlatInvoiceGrid) TableName() string {
	return "sales_flat_invoice_grid"
}

// handler create
func initRoutersSalesFlatInvoiceGrid(r *gin.Engine, salesflatinvoicegrid string) {
	route := r.Group("salesflatinvoicegrid")
	route.GET("/", getSalesFlatInvoiceGrids)
	route.GET("/:id", getSalesFlatInvoiceGrid)
	route.POST("/", createSalesFlatInvoiceGrid)
	route.PUT("/:id", updateSalesFlatInvoiceGrid)
	route.DELETE("/:id", deleteSalesFlatInvoiceGrid)
}

func getSalesFlatInvoiceGrids(c *gin.Context) {
	var salesFlatInvoiceGrids []SalesFlatInvoiceGrid
	if err := g.Find(&salesFlatInvoiceGrids).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatInvoiceGrids)
	}
}

func getSalesFlatInvoiceGrid(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatInvoiceGrid SalesFlatInvoiceGrid
	if err := g.Where("id = ?", id).First(&salesFlatInvoiceGrid).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatInvoiceGrid)
	}
}

func createSalesFlatInvoiceGrid(c *gin.Context) {
	var salesFlatInvoiceGrid SalesFlatInvoiceGrid
	c.BindJSON(&salesFlatInvoiceGrid)
	g.Create(&salesFlatInvoiceGrid)
	c.JSON(200, salesFlatInvoiceGrid)
}

func updateSalesFlatInvoiceGrid(c *gin.Context) {
	var salesFlatInvoiceGrid SalesFlatInvoiceGrid
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatInvoiceGrid).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatInvoiceGrid)
	g.Save(&salesFlatInvoiceGrid)
	c.JSON(200, salesFlatInvoiceGrid)
}
func deleteSalesFlatInvoiceGrid(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatInvoiceGrid SalesFlatInvoiceGrid
	d := g.Where("id = ?", id).Delete(&salesFlatInvoiceGrid)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
