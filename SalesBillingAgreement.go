package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesBillingAgreement :
type SalesBillingAgreement struct {
	AgreementId    uint16     `gorm:"column:agreement_id;primary_key" form:"agreement_id;primary_key" json:"agreement_id;primary_key"`
	CustomerId     uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	MethodCode     string     `gorm:"column:method_code" form:"method_code" json:"method_code"`
	ReferenceId    string     `gorm:"column:reference_id" form:"reference_id" json:"reference_id"`
	Status         string     `gorm:"column:status" form:"status" json:"status"`
	CreatedAt      *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	StoreId        uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	AgreementLabel string     `gorm:"column:agreement_label" form:"agreement_label" json:"agreement_label"`
}

// TableName :
func (*SalesBillingAgreement) TableName() string {
	return "sales_billing_agreement"
}

// handler create
func initRoutersSalesBillingAgreement(r *gin.Engine, salesbillingagreement string) {
	route := r.Group("salesbillingagreement")
	route.GET("/", getSalesBillingAgreements)
	route.GET("/:id", getSalesBillingAgreement)
	route.POST("/", createSalesBillingAgreement)
	route.PUT("/:id", updateSalesBillingAgreement)
	route.DELETE("/:id", deleteSalesBillingAgreement)
}

func getSalesBillingAgreements(c *gin.Context) {
	var salesBillingAgreements []SalesBillingAgreement
	if err := g.Find(&salesBillingAgreements).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesBillingAgreements)
	}
}

func getSalesBillingAgreement(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesBillingAgreement SalesBillingAgreement
	if err := g.Where("id = ?", id).First(&salesBillingAgreement).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesBillingAgreement)
	}
}

func createSalesBillingAgreement(c *gin.Context) {
	var salesBillingAgreement SalesBillingAgreement
	c.BindJSON(&salesBillingAgreement)
	g.Create(&salesBillingAgreement)
	c.JSON(200, salesBillingAgreement)
}

func updateSalesBillingAgreement(c *gin.Context) {
	var salesBillingAgreement SalesBillingAgreement
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesBillingAgreement).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesBillingAgreement)
	g.Save(&salesBillingAgreement)
	c.JSON(200, salesBillingAgreement)
}
func deleteSalesBillingAgreement(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesBillingAgreement SalesBillingAgreement
	d := g.Where("id = ?", id).Delete(&salesBillingAgreement)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
