package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatShipmentGrid :
type SalesFlatShipmentGrid struct {
	EntityId         uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	StoreId          uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	TotalQty         float64    `gorm:"column:total_qty" form:"total_qty" json:"total_qty"`
	OrderId          uint16     `gorm:"column:order_id" form:"order_id" json:"order_id"`
	ShipmentStatus   uint16     `gorm:"column:shipment_status" form:"shipment_status" json:"shipment_status"`
	IncrementId      string     `gorm:"column:increment_id" form:"increment_id" json:"increment_id"`
	OrderIncrementId string     `gorm:"column:order_increment_id" form:"order_increment_id" json:"order_increment_id"`
	CreatedAt        *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	OrderCreatedAt   *time.Time `gorm:"column:order_created_at" form:"order_created_at" json:"order_created_at"`
	ShippingName     string     `gorm:"column:shipping_name" form:"shipping_name" json:"shipping_name"`
}

// TableName :
func (*SalesFlatShipmentGrid) TableName() string {
	return "sales_flat_shipment_grid"
}

// handler create
func initRoutersSalesFlatShipmentGrid(r *gin.Engine, salesflatshipmentgrid string) {
	route := r.Group("salesflatshipmentgrid")
	route.GET("/", getSalesFlatShipmentGrids)
	route.GET("/:id", getSalesFlatShipmentGrid)
	route.POST("/", createSalesFlatShipmentGrid)
	route.PUT("/:id", updateSalesFlatShipmentGrid)
	route.DELETE("/:id", deleteSalesFlatShipmentGrid)
}

func getSalesFlatShipmentGrids(c *gin.Context) {
	var salesFlatShipmentGrids []SalesFlatShipmentGrid
	if err := g.Find(&salesFlatShipmentGrids).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatShipmentGrids)
	}
}

func getSalesFlatShipmentGrid(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatShipmentGrid SalesFlatShipmentGrid
	if err := g.Where("id = ?", id).First(&salesFlatShipmentGrid).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatShipmentGrid)
	}
}

func createSalesFlatShipmentGrid(c *gin.Context) {
	var salesFlatShipmentGrid SalesFlatShipmentGrid
	c.BindJSON(&salesFlatShipmentGrid)
	g.Create(&salesFlatShipmentGrid)
	c.JSON(200, salesFlatShipmentGrid)
}

func updateSalesFlatShipmentGrid(c *gin.Context) {
	var salesFlatShipmentGrid SalesFlatShipmentGrid
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatShipmentGrid).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatShipmentGrid)
	g.Save(&salesFlatShipmentGrid)
	c.JSON(200, salesFlatShipmentGrid)
}
func deleteSalesFlatShipmentGrid(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatShipmentGrid SalesFlatShipmentGrid
	d := g.Where("id = ?", id).Delete(&salesFlatShipmentGrid)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
