package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CataloginventoryStockStatu :
type CataloginventoryStockStatu struct {
	ProductId   uint16  `gorm:"column:product_id;primary_key" form:"product_id;primary_key" json:"product_id;primary_key"`
	WebsiteId   uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	StockId     uint16  `gorm:"column:stock_id;primary_key" form:"stock_id;primary_key" json:"stock_id;primary_key"`
	Qty         float64 `gorm:"column:qty" form:"qty" json:"qty"`
	StockStatus uint16  `gorm:"column:stock_status" form:"stock_status" json:"stock_status"`
}

// TableName :
func (*CataloginventoryStockStatu) TableName() string {
	return "cataloginventory_stock_status"
}

// handler create
func initRoutersCataloginventoryStockStatu(r *gin.Engine, cataloginventorystockstatu string) {
	route := r.Group("cataloginventorystockstatu")
	route.GET("/", getCataloginventoryStockStatus)
	route.GET("/:id", getCataloginventoryStockStatu)
	route.POST("/", createCataloginventoryStockStatu)
	route.PUT("/:id", updateCataloginventoryStockStatu)
	route.DELETE("/:id", deleteCataloginventoryStockStatu)
}

func getCataloginventoryStockStatus(c *gin.Context) {
	var cataloginventoryStockStatus []CataloginventoryStockStatu
	if err := g.Find(&cataloginventoryStockStatus).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cataloginventoryStockStatus)
	}
}

func getCataloginventoryStockStatu(c *gin.Context) {
	id := c.Params.ByName("id")
	var cataloginventoryStockStatu CataloginventoryStockStatu
	if err := g.Where("id = ?", id).First(&cataloginventoryStockStatu).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cataloginventoryStockStatu)
	}
}

func createCataloginventoryStockStatu(c *gin.Context) {
	var cataloginventoryStockStatu CataloginventoryStockStatu
	c.BindJSON(&cataloginventoryStockStatu)
	g.Create(&cataloginventoryStockStatu)
	c.JSON(200, cataloginventoryStockStatu)
}

func updateCataloginventoryStockStatu(c *gin.Context) {
	var cataloginventoryStockStatu CataloginventoryStockStatu
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&cataloginventoryStockStatu).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&cataloginventoryStockStatu)
	g.Save(&cataloginventoryStockStatu)
	c.JSON(200, cataloginventoryStockStatu)
}
func deleteCataloginventoryStockStatu(c *gin.Context) {
	id := c.Params.ByName("id")
	var cataloginventoryStockStatu CataloginventoryStockStatu
	d := g.Where("id = ?", id).Delete(&cataloginventoryStockStatu)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
