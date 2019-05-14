package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatCreditmemoGrid :
type SalesFlatCreditmemoGrid struct {
	EntityId           uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	StoreId            uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	StoreToOrderRate   float64    `gorm:"column:store_to_order_rate" form:"store_to_order_rate" json:"store_to_order_rate"`
	BaseToOrderRate    float64    `gorm:"column:base_to_order_rate" form:"base_to_order_rate" json:"base_to_order_rate"`
	GrandTotal         float64    `gorm:"column:grand_total" form:"grand_total" json:"grand_total"`
	StoreToBaseRate    float64    `gorm:"column:store_to_base_rate" form:"store_to_base_rate" json:"store_to_base_rate"`
	BaseToGlobalRate   float64    `gorm:"column:base_to_global_rate" form:"base_to_global_rate" json:"base_to_global_rate"`
	BaseGrandTotal     float64    `gorm:"column:base_grand_total" form:"base_grand_total" json:"base_grand_total"`
	OrderId            uint16     `gorm:"column:order_id" form:"order_id" json:"order_id"`
	CreditmemoStatus   uint16     `gorm:"column:creditmemo_status" form:"creditmemo_status" json:"creditmemo_status"`
	State              uint16     `gorm:"column:state" form:"state" json:"state"`
	InvoiceId          uint16     `gorm:"column:invoice_id" form:"invoice_id" json:"invoice_id"`
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
func (*SalesFlatCreditmemoGrid) TableName() string {
	return "sales_flat_creditmemo_grid"
}

// handler create
func initRoutersSalesFlatCreditmemoGrid(r *gin.Engine, salesflatcreditmemogrid string) {
	route := r.Group("salesflatcreditmemogrid")
	route.GET("/", getSalesFlatCreditmemoGrids)
	route.GET("/:id", getSalesFlatCreditmemoGrid)
	route.POST("/", createSalesFlatCreditmemoGrid)
	route.PUT("/:id", updateSalesFlatCreditmemoGrid)
	route.DELETE("/:id", deleteSalesFlatCreditmemoGrid)
}

func getSalesFlatCreditmemoGrids(c *gin.Context) {
	var salesFlatCreditmemoGrids []SalesFlatCreditmemoGrid
	if err := g.Find(&salesFlatCreditmemoGrids).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatCreditmemoGrids)
	}
}

func getSalesFlatCreditmemoGrid(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatCreditmemoGrid SalesFlatCreditmemoGrid
	if err := g.Where("id = ?", id).First(&salesFlatCreditmemoGrid).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatCreditmemoGrid)
	}
}

func createSalesFlatCreditmemoGrid(c *gin.Context) {
	var salesFlatCreditmemoGrid SalesFlatCreditmemoGrid
	c.BindJSON(&salesFlatCreditmemoGrid)
	g.Create(&salesFlatCreditmemoGrid)
	c.JSON(200, salesFlatCreditmemoGrid)
}

func updateSalesFlatCreditmemoGrid(c *gin.Context) {
	var salesFlatCreditmemoGrid SalesFlatCreditmemoGrid
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatCreditmemoGrid).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatCreditmemoGrid)
	g.Save(&salesFlatCreditmemoGrid)
	c.JSON(200, salesFlatCreditmemoGrid)
}
func deleteSalesFlatCreditmemoGrid(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatCreditmemoGrid SalesFlatCreditmemoGrid
	d := g.Where("id = ?", id).Delete(&salesFlatCreditmemoGrid)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
