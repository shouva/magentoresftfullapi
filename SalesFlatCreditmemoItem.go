package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesFlatCreditmemoItem :
type SalesFlatCreditmemoItem struct {
	EntityId                  uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	ParentId                  uint16  `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	BasePrice                 float64 `gorm:"column:base_price" form:"base_price" json:"base_price"`
	TaxAmount                 float64 `gorm:"column:tax_amount" form:"tax_amount" json:"tax_amount"`
	BaseRowTotal              float64 `gorm:"column:base_row_total" form:"base_row_total" json:"base_row_total"`
	DiscountAmount            float64 `gorm:"column:discount_amount" form:"discount_amount" json:"discount_amount"`
	RowTotal                  float64 `gorm:"column:row_total" form:"row_total" json:"row_total"`
	BaseDiscountAmount        float64 `gorm:"column:base_discount_amount" form:"base_discount_amount" json:"base_discount_amount"`
	PriceInclTax              float64 `gorm:"column:price_incl_tax" form:"price_incl_tax" json:"price_incl_tax"`
	BaseTaxAmount             float64 `gorm:"column:base_tax_amount" form:"base_tax_amount" json:"base_tax_amount"`
	BasePriceInclTax          float64 `gorm:"column:base_price_incl_tax" form:"base_price_incl_tax" json:"base_price_incl_tax"`
	Qty                       float64 `gorm:"column:qty" form:"qty" json:"qty"`
	BaseCost                  float64 `gorm:"column:base_cost" form:"base_cost" json:"base_cost"`
	Price                     float64 `gorm:"column:price" form:"price" json:"price"`
	BaseRowTotalInclTax       float64 `gorm:"column:base_row_total_incl_tax" form:"base_row_total_incl_tax" json:"base_row_total_incl_tax"`
	RowTotalInclTax           float64 `gorm:"column:row_total_incl_tax" form:"row_total_incl_tax" json:"row_total_incl_tax"`
	ProductId                 uint16  `gorm:"column:product_id" form:"product_id" json:"product_id"`
	OrderItemId               uint16  `gorm:"column:order_item_id" form:"order_item_id" json:"order_item_id"`
	AdditionalData            string  `gorm:"column:additional_data" form:"additional_data" json:"additional_data"`
	Description               string  `gorm:"column:description" form:"description" json:"description"`
	Sku                       string  `gorm:"column:sku" form:"sku" json:"sku"`
	Name                      string  `gorm:"column:name" form:"name" json:"name"`
	HiddenTaxAmount           float64 `gorm:"column:hidden_tax_amount" form:"hidden_tax_amount" json:"hidden_tax_amount"`
	BaseHiddenTaxAmount       float64 `gorm:"column:base_hidden_tax_amount" form:"base_hidden_tax_amount" json:"base_hidden_tax_amount"`
	WeeeTaxDisposition        float64 `gorm:"column:weee_tax_disposition" form:"weee_tax_disposition" json:"weee_tax_disposition"`
	WeeeTaxRowDisposition     float64 `gorm:"column:weee_tax_row_disposition" form:"weee_tax_row_disposition" json:"weee_tax_row_disposition"`
	BaseWeeeTaxDisposition    float64 `gorm:"column:base_weee_tax_disposition" form:"base_weee_tax_disposition" json:"base_weee_tax_disposition"`
	BaseWeeeTaxRowDisposition float64 `gorm:"column:base_weee_tax_row_disposition" form:"base_weee_tax_row_disposition" json:"base_weee_tax_row_disposition"`
	WeeeTaxApplied            string  `gorm:"column:weee_tax_applied" form:"weee_tax_applied" json:"weee_tax_applied"`
	BaseWeeeTaxAppliedAmount  float64 `gorm:"column:base_weee_tax_applied_amount" form:"base_weee_tax_applied_amount" json:"base_weee_tax_applied_amount"`
	BaseWeeeTaxAppliedRowAmnt float64 `gorm:"column:base_weee_tax_applied_row_amnt" form:"base_weee_tax_applied_row_amnt" json:"base_weee_tax_applied_row_amnt"`
	WeeeTaxAppliedAmount      float64 `gorm:"column:weee_tax_applied_amount" form:"weee_tax_applied_amount" json:"weee_tax_applied_amount"`
	WeeeTaxAppliedRowAmount   float64 `gorm:"column:weee_tax_applied_row_amount" form:"weee_tax_applied_row_amount" json:"weee_tax_applied_row_amount"`
}

// TableName :
func (*SalesFlatCreditmemoItem) TableName() string {
	return "sales_flat_creditmemo_item"
}

// handler create
func initRoutersSalesFlatCreditmemoItem(r *gin.Engine, salesflatcreditmemoitem string) {
	route := r.Group("salesflatcreditmemoitem")
	route.GET("/", getSalesFlatCreditmemoItems)
	route.GET("/:id", getSalesFlatCreditmemoItem)
	route.POST("/", createSalesFlatCreditmemoItem)
	route.PUT("/:id", updateSalesFlatCreditmemoItem)
	route.DELETE("/:id", deleteSalesFlatCreditmemoItem)
}

func getSalesFlatCreditmemoItems(c *gin.Context) {
	var salesFlatCreditmemoItems []SalesFlatCreditmemoItem
	if err := g.Find(&salesFlatCreditmemoItems).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatCreditmemoItems)
	}
}

func getSalesFlatCreditmemoItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatCreditmemoItem SalesFlatCreditmemoItem
	if err := g.Where("id = ?", id).First(&salesFlatCreditmemoItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatCreditmemoItem)
	}
}

func createSalesFlatCreditmemoItem(c *gin.Context) {
	var salesFlatCreditmemoItem SalesFlatCreditmemoItem
	c.BindJSON(&salesFlatCreditmemoItem)
	g.Create(&salesFlatCreditmemoItem)
	c.JSON(200, salesFlatCreditmemoItem)
}

func updateSalesFlatCreditmemoItem(c *gin.Context) {
	var salesFlatCreditmemoItem SalesFlatCreditmemoItem
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatCreditmemoItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatCreditmemoItem)
	g.Save(&salesFlatCreditmemoItem)
	c.JSON(200, salesFlatCreditmemoItem)
}
func deleteSalesFlatCreditmemoItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatCreditmemoItem SalesFlatCreditmemoItem
	d := g.Where("id = ?", id).Delete(&salesFlatCreditmemoItem)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
