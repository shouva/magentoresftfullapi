package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatInvoice :
type SalesFlatInvoice struct {
	EntityId                  uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	StoreId                   uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	BaseGrandTotal            float64    `gorm:"column:base_grand_total" form:"base_grand_total" json:"base_grand_total"`
	ShippingTaxAmount         float64    `gorm:"column:shipping_tax_amount" form:"shipping_tax_amount" json:"shipping_tax_amount"`
	TaxAmount                 float64    `gorm:"column:tax_amount" form:"tax_amount" json:"tax_amount"`
	BaseTaxAmount             float64    `gorm:"column:base_tax_amount" form:"base_tax_amount" json:"base_tax_amount"`
	StoreToOrderRate          float64    `gorm:"column:store_to_order_rate" form:"store_to_order_rate" json:"store_to_order_rate"`
	BaseShippingTaxAmount     float64    `gorm:"column:base_shipping_tax_amount" form:"base_shipping_tax_amount" json:"base_shipping_tax_amount"`
	BaseDiscountAmount        float64    `gorm:"column:base_discount_amount" form:"base_discount_amount" json:"base_discount_amount"`
	BaseToOrderRate           float64    `gorm:"column:base_to_order_rate" form:"base_to_order_rate" json:"base_to_order_rate"`
	GrandTotal                float64    `gorm:"column:grand_total" form:"grand_total" json:"grand_total"`
	ShippingAmount            float64    `gorm:"column:shipping_amount" form:"shipping_amount" json:"shipping_amount"`
	SubtotalInclTax           float64    `gorm:"column:subtotal_incl_tax" form:"subtotal_incl_tax" json:"subtotal_incl_tax"`
	BaseSubtotalInclTax       float64    `gorm:"column:base_subtotal_incl_tax" form:"base_subtotal_incl_tax" json:"base_subtotal_incl_tax"`
	StoreToBaseRate           float64    `gorm:"column:store_to_base_rate" form:"store_to_base_rate" json:"store_to_base_rate"`
	BaseShippingAmount        float64    `gorm:"column:base_shipping_amount" form:"base_shipping_amount" json:"base_shipping_amount"`
	TotalQty                  float64    `gorm:"column:total_qty" form:"total_qty" json:"total_qty"`
	BaseToGlobalRate          float64    `gorm:"column:base_to_global_rate" form:"base_to_global_rate" json:"base_to_global_rate"`
	Subtotal                  float64    `gorm:"column:subtotal" form:"subtotal" json:"subtotal"`
	BaseSubtotal              float64    `gorm:"column:base_subtotal" form:"base_subtotal" json:"base_subtotal"`
	DiscountAmount            float64    `gorm:"column:discount_amount" form:"discount_amount" json:"discount_amount"`
	BillingAddressId          uint16     `gorm:"column:billing_address_id" form:"billing_address_id" json:"billing_address_id"`
	IsUsedForRefund           uint16     `gorm:"column:is_used_for_refund" form:"is_used_for_refund" json:"is_used_for_refund"`
	OrderId                   uint16     `gorm:"column:order_id" form:"order_id" json:"order_id"`
	EmailSent                 uint16     `gorm:"column:email_sent" form:"email_sent" json:"email_sent"`
	CanVoidFlag               uint16     `gorm:"column:can_void_flag" form:"can_void_flag" json:"can_void_flag"`
	State                     uint16     `gorm:"column:state" form:"state" json:"state"`
	ShippingAddressId         uint16     `gorm:"column:shipping_address_id" form:"shipping_address_id" json:"shipping_address_id"`
	StoreCurrencyCode         string     `gorm:"column:store_currency_code" form:"store_currency_code" json:"store_currency_code"`
	TransactionId             string     `gorm:"column:transaction_id" form:"transaction_id" json:"transaction_id"`
	OrderCurrencyCode         string     `gorm:"column:order_currency_code" form:"order_currency_code" json:"order_currency_code"`
	BaseCurrencyCode          string     `gorm:"column:base_currency_code" form:"base_currency_code" json:"base_currency_code"`
	GlobalCurrencyCode        string     `gorm:"column:global_currency_code" form:"global_currency_code" json:"global_currency_code"`
	IncrementId               string     `gorm:"column:increment_id" form:"increment_id" json:"increment_id"`
	CreatedAt                 *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt                 *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	HiddenTaxAmount           float64    `gorm:"column:hidden_tax_amount" form:"hidden_tax_amount" json:"hidden_tax_amount"`
	BaseHiddenTaxAmount       float64    `gorm:"column:base_hidden_tax_amount" form:"base_hidden_tax_amount" json:"base_hidden_tax_amount"`
	ShippingHiddenTaxAmount   float64    `gorm:"column:shipping_hidden_tax_amount" form:"shipping_hidden_tax_amount" json:"shipping_hidden_tax_amount"`
	BaseShippingHiddenTaxAmnt float64    `gorm:"column:base_shipping_hidden_tax_amnt" form:"base_shipping_hidden_tax_amnt" json:"base_shipping_hidden_tax_amnt"`
	ShippingInclTax           float64    `gorm:"column:shipping_incl_tax" form:"shipping_incl_tax" json:"shipping_incl_tax"`
	BaseShippingInclTax       float64    `gorm:"column:base_shipping_incl_tax" form:"base_shipping_incl_tax" json:"base_shipping_incl_tax"`
	BaseTotalRefunded         float64    `gorm:"column:base_total_refunded" form:"base_total_refunded" json:"base_total_refunded"`
	DiscountDescription       string     `gorm:"column:discount_description" form:"discount_description" json:"discount_description"`
}

// TableName :
func (*SalesFlatInvoice) TableName() string {
	return "sales_flat_invoice"
}

// handler create
func initRoutersSalesFlatInvoice(r *gin.Engine, salesflatinvoice string) {
	route := r.Group("salesflatinvoice")
	route.GET("/", getSalesFlatInvoices)
	route.GET("/:id", getSalesFlatInvoice)
	route.POST("/", createSalesFlatInvoice)
	route.PUT("/:id", updateSalesFlatInvoice)
	route.DELETE("/:id", deleteSalesFlatInvoice)
}

func getSalesFlatInvoices(c *gin.Context) {
	var salesFlatInvoices []SalesFlatInvoice
	if err := g.Find(&salesFlatInvoices).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatInvoices)
	}
}

func getSalesFlatInvoice(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatInvoice SalesFlatInvoice
	if err := g.Where("id = ?", id).First(&salesFlatInvoice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatInvoice)
	}
}

func createSalesFlatInvoice(c *gin.Context) {
	var salesFlatInvoice SalesFlatInvoice
	c.BindJSON(&salesFlatInvoice)
	g.Create(&salesFlatInvoice)
	c.JSON(200, salesFlatInvoice)
}

func updateSalesFlatInvoice(c *gin.Context) {
	var salesFlatInvoice SalesFlatInvoice
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatInvoice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatInvoice)
	g.Save(&salesFlatInvoice)
	c.JSON(200, salesFlatInvoice)
}
func deleteSalesFlatInvoice(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatInvoice SalesFlatInvoice
	d := g.Where("id = ?", id).Delete(&salesFlatInvoice)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
