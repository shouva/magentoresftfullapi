package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreUrlRewrite :
type CoreUrlRewrite struct {
	UrlRewriteId uint16 `gorm:"column:url_rewrite_id;primary_key" form:"url_rewrite_id;primary_key" json:"url_rewrite_id;primary_key"`
	StoreId      uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	IdPath       string `gorm:"column:id_path" form:"id_path" json:"id_path"`
	RequestPath  string `gorm:"column:request_path" form:"request_path" json:"request_path"`
	TargetPath   string `gorm:"column:target_path" form:"target_path" json:"target_path"`
	IsSystem     uint16 `gorm:"column:is_system" form:"is_system" json:"is_system"`
	Options      string `gorm:"column:options" form:"options" json:"options"`
	Description  string `gorm:"column:description" form:"description" json:"description"`
	CategoryId   uint16 `gorm:"column:category_id" form:"category_id" json:"category_id"`
	ProductId    uint16 `gorm:"column:product_id" form:"product_id" json:"product_id"`
}

// TableName :
func (*CoreUrlRewrite) TableName() string {
	return "core_url_rewrite"
}

// handler create
func initRoutersCoreUrlRewrite(r *gin.Engine, coreurlrewrite string) {
	route := r.Group("coreurlrewrite")
	route.GET("/", getCoreUrlRewrites)
	route.GET("/:id", getCoreUrlRewrite)
	route.POST("/", createCoreUrlRewrite)
	route.PUT("/:id", updateCoreUrlRewrite)
	route.DELETE("/:id", deleteCoreUrlRewrite)
}

func getCoreUrlRewrites(c *gin.Context) {
	var coreUrlRewrites []CoreUrlRewrite
	if err := g.Find(&coreUrlRewrites).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreUrlRewrites)
	}
}

func getCoreUrlRewrite(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreUrlRewrite CoreUrlRewrite
	if err := g.Where("id = ?", id).First(&coreUrlRewrite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreUrlRewrite)
	}
}

func createCoreUrlRewrite(c *gin.Context) {
	var coreUrlRewrite CoreUrlRewrite
	c.BindJSON(&coreUrlRewrite)
	g.Create(&coreUrlRewrite)
	c.JSON(200, coreUrlRewrite)
}

func updateCoreUrlRewrite(c *gin.Context) {
	var coreUrlRewrite CoreUrlRewrite
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreUrlRewrite).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreUrlRewrite)
	g.Save(&coreUrlRewrite)
	c.JSON(200, coreUrlRewrite)
}
func deleteCoreUrlRewrite(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreUrlRewrite CoreUrlRewrite
	d := g.Where("id = ?", id).Delete(&coreUrlRewrite)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
