package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatOrderItem :
type SalesFlatOrderItem struct {
	ItemId                    uint16     `gorm:"column:item_id;primary_key" form:"item_id;primary_key" json:"item_id;primary_key"`
	OrderId                   uint16     `gorm:"column:order_id" form:"order_id" json:"order_id"`
	ParentItemId              uint16     `gorm:"column:parent_item_id" form:"parent_item_id" json:"parent_item_id"`
	QuoteItemId               uint16     `gorm:"column:quote_item_id" form:"quote_item_id" json:"quote_item_id"`
	StoreId                   uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	CreatedAt                 *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt                 *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	ProductId                 uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	ProductType               string     `gorm:"column:product_type" form:"product_type" json:"product_type"`
	ProductOptions            string     `gorm:"column:product_options" form:"product_options" json:"product_options"`
	Weight                    float64    `gorm:"column:weight" form:"weight" json:"weight"`
	IsVirtual                 uint16     `gorm:"column:is_virtual" form:"is_virtual" json:"is_virtual"`
	Sku                       string     `gorm:"column:sku" form:"sku" json:"sku"`
	Name                      string     `gorm:"column:name" form:"name" json:"name"`
	Description               string     `gorm:"column:description" form:"description" json:"description"`
	AppliedRuleIds            string     `gorm:"column:applied_rule_ids" form:"applied_rule_ids" json:"applied_rule_ids"`
	AdditionalData            string     `gorm:"column:additional_data" form:"additional_data" json:"additional_data"`
	FreeShipping              uint16     `gorm:"column:free_shipping" form:"free_shipping" json:"free_shipping"`
	IsQtyDecimal              uint16     `gorm:"column:is_qty_decimal" form:"is_qty_decimal" json:"is_qty_decimal"`
	NoDiscount                uint16     `gorm:"column:no_discount" form:"no_discount" json:"no_discount"`
	QtyBackordered            float64    `gorm:"column:qty_backordered" form:"qty_backordered" json:"qty_backordered"`
	QtyCanceled               float64    `gorm:"column:qty_canceled" form:"qty_canceled" json:"qty_canceled"`
	QtyInvoiced               float64    `gorm:"column:qty_invoiced" form:"qty_invoiced" json:"qty_invoiced"`
	QtyOrdered                float64    `gorm:"column:qty_ordered" form:"qty_ordered" json:"qty_ordered"`
	QtyRefunded               float64    `gorm:"column:qty_refunded" form:"qty_refunded" json:"qty_refunded"`
	QtyShipped                float64    `gorm:"column:qty_shipped" form:"qty_shipped" json:"qty_shipped"`
	BaseCost                  float64    `gorm:"column:base_cost" form:"base_cost" json:"base_cost"`
	Price                     float64    `gorm:"column:price" form:"price" json:"price"`
	BasePrice                 float64    `gorm:"column:base_price" form:"base_price" json:"base_price"`
	OriginalPrice             float64    `gorm:"column:original_price" form:"original_price" json:"original_price"`
	BaseOriginalPrice         float64    `gorm:"column:base_original_price" form:"base_original_price" json:"base_original_price"`
	TaxPercent                float64    `gorm:"column:tax_percent" form:"tax_percent" json:"tax_percent"`
	TaxAmount                 float64    `gorm:"column:tax_amount" form:"tax_amount" json:"tax_amount"`
	BaseTaxAmount             float64    `gorm:"column:base_tax_amount" form:"base_tax_amount" json:"base_tax_amount"`
	TaxInvoiced               float64    `gorm:"column:tax_invoiced" form:"tax_invoiced" json:"tax_invoiced"`
	BaseTaxInvoiced           float64    `gorm:"column:base_tax_invoiced" form:"base_tax_invoiced" json:"base_tax_invoiced"`
	DiscountPercent           float64    `gorm:"column:discount_percent" form:"discount_percent" json:"discount_percent"`
	DiscountAmount            float64    `gorm:"column:discount_amount" form:"discount_amount" json:"discount_amount"`
	BaseDiscountAmount        float64    `gorm:"column:base_discount_amount" form:"base_discount_amount" json:"base_discount_amount"`
	DiscountInvoiced          float64    `gorm:"column:discount_invoiced" form:"discount_invoiced" json:"discount_invoiced"`
	BaseDiscountInvoiced      float64    `gorm:"column:base_discount_invoiced" form:"base_discount_invoiced" json:"base_discount_invoiced"`
	AmountRefunded            float64    `gorm:"column:amount_refunded" form:"amount_refunded" json:"amount_refunded"`
	BaseAmountRefunded        float64    `gorm:"column:base_amount_refunded" form:"base_amount_refunded" json:"base_amount_refunded"`
	RowTotal                  float64    `gorm:"column:row_total" form:"row_total" json:"row_total"`
	BaseRowTotal              float64    `gorm:"column:base_row_total" form:"base_row_total" json:"base_row_total"`
	RowInvoiced               float64    `gorm:"column:row_invoiced" form:"row_invoiced" json:"row_invoiced"`
	BaseRowInvoiced           float64    `gorm:"column:base_row_invoiced" form:"base_row_invoiced" json:"base_row_invoiced"`
	RowWeight                 float64    `gorm:"column:row_weight" form:"row_weight" json:"row_weight"`
	BaseTaxBeforeDiscount     float64    `gorm:"column:base_tax_before_discount" form:"base_tax_before_discount" json:"base_tax_before_discount"`
	TaxBeforeDiscount         float64    `gorm:"column:tax_before_discount" form:"tax_before_discount" json:"tax_before_discount"`
	ExtOrderItemId            string     `gorm:"column:ext_order_item_id" form:"ext_order_item_id" json:"ext_order_item_id"`
	LockedDoInvoice           uint16     `gorm:"column:locked_do_invoice" form:"locked_do_invoice" json:"locked_do_invoice"`
	LockedDoShip              uint16     `gorm:"column:locked_do_ship" form:"locked_do_ship" json:"locked_do_ship"`
	PriceInclTax              float64    `gorm:"column:price_incl_tax" form:"price_incl_tax" json:"price_incl_tax"`
	BasePriceInclTax          float64    `gorm:"column:base_price_incl_tax" form:"base_price_incl_tax" json:"base_price_incl_tax"`
	RowTotalInclTax           float64    `gorm:"column:row_total_incl_tax" form:"row_total_incl_tax" json:"row_total_incl_tax"`
	BaseRowTotalInclTax       float64    `gorm:"column:base_row_total_incl_tax" form:"base_row_total_incl_tax" json:"base_row_total_incl_tax"`
	HiddenTaxAmount           float64    `gorm:"column:hidden_tax_amount" form:"hidden_tax_amount" json:"hidden_tax_amount"`
	BaseHiddenTaxAmount       float64    `gorm:"column:base_hidden_tax_amount" form:"base_hidden_tax_amount" json:"base_hidden_tax_amount"`
	HiddenTaxInvoiced         float64    `gorm:"column:hidden_tax_invoiced" form:"hidden_tax_invoiced" json:"hidden_tax_invoiced"`
	BaseHiddenTaxInvoiced     float64    `gorm:"column:base_hidden_tax_invoiced" form:"base_hidden_tax_invoiced" json:"base_hidden_tax_invoiced"`
	HiddenTaxRefunded         float64    `gorm:"column:hidden_tax_refunded" form:"hidden_tax_refunded" json:"hidden_tax_refunded"`
	BaseHiddenTaxRefunded     float64    `gorm:"column:base_hidden_tax_refunded" form:"base_hidden_tax_refunded" json:"base_hidden_tax_refunded"`
	IsNominal                 uint16     `gorm:"column:is_nominal" form:"is_nominal" json:"is_nominal"`
	TaxCanceled               float64    `gorm:"column:tax_canceled" form:"tax_canceled" json:"tax_canceled"`
	HiddenTaxCanceled         float64    `gorm:"column:hidden_tax_canceled" form:"hidden_tax_canceled" json:"hidden_tax_canceled"`
	TaxRefunded               float64    `gorm:"column:tax_refunded" form:"tax_refunded" json:"tax_refunded"`
	BaseTaxRefunded           float64    `gorm:"column:base_tax_refunded" form:"base_tax_refunded" json:"base_tax_refunded"`
	DiscountRefunded          float64    `gorm:"column:discount_refunded" form:"discount_refunded" json:"discount_refunded"`
	BaseDiscountRefunded      float64    `gorm:"column:base_discount_refunded" form:"base_discount_refunded" json:"base_discount_refunded"`
	GiftMessageId             uint16     `gorm:"column:gift_message_id" form:"gift_message_id" json:"gift_message_id"`
	GiftMessageAvailable      uint16     `gorm:"column:gift_message_available" form:"gift_message_available" json:"gift_message_available"`
	BaseWeeeTaxAppliedAmount  float64    `gorm:"column:base_weee_tax_applied_amount" form:"base_weee_tax_applied_amount" json:"base_weee_tax_applied_amount"`
	BaseWeeeTaxAppliedRowAmnt float64    `gorm:"column:base_weee_tax_applied_row_amnt" form:"base_weee_tax_applied_row_amnt" json:"base_weee_tax_applied_row_amnt"`
	WeeeTaxAppliedAmount      float64    `gorm:"column:weee_tax_applied_amount" form:"weee_tax_applied_amount" json:"weee_tax_applied_amount"`
	WeeeTaxAppliedRowAmount   float64    `gorm:"column:weee_tax_applied_row_amount" form:"weee_tax_applied_row_amount" json:"weee_tax_applied_row_amount"`
	WeeeTaxApplied            string     `gorm:"column:weee_tax_applied" form:"weee_tax_applied" json:"weee_tax_applied"`
	WeeeTaxDisposition        float64    `gorm:"column:weee_tax_disposition" form:"weee_tax_disposition" json:"weee_tax_disposition"`
	WeeeTaxRowDisposition     float64    `gorm:"column:weee_tax_row_disposition" form:"weee_tax_row_disposition" json:"weee_tax_row_disposition"`
	BaseWeeeTaxDisposition    float64    `gorm:"column:base_weee_tax_disposition" form:"base_weee_tax_disposition" json:"base_weee_tax_disposition"`
	BaseWeeeTaxRowDisposition float64    `gorm:"column:base_weee_tax_row_disposition" form:"base_weee_tax_row_disposition" json:"base_weee_tax_row_disposition"`
}

// TableName :
func (*SalesFlatOrderItem) TableName() string {
	return "sales_flat_order_item"
}

// handler create
func initRoutersSalesFlatOrderItem(r *gin.Engine, salesflatorderitem string) {
	route := r.Group("salesflatorderitem")
	route.GET("/", getSalesFlatOrderItems)
	route.GET("/:id", getSalesFlatOrderItem)
	route.POST("/", createSalesFlatOrderItem)
	route.PUT("/:id", updateSalesFlatOrderItem)
	route.DELETE("/:id", deleteSalesFlatOrderItem)
}

func getSalesFlatOrderItems(c *gin.Context) {
	var salesFlatOrderItems []SalesFlatOrderItem
	if err := g.Find(&salesFlatOrderItems).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatOrderItems)
	}
}

func getSalesFlatOrderItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatOrderItem SalesFlatOrderItem
	if err := g.Where("id = ?", id).First(&salesFlatOrderItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatOrderItem)
	}
}

func createSalesFlatOrderItem(c *gin.Context) {
	var salesFlatOrderItem SalesFlatOrderItem
	c.BindJSON(&salesFlatOrderItem)
	g.Create(&salesFlatOrderItem)
	c.JSON(200, salesFlatOrderItem)
}

func updateSalesFlatOrderItem(c *gin.Context) {
	var salesFlatOrderItem SalesFlatOrderItem
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatOrderItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatOrderItem)
	g.Save(&salesFlatOrderItem)
	c.JSON(200, salesFlatOrderItem)
}
func deleteSalesFlatOrderItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatOrderItem SalesFlatOrderItem
	d := g.Where("id = ?", id).Delete(&salesFlatOrderItem)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
