package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatQuoteAddressItem :
type SalesFlatQuoteAddressItem struct {
	AddressItemId        uint16     `gorm:"column:address_item_id;primary_key" form:"address_item_id;primary_key" json:"address_item_id;primary_key"`
	ParentItemId         uint16     `gorm:"column:parent_item_id" form:"parent_item_id" json:"parent_item_id"`
	QuoteAddressId       uint16     `gorm:"column:quote_address_id" form:"quote_address_id" json:"quote_address_id"`
	QuoteItemId          uint16     `gorm:"column:quote_item_id" form:"quote_item_id" json:"quote_item_id"`
	CreatedAt            *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt            *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	AppliedRuleIds       string     `gorm:"column:applied_rule_ids" form:"applied_rule_ids" json:"applied_rule_ids"`
	AdditionalData       string     `gorm:"column:additional_data" form:"additional_data" json:"additional_data"`
	Weight               float64    `gorm:"column:weight" form:"weight" json:"weight"`
	Qty                  float64    `gorm:"column:qty" form:"qty" json:"qty"`
	DiscountAmount       float64    `gorm:"column:discount_amount" form:"discount_amount" json:"discount_amount"`
	TaxAmount            float64    `gorm:"column:tax_amount" form:"tax_amount" json:"tax_amount"`
	RowTotal             float64    `gorm:"column:row_total" form:"row_total" json:"row_total"`
	BaseRowTotal         float64    `gorm:"column:base_row_total" form:"base_row_total" json:"base_row_total"`
	RowTotalWithDiscount float64    `gorm:"column:row_total_with_discount" form:"row_total_with_discount" json:"row_total_with_discount"`
	BaseDiscountAmount   float64    `gorm:"column:base_discount_amount" form:"base_discount_amount" json:"base_discount_amount"`
	BaseTaxAmount        float64    `gorm:"column:base_tax_amount" form:"base_tax_amount" json:"base_tax_amount"`
	RowWeight            float64    `gorm:"column:row_weight" form:"row_weight" json:"row_weight"`
	ProductId            uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	SuperProductId       uint16     `gorm:"column:super_product_id" form:"super_product_id" json:"super_product_id"`
	ParentProductId      uint16     `gorm:"column:parent_product_id" form:"parent_product_id" json:"parent_product_id"`
	Sku                  string     `gorm:"column:sku" form:"sku" json:"sku"`
	Image                string     `gorm:"column:image" form:"image" json:"image"`
	Name                 string     `gorm:"column:name" form:"name" json:"name"`
	Description          string     `gorm:"column:description" form:"description" json:"description"`
	FreeShipping         uint16     `gorm:"column:free_shipping" form:"free_shipping" json:"free_shipping"`
	IsQtyDecimal         uint16     `gorm:"column:is_qty_decimal" form:"is_qty_decimal" json:"is_qty_decimal"`
	Price                float64    `gorm:"column:price" form:"price" json:"price"`
	DiscountPercent      float64    `gorm:"column:discount_percent" form:"discount_percent" json:"discount_percent"`
	NoDiscount           uint16     `gorm:"column:no_discount" form:"no_discount" json:"no_discount"`
	TaxPercent           float64    `gorm:"column:tax_percent" form:"tax_percent" json:"tax_percent"`
	BasePrice            float64    `gorm:"column:base_price" form:"base_price" json:"base_price"`
	BaseCost             float64    `gorm:"column:base_cost" form:"base_cost" json:"base_cost"`
	PriceInclTax         float64    `gorm:"column:price_incl_tax" form:"price_incl_tax" json:"price_incl_tax"`
	BasePriceInclTax     float64    `gorm:"column:base_price_incl_tax" form:"base_price_incl_tax" json:"base_price_incl_tax"`
	RowTotalInclTax      float64    `gorm:"column:row_total_incl_tax" form:"row_total_incl_tax" json:"row_total_incl_tax"`
	BaseRowTotalInclTax  float64    `gorm:"column:base_row_total_incl_tax" form:"base_row_total_incl_tax" json:"base_row_total_incl_tax"`
	HiddenTaxAmount      float64    `gorm:"column:hidden_tax_amount" form:"hidden_tax_amount" json:"hidden_tax_amount"`
	BaseHiddenTaxAmount  float64    `gorm:"column:base_hidden_tax_amount" form:"base_hidden_tax_amount" json:"base_hidden_tax_amount"`
	GiftMessageId        uint16     `gorm:"column:gift_message_id" form:"gift_message_id" json:"gift_message_id"`
}

// TableName :
func (*SalesFlatQuoteAddressItem) TableName() string {
	return "sales_flat_quote_address_item"
}

// handler create
func initRoutersSalesFlatQuoteAddressItem(r *gin.Engine, salesflatquoteaddressitem string) {
	route := r.Group("salesflatquoteaddressitem")
	route.GET("/", getSalesFlatQuoteAddressItems)
	route.GET("/:id", getSalesFlatQuoteAddressItem)
	route.POST("/", createSalesFlatQuoteAddressItem)
	route.PUT("/:id", updateSalesFlatQuoteAddressItem)
	route.DELETE("/:id", deleteSalesFlatQuoteAddressItem)
}

func getSalesFlatQuoteAddressItems(c *gin.Context) {
	var salesFlatQuoteAddressItems []SalesFlatQuoteAddressItem
	if err := g.Find(&salesFlatQuoteAddressItems).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuoteAddressItems)
	}
}

func getSalesFlatQuoteAddressItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuoteAddressItem SalesFlatQuoteAddressItem
	if err := g.Where("id = ?", id).First(&salesFlatQuoteAddressItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuoteAddressItem)
	}
}

func createSalesFlatQuoteAddressItem(c *gin.Context) {
	var salesFlatQuoteAddressItem SalesFlatQuoteAddressItem
	c.BindJSON(&salesFlatQuoteAddressItem)
	g.Create(&salesFlatQuoteAddressItem)
	c.JSON(200, salesFlatQuoteAddressItem)
}

func updateSalesFlatQuoteAddressItem(c *gin.Context) {
	var salesFlatQuoteAddressItem SalesFlatQuoteAddressItem
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatQuoteAddressItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatQuoteAddressItem)
	g.Save(&salesFlatQuoteAddressItem)
	c.JSON(200, salesFlatQuoteAddressItem)
}
func deleteSalesFlatQuoteAddressItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuoteAddressItem SalesFlatQuoteAddressItem
	d := g.Where("id = ?", id).Delete(&salesFlatQuoteAddressItem)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
