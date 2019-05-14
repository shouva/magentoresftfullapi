package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesruleProductAttribute :
type SalesruleProductAttribute struct {
	RuleId          uint16 `gorm:"column:rule_id;primary_key" form:"rule_id;primary_key" json:"rule_id;primary_key"`
	WebsiteId       uint16 `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	CustomerGroupId uint16 `gorm:"column:customer_group_id;primary_key" form:"customer_group_id;primary_key" json:"customer_group_id;primary_key"`
	AttributeId     uint16 `gorm:"column:attribute_id;primary_key" form:"attribute_id;primary_key" json:"attribute_id;primary_key"`
}

// TableName :
func (*SalesruleProductAttribute) TableName() string {
	return "salesrule_product_attribute"
}

// handler create
func initRoutersSalesruleProductAttribute(r *gin.Engine, salesruleproductattribute string) {
	route := r.Group("salesruleproductattribute")
	route.GET("/", getSalesruleProductAttributes)
	route.GET("/:id", getSalesruleProductAttribute)
	route.POST("/", createSalesruleProductAttribute)
	route.PUT("/:id", updateSalesruleProductAttribute)
	route.DELETE("/:id", deleteSalesruleProductAttribute)
}

func getSalesruleProductAttributes(c *gin.Context) {
	var salesruleProductAttributes []SalesruleProductAttribute
	if err := g.Find(&salesruleProductAttributes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleProductAttributes)
	}
}

func getSalesruleProductAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleProductAttribute SalesruleProductAttribute
	if err := g.Where("id = ?", id).First(&salesruleProductAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleProductAttribute)
	}
}

func createSalesruleProductAttribute(c *gin.Context) {
	var salesruleProductAttribute SalesruleProductAttribute
	c.BindJSON(&salesruleProductAttribute)
	g.Create(&salesruleProductAttribute)
	c.JSON(200, salesruleProductAttribute)
}

func updateSalesruleProductAttribute(c *gin.Context) {
	var salesruleProductAttribute SalesruleProductAttribute
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesruleProductAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesruleProductAttribute)
	g.Save(&salesruleProductAttribute)
	c.JSON(200, salesruleProductAttribute)
}
func deleteSalesruleProductAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleProductAttribute SalesruleProductAttribute
	d := g.Where("id = ?", id).Delete(&salesruleProductAttribute)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
