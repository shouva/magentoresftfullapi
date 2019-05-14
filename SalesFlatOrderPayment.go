package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesFlatOrderPayment :
type SalesFlatOrderPayment struct {
	EntityId                  uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	ParentId                  uint16  `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	BaseShippingCaptured      float64 `gorm:"column:base_shipping_captured" form:"base_shipping_captured" json:"base_shipping_captured"`
	ShippingCaptured          float64 `gorm:"column:shipping_captured" form:"shipping_captured" json:"shipping_captured"`
	AmountRefunded            float64 `gorm:"column:amount_refunded" form:"amount_refunded" json:"amount_refunded"`
	BaseAmountPaid            float64 `gorm:"column:base_amount_paid" form:"base_amount_paid" json:"base_amount_paid"`
	AmountCanceled            float64 `gorm:"column:amount_canceled" form:"amount_canceled" json:"amount_canceled"`
	BaseAmountAuthorized      float64 `gorm:"column:base_amount_authorized" form:"base_amount_authorized" json:"base_amount_authorized"`
	BaseAmountPaidOnline      float64 `gorm:"column:base_amount_paid_online" form:"base_amount_paid_online" json:"base_amount_paid_online"`
	BaseAmountRefundedOnline  float64 `gorm:"column:base_amount_refunded_online" form:"base_amount_refunded_online" json:"base_amount_refunded_online"`
	BaseShippingAmount        float64 `gorm:"column:base_shipping_amount" form:"base_shipping_amount" json:"base_shipping_amount"`
	ShippingAmount            float64 `gorm:"column:shipping_amount" form:"shipping_amount" json:"shipping_amount"`
	AmountPaid                float64 `gorm:"column:amount_paid" form:"amount_paid" json:"amount_paid"`
	AmountAuthorized          float64 `gorm:"column:amount_authorized" form:"amount_authorized" json:"amount_authorized"`
	BaseAmountOrdered         float64 `gorm:"column:base_amount_ordered" form:"base_amount_ordered" json:"base_amount_ordered"`
	BaseShippingRefunded      float64 `gorm:"column:base_shipping_refunded" form:"base_shipping_refunded" json:"base_shipping_refunded"`
	ShippingRefunded          float64 `gorm:"column:shipping_refunded" form:"shipping_refunded" json:"shipping_refunded"`
	BaseAmountRefunded        float64 `gorm:"column:base_amount_refunded" form:"base_amount_refunded" json:"base_amount_refunded"`
	AmountOrdered             float64 `gorm:"column:amount_ordered" form:"amount_ordered" json:"amount_ordered"`
	BaseAmountCanceled        float64 `gorm:"column:base_amount_canceled" form:"base_amount_canceled" json:"base_amount_canceled"`
	QuotePaymentId            uint16  `gorm:"column:quote_payment_id" form:"quote_payment_id" json:"quote_payment_id"`
	AdditionalData            string  `gorm:"column:additional_data" form:"additional_data" json:"additional_data"`
	CcExpMonth                string  `gorm:"column:cc_exp_month" form:"cc_exp_month" json:"cc_exp_month"`
	CcSsStartYear             string  `gorm:"column:cc_ss_start_year" form:"cc_ss_start_year" json:"cc_ss_start_year"`
	EcheckBankName            string  `gorm:"column:echeck_bank_name" form:"echeck_bank_name" json:"echeck_bank_name"`
	Method                    string  `gorm:"column:method" form:"method" json:"method"`
	CcDebugRequestBody        string  `gorm:"column:cc_debug_request_body" form:"cc_debug_request_body" json:"cc_debug_request_body"`
	CcSecureVerify            string  `gorm:"column:cc_secure_verify" form:"cc_secure_verify" json:"cc_secure_verify"`
	ProtectionEligibility     string  `gorm:"column:protection_eligibility" form:"protection_eligibility" json:"protection_eligibility"`
	CcApproval                string  `gorm:"column:cc_approval" form:"cc_approval" json:"cc_approval"`
	CcLast4                   string  `gorm:"column:cc_last4" form:"cc_last4" json:"cc_last4"`
	CcStatusDescription       string  `gorm:"column:cc_status_description" form:"cc_status_description" json:"cc_status_description"`
	EcheckType                string  `gorm:"column:echeck_type" form:"echeck_type" json:"echeck_type"`
	CcDebugResponseSerialized string  `gorm:"column:cc_debug_response_serialized" form:"cc_debug_response_serialized" json:"cc_debug_response_serialized"`
	CcSsStartMonth            string  `gorm:"column:cc_ss_start_month" form:"cc_ss_start_month" json:"cc_ss_start_month"`
	EcheckAccountType         string  `gorm:"column:echeck_account_type" form:"echeck_account_type" json:"echeck_account_type"`
	LastTransId               string  `gorm:"column:last_trans_id" form:"last_trans_id" json:"last_trans_id"`
	CcCidStatus               string  `gorm:"column:cc_cid_status" form:"cc_cid_status" json:"cc_cid_status"`
	CcOwner                   string  `gorm:"column:cc_owner" form:"cc_owner" json:"cc_owner"`
	CcType                    string  `gorm:"column:cc_type" form:"cc_type" json:"cc_type"`
	PoNumber                  string  `gorm:"column:po_number" form:"po_number" json:"po_number"`
	CcExpYear                 string  `gorm:"column:cc_exp_year" form:"cc_exp_year" json:"cc_exp_year"`
	CcStatus                  string  `gorm:"column:cc_status" form:"cc_status" json:"cc_status"`
	EcheckRoutingNumber       string  `gorm:"column:echeck_routing_number" form:"echeck_routing_number" json:"echeck_routing_number"`
	AccountStatus             string  `gorm:"column:account_status" form:"account_status" json:"account_status"`
	AnetTransMethod           string  `gorm:"column:anet_trans_method" form:"anet_trans_method" json:"anet_trans_method"`
	CcDebugResponseBody       string  `gorm:"column:cc_debug_response_body" form:"cc_debug_response_body" json:"cc_debug_response_body"`
	CcSsIssue                 string  `gorm:"column:cc_ss_issue" form:"cc_ss_issue" json:"cc_ss_issue"`
	EcheckAccountName         string  `gorm:"column:echeck_account_name" form:"echeck_account_name" json:"echeck_account_name"`
	CcAvsStatus               string  `gorm:"column:cc_avs_status" form:"cc_avs_status" json:"cc_avs_status"`
	CcNumberEnc               string  `gorm:"column:cc_number_enc" form:"cc_number_enc" json:"cc_number_enc"`
	CcTransId                 string  `gorm:"column:cc_trans_id" form:"cc_trans_id" json:"cc_trans_id"`
	PayboxRequestNumber       string  `gorm:"column:paybox_request_number" form:"paybox_request_number" json:"paybox_request_number"`
	AddressStatus             string  `gorm:"column:address_status" form:"address_status" json:"address_status"`
	AdditionalInformation     string  `gorm:"column:additional_information" form:"additional_information" json:"additional_information"`
}

// TableName :
func (*SalesFlatOrderPayment) TableName() string {
	return "sales_flat_order_payment"
}

// handler create
func initRoutersSalesFlatOrderPayment(r *gin.Engine, salesflatorderpayment string) {
	route := r.Group("salesflatorderpayment")
	route.GET("/", getSalesFlatOrderPayments)
	route.GET("/:id", getSalesFlatOrderPayment)
	route.POST("/", createSalesFlatOrderPayment)
	route.PUT("/:id", updateSalesFlatOrderPayment)
	route.DELETE("/:id", deleteSalesFlatOrderPayment)
}

func getSalesFlatOrderPayments(c *gin.Context) {
	var salesFlatOrderPayments []SalesFlatOrderPayment
	if err := g.Find(&salesFlatOrderPayments).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatOrderPayments)
	}
}

func getSalesFlatOrderPayment(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatOrderPayment SalesFlatOrderPayment
	if err := g.Where("id = ?", id).First(&salesFlatOrderPayment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatOrderPayment)
	}
}

func createSalesFlatOrderPayment(c *gin.Context) {
	var salesFlatOrderPayment SalesFlatOrderPayment
	c.BindJSON(&salesFlatOrderPayment)
	g.Create(&salesFlatOrderPayment)
	c.JSON(200, salesFlatOrderPayment)
}

func updateSalesFlatOrderPayment(c *gin.Context) {
	var salesFlatOrderPayment SalesFlatOrderPayment
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatOrderPayment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatOrderPayment)
	g.Save(&salesFlatOrderPayment)
	c.JSON(200, salesFlatOrderPayment)
}
func deleteSalesFlatOrderPayment(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatOrderPayment SalesFlatOrderPayment
	d := g.Where("id = ?", id).Delete(&salesFlatOrderPayment)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
