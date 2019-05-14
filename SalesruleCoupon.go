package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesruleCoupon :
type SalesruleCoupon struct {
	CouponId         uint16     `gorm:"column:coupon_id;primary_key" form:"coupon_id;primary_key" json:"coupon_id;primary_key"`
	RuleId           uint16     `gorm:"column:rule_id" form:"rule_id" json:"rule_id"`
	Code             string     `gorm:"column:code" form:"code" json:"code"`
	UsageLimit       uint16     `gorm:"column:usage_limit" form:"usage_limit" json:"usage_limit"`
	UsagePerCustomer uint16     `gorm:"column:usage_per_customer" form:"usage_per_customer" json:"usage_per_customer"`
	TimesUsed        uint16     `gorm:"column:times_used" form:"times_used" json:"times_used"`
	ExpirationDate   *time.Time `gorm:"column:expiration_date" form:"expiration_date" json:"expiration_date"`
	IsPrimary        uint16     `gorm:"column:is_primary" form:"is_primary" json:"is_primary"`
	CreatedAt        *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	Type             uint16     `gorm:"column:type" form:"type" json:"type"`
}

// TableName :
func (*SalesruleCoupon) TableName() string {
	return "salesrule_coupon"
}

// handler create
func initRoutersSalesruleCoupon(r *gin.Engine, salesrulecoupon string) {
	route := r.Group("salesrulecoupon")
	route.GET("/", getSalesruleCoupons)
	route.GET("/:id", getSalesruleCoupon)
	route.POST("/", createSalesruleCoupon)
	route.PUT("/:id", updateSalesruleCoupon)
	route.DELETE("/:id", deleteSalesruleCoupon)
}

func getSalesruleCoupons(c *gin.Context) {
	var salesruleCoupons []SalesruleCoupon
	if err := g.Find(&salesruleCoupons).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleCoupons)
	}
}

func getSalesruleCoupon(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleCoupon SalesruleCoupon
	if err := g.Where("id = ?", id).First(&salesruleCoupon).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleCoupon)
	}
}

func createSalesruleCoupon(c *gin.Context) {
	var salesruleCoupon SalesruleCoupon
	c.BindJSON(&salesruleCoupon)
	g.Create(&salesruleCoupon)
	c.JSON(200, salesruleCoupon)
}

func updateSalesruleCoupon(c *gin.Context) {
	var salesruleCoupon SalesruleCoupon
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesruleCoupon).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesruleCoupon)
	g.Save(&salesruleCoupon)
	c.JSON(200, salesruleCoupon)
}
func deleteSalesruleCoupon(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleCoupon SalesruleCoupon
	d := g.Where("id = ?", id).Delete(&salesruleCoupon)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
