package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// XmlconnectImage :
type XmlconnectImage struct {
	ImageId       uint16 `gorm:"column:image_id;primary_key" form:"image_id;primary_key" json:"image_id;primary_key"`
	ApplicationId uint16 `gorm:"column:application_id" form:"application_id" json:"application_id"`
	ImageFile     string `gorm:"column:image_file" form:"image_file" json:"image_file"`
	ImageType     string `gorm:"column:image_type" form:"image_type" json:"image_type"`
	Order         uint16 `gorm:"column:order" form:"order" json:"order"`
}

// TableName :
func (*XmlconnectImage) TableName() string {
	return "xmlconnect_images"
}

// handler create
func initRoutersXmlconnectImage(r *gin.Engine, xmlconnectimage string) {
	route := r.Group("xmlconnectimage")
	route.GET("/", getXmlconnectImages)
	route.GET("/:id", getXmlconnectImage)
	route.POST("/", createXmlconnectImage)
	route.PUT("/:id", updateXmlconnectImage)
	route.DELETE("/:id", deleteXmlconnectImage)
}

func getXmlconnectImages(c *gin.Context) {
	var xmlconnectImages []XmlconnectImage
	if err := g.Find(&xmlconnectImages).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, xmlconnectImages)
	}
}

func getXmlconnectImage(c *gin.Context) {
	id := c.Params.ByName("id")
	var xmlconnectImage XmlconnectImage
	if err := g.Where("id = ?", id).First(&xmlconnectImage).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, xmlconnectImage)
	}
}

func createXmlconnectImage(c *gin.Context) {
	var xmlconnectImage XmlconnectImage
	c.BindJSON(&xmlconnectImage)
	g.Create(&xmlconnectImage)
	c.JSON(200, xmlconnectImage)
}

func updateXmlconnectImage(c *gin.Context) {
	var xmlconnectImage XmlconnectImage
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&xmlconnectImage).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&xmlconnectImage)
	g.Save(&xmlconnectImage)
	c.JSON(200, xmlconnectImage)
}
func deleteXmlconnectImage(c *gin.Context) {
	id := c.Params.ByName("id")
	var xmlconnectImage XmlconnectImage
	d := g.Where("id = ?", id).Delete(&xmlconnectImage)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
