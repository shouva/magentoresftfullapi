package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DataflowBatchImport :
type DataflowBatchImport struct {
	BatchImportId uint64 `gorm:"column:batch_import_id;primary_key" form:"batch_import_id;primary_key" json:"batch_import_id;primary_key"`
	BatchId       uint16 `gorm:"column:batch_id" form:"batch_id" json:"batch_id"`
	BatchData     string `gorm:"column:batch_data" form:"batch_data" json:"batch_data"`
	Status        uint16 `gorm:"column:status" form:"status" json:"status"`
}

// TableName :
func (*DataflowBatchImport) TableName() string {
	return "dataflow_batch_import"
}

// handler create
func initRoutersDataflowBatchImport(r *gin.Engine, dataflowbatchimport string) {
	route := r.Group("dataflowbatchimport")
	route.GET("/", getDataflowBatchImports)
	route.GET("/:id", getDataflowBatchImport)
	route.POST("/", createDataflowBatchImport)
	route.PUT("/:id", updateDataflowBatchImport)
	route.DELETE("/:id", deleteDataflowBatchImport)
}

func getDataflowBatchImports(c *gin.Context) {
	var dataflowBatchImports []DataflowBatchImport
	if err := g.Find(&dataflowBatchImports).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowBatchImports)
	}
}

func getDataflowBatchImport(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowBatchImport DataflowBatchImport
	if err := g.Where("id = ?", id).First(&dataflowBatchImport).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowBatchImport)
	}
}

func createDataflowBatchImport(c *gin.Context) {
	var dataflowBatchImport DataflowBatchImport
	c.BindJSON(&dataflowBatchImport)
	g.Create(&dataflowBatchImport)
	c.JSON(200, dataflowBatchImport)
}

func updateDataflowBatchImport(c *gin.Context) {
	var dataflowBatchImport DataflowBatchImport
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&dataflowBatchImport).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&dataflowBatchImport)
	g.Save(&dataflowBatchImport)
	c.JSON(200, dataflowBatchImport)
}
func deleteDataflowBatchImport(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowBatchImport DataflowBatchImport
	d := g.Where("id = ?", id).Delete(&dataflowBatchImport)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
