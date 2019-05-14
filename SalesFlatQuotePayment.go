package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatQuotePayment :
type SalesFlatQuotePayment struct {
	PaymentId             uint16     `gorm:"column:payment_id;primary_key" form:"payment_id;primary_key" json:"payment_id;primary_key"`
	QuoteId               uint16     `gorm:"column:quote_id" form:"quote_id" json:"quote_id"`
	CreatedAt             *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt             *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	Method                string     `gorm:"column:method" form:"method" json:"method"`
	CcType                string     `gorm:"column:cc_type" form:"cc_type" json:"cc_type"`
	CcNumberEnc           string     `gorm:"column:cc_number_enc" form:"cc_number_enc" json:"cc_number_enc"`
	CcLast4               string     `gorm:"column:cc_last4" form:"cc_last4" json:"cc_last4"`
	CcCidEnc              string     `gorm:"column:cc_cid_enc" form:"cc_cid_enc" json:"cc_cid_enc"`
	CcOwner               string     `gorm:"column:cc_owner" form:"cc_owner" json:"cc_owner"`
	CcExpMonth            uint16     `gorm:"column:cc_exp_month" form:"cc_exp_month" json:"cc_exp_month"`
	CcExpYear             uint16     `gorm:"column:cc_exp_year" form:"cc_exp_year" json:"cc_exp_year"`
	CcSsOwner             string     `gorm:"column:cc_ss_owner" form:"cc_ss_owner" json:"cc_ss_owner"`
	CcSsStartMonth        uint16     `gorm:"column:cc_ss_start_month" form:"cc_ss_start_month" json:"cc_ss_start_month"`
	CcSsStartYear         uint16     `gorm:"column:cc_ss_start_year" form:"cc_ss_start_year" json:"cc_ss_start_year"`
	PoNumber              string     `gorm:"column:po_number" form:"po_number" json:"po_number"`
	AdditionalData        string     `gorm:"column:additional_data" form:"additional_data" json:"additional_data"`
	CcSsIssue             string     `gorm:"column:cc_ss_issue" form:"cc_ss_issue" json:"cc_ss_issue"`
	AdditionalInformation string     `gorm:"column:additional_information" form:"additional_information" json:"additional_information"`
	PaypalPayerId         string     `gorm:"column:paypal_payer_id" form:"paypal_payer_id" json:"paypal_payer_id"`
	PaypalPayerStatus     string     `gorm:"column:paypal_payer_status" form:"paypal_payer_status" json:"paypal_payer_status"`
	PaypalCorrelationId   string     `gorm:"column:paypal_correlation_id" form:"paypal_correlation_id" json:"paypal_correlation_id"`
}

// TableName :
func (*SalesFlatQuotePayment) TableName() string {
	return "sales_flat_quote_payment"
}

// handler create
func initRoutersSalesFlatQuotePayment(r *gin.Engine, salesflatquotepayment string) {
	route := r.Group("salesflatquotepayment")
	route.GET("/", getSalesFlatQuotePayments)
	route.GET("/:id", getSalesFlatQuotePayment)
	route.POST("/", createSalesFlatQuotePayment)
	route.PUT("/:id", updateSalesFlatQuotePayment)
	route.DELETE("/:id", deleteSalesFlatQuotePayment)
}

func getSalesFlatQuotePayments(c *gin.Context) {
	var salesFlatQuotePayments []SalesFlatQuotePayment
	if err := g.Find(&salesFlatQuotePayments).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuotePayments)
	}
}

func getSalesFlatQuotePayment(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuotePayment SalesFlatQuotePayment
	if err := g.Where("id = ?", id).First(&salesFlatQuotePayment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuotePayment)
	}
}

func createSalesFlatQuotePayment(c *gin.Context) {
	var salesFlatQuotePayment SalesFlatQuotePayment
	c.BindJSON(&salesFlatQuotePayment)
	g.Create(&salesFlatQuotePayment)
	c.JSON(200, salesFlatQuotePayment)
}

func updateSalesFlatQuotePayment(c *gin.Context) {
	var salesFlatQuotePayment SalesFlatQuotePayment
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatQuotePayment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatQuotePayment)
	g.Save(&salesFlatQuotePayment)
	c.JSON(200, salesFlatQuotePayment)
}
func deleteSalesFlatQuotePayment(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuotePayment SalesFlatQuotePayment
	d := g.Where("id = ?", id).Delete(&salesFlatQuotePayment)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
