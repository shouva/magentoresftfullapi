package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesFlatShipmentItem :
type SalesFlatShipmentItem struct {
	EntityId       uint16  `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	ParentId       uint16  `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	RowTotal       float64 `gorm:"column:row_total" form:"row_total" json:"row_total"`
	Price          float64 `gorm:"column:price" form:"price" json:"price"`
	Weight         float64 `gorm:"column:weight" form:"weight" json:"weight"`
	Qty            float64 `gorm:"column:qty" form:"qty" json:"qty"`
	ProductId      uint16  `gorm:"column:product_id" form:"product_id" json:"product_id"`
	OrderItemId    uint16  `gorm:"column:order_item_id" form:"order_item_id" json:"order_item_id"`
	AdditionalData string  `gorm:"column:additional_data" form:"additional_data" json:"additional_data"`
	Description    string  `gorm:"column:description" form:"description" json:"description"`
	Name           string  `gorm:"column:name" form:"name" json:"name"`
	Sku            string  `gorm:"column:sku" form:"sku" json:"sku"`
}

// TableName :
func (*SalesFlatShipmentItem) TableName() string {
	return "sales_flat_shipment_item"
}

// handler create
func initRoutersSalesFlatShipmentItem(r *gin.Engine, salesflatshipmentitem string) {
	route := r.Group("salesflatshipmentitem")
	route.GET("/", getSalesFlatShipmentItems)
	route.GET("/:id", getSalesFlatShipmentItem)
	route.POST("/", createSalesFlatShipmentItem)
	route.PUT("/:id", updateSalesFlatShipmentItem)
	route.DELETE("/:id", deleteSalesFlatShipmentItem)
}

func getSalesFlatShipmentItems(c *gin.Context) {
	var salesFlatShipmentItems []SalesFlatShipmentItem
	if err := g.Find(&salesFlatShipmentItems).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatShipmentItems)
	}
}

func getSalesFlatShipmentItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatShipmentItem SalesFlatShipmentItem
	if err := g.Where("id = ?", id).First(&salesFlatShipmentItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatShipmentItem)
	}
}

func createSalesFlatShipmentItem(c *gin.Context) {
	var salesFlatShipmentItem SalesFlatShipmentItem
	c.BindJSON(&salesFlatShipmentItem)
	g.Create(&salesFlatShipmentItem)
	c.JSON(200, salesFlatShipmentItem)
}

func updateSalesFlatShipmentItem(c *gin.Context) {
	var salesFlatShipmentItem SalesFlatShipmentItem
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatShipmentItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatShipmentItem)
	g.Save(&salesFlatShipmentItem)
	c.JSON(200, salesFlatShipmentItem)
}
func deleteSalesFlatShipmentItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatShipmentItem SalesFlatShipmentItem
	d := g.Where("id = ?", id).Delete(&salesFlatShipmentItem)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
