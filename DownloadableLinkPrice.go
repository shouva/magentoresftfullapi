package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DownloadableLinkPrice :
type DownloadableLinkPrice struct {
	PriceId   uint16  `gorm:"column:price_id;primary_key" form:"price_id;primary_key" json:"price_id;primary_key"`
	LinkId    uint16  `gorm:"column:link_id" form:"link_id" json:"link_id"`
	WebsiteId uint16  `gorm:"column:website_id" form:"website_id" json:"website_id"`
	Price     float64 `gorm:"column:price" form:"price" json:"price"`
}

// TableName :
func (*DownloadableLinkPrice) TableName() string {
	return "downloadable_link_price"
}

// handler create
func initRoutersDownloadableLinkPrice(r *gin.Engine, downloadablelinkprice string) {
	route := r.Group("downloadablelinkprice")
	route.GET("/", getDownloadableLinkPrices)
	route.GET("/:id", getDownloadableLinkPrice)
	route.POST("/", createDownloadableLinkPrice)
	route.PUT("/:id", updateDownloadableLinkPrice)
	route.DELETE("/:id", deleteDownloadableLinkPrice)
}

func getDownloadableLinkPrices(c *gin.Context) {
	var downloadableLinkPrices []DownloadableLinkPrice
	if err := g.Find(&downloadableLinkPrices).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableLinkPrices)
	}
}

func getDownloadableLinkPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableLinkPrice DownloadableLinkPrice
	if err := g.Where("id = ?", id).First(&downloadableLinkPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, downloadableLinkPrice)
	}
}

func createDownloadableLinkPrice(c *gin.Context) {
	var downloadableLinkPrice DownloadableLinkPrice
	c.BindJSON(&downloadableLinkPrice)
	g.Create(&downloadableLinkPrice)
	c.JSON(200, downloadableLinkPrice)
}

func updateDownloadableLinkPrice(c *gin.Context) {
	var downloadableLinkPrice DownloadableLinkPrice
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&downloadableLinkPrice).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&downloadableLinkPrice)
	g.Save(&downloadableLinkPrice)
	c.JSON(200, downloadableLinkPrice)
}
func deleteDownloadableLinkPrice(c *gin.Context) {
	id := c.Params.ByName("id")
	var downloadableLinkPrice DownloadableLinkPrice
	d := g.Where("id = ?", id).Delete(&downloadableLinkPrice)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
