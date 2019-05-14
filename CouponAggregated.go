package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CouponAggregated :
type CouponAggregated struct {
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
func (*CouponAggregated) TableName() string {
	return "coupon_aggregated"
}

// handler create
func initRoutersCouponAggregated(r *gin.Engine, couponaggregated string) {
	route := r.Group("couponaggregated")
	route.GET("/", getCouponAggregateds)
	route.GET("/:id", getCouponAggregated)
	route.POST("/", createCouponAggregated)
	route.PUT("/:id", updateCouponAggregated)
	route.DELETE("/:id", deleteCouponAggregated)
}

func getCouponAggregateds(c *gin.Context) {
	var couponAggregateds []CouponAggregated
	if err := g.Find(&couponAggregateds).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, couponAggregateds)
	}
}

func getCouponAggregated(c *gin.Context) {
	id := c.Params.ByName("id")
	var couponAggregated CouponAggregated
	if err := g.Where("id = ?", id).First(&couponAggregated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, couponAggregated)
	}
}

func createCouponAggregated(c *gin.Context) {
	var couponAggregated CouponAggregated
	c.BindJSON(&couponAggregated)
	g.Create(&couponAggregated)
	c.JSON(200, couponAggregated)
}

func updateCouponAggregated(c *gin.Context) {
	var couponAggregated CouponAggregated
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&couponAggregated).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&couponAggregated)
	g.Save(&couponAggregated)
	c.JSON(200, couponAggregated)
}
func deleteCouponAggregated(c *gin.Context) {
	id := c.Params.ByName("id")
	var couponAggregated CouponAggregated
	d := g.Where("id = ?", id).Delete(&couponAggregated)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
