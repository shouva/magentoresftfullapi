package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DataflowImportData :
type DataflowImportData struct {
	ImportId     uint16 `gorm:"column:import_id;primary_key" form:"import_id;primary_key" json:"import_id;primary_key"`
	SessionId    uint16 `gorm:"column:session_id" form:"session_id" json:"session_id"`
	SerialNumber uint16 `gorm:"column:serial_number" form:"serial_number" json:"serial_number"`
	Value        string `gorm:"column:value" form:"value" json:"value"`
	Status       uint16 `gorm:"column:status" form:"status" json:"status"`
}

// TableName :
func (*DataflowImportData) TableName() string {
	return "dataflow_import_data"
}

// handler create
func initRoutersDataflowImportData(r *gin.Engine, dataflowimportdata string) {
	route := r.Group("dataflowimportdata")
	route.GET("/", getDataflowImportDatas)
	route.GET("/:id", getDataflowImportData)
	route.POST("/", createDataflowImportData)
	route.PUT("/:id", updateDataflowImportData)
	route.DELETE("/:id", deleteDataflowImportData)
}

func getDataflowImportDatas(c *gin.Context) {
	var dataflowImportDatas []DataflowImportData
	if err := g.Find(&dataflowImportDatas).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowImportDatas)
	}
}

func getDataflowImportData(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowImportData DataflowImportData
	if err := g.Where("id = ?", id).First(&dataflowImportData).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowImportData)
	}
}

func createDataflowImportData(c *gin.Context) {
	var dataflowImportData DataflowImportData
	c.BindJSON(&dataflowImportData)
	g.Create(&dataflowImportData)
	c.JSON(200, dataflowImportData)
}

func updateDataflowImportData(c *gin.Context) {
	var dataflowImportData DataflowImportData
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&dataflowImportData).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&dataflowImportData)
	g.Save(&dataflowImportData)
	c.JSON(200, dataflowImportData)
}
func deleteDataflowImportData(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowImportData DataflowImportData
	d := g.Where("id = ?", id).Delete(&dataflowImportData)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
