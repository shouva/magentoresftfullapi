package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// PaypalSettlementReport :
type PaypalSettlementReport struct {
	ReportId     uint16     `gorm:"column:report_id;primary_key" form:"report_id;primary_key" json:"report_id;primary_key"`
	ReportDate   *time.Time `gorm:"column:report_date" form:"report_date" json:"report_date"`
	AccountId    string     `gorm:"column:account_id" form:"account_id" json:"account_id"`
	Filename     string     `gorm:"column:filename" form:"filename" json:"filename"`
	LastModified *time.Time `gorm:"column:last_modified" form:"last_modified" json:"last_modified"`
}

// TableName :
func (*PaypalSettlementReport) TableName() string {
	return "paypal_settlement_report"
}

// handler create
func initRoutersPaypalSettlementReport(r *gin.Engine, paypalsettlementreport string) {
	route := r.Group("paypalsettlementreport")
	route.GET("/", getPaypalSettlementReports)
	route.GET("/:id", getPaypalSettlementReport)
	route.POST("/", createPaypalSettlementReport)
	route.PUT("/:id", updatePaypalSettlementReport)
	route.DELETE("/:id", deletePaypalSettlementReport)
}

func getPaypalSettlementReports(c *gin.Context) {
	var paypalSettlementReports []PaypalSettlementReport
	if err := g.Find(&paypalSettlementReports).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, paypalSettlementReports)
	}
}

func getPaypalSettlementReport(c *gin.Context) {
	id := c.Params.ByName("id")
	var paypalSettlementReport PaypalSettlementReport
	if err := g.Where("id = ?", id).First(&paypalSettlementReport).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, paypalSettlementReport)
	}
}

func createPaypalSettlementReport(c *gin.Context) {
	var paypalSettlementReport PaypalSettlementReport
	c.BindJSON(&paypalSettlementReport)
	g.Create(&paypalSettlementReport)
	c.JSON(200, paypalSettlementReport)
}

func updatePaypalSettlementReport(c *gin.Context) {
	var paypalSettlementReport PaypalSettlementReport
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&paypalSettlementReport).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&paypalSettlementReport)
	g.Save(&paypalSettlementReport)
	c.JSON(200, paypalSettlementReport)
}
func deletePaypalSettlementReport(c *gin.Context) {
	id := c.Params.ByName("id")
	var paypalSettlementReport PaypalSettlementReport
	d := g.Where("id = ?", id).Delete(&paypalSettlementReport)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
