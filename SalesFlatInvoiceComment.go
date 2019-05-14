package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatInvoiceComment :
type SalesFlatInvoiceComment struct {
	EntityId           uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	ParentId           uint16     `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	IsCustomerNotified uint16     `gorm:"column:is_customer_notified" form:"is_customer_notified" json:"is_customer_notified"`
	IsVisibleOnFront   uint16     `gorm:"column:is_visible_on_front" form:"is_visible_on_front" json:"is_visible_on_front"`
	Comment            string     `gorm:"column:comment" form:"comment" json:"comment"`
	CreatedAt          *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
}

// TableName :
func (*SalesFlatInvoiceComment) TableName() string {
	return "sales_flat_invoice_comment"
}

// handler create
func initRoutersSalesFlatInvoiceComment(r *gin.Engine, salesflatinvoicecomment string) {
	route := r.Group("salesflatinvoicecomment")
	route.GET("/", getSalesFlatInvoiceComments)
	route.GET("/:id", getSalesFlatInvoiceComment)
	route.POST("/", createSalesFlatInvoiceComment)
	route.PUT("/:id", updateSalesFlatInvoiceComment)
	route.DELETE("/:id", deleteSalesFlatInvoiceComment)
}

func getSalesFlatInvoiceComments(c *gin.Context) {
	var salesFlatInvoiceComments []SalesFlatInvoiceComment
	if err := g.Find(&salesFlatInvoiceComments).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatInvoiceComments)
	}
}

func getSalesFlatInvoiceComment(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatInvoiceComment SalesFlatInvoiceComment
	if err := g.Where("id = ?", id).First(&salesFlatInvoiceComment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatInvoiceComment)
	}
}

func createSalesFlatInvoiceComment(c *gin.Context) {
	var salesFlatInvoiceComment SalesFlatInvoiceComment
	c.BindJSON(&salesFlatInvoiceComment)
	g.Create(&salesFlatInvoiceComment)
	c.JSON(200, salesFlatInvoiceComment)
}

func updateSalesFlatInvoiceComment(c *gin.Context) {
	var salesFlatInvoiceComment SalesFlatInvoiceComment
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatInvoiceComment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatInvoiceComment)
	g.Save(&salesFlatInvoiceComment)
	c.JSON(200, salesFlatInvoiceComment)
}
func deleteSalesFlatInvoiceComment(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatInvoiceComment SalesFlatInvoiceComment
	d := g.Where("id = ?", id).Delete(&salesFlatInvoiceComment)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
