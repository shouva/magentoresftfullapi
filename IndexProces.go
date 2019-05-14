package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// IndexProces :
type IndexProces struct {
	ProcessId   uint16     `gorm:"column:process_id;primary_key" form:"process_id;primary_key" json:"process_id;primary_key"`
	IndexerCode string     `gorm:"column:indexer_code" form:"indexer_code" json:"indexer_code"`
	Status      string     `gorm:"column:status" form:"status" json:"status"`
	StartedAt   *time.Time `gorm:"column:started_at" form:"started_at" json:"started_at"`
	EndedAt     *time.Time `gorm:"column:ended_at" form:"ended_at" json:"ended_at"`
	Mode        string     `gorm:"column:mode" form:"mode" json:"mode"`
}

// TableName :
func (*IndexProces) TableName() string {
	return "index_process"
}

// handler create
func initRoutersIndexProces(r *gin.Engine, indexproces string) {
	route := r.Group("indexproces")
	route.GET("/", getIndexProcess)
	route.GET("/:id", getIndexProces)
	route.POST("/", createIndexProces)
	route.PUT("/:id", updateIndexProces)
	route.DELETE("/:id", deleteIndexProces)
}

func getIndexProcess(c *gin.Context) {
	var indexProcess []IndexProces
	if err := g.Find(&indexProcess).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, indexProcess)
	}
}

func getIndexProces(c *gin.Context) {
	id := c.Params.ByName("id")
	var indexProces IndexProces
	if err := g.Where("id = ?", id).First(&indexProces).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, indexProces)
	}
}

func createIndexProces(c *gin.Context) {
	var indexProces IndexProces
	c.BindJSON(&indexProces)
	g.Create(&indexProces)
	c.JSON(200, indexProces)
}

func updateIndexProces(c *gin.Context) {
	var indexProces IndexProces
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&indexProces).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&indexProces)
	g.Save(&indexProces)
	c.JSON(200, indexProces)
}
func deleteIndexProces(c *gin.Context) {
	id := c.Params.ByName("id")
	var indexProces IndexProces
	d := g.Where("id = ?", id).Delete(&indexProces)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
