package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// WeeeDiscount :
type WeeeDiscount struct {
	EntityId        uint16  `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	WebsiteId       uint16  `gorm:"column:website_id" form:"website_id" json:"website_id"`
	CustomerGroupId uint16  `gorm:"column:customer_group_id" form:"customer_group_id" json:"customer_group_id"`
	Value           float64 `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*WeeeDiscount) TableName() string {
	return "weee_discount"
}

// handler create
func initRoutersWeeeDiscount(r *gin.Engine, weeediscount string) {
	route := r.Group("weeediscount")
	route.GET("/", getWeeeDiscounts)
	route.GET("/:id", getWeeeDiscount)
	route.POST("/", createWeeeDiscount)
	route.PUT("/:id", updateWeeeDiscount)
	route.DELETE("/:id", deleteWeeeDiscount)
}

func getWeeeDiscounts(c *gin.Context) {
	var weeeDiscounts []WeeeDiscount
	if err := g.Find(&weeeDiscounts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, weeeDiscounts)
	}
}

func getWeeeDiscount(c *gin.Context) {
	id := c.Params.ByName("id")
	var weeeDiscount WeeeDiscount
	if err := g.Where("id = ?", id).First(&weeeDiscount).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, weeeDiscount)
	}
}

func createWeeeDiscount(c *gin.Context) {
	var weeeDiscount WeeeDiscount
	c.BindJSON(&weeeDiscount)
	g.Create(&weeeDiscount)
	c.JSON(200, weeeDiscount)
}

func updateWeeeDiscount(c *gin.Context) {
	var weeeDiscount WeeeDiscount
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&weeeDiscount).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&weeeDiscount)
	g.Save(&weeeDiscount)
	c.JSON(200, weeeDiscount)
}
func deleteWeeeDiscount(c *gin.Context) {
	id := c.Params.ByName("id")
	var weeeDiscount WeeeDiscount
	d := g.Where("id = ?", id).Delete(&weeeDiscount)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
