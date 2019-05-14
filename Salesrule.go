package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Salesrule :
type Salesrule struct {
	RuleId               uint16     `gorm:"column:rule_id;primary_key" form:"rule_id;primary_key" json:"rule_id;primary_key"`
	Name                 string     `gorm:"column:name" form:"name" json:"name"`
	Description          string     `gorm:"column:description" form:"description" json:"description"`
	FromDate             *time.Time `gorm:"column:from_date" form:"from_date" json:"from_date"`
	ToDate               *time.Time `gorm:"column:to_date" form:"to_date" json:"to_date"`
	UsesPerCustomer      uint16     `gorm:"column:uses_per_customer" form:"uses_per_customer" json:"uses_per_customer"`
	IsActive             uint16     `gorm:"column:is_active" form:"is_active" json:"is_active"`
	ConditionsSerialized string     `gorm:"column:conditions_serialized" form:"conditions_serialized" json:"conditions_serialized"`
	ActionsSerialized    string     `gorm:"column:actions_serialized" form:"actions_serialized" json:"actions_serialized"`
	StopRulesProcessing  uint16     `gorm:"column:stop_rules_processing" form:"stop_rules_processing" json:"stop_rules_processing"`
	IsAdvanced           uint16     `gorm:"column:is_advanced" form:"is_advanced" json:"is_advanced"`
	ProductIds           string     `gorm:"column:product_ids" form:"product_ids" json:"product_ids"`
	SortOrder            uint16     `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
	SimpleAction         string     `gorm:"column:simple_action" form:"simple_action" json:"simple_action"`
	DiscountAmount       float64    `gorm:"column:discount_amount" form:"discount_amount" json:"discount_amount"`
	DiscountQty          float64    `gorm:"column:discount_qty" form:"discount_qty" json:"discount_qty"`
	DiscountStep         uint16     `gorm:"column:discount_step" form:"discount_step" json:"discount_step"`
	SimpleFreeShipping   uint16     `gorm:"column:simple_free_shipping" form:"simple_free_shipping" json:"simple_free_shipping"`
	ApplyToShipping      uint16     `gorm:"column:apply_to_shipping" form:"apply_to_shipping" json:"apply_to_shipping"`
	TimesUsed            uint16     `gorm:"column:times_used" form:"times_used" json:"times_used"`
	IsRss                uint16     `gorm:"column:is_rss" form:"is_rss" json:"is_rss"`
	CouponType           uint16     `gorm:"column:coupon_type" form:"coupon_type" json:"coupon_type"`
	UseAutoGeneration    uint16     `gorm:"column:use_auto_generation" form:"use_auto_generation" json:"use_auto_generation"`
	UsesPerCoupon        uint16     `gorm:"column:uses_per_coupon" form:"uses_per_coupon" json:"uses_per_coupon"`
}

// TableName :
func (*Salesrule) TableName() string {
	return "salesrule"
}

// handler create
func initRoutersSalesrule(r *gin.Engine, salesrule string) {
	route := r.Group("salesrule")
	route.GET("/", getSalesrules)
	route.GET("/:id", getSalesrule)
	route.POST("/", createSalesrule)
	route.PUT("/:id", updateSalesrule)
	route.DELETE("/:id", deleteSalesrule)
}

func getSalesrules(c *gin.Context) {
	var salesrules []Salesrule
	if err := g.Find(&salesrules).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesrules)
	}
}

func getSalesrule(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesrule Salesrule
	if err := g.Where("id = ?", id).First(&salesrule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesrule)
	}
}

func createSalesrule(c *gin.Context) {
	var salesrule Salesrule
	c.BindJSON(&salesrule)
	g.Create(&salesrule)
	c.JSON(200, salesrule)
}

func updateSalesrule(c *gin.Context) {
	var salesrule Salesrule
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesrule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesrule)
	g.Save(&salesrule)
	c.JSON(200, salesrule)
}
func deleteSalesrule(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesrule Salesrule
	d := g.Where("id = ?", id).Delete(&salesrule)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
