package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesBillingAgreementOrder :
type SalesBillingAgreementOrder struct {
	AgreementId uint16 `gorm:"column:agreement_id;primary_key" form:"agreement_id;primary_key" json:"agreement_id;primary_key"`
	OrderId     uint16 `gorm:"column:order_id;primary_key" form:"order_id;primary_key" json:"order_id;primary_key"`
}

// TableName :
func (*SalesBillingAgreementOrder) TableName() string {
	return "sales_billing_agreement_order"
}

// handler create
func initRoutersSalesBillingAgreementOrder(r *gin.Engine, salesbillingagreementorder string) {
	route := r.Group("salesbillingagreementorder")
	route.GET("/", getSalesBillingAgreementOrders)
	route.GET("/:id", getSalesBillingAgreementOrder)
	route.POST("/", createSalesBillingAgreementOrder)
	route.PUT("/:id", updateSalesBillingAgreementOrder)
	route.DELETE("/:id", deleteSalesBillingAgreementOrder)
}

func getSalesBillingAgreementOrders(c *gin.Context) {
	var salesBillingAgreementOrders []SalesBillingAgreementOrder
	if err := g.Find(&salesBillingAgreementOrders).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesBillingAgreementOrders)
	}
}

func getSalesBillingAgreementOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesBillingAgreementOrder SalesBillingAgreementOrder
	if err := g.Where("id = ?", id).First(&salesBillingAgreementOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesBillingAgreementOrder)
	}
}

func createSalesBillingAgreementOrder(c *gin.Context) {
	var salesBillingAgreementOrder SalesBillingAgreementOrder
	c.BindJSON(&salesBillingAgreementOrder)
	g.Create(&salesBillingAgreementOrder)
	c.JSON(200, salesBillingAgreementOrder)
}

func updateSalesBillingAgreementOrder(c *gin.Context) {
	var salesBillingAgreementOrder SalesBillingAgreementOrder
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesBillingAgreementOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesBillingAgreementOrder)
	g.Save(&salesBillingAgreementOrder)
	c.JSON(200, salesBillingAgreementOrder)
}
func deleteSalesBillingAgreementOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesBillingAgreementOrder SalesBillingAgreementOrder
	d := g.Where("id = ?", id).Delete(&salesBillingAgreementOrder)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
