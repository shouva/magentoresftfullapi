package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogruleProduct :
type CatalogruleProduct struct {
	RuleProductId     uint16  `gorm:"column:rule_product_id;primary_key" form:"rule_product_id;primary_key" json:"rule_product_id;primary_key"`
	RuleId            uint16  `gorm:"column:rule_id" form:"rule_id" json:"rule_id"`
	FromTime          uint16  `gorm:"column:from_time" form:"from_time" json:"from_time"`
	ToTime            uint16  `gorm:"column:to_time" form:"to_time" json:"to_time"`
	CustomerGroupId   uint16  `gorm:"column:customer_group_id" form:"customer_group_id" json:"customer_group_id"`
	ProductId         uint16  `gorm:"column:product_id" form:"product_id" json:"product_id"`
	ActionOperator    string  `gorm:"column:action_operator" form:"action_operator" json:"action_operator"`
	ActionAmount      float64 `gorm:"column:action_amount" form:"action_amount" json:"action_amount"`
	ActionStop        uint16  `gorm:"column:action_stop" form:"action_stop" json:"action_stop"`
	SortOrder         uint16  `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
	WebsiteId         uint16  `gorm:"column:website_id" form:"website_id" json:"website_id"`
	SubSimpleAction   string  `gorm:"column:sub_simple_action" form:"sub_simple_action" json:"sub_simple_action"`
	SubDiscountAmount float64 `gorm:"column:sub_discount_amount" form:"sub_discount_amount" json:"sub_discount_amount"`
}

// TableName :
func (*CatalogruleProduct) TableName() string {
	return "catalogrule_product"
}

// handler create
func initRoutersCatalogruleProduct(r *gin.Engine, catalogruleproduct string) {
	route := r.Group("catalogruleproduct")
	route.GET("/", getCatalogruleProducts)
	route.GET("/:id", getCatalogruleProduct)
	route.POST("/", createCatalogruleProduct)
	route.PUT("/:id", updateCatalogruleProduct)
	route.DELETE("/:id", deleteCatalogruleProduct)
}

func getCatalogruleProducts(c *gin.Context) {
	var catalogruleProducts []CatalogruleProduct
	if err := g.Find(&catalogruleProducts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogruleProducts)
	}
}

func getCatalogruleProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogruleProduct CatalogruleProduct
	if err := g.Where("id = ?", id).First(&catalogruleProduct).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogruleProduct)
	}
}

func createCatalogruleProduct(c *gin.Context) {
	var catalogruleProduct CatalogruleProduct
	c.BindJSON(&catalogruleProduct)
	g.Create(&catalogruleProduct)
	c.JSON(200, catalogruleProduct)
}

func updateCatalogruleProduct(c *gin.Context) {
	var catalogruleProduct CatalogruleProduct
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogruleProduct).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogruleProduct)
	g.Save(&catalogruleProduct)
	c.JSON(200, catalogruleProduct)
}
func deleteCatalogruleProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogruleProduct CatalogruleProduct
	d := g.Where("id = ?", id).Delete(&catalogruleProduct)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
