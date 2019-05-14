package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CataloginventoryStock :
type CataloginventoryStock struct {
	StockId   uint16 `gorm:"column:stock_id;primary_key" form:"stock_id;primary_key" json:"stock_id;primary_key"`
	StockName string `gorm:"column:stock_name" form:"stock_name" json:"stock_name"`
}

// TableName :
func (*CataloginventoryStock) TableName() string {
	return "cataloginventory_stock"
}

// handler create
func initRoutersCataloginventoryStock(r *gin.Engine, cataloginventorystock string) {
	route := r.Group("cataloginventorystock")
	route.GET("/", getCataloginventoryStocks)
	route.GET("/:id", getCataloginventoryStock)
	route.POST("/", createCataloginventoryStock)
	route.PUT("/:id", updateCataloginventoryStock)
	route.DELETE("/:id", deleteCataloginventoryStock)
}

func getCataloginventoryStocks(c *gin.Context) {
	var cataloginventoryStocks []CataloginventoryStock
	if err := g.Find(&cataloginventoryStocks).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cataloginventoryStocks)
	}
}

func getCataloginventoryStock(c *gin.Context) {
	id := c.Params.ByName("id")
	var cataloginventoryStock CataloginventoryStock
	if err := g.Where("id = ?", id).First(&cataloginventoryStock).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cataloginventoryStock)
	}
}

func createCataloginventoryStock(c *gin.Context) {
	var cataloginventoryStock CataloginventoryStock
	c.BindJSON(&cataloginventoryStock)
	g.Create(&cataloginventoryStock)
	c.JSON(200, cataloginventoryStock)
}

func updateCataloginventoryStock(c *gin.Context) {
	var cataloginventoryStock CataloginventoryStock
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&cataloginventoryStock).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&cataloginventoryStock)
	g.Save(&cataloginventoryStock)
	c.JSON(200, cataloginventoryStock)
}
func deleteCataloginventoryStock(c *gin.Context) {
	id := c.Params.ByName("id")
	var cataloginventoryStock CataloginventoryStock
	d := g.Where("id = ?", id).Delete(&cataloginventoryStock)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
