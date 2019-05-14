package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ShippingTablerate :
type ShippingTablerate struct {
	Pk             uint16  `gorm:"column:pk;primary_key" form:"pk;primary_key" json:"pk;primary_key"`
	WebsiteId      uint16  `gorm:"column:website_id" form:"website_id" json:"website_id"`
	DestCountryId  string  `gorm:"column:dest_country_id" form:"dest_country_id" json:"dest_country_id"`
	DestRegionId   uint16  `gorm:"column:dest_region_id" form:"dest_region_id" json:"dest_region_id"`
	DestZip        string  `gorm:"column:dest_zip" form:"dest_zip" json:"dest_zip"`
	ConditionName  string  `gorm:"column:condition_name" form:"condition_name" json:"condition_name"`
	ConditionValue float64 `gorm:"column:condition_value" form:"condition_value" json:"condition_value"`
	Price          float64 `gorm:"column:price" form:"price" json:"price"`
	Cost           float64 `gorm:"column:cost" form:"cost" json:"cost"`
}

// TableName :
func (*ShippingTablerate) TableName() string {
	return "shipping_tablerate"
}

// handler create
func initRoutersShippingTablerate(r *gin.Engine, shippingtablerate string) {
	route := r.Group("shippingtablerate")
	route.GET("/", getShippingTablerates)
	route.GET("/:id", getShippingTablerate)
	route.POST("/", createShippingTablerate)
	route.PUT("/:id", updateShippingTablerate)
	route.DELETE("/:id", deleteShippingTablerate)
}

func getShippingTablerates(c *gin.Context) {
	var shippingTablerates []ShippingTablerate
	if err := g.Find(&shippingTablerates).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, shippingTablerates)
	}
}

func getShippingTablerate(c *gin.Context) {
	id := c.Params.ByName("id")
	var shippingTablerate ShippingTablerate
	if err := g.Where("id = ?", id).First(&shippingTablerate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, shippingTablerate)
	}
}

func createShippingTablerate(c *gin.Context) {
	var shippingTablerate ShippingTablerate
	c.BindJSON(&shippingTablerate)
	g.Create(&shippingTablerate)
	c.JSON(200, shippingTablerate)
}

func updateShippingTablerate(c *gin.Context) {
	var shippingTablerate ShippingTablerate
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&shippingTablerate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&shippingTablerate)
	g.Save(&shippingTablerate)
	c.JSON(200, shippingTablerate)
}
func deleteShippingTablerate(c *gin.Context) {
	id := c.Params.ByName("id")
	var shippingTablerate ShippingTablerate
	d := g.Where("id = ?", id).Delete(&shippingTablerate)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
