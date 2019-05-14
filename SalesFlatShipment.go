package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatShipment :
type SalesFlatShipment struct {
	EntityId          uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	StoreId           uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	TotalWeight       float64    `gorm:"column:total_weight" form:"total_weight" json:"total_weight"`
	TotalQty          float64    `gorm:"column:total_qty" form:"total_qty" json:"total_qty"`
	EmailSent         uint16     `gorm:"column:email_sent" form:"email_sent" json:"email_sent"`
	OrderId           uint16     `gorm:"column:order_id" form:"order_id" json:"order_id"`
	CustomerId        uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	ShippingAddressId uint16     `gorm:"column:shipping_address_id" form:"shipping_address_id" json:"shipping_address_id"`
	BillingAddressId  uint16     `gorm:"column:billing_address_id" form:"billing_address_id" json:"billing_address_id"`
	ShipmentStatus    uint16     `gorm:"column:shipment_status" form:"shipment_status" json:"shipment_status"`
	IncrementId       string     `gorm:"column:increment_id" form:"increment_id" json:"increment_id"`
	CreatedAt         *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	Packages          string     `gorm:"column:packages" form:"packages" json:"packages"`
	ShippingLabel     []byte     `gorm:"column:shipping_label" form:"shipping_label" json:"shipping_label"`
}

// TableName :
func (*SalesFlatShipment) TableName() string {
	return "sales_flat_shipment"
}

// handler create
func initRoutersSalesFlatShipment(r *gin.Engine, salesflatshipment string) {
	route := r.Group("salesflatshipment")
	route.GET("/", getSalesFlatShipments)
	route.GET("/:id", getSalesFlatShipment)
	route.POST("/", createSalesFlatShipment)
	route.PUT("/:id", updateSalesFlatShipment)
	route.DELETE("/:id", deleteSalesFlatShipment)
}

func getSalesFlatShipments(c *gin.Context) {
	var salesFlatShipments []SalesFlatShipment
	if err := g.Find(&salesFlatShipments).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatShipments)
	}
}

func getSalesFlatShipment(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatShipment SalesFlatShipment
	if err := g.Where("id = ?", id).First(&salesFlatShipment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatShipment)
	}
}

func createSalesFlatShipment(c *gin.Context) {
	var salesFlatShipment SalesFlatShipment
	c.BindJSON(&salesFlatShipment)
	g.Create(&salesFlatShipment)
	c.JSON(200, salesFlatShipment)
}

func updateSalesFlatShipment(c *gin.Context) {
	var salesFlatShipment SalesFlatShipment
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatShipment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatShipment)
	g.Save(&salesFlatShipment)
	c.JSON(200, salesFlatShipment)
}
func deleteSalesFlatShipment(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatShipment SalesFlatShipment
	d := g.Where("id = ?", id).Delete(&salesFlatShipment)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
