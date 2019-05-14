package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatShipmentTrack :
type SalesFlatShipmentTrack struct {
	EntityId    uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	ParentId    uint16     `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	Weight      float64    `gorm:"column:weight" form:"weight" json:"weight"`
	Qty         float64    `gorm:"column:qty" form:"qty" json:"qty"`
	OrderId     uint16     `gorm:"column:order_id" form:"order_id" json:"order_id"`
	TrackNumber string     `gorm:"column:track_number" form:"track_number" json:"track_number"`
	Description string     `gorm:"column:description" form:"description" json:"description"`
	Title       string     `gorm:"column:title" form:"title" json:"title"`
	CarrierCode string     `gorm:"column:carrier_code" form:"carrier_code" json:"carrier_code"`
	CreatedAt   *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
}

// TableName :
func (*SalesFlatShipmentTrack) TableName() string {
	return "sales_flat_shipment_track"
}

// handler create
func initRoutersSalesFlatShipmentTrack(r *gin.Engine, salesflatshipmenttrack string) {
	route := r.Group("salesflatshipmenttrack")
	route.GET("/", getSalesFlatShipmentTracks)
	route.GET("/:id", getSalesFlatShipmentTrack)
	route.POST("/", createSalesFlatShipmentTrack)
	route.PUT("/:id", updateSalesFlatShipmentTrack)
	route.DELETE("/:id", deleteSalesFlatShipmentTrack)
}

func getSalesFlatShipmentTracks(c *gin.Context) {
	var salesFlatShipmentTracks []SalesFlatShipmentTrack
	if err := g.Find(&salesFlatShipmentTracks).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatShipmentTracks)
	}
}

func getSalesFlatShipmentTrack(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatShipmentTrack SalesFlatShipmentTrack
	if err := g.Where("id = ?", id).First(&salesFlatShipmentTrack).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatShipmentTrack)
	}
}

func createSalesFlatShipmentTrack(c *gin.Context) {
	var salesFlatShipmentTrack SalesFlatShipmentTrack
	c.BindJSON(&salesFlatShipmentTrack)
	g.Create(&salesFlatShipmentTrack)
	c.JSON(200, salesFlatShipmentTrack)
}

func updateSalesFlatShipmentTrack(c *gin.Context) {
	var salesFlatShipmentTrack SalesFlatShipmentTrack
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatShipmentTrack).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatShipmentTrack)
	g.Save(&salesFlatShipmentTrack)
	c.JSON(200, salesFlatShipmentTrack)
}
func deleteSalesFlatShipmentTrack(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatShipmentTrack SalesFlatShipmentTrack
	d := g.Where("id = ?", id).Delete(&salesFlatShipmentTrack)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
