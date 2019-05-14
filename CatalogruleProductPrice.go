package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CatalogruleProductPrice :
type CatalogruleProductPrice struct {
	RuleProductPriceId uint16     `gorm:"column:rule_product_price_id;primary_key" form:"rule_product_price_id;primary_key" json:"rule_product_price_id;primary_key"`
	RuleDate           *time.Time `gorm:"column:rule_date" form:"rule_date" json:"rule_date"`
	CustomerGroupId    uint16     `gorm:"column:customer_group_id" form:"customer_group_id" json:"customer_group_id"`
	ProductId          uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	RulePrice          float64    `gorm:"column:rule_price" form:"rule_price" json:"rule_price"`
	WebsiteId          uint16     `gorm:"column:website_id" form:"website_id" json:"website_id"`
	LatestStartDate    *time.Time `gorm:"column:latest_start_date" form:"latest_start_date" json:"latest_start_date"`
	EarliestEndDate    *time.Time `gorm:"column:earliest_end_date" form:"earliest_end_date" json:"earliest_end_date"`
}

// TableName :
func (*CatalogruleProductPrice) TableName() string {
	return "catalogrule_product_price"
}

// handler create
func initRoutersCatalogruleProductPrice(r *gin.Engine, catalogruleproductprice string) {
	route := r.Group("catalogruleproductprice")
	route.GET("/", getCatalogruleProductPrices)
	route.GET("/:id", getCatalogruleProductPrice)
	route.POST("/", createCatalogruleProductPrice)
	route.PUT("/:id", updateCatalogruleProductPrice)
	route.DELETE("/:id", deleteCatalogruleProductPrice)
}

func getCatalogruleProductPrices(c *gin.Context) {
	var catalogruleProductPrices []CatalogruleProductPrice
	if err := g.Find(&catalogruleProductPrices).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogruleProductPrices)
	}
}

func getCatalogruleProductPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogruleProductPrice CatalogruleProductPrice
	if err := g.Where("id = ?", id).First(&catalogruleProductPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogruleProductPrice)
	}
}

func createCatalogruleProductPrice(c *gin.Context) {
	var catalogruleProductPrice CatalogruleProductPrice
	c.BindJSON(&catalogruleProductPrice)
	g.Create(&catalogruleProductPrice)
	c.JSON(200, catalogruleProductPrice)
}

func updateCatalogruleProductPrice(c *gin.Context) {
	var catalogruleProductPrice CatalogruleProductPrice
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogruleProductPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogruleProductPrice)
	g.Save(&catalogruleProductPrice)
	c.JSON(200, catalogruleProductPrice)
}
func deleteCatalogruleProductPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogruleProductPrice CatalogruleProductPrice
	d := g.Where("id = ?", id).Delete(&catalogruleProductPrice)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
