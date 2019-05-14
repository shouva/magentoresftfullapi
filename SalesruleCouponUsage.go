package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesruleCouponUsage :
type SalesruleCouponUsage struct {
	CouponId   uint16 `gorm:"column:coupon_id;primary_key" form:"coupon_id;primary_key" json:"coupon_id;primary_key"`
	CustomerId uint16 `gorm:"column:customer_id;primary_key" form:"customer_id;primary_key" json:"customer_id;primary_key"`
	TimesUsed  uint16 `gorm:"column:times_used" form:"times_used" json:"times_used"`
}

// TableName :
func (*SalesruleCouponUsage) TableName() string {
	return "salesrule_coupon_usage"
}

// handler create
func initRoutersSalesruleCouponUsage(r *gin.Engine, salesrulecouponusage string) {
	route := r.Group("salesrulecouponusage")
	route.GET("/", getSalesruleCouponUsages)
	route.GET("/:id", getSalesruleCouponUsage)
	route.POST("/", createSalesruleCouponUsage)
	route.PUT("/:id", updateSalesruleCouponUsage)
	route.DELETE("/:id", deleteSalesruleCouponUsage)
}

func getSalesruleCouponUsages(c *gin.Context) {
	var salesruleCouponUsages []SalesruleCouponUsage
	if err := g.Find(&salesruleCouponUsages).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleCouponUsages)
	}
}

func getSalesruleCouponUsage(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleCouponUsage SalesruleCouponUsage
	if err := g.Where("id = ?", id).First(&salesruleCouponUsage).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleCouponUsage)
	}
}

func createSalesruleCouponUsage(c *gin.Context) {
	var salesruleCouponUsage SalesruleCouponUsage
	c.BindJSON(&salesruleCouponUsage)
	g.Create(&salesruleCouponUsage)
	c.JSON(200, salesruleCouponUsage)
}

func updateSalesruleCouponUsage(c *gin.Context) {
	var salesruleCouponUsage SalesruleCouponUsage
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesruleCouponUsage).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesruleCouponUsage)
	g.Save(&salesruleCouponUsage)
	c.JSON(200, salesruleCouponUsage)
}
func deleteSalesruleCouponUsage(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleCouponUsage SalesruleCouponUsage
	d := g.Where("id = ?", id).Delete(&salesruleCouponUsage)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
