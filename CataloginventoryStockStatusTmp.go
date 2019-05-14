package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CataloginventoryStockStatusTmp :
type CataloginventoryStockStatusTmp struct {
	ProductId   uint16  `gorm:"column:product_id;primary_key" form:"product_id;primary_key" json:"product_id;primary_key"`
	WebsiteId   uint16  `gorm:"column:website_id;primary_key" form:"website_id;primary_key" json:"website_id;primary_key"`
	StockId     uint16  `gorm:"column:stock_id;primary_key" form:"stock_id;primary_key" json:"stock_id;primary_key"`
	Qty         float64 `gorm:"column:qty" form:"qty" json:"qty"`
	StockStatus uint16  `gorm:"column:stock_status" form:"stock_status" json:"stock_status"`
}

// TableName :
func (*CataloginventoryStockStatusTmp) TableName() string {
	return "cataloginventory_stock_status_tmp"
}

// handler create
func initRoutersCataloginventoryStockStatusTmp(r *gin.Engine, cataloginventorystockstatustmp string) {
	route := r.Group("cataloginventorystockstatustmp")
	route.GET("/", getCataloginventoryStockStatusTmps)
	route.GET("/:id", getCataloginventoryStockStatusTmp)
	route.POST("/", createCataloginventoryStockStatusTmp)
	route.PUT("/:id", updateCataloginventoryStockStatusTmp)
	route.DELETE("/:id", deleteCataloginventoryStockStatusTmp)
}

func getCataloginventoryStockStatusTmps(c *gin.Context) {
	var cataloginventoryStockStatusTmps []CataloginventoryStockStatusTmp
	if err := g.Find(&cataloginventoryStockStatusTmps).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cataloginventoryStockStatusTmps)
	}
}

func getCataloginventoryStockStatusTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var cataloginventoryStockStatusTmp CataloginventoryStockStatusTmp
	if err := g.Where("id = ?", id).First(&cataloginventoryStockStatusTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cataloginventoryStockStatusTmp)
	}
}

func createCataloginventoryStockStatusTmp(c *gin.Context) {
	var cataloginventoryStockStatusTmp CataloginventoryStockStatusTmp
	c.BindJSON(&cataloginventoryStockStatusTmp)
	g.Create(&cataloginventoryStockStatusTmp)
	c.JSON(200, cataloginventoryStockStatusTmp)
}

func updateCataloginventoryStockStatusTmp(c *gin.Context) {
	var cataloginventoryStockStatusTmp CataloginventoryStockStatusTmp
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&cataloginventoryStockStatusTmp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&cataloginventoryStockStatusTmp)
	g.Save(&cataloginventoryStockStatusTmp)
	c.JSON(200, cataloginventoryStockStatusTmp)
}
func deleteCataloginventoryStockStatusTmp(c *gin.Context) {
	id := c.Params.ByName("id")
	var cataloginventoryStockStatusTmp CataloginventoryStockStatusTmp
	d := g.Where("id = ?", id).Delete(&cataloginventoryStockStatusTmp)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
