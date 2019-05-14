package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// DataflowProfile :
type DataflowProfile struct {
	ProfileId    uint16     `gorm:"column:profile_id;primary_key" form:"profile_id;primary_key" json:"profile_id;primary_key"`
	Name         string     `gorm:"column:name" form:"name" json:"name"`
	CreatedAt    *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	ActionsXml   string     `gorm:"column:actions_xml" form:"actions_xml" json:"actions_xml"`
	GuiData      string     `gorm:"column:gui_data" form:"gui_data" json:"gui_data"`
	Direction    string     `gorm:"column:direction" form:"direction" json:"direction"`
	EntityType   string     `gorm:"column:entity_type" form:"entity_type" json:"entity_type"`
	StoreId      uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
	DataTransfer string     `gorm:"column:data_transfer" form:"data_transfer" json:"data_transfer"`
}

// TableName :
func (*DataflowProfile) TableName() string {
	return "dataflow_profile"
}

// handler create
func initRoutersDataflowProfile(r *gin.Engine, dataflowprofile string) {
	route := r.Group("dataflowprofile")
	route.GET("/", getDataflowProfiles)
	route.GET("/:id", getDataflowProfile)
	route.POST("/", createDataflowProfile)
	route.PUT("/:id", updateDataflowProfile)
	route.DELETE("/:id", deleteDataflowProfile)
}

func getDataflowProfiles(c *gin.Context) {
	var dataflowProfiles []DataflowProfile
	if err := g.Find(&dataflowProfiles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowProfiles)
	}
}

func getDataflowProfile(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowProfile DataflowProfile
	if err := g.Where("id = ?", id).First(&dataflowProfile).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, dataflowProfile)
	}
}

func createDataflowProfile(c *gin.Context) {
	var dataflowProfile DataflowProfile
	c.BindJSON(&dataflowProfile)
	g.Create(&dataflowProfile)
	c.JSON(200, dataflowProfile)
}

func updateDataflowProfile(c *gin.Context) {
	var dataflowProfile DataflowProfile
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&dataflowProfile).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&dataflowProfile)
	g.Save(&dataflowProfile)
	c.JSON(200, dataflowProfile)
}
func deleteDataflowProfile(c *gin.Context) {
	id := c.Params.ByName("id")
	var dataflowProfile DataflowProfile
	d := g.Where("id = ?", id).Delete(&dataflowProfile)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
