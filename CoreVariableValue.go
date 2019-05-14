package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreVariableValue :
type CoreVariableValue struct {
	ValueId    uint16 `gorm:"column:value_id;primary_key" form:"value_id;primary_key" json:"value_id;primary_key"`
	VariableId uint16 `gorm:"column:variable_id" form:"variable_id" json:"variable_id"`
	StoreId    uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	PlainValue string `gorm:"column:plain_value" form:"plain_value" json:"plain_value"`
	HtmlValue  string `gorm:"column:html_value" form:"html_value" json:"html_value"`
}

// TableName :
func (*CoreVariableValue) TableName() string {
	return "core_variable_value"
}

// handler create
func initRoutersCoreVariableValue(r *gin.Engine, corevariablevalue string) {
	route := r.Group("corevariablevalue")
	route.GET("/", getCoreVariableValues)
	route.GET("/:id", getCoreVariableValue)
	route.POST("/", createCoreVariableValue)
	route.PUT("/:id", updateCoreVariableValue)
	route.DELETE("/:id", deleteCoreVariableValue)
}

func getCoreVariableValues(c *gin.Context) {
	var coreVariableValues []CoreVariableValue
	if err := g.Find(&coreVariableValues).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreVariableValues)
	}
}

func getCoreVariableValue(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreVariableValue CoreVariableValue
	if err := g.Where("id = ?", id).First(&coreVariableValue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreVariableValue)
	}
}

func createCoreVariableValue(c *gin.Context) {
	var coreVariableValue CoreVariableValue
	c.BindJSON(&coreVariableValue)
	g.Create(&coreVariableValue)
	c.JSON(200, coreVariableValue)
}

func updateCoreVariableValue(c *gin.Context) {
	var coreVariableValue CoreVariableValue
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreVariableValue).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreVariableValue)
	g.Save(&coreVariableValue)
	c.JSON(200, coreVariableValue)
}
func deleteCoreVariableValue(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreVariableValue CoreVariableValue
	d := g.Where("id = ?", id).Delete(&coreVariableValue)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
