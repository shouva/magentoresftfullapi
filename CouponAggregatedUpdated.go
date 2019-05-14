package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CouponAggregatedUpdated :
type CouponAggregatedUpdated struct {
	ID                   uint16     `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	Period               *time.Time `gorm:"column:period" form:"period" json:"period"`
	StoreId              uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	OrderStatus          string     `gorm:"column:order_status" form:"order_status" json:"order_status"`
	CouponCode           string     `gorm:"column:coupon_code" form:"coupon_code" json:"coupon_code"`
	CouponUses           uint16     `gorm:"column:coupon_uses" form:"coupon_uses" json:"coupon_uses"`
	SubtotalAmount       float64    `gorm:"column:subtotal_amount" form:"subtotal_amount" json:"subtotal_amount"`
	DiscountAmount       float64    `gorm:"column:discount_amount" form:"discount_amount" json:"discount_amount"`
	TotalAmount          float64    `gorm:"column:total_amount" form:"total_amount" json:"total_amount"`
	SubtotalAmountActual float64    `gorm:"column:subtotal_amount_actual" form:"subtotal_amount_actual" json:"subtotal_amount_actual"`
	DiscountAmountActual float64    `gorm:"column:discount_amount_actual" form:"discount_amount_actual" json:"discount_amount_actual"`
	TotalAmountActual    float64    `gorm:"column:total_amount_actual" form:"total_amount_actual" json:"total_amount_actual"`
	RuleName             string     `gorm:"column:rule_name" form:"rule_name" json:"rule_name"`
}

// TableName :
func (*CouponAggregatedUpdated) TableName() string {
	return "coupon_aggregated_updated"
}

// handler create
func initRoutersCouponAggregatedUpdated(r *gin.Engine, couponaggregatedupdated string) {
	route := r.Group("couponaggregatedupdated")
	route.GET("/", getCouponAggregatedUpdateds)
	route.GET("/:id", getCouponAggregatedUpdated)
	route.POST("/", createCouponAggregatedUpdated)
	route.PUT("/:id", updateCouponAggregatedUpdated)
	route.DELETE("/:id", deleteCouponAggregatedUpdated)
}

func getCouponAggregatedUpdateds(c *gin.Context) {
	var couponAggregatedUpdateds []CouponAggregatedUpdated
	if err := g.Find(&couponAggregatedUpdateds).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, couponAggregatedUpdateds)
	}
}

func getCouponAggregatedUpdated(c *gin.Context) {
	id := c.Params.ByName("id")
	var couponAggregatedUpdated CouponAggregatedUpdated
	if err := g.Where("id = ?", id).First(&couponAggregatedUpdated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, couponAggregatedUpdated)
	}
}

func createCouponAggregatedUpdated(c *gin.Context) {
	var couponAggregatedUpdated CouponAggregatedUpdated
	c.BindJSON(&couponAggregatedUpdated)
	g.Create(&couponAggregatedUpdated)
	c.JSON(200, couponAggregatedUpdated)
}

func updateCouponAggregatedUpdated(c *gin.Context) {
	var couponAggregatedUpdated CouponAggregatedUpdated
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&couponAggregatedUpdated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&couponAggregatedUpdated)
	g.Save(&couponAggregatedUpdated)
	c.JSON(200, couponAggregatedUpdated)
}
func deleteCouponAggregatedUpdated(c *gin.Context) {
	id := c.Params.ByName("id")
	var couponAggregatedUpdated CouponAggregatedUpdated
	d := g.Where("id = ?", id).Delete(&couponAggregatedUpdated)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
