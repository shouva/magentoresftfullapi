package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CatalogruleAffectedProduct :
type CatalogruleAffectedProduct struct {
	ProductId uint16 `gorm:"column:product_id;primary_key" form:"product_id;primary_key" json:"product_id;primary_key"`
}

// TableName :
func (*CatalogruleAffectedProduct) TableName() string {
	return "catalogrule_affected_product"
}

// handler create
func initRoutersCatalogruleAffectedProduct(r *gin.Engine, catalogruleaffectedproduct string) {
	route := r.Group("catalogruleaffectedproduct")
	route.GET("/", getCatalogruleAffectedProducts)
	route.GET("/:id", getCatalogruleAffectedProduct)
	route.POST("/", createCatalogruleAffectedProduct)
	route.PUT("/:id", updateCatalogruleAffectedProduct)
	route.DELETE("/:id", deleteCatalogruleAffectedProduct)
}

func getCatalogruleAffectedProducts(c *gin.Context) {
	var catalogruleAffectedProducts []CatalogruleAffectedProduct
	if err := g.Find(&catalogruleAffectedProducts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogruleAffectedProducts)
	}
}

func getCatalogruleAffectedProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogruleAffectedProduct CatalogruleAffectedProduct
	if err := g.Where("id = ?", id).First(&catalogruleAffectedProduct).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogruleAffectedProduct)
	}
}

func createCatalogruleAffectedProduct(c *gin.Context) {
	var catalogruleAffectedProduct CatalogruleAffectedProduct
	c.BindJSON(&catalogruleAffectedProduct)
	g.Create(&catalogruleAffectedProduct)
	c.JSON(200, catalogruleAffectedProduct)
}

func updateCatalogruleAffectedProduct(c *gin.Context) {
	var catalogruleAffectedProduct CatalogruleAffectedProduct
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogruleAffectedProduct).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogruleAffectedProduct)
	g.Save(&catalogruleAffectedProduct)
	c.JSON(200, catalogruleAffectedProduct)
}
func deleteCatalogruleAffectedProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogruleAffectedProduct CatalogruleAffectedProduct
	d := g.Where("id = ?", id).Delete(&catalogruleAffectedProduct)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
