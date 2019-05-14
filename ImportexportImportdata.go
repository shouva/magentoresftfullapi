package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ImportexportImportdata :
type ImportexportImportdata struct {
	ID       uint16 `gorm:"column:id;primary_key" form:"id;primary_key" json:"id;primary_key"`
	Entity   string `gorm:"column:entity" form:"entity" json:"entity"`
	Behavior string `gorm:"column:behavior" form:"behavior" json:"behavior"`
	Data     string `gorm:"column:data" form:"data" json:"data"`
}

// TableName :
func (*ImportexportImportdata) TableName() string {
	return "importexport_importdata"
}

// handler create
func initRoutersImportexportImportdata(r *gin.Engine, importexportimportdata string) {
	route := r.Group("importexportimportdata")
	route.GET("/", getImportexportImportdatas)
	route.GET("/:id", getImportexportImportdata)
	route.POST("/", createImportexportImportdata)
	route.PUT("/:id", updateImportexportImportdata)
	route.DELETE("/:id", deleteImportexportImportdata)
}

func getImportexportImportdatas(c *gin.Context) {
	var importexportImportdatas []ImportexportImportdata
	if err := g.Find(&importexportImportdatas).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, importexportImportdatas)
	}
}

func getImportexportImportdata(c *gin.Context) {
	id := c.Params.ByName("id")
	var importexportImportdata ImportexportImportdata
	if err := g.Where("id = ?", id).First(&importexportImportdata).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, importexportImportdata)
	}
}

func createImportexportImportdata(c *gin.Context) {
	var importexportImportdata ImportexportImportdata
	c.BindJSON(&importexportImportdata)
	g.Create(&importexportImportdata)
	c.JSON(200, importexportImportdata)
}

func updateImportexportImportdata(c *gin.Context) {
	var importexportImportdata ImportexportImportdata
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&importexportImportdata).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&importexportImportdata)
	g.Save(&importexportImportdata)
	c.JSON(200, importexportImportdata)
}
func deleteImportexportImportdata(c *gin.Context) {
	id := c.Params.ByName("id")
	var importexportImportdata ImportexportImportdata
	d := g.Where("id = ?", id).Delete(&importexportImportdata)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
