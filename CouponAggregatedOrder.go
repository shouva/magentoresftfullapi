package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CouponAggregatedOrder :
type CouponAggregatedOrder struct {
	ID             uint16     `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	Period         *time.Time `gorm:"column:period" form:"period" json:"period"`
	StoreId        uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	OrderStatus    string     `gorm:"column:order_status" form:"order_status" json:"order_status"`
	CouponCode     string     `gorm:"column:coupon_code" form:"coupon_code" json:"coupon_code"`
	CouponUses     uint16     `gorm:"column:coupon_uses" form:"coupon_uses" json:"coupon_uses"`
	SubtotalAmount float64    `gorm:"column:subtotal_amount" form:"subtotal_amount" json:"subtotal_amount"`
	DiscountAmount float64    `gorm:"column:discount_amount" form:"discount_amount" json:"discount_amount"`
	TotalAmount    float64    `gorm:"column:total_amount" form:"total_amount" json:"total_amount"`
	RuleName       string     `gorm:"column:rule_name" form:"rule_name" json:"rule_name"`
}

// TableName :
func (*CouponAggregatedOrder) TableName() string {
	return "coupon_aggregated_order"
}

// handler create
func initRoutersCouponAggregatedOrder(r *gin.Engine, couponaggregatedorder string) {
	route := r.Group("couponaggregatedorder")
	route.GET("/", getCouponAggregatedOrders)
	route.GET("/:id", getCouponAggregatedOrder)
	route.POST("/", createCouponAggregatedOrder)
	route.PUT("/:id", updateCouponAggregatedOrder)
	route.DELETE("/:id", deleteCouponAggregatedOrder)
}

func getCouponAggregatedOrders(c *gin.Context) {
	var couponAggregatedOrders []CouponAggregatedOrder
	if err := g.Find(&couponAggregatedOrders).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, couponAggregatedOrders)
	}
}

func getCouponAggregatedOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var couponAggregatedOrder CouponAggregatedOrder
	if err := g.Where("id = ?", id).First(&couponAggregatedOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, couponAggregatedOrder)
	}
}

func createCouponAggregatedOrder(c *gin.Context) {
	var couponAggregatedOrder CouponAggregatedOrder
	c.BindJSON(&couponAggregatedOrder)
	g.Create(&couponAggregatedOrder)
	c.JSON(200, couponAggregatedOrder)
}

func updateCouponAggregatedOrder(c *gin.Context) {
	var couponAggregatedOrder CouponAggregatedOrder
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&couponAggregatedOrder).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&couponAggregatedOrder)
	g.Save(&couponAggregatedOrder)
	c.JSON(200, couponAggregatedOrder)
}
func deleteCouponAggregatedOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var couponAggregatedOrder CouponAggregatedOrder
	d := g.Where("id = ?", id).Delete(&couponAggregatedOrder)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
