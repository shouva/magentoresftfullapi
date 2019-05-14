package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatShipmentComment :
type SalesFlatShipmentComment struct {
	EntityId           uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	ParentId           uint16     `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	IsCustomerNotified uint16     `gorm:"column:is_customer_notified" form:"is_customer_notified" json:"is_customer_notified"`
	IsVisibleOnFront   uint16     `gorm:"column:is_visible_on_front" form:"is_visible_on_front" json:"is_visible_on_front"`
	Comment            string     `gorm:"column:comment" form:"comment" json:"comment"`
	CreatedAt          *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
}

// TableName :
func (*SalesFlatShipmentComment) TableName() string {
	return "sales_flat_shipment_comment"
}

// handler create
func initRoutersSalesFlatShipmentComment(r *gin.Engine, salesflatshipmentcomment string) {
	route := r.Group("salesflatshipmentcomment")
	route.GET("/", getSalesFlatShipmentComments)
	route.GET("/:id", getSalesFlatShipmentComment)
	route.POST("/", createSalesFlatShipmentComment)
	route.PUT("/:id", updateSalesFlatShipmentComment)
	route.DELETE("/:id", deleteSalesFlatShipmentComment)
}

func getSalesFlatShipmentComments(c *gin.Context) {
	var salesFlatShipmentComments []SalesFlatShipmentComment
	if err := g.Find(&salesFlatShipmentComments).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatShipmentComments)
	}
}

func getSalesFlatShipmentComment(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatShipmentComment SalesFlatShipmentComment
	if err := g.Where("id = ?", id).First(&salesFlatShipmentComment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatShipmentComment)
	}
}

func createSalesFlatShipmentComment(c *gin.Context) {
	var salesFlatShipmentComment SalesFlatShipmentComment
	c.BindJSON(&salesFlatShipmentComment)
	g.Create(&salesFlatShipmentComment)
	c.JSON(200, salesFlatShipmentComment)
}

func updateSalesFlatShipmentComment(c *gin.Context) {
	var salesFlatShipmentComment SalesFlatShipmentComment
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatShipmentComment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatShipmentComment)
	g.Save(&salesFlatShipmentComment)
	c.JSON(200, salesFlatShipmentComment)
}
func deleteSalesFlatShipmentComment(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatShipmentComment SalesFlatShipmentComment
	d := g.Where("id = ?", id).Delete(&salesFlatShipmentComment)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
