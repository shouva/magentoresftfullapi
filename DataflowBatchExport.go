package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DataflowBatchExport :
type DataflowBatchExport struct {
	BatchExportId uint64 `gorm:"column:batch_export_id;primary_key" form:"batch_export_id;primary_key" json:"batch_export_id;primary_key"`
	BatchId       uint16 `gorm:"column:batch_id" form:"batch_id" json:"batch_id"`
	BatchData     string `gorm:"column:batch_data" form:"batch_data" json:"batch_data"`
	Status        uint16 `gorm:"column:status" form:"status" json:"status"`
}

// TableName :
func (*DataflowBatchExport) TableName() string {
	return "dataflow_batch_export"
}

// handler create
func initRoutersDataflowBatchExport(r *gin.Engine, dataflowbatchexport string) {
	route := r.Group("dataflowbatchexport")
	route.GET("/", getDataflowBatchExports)
	route.GET("/:id", getDataflowBatchExport)
	route.POST("/", createDataflowBatchExport)
	route.PUT("/:id", updateDataflowBatchExport)
	route.DELETE("/:id", deleteDataflowBatchExport)
}

func getDataflowBatchExports(c *gin.Context) {
	var dataflowBatchExports []DataflowBatchExport
	if err := g.Find(&dataflowBatchExports).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowBatchExports)
	}
}

func getDataflowBatchExport(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowBatchExport DataflowBatchExport
	if err := g.Where("id = ?", id).First(&dataflowBatchExport).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowBatchExport)
	}
}

func createDataflowBatchExport(c *gin.Context) {
	var dataflowBatchExport DataflowBatchExport
	c.BindJSON(&dataflowBatchExport)
	g.Create(&dataflowBatchExport)
	c.JSON(200, dataflowBatchExport)
}

func updateDataflowBatchExport(c *gin.Context) {
	var dataflowBatchExport DataflowBatchExport
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&dataflowBatchExport).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&dataflowBatchExport)
	g.Save(&dataflowBatchExport)
	c.JSON(200, dataflowBatchExport)
}
func deleteDataflowBatchExport(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowBatchExport DataflowBatchExport
	d := g.Where("id = ?", id).Delete(&dataflowBatchExport)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
