package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatOrder :
type SalesFlatOrder struct {
	EntityId                   uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	State                      string     `gorm:"column:state" form:"state" json:"state"`
	Status                     string     `gorm:"column:status" form:"status" json:"status"`
	CouponCode                 string     `gorm:"column:coupon_code" form:"coupon_code" json:"coupon_code"`
	ProtectCode                string     `gorm:"column:protect_code" form:"protect_code" json:"protect_code"`
	ShippingDescription        string     `gorm:"column:shipping_description" form:"shipping_description" json:"shipping_description"`
	IsVirtual                  uint16     `gorm:"column:is_virtual" form:"is_virtual" json:"is_virtual"`
	StoreId                    uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	CustomerId                 uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	BaseDiscountAmount         float64    `gorm:"column:base_discount_amount" form:"base_discount_amount" json:"base_discount_amount"`
	BaseDiscountCanceled       float64    `gorm:"column:base_discount_canceled" form:"base_discount_canceled" json:"base_discount_canceled"`
	BaseDiscountInvoiced       float64    `gorm:"column:base_discount_invoiced" form:"base_discount_invoiced" json:"base_discount_invoiced"`
	BaseDiscountRefunded       float64    `gorm:"column:base_discount_refunded" form:"base_discount_refunded" json:"base_discount_refunded"`
	BaseGrandTotal             float64    `gorm:"column:base_grand_total" form:"base_grand_total" json:"base_grand_total"`
	BaseShippingAmount         float64    `gorm:"column:base_shipping_amount" form:"base_shipping_amount" json:"base_shipping_amount"`
	BaseShippingCanceled       float64    `gorm:"column:base_shipping_canceled" form:"base_shipping_canceled" json:"base_shipping_canceled"`
	BaseShippingInvoiced       float64    `gorm:"column:base_shipping_invoiced" form:"base_shipping_invoiced" json:"base_shipping_invoiced"`
	BaseShippingRefunded       float64    `gorm:"column:base_shipping_refunded" form:"base_shipping_refunded" json:"base_shipping_refunded"`
	BaseShippingTaxAmount      float64    `gorm:"column:base_shipping_tax_amount" form:"base_shipping_tax_amount" json:"base_shipping_tax_amount"`
	BaseShippingTaxRefunded    float64    `gorm:"column:base_shipping_tax_refunded" form:"base_shipping_tax_refunded" json:"base_shipping_tax_refunded"`
	BaseSubtotal               float64    `gorm:"column:base_subtotal" form:"base_subtotal" json:"base_subtotal"`
	BaseSubtotalCanceled       float64    `gorm:"column:base_subtotal_canceled" form:"base_subtotal_canceled" json:"base_subtotal_canceled"`
	BaseSubtotalInvoiced       float64    `gorm:"column:base_subtotal_invoiced" form:"base_subtotal_invoiced" json:"base_subtotal_invoiced"`
	BaseSubtotalRefunded       float64    `gorm:"column:base_subtotal_refunded" form:"base_subtotal_refunded" json:"base_subtotal_refunded"`
	BaseTaxAmount              float64    `gorm:"column:base_tax_amount" form:"base_tax_amount" json:"base_tax_amount"`
	BaseTaxCanceled            float64    `gorm:"column:base_tax_canceled" form:"base_tax_canceled" json:"base_tax_canceled"`
	BaseTaxInvoiced            float64    `gorm:"column:base_tax_invoiced" form:"base_tax_invoiced" json:"base_tax_invoiced"`
	BaseTaxRefunded            float64    `gorm:"column:base_tax_refunded" form:"base_tax_refunded" json:"base_tax_refunded"`
	BaseToGlobalRate           float64    `gorm:"column:base_to_global_rate" form:"base_to_global_rate" json:"base_to_global_rate"`
	BaseToOrderRate            float64    `gorm:"column:base_to_order_rate" form:"base_to_order_rate" json:"base_to_order_rate"`
	BaseTotalCanceled          float64    `gorm:"column:base_total_canceled" form:"base_total_canceled" json:"base_total_canceled"`
	BaseTotalInvoiced          float64    `gorm:"column:base_total_invoiced" form:"base_total_invoiced" json:"base_total_invoiced"`
	BaseTotalInvoicedCost      float64    `gorm:"column:base_total_invoiced_cost" form:"base_total_invoiced_cost" json:"base_total_invoiced_cost"`
	BaseTotalOfflineRefunded   float64    `gorm:"column:base_total_offline_refunded" form:"base_total_offline_refunded" json:"base_total_offline_refunded"`
	BaseTotalOnlineRefunded    float64    `gorm:"column:base_total_online_refunded" form:"base_total_online_refunded" json:"base_total_online_refunded"`
	BaseTotalPaid              float64    `gorm:"column:base_total_paid" form:"base_total_paid" json:"base_total_paid"`
	BaseTotalQtyOrdered        float64    `gorm:"column:base_total_qty_ordered" form:"base_total_qty_ordered" json:"base_total_qty_ordered"`
	BaseTotalRefunded          float64    `gorm:"column:base_total_refunded" form:"base_total_refunded" json:"base_total_refunded"`
	DiscountAmount             float64    `gorm:"column:discount_amount" form:"discount_amount" json:"discount_amount"`
	DiscountCanceled           float64    `gorm:"column:discount_canceled" form:"discount_canceled" json:"discount_canceled"`
	DiscountInvoiced           float64    `gorm:"column:discount_invoiced" form:"discount_invoiced" json:"discount_invoiced"`
	DiscountRefunded           float64    `gorm:"column:discount_refunded" form:"discount_refunded" json:"discount_refunded"`
	GrandTotal                 float64    `gorm:"column:grand_total" form:"grand_total" json:"grand_total"`
	ShippingAmount             float64    `gorm:"column:shipping_amount" form:"shipping_amount" json:"shipping_amount"`
	ShippingCanceled           float64    `gorm:"column:shipping_canceled" form:"shipping_canceled" json:"shipping_canceled"`
	ShippingInvoiced           float64    `gorm:"column:shipping_invoiced" form:"shipping_invoiced" json:"shipping_invoiced"`
	ShippingRefunded           float64    `gorm:"column:shipping_refunded" form:"shipping_refunded" json:"shipping_refunded"`
	ShippingTaxAmount          float64    `gorm:"column:shipping_tax_amount" form:"shipping_tax_amount" json:"shipping_tax_amount"`
	ShippingTaxRefunded        float64    `gorm:"column:shipping_tax_refunded" form:"shipping_tax_refunded" json:"shipping_tax_refunded"`
	StoreToBaseRate            float64    `gorm:"column:store_to_base_rate" form:"store_to_base_rate" json:"store_to_base_rate"`
	StoreToOrderRate           float64    `gorm:"column:store_to_order_rate" form:"store_to_order_rate" json:"store_to_order_rate"`
	Subtotal                   float64    `gorm:"column:subtotal" form:"subtotal" json:"subtotal"`
	SubtotalCanceled           float64    `gorm:"column:subtotal_canceled" form:"subtotal_canceled" json:"subtotal_canceled"`
	SubtotalInvoiced           float64    `gorm:"column:subtotal_invoiced" form:"subtotal_invoiced" json:"subtotal_invoiced"`
	SubtotalRefunded           float64    `gorm:"column:subtotal_refunded" form:"subtotal_refunded" json:"subtotal_refunded"`
	TaxAmount                  float64    `gorm:"column:tax_amount" form:"tax_amount" json:"tax_amount"`
	TaxCanceled                float64    `gorm:"column:tax_canceled" form:"tax_canceled" json:"tax_canceled"`
	TaxInvoiced                float64    `gorm:"column:tax_invoiced" form:"tax_invoiced" json:"tax_invoiced"`
	TaxRefunded                float64    `gorm:"column:tax_refunded" form:"tax_refunded" json:"tax_refunded"`
	TotalCanceled              float64    `gorm:"column:total_canceled" form:"total_canceled" json:"total_canceled"`
	TotalInvoiced              float64    `gorm:"column:total_invoiced" form:"total_invoiced" json:"total_invoiced"`
	TotalOfflineRefunded       float64    `gorm:"column:total_offline_refunded" form:"total_offline_refunded" json:"total_offline_refunded"`
	TotalOnlineRefunded        float64    `gorm:"column:total_online_refunded" form:"total_online_refunded" json:"total_online_refunded"`
	TotalPaid                  float64    `gorm:"column:total_paid" form:"total_paid" json:"total_paid"`
	TotalQtyOrdered            float64    `gorm:"column:total_qty_ordered" form:"total_qty_ordered" json:"total_qty_ordered"`
	TotalRefunded              float64    `gorm:"column:total_refunded" form:"total_refunded" json:"total_refunded"`
	CanShipPartially           uint16     `gorm:"column:can_ship_partially" form:"can_ship_partially" json:"can_ship_partially"`
	CanShipPartiallyItem       uint16     `gorm:"column:can_ship_partially_item" form:"can_ship_partially_item" json:"can_ship_partially_item"`
	CustomerIsGuest            uint16     `gorm:"column:customer_is_guest" form:"customer_is_guest" json:"customer_is_guest"`
	CustomerNoteNotify         uint16     `gorm:"column:customer_note_notify" form:"customer_note_notify" json:"customer_note_notify"`
	BillingAddressId           uint16     `gorm:"column:billing_address_id" form:"billing_address_id" json:"billing_address_id"`
	CustomerGroupId            uint16     `gorm:"column:customer_group_id" form:"customer_group_id" json:"customer_group_id"`
	EditIncrement              uint16     `gorm:"column:edit_increment" form:"edit_increment" json:"edit_increment"`
	EmailSent                  uint16     `gorm:"column:email_sent" form:"email_sent" json:"email_sent"`
	ForcedShipmentWithInvoice  uint16     `gorm:"column:forced_shipment_with_invoice" form:"forced_shipment_with_invoice" json:"forced_shipment_with_invoice"`
	PaymentAuthExpiration      uint16     `gorm:"column:payment_auth_expiration" form:"payment_auth_expiration" json:"payment_auth_expiration"`
	QuoteAddressId             uint16     `gorm:"column:quote_address_id" form:"quote_address_id" json:"quote_address_id"`
	QuoteId                    uint16     `gorm:"column:quote_id" form:"quote_id" json:"quote_id"`
	ShippingAddressId          uint16     `gorm:"column:shipping_address_id" form:"shipping_address_id" json:"shipping_address_id"`
	AdjustmentNegative         float64    `gorm:"column:adjustment_negative" form:"adjustment_negative" json:"adjustment_negative"`
	AdjustmentPositive         float64    `gorm:"column:adjustment_positive" form:"adjustment_positive" json:"adjustment_positive"`
	BaseAdjustmentNegative     float64    `gorm:"column:base_adjustment_negative" form:"base_adjustment_negative" json:"base_adjustment_negative"`
	BaseAdjustmentPositive     float64    `gorm:"column:base_adjustment_positive" form:"base_adjustment_positive" json:"base_adjustment_positive"`
	BaseShippingDiscountAmount float64    `gorm:"column:base_shipping_discount_amount" form:"base_shipping_discount_amount" json:"base_shipping_discount_amount"`
	BaseSubtotalInclTax        float64    `gorm:"column:base_subtotal_incl_tax" form:"base_subtotal_incl_tax" json:"base_subtotal_incl_tax"`
	BaseTotalDue               float64    `gorm:"column:base_total_due" form:"base_total_due" json:"base_total_due"`
	PaymentAuthorizationAmount float64    `gorm:"column:payment_authorization_amount" form:"payment_authorization_amount" json:"payment_authorization_amount"`
	ShippingDiscountAmount     float64    `gorm:"column:shipping_discount_amount" form:"shipping_discount_amount" json:"shipping_discount_amount"`
	SubtotalInclTax            float64    `gorm:"column:subtotal_incl_tax" form:"subtotal_incl_tax" json:"subtotal_incl_tax"`
	TotalDue                   float64    `gorm:"column:total_due" form:"total_due" json:"total_due"`
	Weight                     float64    `gorm:"column:weight" form:"weight" json:"weight"`
	CustomerDob                *time.Time `gorm:"column:customer_dob" form:"customer_dob" json:"customer_dob"`
	IncrementId                string     `gorm:"column:increment_id" form:"increment_id" json:"increment_id"`
	AppliedRuleIds             string     `gorm:"column:applied_rule_ids" form:"applied_rule_ids" json:"applied_rule_ids"`
	BaseCurrencyCode           string     `gorm:"column:base_currency_code" form:"base_currency_code" json:"base_currency_code"`
	CustomerEmail              string     `gorm:"column:customer_email" form:"customer_email" json:"customer_email"`
	CustomerFirstname          string     `gorm:"column:customer_firstname" form:"customer_firstname" json:"customer_firstname"`
	CustomerLastname           string     `gorm:"column:customer_lastname" form:"customer_lastname" json:"customer_lastname"`
	CustomerMiddlename         string     `gorm:"column:customer_middlename" form:"customer_middlename" json:"customer_middlename"`
	CustomerPrefix             string     `gorm:"column:customer_prefix" form:"customer_prefix" json:"customer_prefix"`
	CustomerSuffix             string     `gorm:"column:customer_suffix" form:"customer_suffix" json:"customer_suffix"`
	CustomerTaxvat             string     `gorm:"column:customer_taxvat" form:"customer_taxvat" json:"customer_taxvat"`
	DiscountDescription        string     `gorm:"column:discount_description" form:"discount_description" json:"discount_description"`
	ExtCustomerId              string     `gorm:"column:ext_customer_id" form:"ext_customer_id" json:"ext_customer_id"`
	ExtOrderId                 string     `gorm:"column:ext_order_id" form:"ext_order_id" json:"ext_order_id"`
	GlobalCurrencyCode         string     `gorm:"column:global_currency_code" form:"global_currency_code" json:"global_currency_code"`
	HoldBeforeState            string     `gorm:"column:hold_before_state" form:"hold_before_state" json:"hold_before_state"`
	HoldBeforeStatus           string     `gorm:"column:hold_before_status" form:"hold_before_status" json:"hold_before_status"`
	OrderCurrencyCode          string     `gorm:"column:order_currency_code" form:"order_currency_code" json:"order_currency_code"`
	OriginalIncrementId        string     `gorm:"column:original_increment_id" form:"original_increment_id" json:"original_increment_id"`
	RelationChildId            string     `gorm:"column:relation_child_id" form:"relation_child_id" json:"relation_child_id"`
	RelationChildRealId        string     `gorm:"column:relation_child_real_id" form:"relation_child_real_id" json:"relation_child_real_id"`
	RelationParentId           string     `gorm:"column:relation_parent_id" form:"relation_parent_id" json:"relation_parent_id"`
	RelationParentRealId       string     `gorm:"column:relation_parent_real_id" form:"relation_parent_real_id" json:"relation_parent_real_id"`
	RemoteIp                   string     `gorm:"column:remote_ip" form:"remote_ip" json:"remote_ip"`
	ShippingMethod             string     `gorm:"column:shipping_method" form:"shipping_method" json:"shipping_method"`
	StoreCurrencyCode          string     `gorm:"column:store_currency_code" form:"store_currency_code" json:"store_currency_code"`
	StoreName                  string     `gorm:"column:store_name" form:"store_name" json:"store_name"`
	XForwardedFor              string     `gorm:"column:x_forwarded_for" form:"x_forwarded_for" json:"x_forwarded_for"`
	CustomerNote               string     `gorm:"column:customer_note" form:"customer_note" json:"customer_note"`
	CreatedAt                  *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt                  *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	TotalItemCount             uint16     `gorm:"column:total_item_count" form:"total_item_count" json:"total_item_count"`
	CustomerGender             uint16     `gorm:"column:customer_gender" form:"customer_gender" json:"customer_gender"`
	HiddenTaxAmount            float64    `gorm:"column:hidden_tax_amount" form:"hidden_tax_amount" json:"hidden_tax_amount"`
	BaseHiddenTaxAmount        float64    `gorm:"column:base_hidden_tax_amount" form:"base_hidden_tax_amount" json:"base_hidden_tax_amount"`
	ShippingHiddenTaxAmount    float64    `gorm:"column:shipping_hidden_tax_amount" form:"shipping_hidden_tax_amount" json:"shipping_hidden_tax_amount"`
	BaseShippingHiddenTaxAmnt  float64    `gorm:"column:base_shipping_hidden_tax_amnt" form:"base_shipping_hidden_tax_amnt" json:"base_shipping_hidden_tax_amnt"`
	HiddenTaxInvoiced          float64    `gorm:"column:hidden_tax_invoiced" form:"hidden_tax_invoiced" json:"hidden_tax_invoiced"`
	BaseHiddenTaxInvoiced      float64    `gorm:"column:base_hidden_tax_invoiced" form:"base_hidden_tax_invoiced" json:"base_hidden_tax_invoiced"`
	HiddenTaxRefunded          float64    `gorm:"column:hidden_tax_refunded" form:"hidden_tax_refunded" json:"hidden_tax_refunded"`
	BaseHiddenTaxRefunded      float64    `gorm:"column:base_hidden_tax_refunded" form:"base_hidden_tax_refunded" json:"base_hidden_tax_refunded"`
	ShippingInclTax            float64    `gorm:"column:shipping_incl_tax" form:"shipping_incl_tax" json:"shipping_incl_tax"`
	BaseShippingInclTax        float64    `gorm:"column:base_shipping_incl_tax" form:"base_shipping_incl_tax" json:"base_shipping_incl_tax"`
	CouponRuleName             string     `gorm:"column:coupon_rule_name" form:"coupon_rule_name" json:"coupon_rule_name"`
	PaypalIpnCustomerNotified  uint16     `gorm:"column:paypal_ipn_customer_notified" form:"paypal_ipn_customer_notified" json:"paypal_ipn_customer_notified"`
	GiftMessageId              uint16     `gorm:"column:gift_message_id" form:"gift_message_id" json:"gift_message_id"`
	GemgentoId                 uint16     `gorm:"column:gemgento_id" form:"gemgento_id" json:"gemgento_id"`
}

// TableName :
func (*SalesFlatOrder) TableName() string {
	return "sales_flat_order"
}

// handler create
func initRoutersSalesFlatOrder(r *gin.Engine, salesflatorder string) {
	route := r.Group("salesflatorder")
	route.GET("/", getSalesFlatOrders)
	route.GET("/:id", getSalesFlatOrder)
	route.POST("/", createSalesFlatOrder)
	route.PUT("/:id", updateSalesFlatOrder)
	route.DELETE("/:id", deleteSalesFlatOrder)
}

func getSalesFlatOrders(c *gin.Context) {
	var salesFlatOrders []SalesFlatOrder
	if err := g.Find(&salesFlatOrders).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatOrders)
	}
}

func getSalesFlatOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatOrder SalesFlatOrder
	if err := g.Where("id = ?", id).First(&salesFlatOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatOrder)
	}
}

func createSalesFlatOrder(c *gin.Context) {
	var salesFlatOrder SalesFlatOrder
	c.BindJSON(&salesFlatOrder)
	g.Create(&salesFlatOrder)
	c.JSON(200, salesFlatOrder)
}

func updateSalesFlatOrder(c *gin.Context) {
	var salesFlatOrder SalesFlatOrder
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatOrder)
	g.Save(&salesFlatOrder)
	c.JSON(200, salesFlatOrder)
}
func deleteSalesFlatOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatOrder SalesFlatOrder
	d := g.Where("id = ?", id).Delete(&salesFlatOrder)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
