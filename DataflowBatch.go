package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// DataflowBatch :
type DataflowBatch struct {
	BatchId   uint16     `gorm:"column:batch_id;primary_key" form:"batch_id;primary_key" json:"batch_id;primary_key"`
	ProfileId uint16     `gorm:"column:profile_id" form:"profile_id" json:"profile_id"`
	StoreId   uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Adapter   string     `gorm:"column:adapter" form:"adapter" json:"adapter"`
	Params    string     `gorm:"column:params" form:"params" json:"params"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
}

// TableName :
func (*DataflowBatch) TableName() string {
	return "dataflow_batch"
}

// handler create
func initRoutersDataflowBatch(r *gin.Engine, dataflowbatch string) {
	route := r.Group("dataflowbatch")
	route.GET("/", getDataflowBatchs)
	route.GET("/:id", getDataflowBatch)
	route.POST("/", createDataflowBatch)
	route.PUT("/:id", updateDataflowBatch)
	route.DELETE("/:id", deleteDataflowBatch)
}

func getDataflowBatchs(c *gin.Context) {
	var dataflowBatchs []DataflowBatch
	if err := g.Find(&dataflowBatchs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowBatchs)
	}
}

func getDataflowBatch(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowBatch DataflowBatch
	if err := g.Where("id = ?", id).First(&dataflowBatch).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowBatch)
	}
}

func createDataflowBatch(c *gin.Context) {
	var dataflowBatch DataflowBatch
	c.BindJSON(&dataflowBatch)
	g.Create(&dataflowBatch)
	c.JSON(200, dataflowBatch)
}

func updateDataflowBatch(c *gin.Context) {
	var dataflowBatch DataflowBatch
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&dataflowBatch).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&dataflowBatch)
	g.Save(&dataflowBatch)
	c.JSON(200, dataflowBatch)
}
func deleteDataflowBatch(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowBatch DataflowBatch
	d := g.Where("id = ?", id).Delete(&dataflowBatch)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
