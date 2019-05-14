package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CataloginventoryStockItem :
type CataloginventoryStockItem struct {
	ItemId                  uint16     `gorm:"column:item_id;primary_key" form:"item_id;primary_key" json:"item_id;primary_key"`
	ProductId               uint16     `gorm:"column:product_id" form:"product_id" json:"product_id"`
	StockId                 uint16     `gorm:"column:stock_id" form:"stock_id" json:"stock_id"`
	Qty                     float64    `gorm:"column:qty" form:"qty" json:"qty"`
	MinQty                  float64    `gorm:"column:min_qty" form:"min_qty" json:"min_qty"`
	UseConfigMinQty         uint16     `gorm:"column:use_config_min_qty" form:"use_config_min_qty" json:"use_config_min_qty"`
	IsQtyDecimal            uint16     `gorm:"column:is_qty_decimal" form:"is_qty_decimal" json:"is_qty_decimal"`
	Backorders              uint16     `gorm:"column:backorders" form:"backorders" json:"backorders"`
	UseConfigBackorders     uint16     `gorm:"column:use_config_backorders" form:"use_config_backorders" json:"use_config_backorders"`
	MinSaleQty              float64    `gorm:"column:min_sale_qty" form:"min_sale_qty" json:"min_sale_qty"`
	UseConfigMinSaleQty     uint16     `gorm:"column:use_config_min_sale_qty" form:"use_config_min_sale_qty" json:"use_config_min_sale_qty"`
	MaxSaleQty              float64    `gorm:"column:max_sale_qty" form:"max_sale_qty" json:"max_sale_qty"`
	UseConfigMaxSaleQty     uint16     `gorm:"column:use_config_max_sale_qty" form:"use_config_max_sale_qty" json:"use_config_max_sale_qty"`
	IsInStock               uint16     `gorm:"column:is_in_stock" form:"is_in_stock" json:"is_in_stock"`
	LowStockDate            *time.Time `gorm:"column:low_stock_date" form:"low_stock_date" json:"low_stock_date"`
	NotifyStockQty          float64    `gorm:"column:notify_stock_qty" form:"notify_stock_qty" json:"notify_stock_qty"`
	UseConfigNotifyStockQty uint16     `gorm:"column:use_config_notify_stock_qty" form:"use_config_notify_stock_qty" json:"use_config_notify_stock_qty"`
	ManageStock             uint16     `gorm:"column:manage_stock" form:"manage_stock" json:"manage_stock"`
	UseConfigManageStock    uint16     `gorm:"column:use_config_manage_stock" form:"use_config_manage_stock" json:"use_config_manage_stock"`
	StockStatusChangedAuto  uint16     `gorm:"column:stock_status_changed_auto" form:"stock_status_changed_auto" json:"stock_status_changed_auto"`
	UseConfigQtyIncrements  uint16     `gorm:"column:use_config_qty_increments" form:"use_config_qty_increments" json:"use_config_qty_increments"`
	QtyIncrements           float64    `gorm:"column:qty_increments" form:"qty_increments" json:"qty_increments"`
	UseConfigEnableQtyInc   uint16     `gorm:"column:use_config_enable_qty_inc" form:"use_config_enable_qty_inc" json:"use_config_enable_qty_inc"`
	EnableQtyIncrements     uint16     `gorm:"column:enable_qty_increments" form:"enable_qty_increments" json:"enable_qty_increments"`
	IsDecimalDivided        uint16     `gorm:"column:is_decimal_divided" form:"is_decimal_divided" json:"is_decimal_divided"`
}

// TableName :
func (*CataloginventoryStockItem) TableName() string {
	return "cataloginventory_stock_item"
}

// handler create
func initRoutersCataloginventoryStockItem(r *gin.Engine, cataloginventorystockitem string) {
	route := r.Group("cataloginventorystockitem")
	route.GET("/", getCataloginventoryStockItems)
	route.GET("/:id", getCataloginventoryStockItem)
	route.POST("/", createCataloginventoryStockItem)
	route.PUT("/:id", updateCataloginventoryStockItem)
	route.DELETE("/:id", deleteCataloginventoryStockItem)
}

func getCataloginventoryStockItems(c *gin.Context) {
	var cataloginventoryStockItems []CataloginventoryStockItem
	if err := g.Find(&cataloginventoryStockItems).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cataloginventoryStockItems)
	}
}

func getCataloginventoryStockItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var cataloginventoryStockItem CataloginventoryStockItem
	if err := g.Where("id = ?", id).First(&cataloginventoryStockItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cataloginventoryStockItem)
	}
}

func createCataloginventoryStockItem(c *gin.Context) {
	var cataloginventoryStockItem CataloginventoryStockItem
	c.BindJSON(&cataloginventoryStockItem)
	g.Create(&cataloginventoryStockItem)
	c.JSON(200, cataloginventoryStockItem)
}

func updateCataloginventoryStockItem(c *gin.Context) {
	var cataloginventoryStockItem CataloginventoryStockItem
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&cataloginventoryStockItem).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&cataloginventoryStockItem)
	g.Save(&cataloginventoryStockItem)
	c.JSON(200, cataloginventoryStockItem)
}
func deleteCataloginventoryStockItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var cataloginventoryStockItem CataloginventoryStockItem
	d := g.Where("id = ?", id).Delete(&cataloginventoryStockItem)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
