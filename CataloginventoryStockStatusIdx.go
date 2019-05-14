package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CataloginventoryStockStatusIdx :
type CataloginventoryStockStatusIdx struct {
	ProductId   uint16  `gorm:"column:product_id;primary_key" form:"product_id;primary_key" json:"product_id;primary_key"`
	WebsiteId   uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	StockId     uint16  `gorm:"column:stock_id;primary_key" form:"stock_id;primary_key" json:"stock_id;primary_key"`
	Qty         float64 `gorm:"column:qty" form:"qty" json:"qty"`
	StockStatus uint16  `gorm:"column:stock_status" form:"stock_status" json:"stock_status"`
}

// TableName :
func (*CataloginventoryStockStatusIdx) TableName() string {
	return "cataloginventory_stock_status_idx"
}

// handler create
func initRoutersCataloginventoryStockStatusIdx(r *gin.Engine, cataloginventorystockstatusidx string) {
	route := r.Group("cataloginventorystockstatusidx")
	route.GET("/", getCataloginventoryStockStatusIdxs)
	route.GET("/:id", getCataloginventoryStockStatusIdx)
	route.POST("/", createCataloginventoryStockStatusIdx)
	route.PUT("/:id", updateCataloginventoryStockStatusIdx)
	route.DELETE("/:id", deleteCataloginventoryStockStatusIdx)
}

func getCataloginventoryStockStatusIdxs(c *gin.Context) {
	var cataloginventoryStockStatusIdxs []CataloginventoryStockStatusIdx
	if err := g.Find(&cataloginventoryStockStatusIdxs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cataloginventoryStockStatusIdxs)
	}
}

func getCataloginventoryStockStatusIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var cataloginventoryStockStatusIdx CataloginventoryStockStatusIdx
	if err := g.Where("id = ?", id).First(&cataloginventoryStockStatusIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cataloginventoryStockStatusIdx)
	}
}

func createCataloginventoryStockStatusIdx(c *gin.Context) {
	var cataloginventoryStockStatusIdx CataloginventoryStockStatusIdx
	c.BindJSON(&cataloginventoryStockStatusIdx)
	g.Create(&cataloginventoryStockStatusIdx)
	c.JSON(200, cataloginventoryStockStatusIdx)
}

func updateCataloginventoryStockStatusIdx(c *gin.Context) {
	var cataloginventoryStockStatusIdx CataloginventoryStockStatusIdx
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&cataloginventoryStockStatusIdx).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&cataloginventoryStockStatusIdx)
	g.Save(&cataloginventoryStockStatusIdx)
	c.JSON(200, cataloginventoryStockStatusIdx)
}
func deleteCataloginventoryStockStatusIdx(c *gin.Context) {
	id := c.Params.ByName("id")
	var cataloginventoryStockStatusIdx CataloginventoryStockStatusIdx
	d := g.Where("id = ?", id).Delete(&cataloginventoryStockStatusIdx)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
