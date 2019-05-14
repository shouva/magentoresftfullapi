package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesPaymentTransaction :
type SalesPaymentTransaction struct {
	TransactionId         uint16     `gorm:"column:transaction_id;primary_key" form:"transaction_id;primary_key" json:"transaction_id;primary_key"`
	ParentId              uint16     `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	OrderId               uint16     `gorm:"column:order_id" form:"order_id" json:"order_id"`
	PaymentId             uint16     `gorm:"column:payment_id" form:"payment_id" json:"payment_id"`
	TxnId                 string     `gorm:"column:txn_id" form:"txn_id" json:"txn_id"`
	ParentTxnId           string     `gorm:"column:parent_txn_id" form:"parent_txn_id" json:"parent_txn_id"`
	TxnType               string     `gorm:"column:txn_type" form:"txn_type" json:"txn_type"`
	IsClosed              uint16     `gorm:"column:is_closed" form:"is_closed" json:"is_closed"`
	AdditionalInformation []byte     `gorm:"column:additional_information" form:"additional_information" json:"additional_information"`
	CreatedAt             *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
}

// TableName :
func (*SalesPaymentTransaction) TableName() string {
	return "sales_payment_transaction"
}

// handler create
func initRoutersSalesPaymentTransaction(r *gin.Engine, salespaymenttransaction string) {
	route := r.Group("salespaymenttransaction")
	route.GET("/", getSalesPaymentTransactions)
	route.GET("/:id", getSalesPaymentTransaction)
	route.POST("/", createSalesPaymentTransaction)
	route.PUT("/:id", updateSalesPaymentTransaction)
	route.DELETE("/:id", deleteSalesPaymentTransaction)
}

func getSalesPaymentTransactions(c *gin.Context) {
	var salesPaymentTransactions []SalesPaymentTransaction
	if err := g.Find(&salesPaymentTransactions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesPaymentTransactions)
	}
}

func getSalesPaymentTransaction(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesPaymentTransaction SalesPaymentTransaction
	if err := g.Where("id = ?", id).First(&salesPaymentTransaction).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesPaymentTransaction)
	}
}

func createSalesPaymentTransaction(c *gin.Context) {
	var salesPaymentTransaction SalesPaymentTransaction
	c.BindJSON(&salesPaymentTransaction)
	g.Create(&salesPaymentTransaction)
	c.JSON(200, salesPaymentTransaction)
}

func updateSalesPaymentTransaction(c *gin.Context) {
	var salesPaymentTransaction SalesPaymentTransaction
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesPaymentTransaction).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesPaymentTransaction)
	g.Save(&salesPaymentTransaction)
	c.JSON(200, salesPaymentTransaction)
}
func deleteSalesPaymentTransaction(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesPaymentTransaction SalesPaymentTransaction
	d := g.Where("id = ?", id).Delete(&salesPaymentTransaction)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
