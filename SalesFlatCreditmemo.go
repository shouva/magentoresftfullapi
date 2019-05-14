package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatCreditmemo :
type SalesFlatCreditmemo struct {
	EntityId                  uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	StoreId                   uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	AdjustmentPositive        float64    `gorm:"column:adjustment_positive" form:"adjustment_positive" json:"adjustment_positive"`
	BaseShippingTaxAmount     float64    `gorm:"column:base_shipping_tax_amount" form:"base_shipping_tax_amount" json:"base_shipping_tax_amount"`
	StoreToOrderRate          float64    `gorm:"column:store_to_order_rate" form:"store_to_order_rate" json:"store_to_order_rate"`
	BaseDiscountAmount        float64    `gorm:"column:base_discount_amount" form:"base_discount_amount" json:"base_discount_amount"`
	BaseToOrderRate           float64    `gorm:"column:base_to_order_rate" form:"base_to_order_rate" json:"base_to_order_rate"`
	GrandTotal                float64    `gorm:"column:grand_total" form:"grand_total" json:"grand_total"`
	BaseAdjustmentNegative    float64    `gorm:"column:base_adjustment_negative" form:"base_adjustment_negative" json:"base_adjustment_negative"`
	BaseSubtotalInclTax       float64    `gorm:"column:base_subtotal_incl_tax" form:"base_subtotal_incl_tax" json:"base_subtotal_incl_tax"`
	ShippingAmount            float64    `gorm:"column:shipping_amount" form:"shipping_amount" json:"shipping_amount"`
	SubtotalInclTax           float64    `gorm:"column:subtotal_incl_tax" form:"subtotal_incl_tax" json:"subtotal_incl_tax"`
	AdjustmentNegative        float64    `gorm:"column:adjustment_negative" form:"adjustment_negative" json:"adjustment_negative"`
	BaseShippingAmount        float64    `gorm:"column:base_shipping_amount" form:"base_shipping_amount" json:"base_shipping_amount"`
	StoreToBaseRate           float64    `gorm:"column:store_to_base_rate" form:"store_to_base_rate" json:"store_to_base_rate"`
	BaseToGlobalRate          float64    `gorm:"column:base_to_global_rate" form:"base_to_global_rate" json:"base_to_global_rate"`
	BaseAdjustment            float64    `gorm:"column:base_adjustment" form:"base_adjustment" json:"base_adjustment"`
	BaseSubtotal              float64    `gorm:"column:base_subtotal" form:"base_subtotal" json:"base_subtotal"`
	DiscountAmount            float64    `gorm:"column:discount_amount" form:"discount_amount" json:"discount_amount"`
	Subtotal                  float64    `gorm:"column:subtotal" form:"subtotal" json:"subtotal"`
	Adjustment                float64    `gorm:"column:adjustment" form:"adjustment" json:"adjustment"`
	BaseGrandTotal            float64    `gorm:"column:base_grand_total" form:"base_grand_total" json:"base_grand_total"`
	BaseAdjustmentPositive    float64    `gorm:"column:base_adjustment_positive" form:"base_adjustment_positive" json:"base_adjustment_positive"`
	BaseTaxAmount             float64    `gorm:"column:base_tax_amount" form:"base_tax_amount" json:"base_tax_amount"`
	ShippingTaxAmount         float64    `gorm:"column:shipping_tax_amount" form:"shipping_tax_amount" json:"shipping_tax_amount"`
	TaxAmount                 float64    `gorm:"column:tax_amount" form:"tax_amount" json:"tax_amount"`
	OrderId                   uint16     `gorm:"column:order_id" form:"order_id" json:"order_id"`
	EmailSent                 uint16     `gorm:"column:email_sent" form:"email_sent" json:"email_sent"`
	CreditmemoStatus          uint16     `gorm:"column:creditmemo_status" form:"creditmemo_status" json:"creditmemo_status"`
	State                     uint16     `gorm:"column:state" form:"state" json:"state"`
	ShippingAddressId         uint16     `gorm:"column:shipping_address_id" form:"shipping_address_id" json:"shipping_address_id"`
	BillingAddressId          uint16     `gorm:"column:billing_address_id" form:"billing_address_id" json:"billing_address_id"`
	InvoiceId                 uint16     `gorm:"column:invoice_id" form:"invoice_id" json:"invoice_id"`
	StoreCurrencyCode         string     `gorm:"column:store_currency_code" form:"store_currency_code" json:"store_currency_code"`
	OrderCurrencyCode         string     `gorm:"column:order_currency_code" form:"order_currency_code" json:"order_currency_code"`
	BaseCurrencyCode          string     `gorm:"column:base_currency_code" form:"base_currency_code" json:"base_currency_code"`
	GlobalCurrencyCode        string     `gorm:"column:global_currency_code" form:"global_currency_code" json:"global_currency_code"`
	TransactionId             string     `gorm:"column:transaction_id" form:"transaction_id" json:"transaction_id"`
	IncrementId               string     `gorm:"column:increment_id" form:"increment_id" json:"increment_id"`
	CreatedAt                 *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt                 *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	HiddenTaxAmount           float64    `gorm:"column:hidden_tax_amount" form:"hidden_tax_amount" json:"hidden_tax_amount"`
	BaseHiddenTaxAmount       float64    `gorm:"column:base_hidden_tax_amount" form:"base_hidden_tax_amount" json:"base_hidden_tax_amount"`
	ShippingHiddenTaxAmount   float64    `gorm:"column:shipping_hidden_tax_amount" form:"shipping_hidden_tax_amount" json:"shipping_hidden_tax_amount"`
	BaseShippingHiddenTaxAmnt float64    `gorm:"column:base_shipping_hidden_tax_amnt" form:"base_shipping_hidden_tax_amnt" json:"base_shipping_hidden_tax_amnt"`
	ShippingInclTax           float64    `gorm:"column:shipping_incl_tax" form:"shipping_incl_tax" json:"shipping_incl_tax"`
	BaseShippingInclTax       float64    `gorm:"column:base_shipping_incl_tax" form:"base_shipping_incl_tax" json:"base_shipping_incl_tax"`
	DiscountDescription       string     `gorm:"column:discount_description" form:"discount_description" json:"discount_description"`
}

// TableName :
func (*SalesFlatCreditmemo) TableName() string {
	return "sales_flat_creditmemo"
}

// handler create
func initRoutersSalesFlatCreditmemo(r *gin.Engine, salesflatcreditmemo string) {
	route := r.Group("salesflatcreditmemo")
	route.GET("/", getSalesFlatCreditmemos)
	route.GET("/:id", getSalesFlatCreditmemo)
	route.POST("/", createSalesFlatCreditmemo)
	route.PUT("/:id", updateSalesFlatCreditmemo)
	route.DELETE("/:id", deleteSalesFlatCreditmemo)
}

func getSalesFlatCreditmemos(c *gin.Context) {
	var salesFlatCreditmemos []SalesFlatCreditmemo
	if err := g.Find(&salesFlatCreditmemos).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatCreditmemos)
	}
}

func getSalesFlatCreditmemo(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatCreditmemo SalesFlatCreditmemo
	if err := g.Where("id = ?", id).First(&salesFlatCreditmemo).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatCreditmemo)
	}
}

func createSalesFlatCreditmemo(c *gin.Context) {
	var salesFlatCreditmemo SalesFlatCreditmemo
	c.BindJSON(&salesFlatCreditmemo)
	g.Create(&salesFlatCreditmemo)
	c.JSON(200, salesFlatCreditmemo)
}

func updateSalesFlatCreditmemo(c *gin.Context) {
	var salesFlatCreditmemo SalesFlatCreditmemo
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatCreditmemo).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatCreditmemo)
	g.Save(&salesFlatCreditmemo)
	c.JSON(200, salesFlatCreditmemo)
}
func deleteSalesFlatCreditmemo(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatCreditmemo SalesFlatCreditmemo
	d := g.Where("id = ?", id).Delete(&salesFlatCreditmemo)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
