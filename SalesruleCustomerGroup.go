package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesruleCustomerGroup :
type SalesruleCustomerGroup struct {
	RuleId          uint16 `gorm:"column:rule_id;primary_key" form:"rule_id;primary_key" json:"rule_id;primary_key"`
	CustomerGroupId uint16 `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
}

// TableName :
func (*SalesruleCustomerGroup) TableName() string {
	return "salesrule_customer_group"
}

// handler create
func initRoutersSalesruleCustomerGroup(r *gin.Engine, salesrulecustomergroup string) {
	route := r.Group("salesrulecustomergroup")
	route.GET("/", getSalesruleCustomerGroups)
	route.GET("/:id", getSalesruleCustomerGroup)
	route.POST("/", createSalesruleCustomerGroup)
	route.PUT("/:id", updateSalesruleCustomerGroup)
	route.DELETE("/:id", deleteSalesruleCustomerGroup)
}

func getSalesruleCustomerGroups(c *gin.Context) {
	var salesruleCustomerGroups []SalesruleCustomerGroup
	if err := g.Find(&salesruleCustomerGroups).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleCustomerGroups)
	}
}

func getSalesruleCustomerGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleCustomerGroup SalesruleCustomerGroup
	if err := g.Where("id = ?", id).First(&salesruleCustomerGroup).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleCustomerGroup)
	}
}

func createSalesruleCustomerGroup(c *gin.Context) {
	var salesruleCustomerGroup SalesruleCustomerGroup
	c.BindJSON(&salesruleCustomerGroup)
	g.Create(&salesruleCustomerGroup)
	c.JSON(200, salesruleCustomerGroup)
}

func updateSalesruleCustomerGroup(c *gin.Context) {
	var salesruleCustomerGroup SalesruleCustomerGroup
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesruleCustomerGroup).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesruleCustomerGroup)
	g.Save(&salesruleCustomerGroup)
	c.JSON(200, salesruleCustomerGroup)
}
func deleteSalesruleCustomerGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleCustomerGroup SalesruleCustomerGroup
	d := g.Where("id = ?", id).Delete(&salesruleCustomerGroup)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
