package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TaxClas :
type TaxClas struct {
	ClassId   uint16 `gorm:"column:class_id;primary_key" form:"class_id;primary_key" json:"class_id;primary_key"`
	ClassName string `gorm:"column:class_name" form:"class_name" json:"class_name"`
	ClassType string `gorm:"column:class_type" form:"class_type" json:"class_type"`
}

// TableName :
func (*TaxClas) TableName() string {
	return "tax_class"
}

// handler create
func initRoutersTaxClas(r *gin.Engine, taxclas string) {
	route := r.Group("taxclas")
	route.GET("/", getTaxClass)
	route.GET("/:id", getTaxClas)
	route.POST("/", createTaxClas)
	route.PUT("/:id", updateTaxClas)
	route.DELETE("/:id", deleteTaxClas)
}

func getTaxClass(c *gin.Context) {
	var taxClass []TaxClas
	if err := g.Find(&taxClass).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxClass)
	}
}

func getTaxClas(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxClas TaxClas
	if err := g.Where("id = ?", id).First(&taxClas).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, taxClas)
	}
}

func createTaxClas(c *gin.Context) {
	var taxClas TaxClas
	c.BindJSON(&taxClas)
	g.Create(&taxClas)
	c.JSON(200, taxClas)
}

func updateTaxClas(c *gin.Context) {
	var taxClas TaxClas
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&taxClas).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&taxClas)
	g.Save(&taxClas)
	c.JSON(200, taxClas)
}
func deleteTaxClas(c *gin.Context) {
	id := c.Params.ByName("id")
	var taxClas TaxClas
	d := g.Where("id = ?", id).Delete(&taxClas)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
