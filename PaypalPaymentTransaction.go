package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// PaypalPaymentTransaction :
type PaypalPaymentTransaction struct {
	TransactionId         uint16     `gorm:"column:transaction_id;primary_key" form:"transaction_id;primary_key" json:"transaction_id;primary_key"`
	TxnId                 string     `gorm:"column:txn_id" form:"txn_id" json:"txn_id"`
	AdditionalInformation []byte     `gorm:"column:additional_information" form:"additional_information" json:"additional_information"`
	CreatedAt             *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
}

// TableName :
func (*PaypalPaymentTransaction) TableName() string {
	return "paypal_payment_transaction"
}

// handler create
func initRoutersPaypalPaymentTransaction(r *gin.Engine, paypalpaymenttransaction string) {
	route := r.Group("paypalpaymenttransaction")
	route.GET("/", getPaypalPaymentTransactions)
	route.GET("/:id", getPaypalPaymentTransaction)
	route.POST("/", createPaypalPaymentTransaction)
	route.PUT("/:id", updatePaypalPaymentTransaction)
	route.DELETE("/:id", deletePaypalPaymentTransaction)
}

func getPaypalPaymentTransactions(c *gin.Context) {
	var paypalPaymentTransactions []PaypalPaymentTransaction
	if err := g.Find(&paypalPaymentTransactions).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, paypalPaymentTransactions)
	}
}

func getPaypalPaymentTransaction(c *gin.Context) {
	id := c.Params.ByName("id")
	var paypalPaymentTransaction PaypalPaymentTransaction
	if err := g.Where("id = ?", id).First(&paypalPaymentTransaction).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, paypalPaymentTransaction)
	}
}

func createPaypalPaymentTransaction(c *gin.Context) {
	var paypalPaymentTransaction PaypalPaymentTransaction
	c.BindJSON(&paypalPaymentTransaction)
	g.Create(&paypalPaymentTransaction)
	c.JSON(200, paypalPaymentTransaction)
}

func updatePaypalPaymentTransaction(c *gin.Context) {
	var paypalPaymentTransaction PaypalPaymentTransaction
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&paypalPaymentTransaction).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&paypalPaymentTransaction)
	g.Save(&paypalPaymentTransaction)
	c.JSON(200, paypalPaymentTransaction)
}
func deletePaypalPaymentTransaction(c *gin.Context) {
	id := c.Params.ByName("id")
	var paypalPaymentTransaction PaypalPaymentTransaction
	d := g.Where("id = ?", id).Delete(&paypalPaymentTransaction)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
