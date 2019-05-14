package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatQuote :
type SalesFlatQuote struct {
	EntityId                 uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	StoreId                  uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	CreatedAt                *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt                *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	ConvertedAt              *time.Time `gorm:"column:converted_at" form:"converted_at" json:"converted_at"`
	IsActive                 uint16     `gorm:"column:is_active" form:"is_active" json:"is_active"`
	IsVirtual                uint16     `gorm:"column:is_virtual" form:"is_virtual" json:"is_virtual"`
	IsMultiShipping          uint16     `gorm:"column:is_multi_shipping" form:"is_multi_shipping" json:"is_multi_shipping"`
	ItemsCount               uint16     `gorm:"column:items_count" form:"items_count" json:"items_count"`
	ItemsQty                 float64    `gorm:"column:items_qty" form:"items_qty" json:"items_qty"`
	OrigOrderId              uint16     `gorm:"column:orig_order_id" form:"orig_order_id" json:"orig_order_id"`
	StoreToBaseRate          float64    `gorm:"column:store_to_base_rate" form:"store_to_base_rate" json:"store_to_base_rate"`
	StoreToQuoteRate         float64    `gorm:"column:store_to_quote_rate" form:"store_to_quote_rate" json:"store_to_quote_rate"`
	BaseCurrencyCode         string     `gorm:"column:base_currency_code" form:"base_currency_code" json:"base_currency_code"`
	StoreCurrencyCode        string     `gorm:"column:store_currency_code" form:"store_currency_code" json:"store_currency_code"`
	QuoteCurrencyCode        string     `gorm:"column:quote_currency_code" form:"quote_currency_code" json:"quote_currency_code"`
	GrandTotal               float64    `gorm:"column:grand_total" form:"grand_total" json:"grand_total"`
	BaseGrandTotal           float64    `gorm:"column:base_grand_total" form:"base_grand_total" json:"base_grand_total"`
	CheckoutMethod           string     `gorm:"column:checkout_method" form:"checkout_method" json:"checkout_method"`
	CustomerId               uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	CustomerTaxClassId       uint16     `gorm:"column:customer_tax_class_id" form:"customer_tax_class_id" json:"customer_tax_class_id"`
	CustomerGroupId          uint16     `gorm:"column:customer_group_id" form:"customer_group_id" json:"customer_group_id"`
	CustomerEmail            string     `gorm:"column:customer_email" form:"customer_email" json:"customer_email"`
	CustomerPrefix           string     `gorm:"column:customer_prefix" form:"customer_prefix" json:"customer_prefix"`
	CustomerFirstname        string     `gorm:"column:customer_firstname" form:"customer_firstname" json:"customer_firstname"`
	CustomerMiddlename       string     `gorm:"column:customer_middlename" form:"customer_middlename" json:"customer_middlename"`
	CustomerLastname         string     `gorm:"column:customer_lastname" form:"customer_lastname" json:"customer_lastname"`
	CustomerSuffix           string     `gorm:"column:customer_suffix" form:"customer_suffix" json:"customer_suffix"`
	CustomerDob              *time.Time `gorm:"column:customer_dob" form:"customer_dob" json:"customer_dob"`
	CustomerNote             string     `gorm:"column:customer_note" form:"customer_note" json:"customer_note"`
	CustomerNoteNotify       uint16     `gorm:"column:customer_note_notify" form:"customer_note_notify" json:"customer_note_notify"`
	CustomerIsGuest          uint16     `gorm:"column:customer_is_guest" form:"customer_is_guest" json:"customer_is_guest"`
	RemoteIp                 string     `gorm:"column:remote_ip" form:"remote_ip" json:"remote_ip"`
	AppliedRuleIds           string     `gorm:"column:applied_rule_ids" form:"applied_rule_ids" json:"applied_rule_ids"`
	ReservedOrderId          string     `gorm:"column:reserved_order_id" form:"reserved_order_id" json:"reserved_order_id"`
	PasswordHash             string     `gorm:"column:password_hash" form:"password_hash" json:"password_hash"`
	CouponCode               string     `gorm:"column:coupon_code" form:"coupon_code" json:"coupon_code"`
	GlobalCurrencyCode       string     `gorm:"column:global_currency_code" form:"global_currency_code" json:"global_currency_code"`
	BaseToGlobalRate         float64    `gorm:"column:base_to_global_rate" form:"base_to_global_rate" json:"base_to_global_rate"`
	BaseToQuoteRate          float64    `gorm:"column:base_to_quote_rate" form:"base_to_quote_rate" json:"base_to_quote_rate"`
	CustomerTaxvat           string     `gorm:"column:customer_taxvat" form:"customer_taxvat" json:"customer_taxvat"`
	CustomerGender           string     `gorm:"column:customer_gender" form:"customer_gender" json:"customer_gender"`
	Subtotal                 float64    `gorm:"column:subtotal" form:"subtotal" json:"subtotal"`
	BaseSubtotal             float64    `gorm:"column:base_subtotal" form:"base_subtotal" json:"base_subtotal"`
	SubtotalWithDiscount     float64    `gorm:"column:subtotal_with_discount" form:"subtotal_with_discount" json:"subtotal_with_discount"`
	BaseSubtotalWithDiscount float64    `gorm:"column:base_subtotal_with_discount" form:"base_subtotal_with_discount" json:"base_subtotal_with_discount"`
	IsChanged                uint16     `gorm:"column:is_changed" form:"is_changed" json:"is_changed"`
	TriggerRecollect         uint16     `gorm:"column:trigger_recollect" form:"trigger_recollect" json:"trigger_recollect"`
	ExtShippingInfo          string     `gorm:"column:ext_shipping_info" form:"ext_shipping_info" json:"ext_shipping_info"`
	GiftMessageId            uint16     `gorm:"column:gift_message_id" form:"gift_message_id" json:"gift_message_id"`
	IsPersistent             uint16     `gorm:"column:is_persistent" form:"is_persistent" json:"is_persistent"`
	GemgentoId               uint16     `gorm:"column:gemgento_id" form:"gemgento_id" json:"gemgento_id"`
}

// TableName :
func (*SalesFlatQuote) TableName() string {
	return "sales_flat_quote"
}

// handler create
func initRoutersSalesFlatQuote(r *gin.Engine, salesflatquote string) {
	route := r.Group("salesflatquote")
	route.GET("/", getSalesFlatQuotes)
	route.GET("/:id", getSalesFlatQuote)
	route.POST("/", createSalesFlatQuote)
	route.PUT("/:id", updateSalesFlatQuote)
	route.DELETE("/:id", deleteSalesFlatQuote)
}

func getSalesFlatQuotes(c *gin.Context) {
	var salesFlatQuotes []SalesFlatQuote
	if err := g.Find(&salesFlatQuotes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuotes)
	}
}

func getSalesFlatQuote(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuote SalesFlatQuote
	if err := g.Where("id = ?", id).First(&salesFlatQuote).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuote)
	}
}

func createSalesFlatQuote(c *gin.Context) {
	var salesFlatQuote SalesFlatQuote
	c.BindJSON(&salesFlatQuote)
	g.Create(&salesFlatQuote)
	c.JSON(200, salesFlatQuote)
}

func updateSalesFlatQuote(c *gin.Context) {
	var salesFlatQuote SalesFlatQuote
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatQuote).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatQuote)
	g.Save(&salesFlatQuote)
	c.JSON(200, salesFlatQuote)
}
func deleteSalesFlatQuote(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuote SalesFlatQuote
	d := g.Where("id = ?", id).Delete(&salesFlatQuote)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
