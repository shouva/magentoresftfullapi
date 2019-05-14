package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// WeeeTax :
type WeeeTax struct {
	ValueId      uint16  `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	WebsiteId    uint16  `gorm:"column:website_id" form:"website_id" json:"website_id"`
	EntityId     uint16  `gorm:"column:entity_id" form:"entity_id" json:"entity_id"`
	Country      string  `gorm:"column:country" form:"country" json:"country"`
	Value        float64 `gorm:"column:value" form:"value" json:"value"`
	State        string  `gorm:"column:state" form:"state" json:"state"`
	AttributeId  uint16  `gorm:"column:attribute_id" form:"attribute_id" json:"attribute_id"`
	EntityTypeId uint16  `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
}

// TableName :
func (*WeeeTax) TableName() string {
	return "weee_tax"
}

// handler create
func initRoutersWeeeTax(r *gin.Engine, weeetax string) {
	route := r.Group("weeetax")
	route.GET("/", getWeeeTaxs)
	route.GET("/:id", getWeeeTax)
	route.POST("/", createWeeeTax)
	route.PUT("/:id", updateWeeeTax)
	route.DELETE("/:id", deleteWeeeTax)
}

func getWeeeTaxs(c *gin.Context) {
	var weeeTaxs []WeeeTax
	if err := g.Find(&weeeTaxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, weeeTaxs)
	}
}

func getWeeeTax(c *gin.Context) {
	id := c.Params.ByName("id")
	var weeeTax WeeeTax
	if err := g.Where("id = ?", id).First(&weeeTax).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, weeeTax)
	}
}

func createWeeeTax(c *gin.Context) {
	var weeeTax WeeeTax
	c.BindJSON(&weeeTax)
	g.Create(&weeeTax)
	c.JSON(200, weeeTax)
}

func updateWeeeTax(c *gin.Context) {
	var weeeTax WeeeTax
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&weeeTax).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&weeeTax)
	g.Save(&weeeTax)
	c.JSON(200, weeeTax)
}
func deleteWeeeTax(c *gin.Context) {
	id := c.Params.ByName("id")
	var weeeTax WeeeTax
	d := g.Where("id = ?", id).Delete(&weeeTax)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
