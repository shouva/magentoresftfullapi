package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SalesruleLabel :
type SalesruleLabel struct {
	LabelId uint16 `gorm:"column:label_id;primary_key" form:"label_id;primary_key" json:"label_id;primary_key"`
	RuleId  uint16 `gorm:"column:rule_id" form:"rule_id" json:"rule_id"`
	StoreId uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Label   string `gorm:"column:label" form:"label" json:"label"`
}

// TableName :
func (*SalesruleLabel) TableName() string {
	return "salesrule_label"
}

// handler create
func initRoutersSalesruleLabel(r *gin.Engine, salesrulelabel string) {
	route := r.Group("salesrulelabel")
	route.GET("/", getSalesruleLabels)
	route.GET("/:id", getSalesruleLabel)
	route.POST("/", createSalesruleLabel)
	route.PUT("/:id", updateSalesruleLabel)
	route.DELETE("/:id", deleteSalesruleLabel)
}

func getSalesruleLabels(c *gin.Context) {
	var salesruleLabels []SalesruleLabel
	if err := g.Find(&salesruleLabels).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleLabels)
	}
}

func getSalesruleLabel(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleLabel SalesruleLabel
	if err := g.Where("id = ?", id).First(&salesruleLabel).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, salesruleLabel)
	}
}

func createSalesruleLabel(c *gin.Context) {
	var salesruleLabel SalesruleLabel
	c.BindJSON(&salesruleLabel)
	g.Create(&salesruleLabel)
	c.JSON(200, salesruleLabel)
}

func updateSalesruleLabel(c *gin.Context) {
	var salesruleLabel SalesruleLabel
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&salesruleLabel).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&salesruleLabel)
	g.Save(&salesruleLabel)
	c.JSON(200, salesruleLabel)
}
func deleteSalesruleLabel(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesruleLabel SalesruleLabel
	d := g.Where("id = ?", id).Delete(&salesruleLabel)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
