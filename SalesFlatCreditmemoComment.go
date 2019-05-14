package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SalesFlatCreditmemoComment :
type SalesFlatCreditmemoComment struct {
	EntityId           uint16     `gorm:"column:entity_id;primary_key" form:"entity_id;primary_key" json:"entity_id;primary_key"`
	ParentId           uint16     `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`
	IsCustomerNotified uint16     `gorm:"column:is_customer_notified" form:"is_customer_notified" json:"is_customer_notified"`
	IsVisibleOnFront   uint16     `gorm:"column:is_visible_on_front" form:"is_visible_on_front" json:"is_visible_on_front"`
	Comment            string     `gorm:"column:comment" form:"comment" json:"comment"`
	CreatedAt          *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
}

// TableName :
func (*SalesFlatCreditmemoComment) TableName() string {
	return "sales_flat_creditmemo_comment"
}

// handler create
func initRoutersSalesFlatCreditmemoComment(r *gin.Engine, salesflatcreditmemocomment string) {
	route := r.Group("salesflatcreditmemocomment")
	route.GET("/", getSalesFlatCreditmemoComments)
	route.GET("/:id", getSalesFlatCreditmemoComment)
	route.POST("/", createSalesFlatCreditmemoComment)
	route.PUT("/:id", updateSalesFlatCreditmemoComment)
	route.DELETE("/:id", deleteSalesFlatCreditmemoComment)
}

func getSalesFlatCreditmemoComments(c *gin.Context) {
	var salesFlatCreditmemoComments []SalesFlatCreditmemoComment
	if err := g.Find(&salesFlatCreditmemoComments).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatCreditmemoComments)
	}
}

func getSalesFlatCreditmemoComment(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatCreditmemoComment SalesFlatCreditmemoComment
	if err := g.Where("id = ?", id).First(&salesFlatCreditmemoComment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesFlatCreditmemoComment)
	}
}

func createSalesFlatCreditmemoComment(c *gin.Context) {
	var salesFlatCreditmemoComment SalesFlatCreditmemoComment
	c.BindJSON(&salesFlatCreditmemoComment)
	g.Create(&salesFlatCreditmemoComment)
	c.JSON(200, salesFlatCreditmemoComment)
}

func updateSalesFlatCreditmemoComment(c *gin.Context) {
	var salesFlatCreditmemoComment SalesFlatCreditmemoComment
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesFlatCreditmemoComment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesFlatCreditmemoComment)
	g.Save(&salesFlatCreditmemoComment)
	c.JSON(200, salesFlatCreditmemoComment)
}
func deleteSalesFlatCreditmemoComment(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesFlatCreditmemoComment SalesFlatCreditmemoComment
	d := g.Where("id = ?", id).Delete(&salesFlatCreditmemoComment)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
