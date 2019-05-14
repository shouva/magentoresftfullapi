package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesRecurringProfile :
type SalesRecurringProfile struct {
	ProfileId            uint16     `gorm:"column:profile_id;primary_key" form:"profile_id;primary_key" json:"profile_id;primary_key"`
	State                string     `gorm:"column:state" form:"state" json:"state"`
	CustomerId           uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	StoreId              uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	MethodCode           string     `gorm:"column:method_code" form:"method_code" json:"method_code"`
	CreatedAt            *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt            *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	ReferenceId          string     `gorm:"column:reference_id" form:"reference_id" json:"reference_id"`
	SubscriberName       string     `gorm:"column:subscriber_name" form:"subscriber_name" json:"subscriber_name"`
	StartDatetime        *time.Time `gorm:"column:start_datetime" form:"start_datetime" json:"start_datetime"`
	InternalReferenceId  string     `gorm:"column:internal_reference_id" form:"internal_reference_id" json:"internal_reference_id"`
	ScheduleDescription  string     `gorm:"column:schedule_description" form:"schedule_description" json:"schedule_description"`
	SuspensionThreshold  uint16     `gorm:"column:suspension_threshold" form:"suspension_threshold" json:"suspension_threshold"`
	BillFailedLater      uint16     `gorm:"column:bill_failed_later" form:"bill_failed_later" json:"bill_failed_later"`
	PeriodUnit           string     `gorm:"column:period_unit" form:"period_unit" json:"period_unit"`
	PeriodFrequency      uint16     `gorm:"column:period_frequency" form:"period_frequency" json:"period_frequency"`
	PeriodMaxCycles      uint16     `gorm:"column:period_max_cycles" form:"period_max_cycles" json:"period_max_cycles"`
	BillingAmount        float64    `gorm:"column:billing_amount" form:"billing_amount" json:"billing_amount"`
	TrialPeriodUnit      string     `gorm:"column:trial_period_unit" form:"trial_period_unit" json:"trial_period_unit"`
	TrialPeriodFrequency uint16     `gorm:"column:trial_period_frequency" form:"trial_period_frequency" json:"trial_period_frequency"`
	TrialPeriodMaxCycles uint16     `gorm:"column:trial_period_max_cycles" form:"trial_period_max_cycles" json:"trial_period_max_cycles"`
	TrialBillingAmount   string     `gorm:"column:trial_billing_amount" form:"trial_billing_amount" json:"trial_billing_amount"`
	CurrencyCode         string     `gorm:"column:currency_code" form:"currency_code" json:"currency_code"`
	ShippingAmount       float64    `gorm:"column:shipping_amount" form:"shipping_amount" json:"shipping_amount"`
	TaxAmount            float64    `gorm:"column:tax_amount" form:"tax_amount" json:"tax_amount"`
	InitAmount           float64    `gorm:"column:init_amount" form:"init_amount" json:"init_amount"`
	InitMayFail          uint16     `gorm:"column:init_may_fail" form:"init_may_fail" json:"init_may_fail"`
	OrderInfo            string     `gorm:"column:order_info" form:"order_info" json:"order_info"`
	OrderItemInfo        string     `gorm:"column:order_item_info" form:"order_item_info" json:"order_item_info"`
	BillingAddressInfo   string     `gorm:"column:billing_address_info" form:"billing_address_info" json:"billing_address_info"`
	ShippingAddressInfo  string     `gorm:"column:shipping_address_info" form:"shipping_address_info" json:"shipping_address_info"`
	ProfileVendorInfo    string     `gorm:"column:profile_vendor_info" form:"profile_vendor_info" json:"profile_vendor_info"`
	AdditionalInfo       string     `gorm:"column:additional_info" form:"additional_info" json:"additional_info"`
}

// TableName :
func (*SalesRecurringProfile) TableName() string {
	return "sales_recurring_profile"
}

// handler create
func initRoutersSalesRecurringProfile(r *gin.Engine, salesrecurringprofile string) {
	route := r.Group("salesrecurringprofile")
	route.GET("/", getSalesRecurringProfiles)
	route.GET("/:id", getSalesRecurringProfile)
	route.POST("/", createSalesRecurringProfile)
	route.PUT("/:id", updateSalesRecurringProfile)
	route.DELETE("/:id", deleteSalesRecurringProfile)
}

func getSalesRecurringProfiles(c *gin.Context) {
	var salesRecurringProfiles []SalesRecurringProfile
	if err := g.Find(&salesRecurringProfiles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesRecurringProfiles)
	}
}

func getSalesRecurringProfile(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesRecurringProfile SalesRecurringProfile
	if err := g.Where("id = ?", id).First(&salesRecurringProfile).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesRecurringProfile)
	}
}

func createSalesRecurringProfile(c *gin.Context) {
	var salesRecurringProfile SalesRecurringProfile
	c.BindJSON(&salesRecurringProfile)
	g.Create(&salesRecurringProfile)
	c.JSON(200, salesRecurringProfile)
}

func updateSalesRecurringProfile(c *gin.Context) {
	var salesRecurringProfile SalesRecurringProfile
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesRecurringProfile).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesRecurringProfile)
	g.Save(&salesRecurringProfile)
	c.JSON(200, salesRecurringProfile)
}
func deleteSalesRecurringProfile(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesRecurringProfile SalesRecurringProfile
	d := g.Where("id = ?", id).Delete(&salesRecurringProfile)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
