package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// XmlconnectConfigData :
type XmlconnectConfigData struct {
	ApplicationId uint16 `gorm:"column:application_id;primary_key" form:"application_id;primary_key" json:"application_id;primary_key"`
	Category      string `gorm:"column:category;primary_key" form:"category;primary_key" json:"category;primary_key"`
	Path          string `gorm:"column:path;primary_key" form:"path;primary_key" json:"path;primary_key"`
	Value         string `gorm:"column:value" form:"value" json:"value"`
}

// TableName :
func (*XmlconnectConfigData) TableName() string {
	return "xmlconnect_config_data"
}

// handler create
func initRoutersXmlconnectConfigData(r *gin.Engine, xmlconnectconfigdata string) {
	route := r.Group("xmlconnectconfigdata")
	route.GET("/", getXmlconnectConfigDatas)
	route.GET("/:id", getXmlconnectConfigData)
	route.POST("/", createXmlconnectConfigData)
	route.PUT("/:id", updateXmlconnectConfigData)
	route.DELETE("/:id", deleteXmlconnectConfigData)
}

func getXmlconnectConfigDatas(c *gin.Context) {
	var xmlconnectConfigDatas []XmlconnectConfigData
	if err := g.Find(&xmlconnectConfigDatas).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, xmlconnectConfigDatas)
	}
}

func getXmlconnectConfigData(c *gin.Context) {
	id := c.Params.ByName("id")
	var xmlconnectConfigData XmlconnectConfigData
	if err := g.Where("id = ?", id).First(&xmlconnectConfigData).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, xmlconnectConfigData)
	}
}

func createXmlconnectConfigData(c *gin.Context) {
	var xmlconnectConfigData XmlconnectConfigData
	c.BindJSON(&xmlconnectConfigData)
	g.Create(&xmlconnectConfigData)
	c.JSON(200, xmlconnectConfigData)
}

func updateXmlconnectConfigData(c *gin.Context) {
	var xmlconnectConfigData XmlconnectConfigData
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&xmlconnectConfigData).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&xmlconnectConfigData)
	g.Save(&xmlconnectConfigData)
	c.JSON(200, xmlconnectConfigData)
}
func deleteXmlconnectConfigData(c *gin.Context) {
	id := c.Params.ByName("id")
	var xmlconnectConfigData XmlconnectConfigData
	d := g.Where("id = ?", id).Delete(&xmlconnectConfigData)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
