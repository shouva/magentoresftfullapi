package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesOrderStatusState :
type SalesOrderStatusState struct {
	Status    string `gorm:"column:status;primary_key" form:"status;primary_key" json:"status;primary_key"`
	State     string `gorm:"column:state;primary_key" form:"state;primary_key" json:"state;primary_key"`
	IsDefault uint16 `gorm:"column:is_default" form:"is_default" json:"is_default"`
}

// TableName :
func (*SalesOrderStatusState) TableName() string {
	return "sales_order_status_state"
}

// handler create
func initRoutersSalesOrderStatusState(r *gin.Engine, salesorderstatusstate string) {
	route := r.Group("salesorderstatusstate")
	route.GET("/", getSalesOrderStatusStates)
	route.GET("/:id", getSalesOrderStatusState)
	route.POST("/", createSalesOrderStatusState)
	route.PUT("/:id", updateSalesOrderStatusState)
	route.DELETE("/:id", deleteSalesOrderStatusState)
}

func getSalesOrderStatusStates(c *gin.Context) {
	var salesOrderStatusStates []SalesOrderStatusState
	if err := g.Find(&salesOrderStatusStates).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderStatusStates)
	}
}

func getSalesOrderStatusState(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderStatusState SalesOrderStatusState
	if err := g.Where("id = ?", id).First(&salesOrderStatusState).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesOrderStatusState)
	}
}

func createSalesOrderStatusState(c *gin.Context) {
	var salesOrderStatusState SalesOrderStatusState
	c.BindJSON(&salesOrderStatusState)
	g.Create(&salesOrderStatusState)
	c.JSON(200, salesOrderStatusState)
}

func updateSalesOrderStatusState(c *gin.Context) {
	var salesOrderStatusState SalesOrderStatusState
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesOrderStatusState).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesOrderStatusState)
	g.Save(&salesOrderStatusState)
	c.JSON(200, salesOrderStatusState)
}
func deleteSalesOrderStatusState(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesOrderStatusState SalesOrderStatusState
	d := g.Where("id = ?", id).Delete(&salesOrderStatusState)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
