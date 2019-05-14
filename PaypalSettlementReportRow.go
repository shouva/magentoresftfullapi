package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// PaypalSettlementReportRow :
type PaypalSettlementReportRow struct {
	RowId                     uint16     `gorm:"column:row_id;primary_key" form:"row_id;primary_key" json:"row_id;primary_key"`
	ReportId                  uint16     `gorm:"column:report_id" form:"report_id" json:"report_id"`
	TransactionId             string     `gorm:"column:transaction_id" form:"transaction_id" json:"transaction_id"`
	InvoiceId                 string     `gorm:"column:invoice_id" form:"invoice_id" json:"invoice_id"`
	PaypalReferenceId         string     `gorm:"column:paypal_reference_id" form:"paypal_reference_id" json:"paypal_reference_id"`
	PaypalReferenceIdType     string     `gorm:"column:paypal_reference_id_type" form:"paypal_reference_id_type" json:"paypal_reference_id_type"`
	TransactionEventCode      string     `gorm:"column:transaction_event_code" form:"transaction_event_code" json:"transaction_event_code"`
	TransactionInitiationDate *time.Time `gorm:"column:transaction_initiation_date" form:"transaction_initiation_date" json:"transaction_initiation_date"`
	TransactionCompletionDate *time.Time `gorm:"column:transaction_completion_date" form:"transaction_completion_date" json:"transaction_completion_date"`
	TransactionDebitOrCredit  string     `gorm:"column:transaction_debit_or_credit" form:"transaction_debit_or_credit" json:"transaction_debit_or_credit"`
	GrossTransactionAmount    float64    `gorm:"column:gross_transaction_amount" form:"gross_transaction_amount" json:"gross_transaction_amount"`
	GrossTransactionCurrency  string     `gorm:"column:gross_transaction_currency" form:"gross_transaction_currency" json:"gross_transaction_currency"`
	FeeDebitOrCredit          string     `gorm:"column:fee_debit_or_credit" form:"fee_debit_or_credit" json:"fee_debit_or_credit"`
	FeeAmount                 float64    `gorm:"column:fee_amount" form:"fee_amount" json:"fee_amount"`
	FeeCurrency               string     `gorm:"column:fee_currency" form:"fee_currency" json:"fee_currency"`
	CustomField               string     `gorm:"column:custom_field" form:"custom_field" json:"custom_field"`
	ConsumerId                string     `gorm:"column:consumer_id" form:"consumer_id" json:"consumer_id"`
	PaymentTrackingId         string     `gorm:"column:payment_tracking_id" form:"payment_tracking_id" json:"payment_tracking_id"`
	StoreId                   string     `gorm:"column:store_id" form:"store_id" json:"store_id"`
}

// TableName :
func (*PaypalSettlementReportRow) TableName() string {
	return "paypal_settlement_report_row"
}

// handler create
func initRoutersPaypalSettlementReportRow(r *gin.Engine, paypalsettlementreportrow string) {
	route := r.Group("paypalsettlementreportrow")
	route.GET("/", getPaypalSettlementReportRows)
	route.GET("/:id", getPaypalSettlementReportRow)
	route.POST("/", createPaypalSettlementReportRow)
	route.PUT("/:id", updatePaypalSettlementReportRow)
	route.DELETE("/:id", deletePaypalSettlementReportRow)
}

func getPaypalSettlementReportRows(c *gin.Context) {
	var paypalSettlementReportRows []PaypalSettlementReportRow
	if err := g.Find(&paypalSettlementReportRows).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, paypalSettlementReportRows)
	}
}

func getPaypalSettlementReportRow(c *gin.Context) {
	id := c.Params.ByName("id")
	var paypalSettlementReportRow PaypalSettlementReportRow
	if err := g.Where("id = ?", id).First(&paypalSettlementReportRow).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, paypalSettlementReportRow)
	}
}

func createPaypalSettlementReportRow(c *gin.Context) {
	var paypalSettlementReportRow PaypalSettlementReportRow
	c.BindJSON(&paypalSettlementReportRow)
	g.Create(&paypalSettlementReportRow)
	c.JSON(200, paypalSettlementReportRow)
}

func updatePaypalSettlementReportRow(c *gin.Context) {
	var paypalSettlementReportRow PaypalSettlementReportRow
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&paypalSettlementReportRow).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&paypalSettlementReportRow)
	g.Save(&paypalSettlementReportRow)
	c.JSON(200, paypalSettlementReportRow)
}
func deletePaypalSettlementReportRow(c *gin.Context) {
	id := c.Params.ByName("id")
	var paypalSettlementReportRow PaypalSettlementReportRow
	d := g.Where("id = ?", id).Delete(&paypalSettlementReportRow)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
