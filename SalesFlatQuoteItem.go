package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatQuoteItem :
type SalesFlatQuoteItem struct {
	ItemId                    uint16     `gorm:"column:item_id;primary_key" form:"item_id;primary_key" json:"item_id;primary_key"`
	QuoteId                   uint16     `gorm:"column:quote_id" form:"quote_id" json:"quote_id"`
	CreatedAt                 *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt                 *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	ProductId                 uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	StoreId                   uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	ParentItemId              uint16     `gorm:"column:parent_item_id" form:"parent_item_id" json:"parent_item_id"`
	IsVirtual                 uint16     `gorm:"column:is_virtual" form:"is_virtual" json:"is_virtual"`
	Sku                       string     `gorm:"column:sku" form:"sku" json:"sku"`
	Name                      string     `gorm:"column:name" form:"name" json:"name"`
	Description               string     `gorm:"column:description" form:"description" json:"description"`
	AppliedRuleIds            string     `gorm:"column:applied_rule_ids" form:"applied_rule_ids" json:"applied_rule_ids"`
	AdditionalData            string     `gorm:"column:additional_data" form:"additional_data" json:"additional_data"`
	FreeShipping              uint16     `gorm:"column:free_shipping" form:"free_shipping" json:"free_shipping"`
	IsQtyDecimal              uint16     `gorm:"column:is_qty_decimal" form:"is_qty_decimal" json:"is_qty_decimal"`
	NoDiscount                uint16     `gorm:"column:no_discount" form:"no_discount" json:"no_discount"`
	Weight                    float64    `gorm:"column:weight" form:"weight" json:"weight"`
	Qty                       float64    `gorm:"column:qty" form:"qty" json:"qty"`
	Price                     float64    `gorm:"column:price" form:"price" json:"price"`
	BasePrice                 float64    `gorm:"column:base_price" form:"base_price" json:"base_price"`
	CustomPrice               float64    `gorm:"column:custom_price" form:"custom_price" json:"custom_price"`
	DiscountPercent           float64    `gorm:"column:discount_percent" form:"discount_percent" json:"discount_percent"`
	DiscountAmount            float64    `gorm:"column:discount_amount" form:"discount_amount" json:"discount_amount"`
	BaseDiscountAmount        float64    `gorm:"column:base_discount_amount" form:"base_discount_amount" json:"base_discount_amount"`
	TaxPercent                float64    `gorm:"column:tax_percent" form:"tax_percent" json:"tax_percent"`
	TaxAmount                 float64    `gorm:"column:tax_amount" form:"tax_amount" json:"tax_amount"`
	BaseTaxAmount             float64    `gorm:"column:base_tax_amount" form:"base_tax_amount" json:"base_tax_amount"`
	RowTotal                  float64    `gorm:"column:row_total" form:"row_total" json:"row_total"`
	BaseRowTotal              float64    `gorm:"column:base_row_total" form:"base_row_total" json:"base_row_total"`
	RowTotalWithDiscount      float64    `gorm:"column:row_total_with_discount" form:"row_total_with_discount" json:"row_total_with_discount"`
	RowWeight                 float64    `gorm:"column:row_weight" form:"row_weight" json:"row_weight"`
	ProductType               string     `gorm:"column:product_type" form:"product_type" json:"product_type"`
	BaseTaxBeforeDiscount     float64    `gorm:"column:base_tax_before_discount" form:"base_tax_before_discount" json:"base_tax_before_discount"`
	TaxBeforeDiscount         float64    `gorm:"column:tax_before_discount" form:"tax_before_discount" json:"tax_before_discount"`
	OriginalCustomPrice       float64    `gorm:"column:original_custom_price" form:"original_custom_price" json:"original_custom_price"`
	RedirectUrl               string     `gorm:"column:redirect_url" form:"redirect_url" json:"redirect_url"`
	BaseCost                  float64    `gorm:"column:base_cost" form:"base_cost" json:"base_cost"`
	PriceInclTax              float64    `gorm:"column:price_incl_tax" form:"price_incl_tax" json:"price_incl_tax"`
	BasePriceInclTax          float64    `gorm:"column:base_price_incl_tax" form:"base_price_incl_tax" json:"base_price_incl_tax"`
	RowTotalInclTax           float64    `gorm:"column:row_total_incl_tax" form:"row_total_incl_tax" json:"row_total_incl_tax"`
	BaseRowTotalInclTax       float64    `gorm:"column:base_row_total_incl_tax" form:"base_row_total_incl_tax" json:"base_row_total_incl_tax"`
	HiddenTaxAmount           float64    `gorm:"column:hidden_tax_amount" form:"hidden_tax_amount" json:"hidden_tax_amount"`
	BaseHiddenTaxAmount       float64    `gorm:"column:base_hidden_tax_amount" form:"base_hidden_tax_amount" json:"base_hidden_tax_amount"`
	GiftMessageId             uint16     `gorm:"column:gift_message_id" form:"gift_message_id" json:"gift_message_id"`
	WeeeTaxDisposition        float64    `gorm:"column:weee_tax_disposition" form:"weee_tax_disposition" json:"weee_tax_disposition"`
	WeeeTaxRowDisposition     float64    `gorm:"column:weee_tax_row_disposition" form:"weee_tax_row_disposition" json:"weee_tax_row_disposition"`
	BaseWeeeTaxDisposition    float64    `gorm:"column:base_weee_tax_disposition" form:"base_weee_tax_disposition" json:"base_weee_tax_disposition"`
	BaseWeeeTaxRowDisposition float64    `gorm:"column:base_weee_tax_row_disposition" form:"base_weee_tax_row_disposition" json:"base_weee_tax_row_disposition"`
	WeeeTaxApplied            string     `gorm:"column:weee_tax_applied" form:"weee_tax_applied" json:"weee_tax_applied"`
	WeeeTaxAppliedAmount      float64    `gorm:"column:weee_tax_applied_amount" form:"weee_tax_applied_amount" json:"weee_tax_applied_amount"`
	WeeeTaxAppliedRowAmount   float64    `gorm:"column:weee_tax_applied_row_amount" form:"weee_tax_applied_row_amount" json:"weee_tax_applied_row_amount"`
	BaseWeeeTaxAppliedAmount  float64    `gorm:"column:base_weee_tax_applied_amount" form:"base_weee_tax_applied_amount" json:"base_weee_tax_applied_amount"`
	BaseWeeeTaxAppliedRowAmnt float64    `gorm:"column:base_weee_tax_applied_row_amnt" form:"base_weee_tax_applied_row_amnt" json:"base_weee_tax_applied_row_amnt"`
}

// TableName :
func (*SalesFlatQuoteItem) TableName() string {
	return "sales_flat_quote_item"
}

// handler create
func initRoutersSalesFlatQuoteItem(r *gin.Engine, salesflatquoteitem string) {
	route := r.Group("salesflatquoteitem")
	route.GET("/", getSalesFlatQuoteItems)
	route.GET("/:id", getSalesFlatQuoteItem)
	route.POST("/", createSalesFlatQuoteItem)
	route.PUT("/:id", updateSalesFlatQuoteItem)
	route.DELETE("/:id", deleteSalesFlatQuoteItem)
}

func getSalesFlatQuoteItems(c *gin.Context) {
	var salesFlatQuoteItems []SalesFlatQuoteItem
	if err := g.Find(&salesFlatQuoteItems).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuoteItems)
	}
}

func getSalesFlatQuoteItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuoteItem SalesFlatQuoteItem
	if err := g.Where("id = ?", id).First(&salesFlatQuoteItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatQuoteItem)
	}
}

func createSalesFlatQuoteItem(c *gin.Context) {
	var salesFlatQuoteItem SalesFlatQuoteItem
	c.BindJSON(&salesFlatQuoteItem)
	g.Create(&salesFlatQuoteItem)
	c.JSON(200, salesFlatQuoteItem)
}

func updateSalesFlatQuoteItem(c *gin.Context) {
	var salesFlatQuoteItem SalesFlatQuoteItem
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatQuoteItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatQuoteItem)
	g.Save(&salesFlatQuoteItem)
	c.JSON(200, salesFlatQuoteItem)
}
func deleteSalesFlatQuoteItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatQuoteItem SalesFlatQuoteItem
	d := g.Where("id = ?", id).Delete(&salesFlatQuoteItem)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
