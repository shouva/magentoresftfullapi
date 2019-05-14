package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesruleCustomer :
type SalesruleCustomer struct {
	RuleCustomerId uint16 `gorm:"column:rule_customer_id;primary_key" form:"rule_customer_id;primary_key" json:"rule_customer_id;primary_key"`
	RuleId         uint16 `gorm:"column:rule_id" form:"rule_id" json:"rule_id"`
	CustomerId     uint16 `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	TimesUsed      uint16 `gorm:"column:times_used" form:"times_used" json:"times_used"`
}

// TableName :
func (*SalesruleCustomer) TableName() string {
	return "salesrule_customer"
}

// handler create
func initRoutersSalesruleCustomer(r *gin.Engine, salesrulecustomer string) {
	route := r.Group("salesrulecustomer")
	route.GET("/", getSalesruleCustomers)
	route.GET("/:id", getSalesruleCustomer)
	route.POST("/", createSalesruleCustomer)
	route.PUT("/:id", updateSalesruleCustomer)
	route.DELETE("/:id", deleteSalesruleCustomer)
}

func getSalesruleCustomers(c *gin.Context) {
	var salesruleCustomers []SalesruleCustomer
	if err := g.Find(&salesruleCustomers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleCustomers)
	}
}

func getSalesruleCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleCustomer SalesruleCustomer
	if err := g.Where("id = ?", id).First(&salesruleCustomer).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleCustomer)
	}
}

func createSalesruleCustomer(c *gin.Context) {
	var salesruleCustomer SalesruleCustomer
	c.BindJSON(&salesruleCustomer)
	g.Create(&salesruleCustomer)
	c.JSON(200, salesruleCustomer)
}

func updateSalesruleCustomer(c *gin.Context) {
	var salesruleCustomer SalesruleCustomer
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesruleCustomer).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesruleCustomer)
	g.Save(&salesruleCustomer)
	c.JSON(200, salesruleCustomer)
}
func deleteSalesruleCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleCustomer SalesruleCustomer
	d := g.Where("id = ?", id).Delete(&salesruleCustomer)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
