package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreVariable :
type CoreVariable struct {
	VariableId uint16 `gorm:"column:variable_id;primary_key" form:"variable_id;primary_key" json:"variable_id;primary_key"`
	Code       string `gorm:"column:code" form:"code" json:"code"`
	Name       string `gorm:"column:name" form:"name" json:"name"`
}

// TableName :
func (*CoreVariable) TableName() string {
	return "core_variable"
}

// handler create
func initRoutersCoreVariable(r *gin.Engine, corevariable string) {
	route := r.Group("corevariable")
	route.GET("/", getCoreVariables)
	route.GET("/:id", getCoreVariable)
	route.POST("/", createCoreVariable)
	route.PUT("/:id", updateCoreVariable)
	route.DELETE("/:id", deleteCoreVariable)
}

func getCoreVariables(c *gin.Context) {
	var coreVariables []CoreVariable
	if err := g.Find(&coreVariables).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreVariables)
	}
}

func getCoreVariable(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreVariable CoreVariable
	if err := g.Where("id = ?", id).First(&coreVariable).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreVariable)
	}
}

func createCoreVariable(c *gin.Context) {
	var coreVariable CoreVariable
	c.BindJSON(&coreVariable)
	g.Create(&coreVariable)
	c.JSON(200, coreVariable)
}

func updateCoreVariable(c *gin.Context) {
	var coreVariable CoreVariable
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreVariable).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreVariable)
	g.Save(&coreVariable)
	c.JSON(200, coreVariable)
}
func deleteCoreVariable(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreVariable CoreVariable
	d := g.Where("id = ?", id).Delete(&coreVariable)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
