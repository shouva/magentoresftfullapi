package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatQuoteAddres :
type SalesFlatQuoteAddres struct {
	AddressId                  uint16     `gorm:"column:address_id;primary_key" form:"address_id;primary_key" json:"address_id;primary_key"`
	QuoteId                    uint16     `gorm:"column:quote_id" form:"quote_id" json:"quote_id"`
	CreatedAt                  *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt                  *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	CustomerId                 uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	SaveInAddressBook          uint16     `gorm:"column:save_in_address_book" form:"save_in_address_book" json:"save_in_address_book"`
	CustomerAddressId          uint16     `gorm:"column:customer_address_id" form:"customer_address_id" json:"customer_address_id"`
	AddressType                string     `gorm:"column:address_type" form:"address_type" json:"address_type"`
	Email                      string     `gorm:"column:email" form:"email" json:"email"`
	Prefix                     string     `gorm:"column:prefix" form:"prefix" json:"prefix"`
	Firstname                  string     `gorm:"column:firstname" form:"firstname" json:"firstname"`
	Middlename                 string     `gorm:"column:middlename" form:"middlename" json:"middlename"`
	Lastname                   string     `gorm:"column:lastname" form:"lastname" json:"lastname"`
	Suffix                     string     `gorm:"column:suffix" form:"suffix" json:"suffix"`
	Company                    string     `gorm:"column:company" form:"company" json:"company"`
	Street                     string     `gorm:"column:street" form:"street" json:"street"`
	City                       string     `gorm:"column:city" form:"city" json:"city"`
	Region                     string     `gorm:"column:region" form:"region" json:"region"`
	RegionId                   uint16     `gorm:"column:region_id" form:"region_id" json:"region_id"`
	Postcode                   string     `gorm:"column:postcode" form:"postcode" json:"postcode"`
	CountryId                  string     `gorm:"column:country_id" form:"country_id" json:"country_id"`
	Telephone                  string     `gorm:"column:telephone" form:"telephone" json:"telephone"`
	Fax                        string     `gorm:"column:fax" form:"fax" json:"fax"`
	SameAsBilling              uint16     `gorm:"column:same_as_billing" form:"same_as_billing" json:"same_as_billing"`
	FreeShipping               uint16     `gorm:"column:free_shipping" form:"free_shipping" json:"free_shipping"`
	CollectShippingRates       uint16     `gorm:"column:collect_shipping_rates" form:"collect_shipping_rates" json:"collect_shipping_rates"`
	ShippingMethod             string     `gorm:"column:shipping_method" form:"shipping_method" json:"shipping_method"`
	ShippingDescription        string     `gorm:"column:shipping_description" form:"shipping_description" json:"shipping_description"`
	Weight                     float64    `gorm:"column:weight" form:"weight" json:"weight"`
	Subtotal                   float64    `gorm:"column:subtotal" form:"subtotal" json:"subtotal"`
	BaseSubtotal               float64    `gorm:"column:base_subtotal" form:"base_subtotal" json:"base_subtotal"`
	SubtotalWithDiscount       float64    `gorm:"column:subtotal_with_discount" form:"subtotal_with_discount" json:"subtotal_with_discount"`
	BaseSubtotalWithDiscount   float64    `gorm:"column:base_subtotal_with_discount" form:"base_subtotal_with_discount" json:"base_subtotal_with_discount"`
	TaxAmount                  float64    `gorm:"column:tax_amount" form:"tax_amount" json:"tax_amount"`
	BaseTaxAmount              float64    `gorm:"column:base_tax_amount" form:"base_tax_amount" json:"base_tax_amount"`
	ShippingAmount             float64    `gorm:"column:shipping_amount" form:"shipping_amount" json:"shipping_amount"`
	BaseShippingAmount         float64    `gorm:"column:base_shipping_amount" form:"base_shipping_amount" json:"base_shipping_amount"`
	ShippingTaxAmount          float64    `gorm:"column:shipping_tax_amount" form:"shipping_tax_amount" json:"shipping_tax_amount"`
	BaseShippingTaxAmount      float64    `gorm:"column:base_shipping_tax_amount" form:"base_shipping_tax_amount" json:"base_shipping_tax_amount"`
	DiscountAmount             float64    `gorm:"column:discount_amount" form:"discount_amount" json:"discount_amount"`
	BaseDiscountAmount         float64    `gorm:"column:base_discount_amount" form:"base_discount_amount" json:"base_discount_amount"`
	GrandTotal                 float64    `gorm:"column:grand_total" form:"grand_total" json:"grand_total"`
	BaseGrandTotal             float64    `gorm:"column:base_grand_total" form:"base_grand_total" json:"base_grand_total"`
	CustomerNotes              string     `gorm:"column:customer_notes" form:"customer_notes" json:"customer_notes"`
	AppliedTaxes               string     `gorm:"column:applied_taxes" form:"applied_taxes" json:"applied_taxes"`
	DiscountDescription        string     `gorm:"column:discount_description" form:"discount_description" json:"discount_description"`
	ShippingDiscountAmount     float64    `gorm:"column:shipping_discount_amount" form:"shipping_discount_amount" json:"shipping_discount_amount"`
	BaseShippingDiscountAmount float64    `gorm:"column:base_shipping_discount_amount" form:"base_shipping_discount_amount" json:"base_shipping_discount_amount"`
	SubtotalInclTax            float64    `gorm:"column:subtotal_incl_tax" form:"subtotal_incl_tax" json:"subtotal_incl_tax"`
	BaseSubtotalTotalInclTax   float64    `gorm:"column:base_subtotal_total_incl_tax" form:"base_subtotal_total_incl_tax" json:"base_subtotal_total_incl_tax"`
	HiddenTaxAmount            float64    `gorm:"column:hidden_tax_amount" form:"hidden_tax_amount" json:"hidden_tax_amount"`
	BaseHiddenTaxAmount        float64    `gorm:"column:base_hidden_tax_amount" form:"base_hidden_tax_amount" json:"base_hidden_tax_amount"`
	ShippingHiddenTaxAmount    float64    `gorm:"column:shipping_hidden_tax_amount" form:"shipping_hidden_tax_amount" json:"shipping_hidden_tax_amount"`
	BaseShippingHiddenTaxAmnt  float64    `gorm:"column:base_shipping_hidden_tax_amnt" form:"base_shipping_hidden_tax_amnt" json:"base_shipping_hidden_tax_amnt"`
	ShippingInclTax            float64    `gorm:"column:shipping_incl_tax" form:"shipping_incl_tax" json:"shipping_incl_tax"`
	BaseShippingInclTax        float64    `gorm:"column:base_shipping_incl_tax" form:"base_shipping_incl_tax" json:"base_shipping_incl_tax"`
	VatId                      string     `gorm:"column:vat_id" form:"vat_id" json:"vat_id"`
	VatIsValid                 uint16     `gorm:"column:vat_is_valid" form:"vat_is_valid" json:"vat_is_valid"`
	VatRequestId               string     `gorm:"column:vat_request_id" form:"vat_request_id" json:"vat_request_id"`
	VatRequestDate             string     `gorm:"column:vat_request_date" form:"vat_request_date" json:"vat_request_date"`
	VatRequestSuccess          uint16     `gorm:"column:vat_request_success" form:"vat_request_success" json:"vat_request_success"`
	GiftMessageId              uint16     `gorm:"column:gift_message_id" form:"gift_message_id" json:"gift_message_id"`
}

// TableName :
func (*SalesFlatQuoteAddres) TableName() string {
	return "sales_flat_quote_address"
}

// handler create
func initRoutersSalesFlatQuoteAddres(r *gin.Engine, salesflatquoteaddres string) {
	route := r.Group("salesflatquoteaddres")
	route.GET("/", getSalesFlatQuoteAddress)
	route.GET("/:id", getSalesFlatQuoteAddres)
	route.POST("/", createSalesFlatQuoteAddres)
	route.PUT("/:id", updateSalesFlatQuoteAddres)
	route.DELETE("/:id", deleteSalesFlatQuoteAddres)
}

func getSalesFlatQuoteAddress(c *gin.Context) {
	var salesFlatQuoteAddress []SalesFlatQuoteAddres
	if err := g.Find(&salesFlatQuoteAddress).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuoteAddress)
	}
}

func getSalesFlatQuoteAddres(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuoteAddres SalesFlatQuoteAddres
	if err := g.Where("id = ?", id).First(&salesFlatQuoteAddres).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuoteAddres)
	}
}

func createSalesFlatQuoteAddres(c *gin.Context) {
	var salesFlatQuoteAddres SalesFlatQuoteAddres
	c.BindJSON(&salesFlatQuoteAddres)
	g.Create(&salesFlatQuoteAddres)
	c.JSON(200, salesFlatQuoteAddres)
}

func updateSalesFlatQuoteAddres(c *gin.Context) {
	var salesFlatQuoteAddres SalesFlatQuoteAddres
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatQuoteAddres).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatQuoteAddres)
	g.Save(&salesFlatQuoteAddres)
	c.JSON(200, salesFlatQuoteAddres)
}
func deleteSalesFlatQuoteAddres(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuoteAddres SalesFlatQuoteAddres
	d := g.Where("id = ?", id).Delete(&salesFlatQuoteAddres)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
