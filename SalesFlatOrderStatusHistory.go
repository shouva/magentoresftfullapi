package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatOrderStatusHistory :
type SalesFlatOrderStatusHistory struct {
	EntityId           uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	ParentId           uint16     `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	IsCustomerNotified uint16     `gorm:"column:is_customer_notified" form:"is_customer_notified" json:"is_customer_notified"`
	IsVisibleOnFront   uint16     `gorm:"column:is_visible_on_front" form:"is_visible_on_front" json:"is_visible_on_front"`
	Comment            string     `gorm:"column:comment" form:"comment" json:"comment"`
	Status             string     `gorm:"column:status" form:"status" json:"status"`
	CreatedAt          *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	EntityName         string     `gorm:"column:entity_name" form:"entity_name" json:"entity_name"`
}

// TableName :
func (*SalesFlatOrderStatusHistory) TableName() string {
	return "sales_flat_order_status_history"
}

// handler create
func initRoutersSalesFlatOrderStatusHistory(r *gin.Engine, salesflatorderstatushistory string) {
	route := r.Group("salesflatorderstatushistory")
	route.GET("/", getSalesFlatOrderStatusHistorys)
	route.GET("/:id", getSalesFlatOrderStatusHistory)
	route.POST("/", createSalesFlatOrderStatusHistory)
	route.PUT("/:id", updateSalesFlatOrderStatusHistory)
	route.DELETE("/:id", deleteSalesFlatOrderStatusHistory)
}

func getSalesFlatOrderStatusHistorys(c *gin.Context) {
	var salesFlatOrderStatusHistorys []SalesFlatOrderStatusHistory
	if err := g.Find(&salesFlatOrderStatusHistorys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatOrderStatusHistorys)
	}
}

func getSalesFlatOrderStatusHistory(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatOrderStatusHistory SalesFlatOrderStatusHistory
	if err := g.Where("id = ?", id).First(&salesFlatOrderStatusHistory).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatOrderStatusHistory)
	}
}

func createSalesFlatOrderStatusHistory(c *gin.Context) {
	var salesFlatOrderStatusHistory SalesFlatOrderStatusHistory
	c.BindJSON(&salesFlatOrderStatusHistory)
	g.Create(&salesFlatOrderStatusHistory)
	c.JSON(200, salesFlatOrderStatusHistory)
}

func updateSalesFlatOrderStatusHistory(c *gin.Context) {
	var salesFlatOrderStatusHistory SalesFlatOrderStatusHistory
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatOrderStatusHistory).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatOrderStatusHistory)
	g.Save(&salesFlatOrderStatusHistory)
	c.JSON(200, salesFlatOrderStatusHistory)
}
func deleteSalesFlatOrderStatusHistory(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatOrderStatusHistory SalesFlatOrderStatusHistory
	d := g.Where("id = ?", id).Delete(&salesFlatOrderStatusHistory)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
