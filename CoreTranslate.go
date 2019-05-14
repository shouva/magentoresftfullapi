package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CoreTranslate :
type CoreTranslate struct {
	KeyId     uint16 `gorm:"column:key_id;primary_key" form:"key_id;primary_key" json:"key_id;primary_key"`
	String    string `gorm:"column:string" form:"string" json:"string"`
	StoreId   uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Translate string `gorm:"column:translate" form:"translate" json:"translate"`
	Locale    string `gorm:"column:locale" form:"locale" json:"locale"`
	CrcString uint64 `gorm:"column:crc_string" form:"crc_string" json:"crc_string"`
}

// TableName :
func (*CoreTranslate) TableName() string {
	return "core_translate"
}

// handler create
func initRoutersCoreTranslate(r *gin.Engine, coretranslate string) {
	route := r.Group("coretranslate")
	route.GET("/", getCoreTranslates)
	route.GET("/:id", getCoreTranslate)
	route.POST("/", createCoreTranslate)
	route.PUT("/:id", updateCoreTranslate)
	route.DELETE("/:id", deleteCoreTranslate)
}

func getCoreTranslates(c *gin.Context) {
	var coreTranslates []CoreTranslate
	if err := g.Find(&coreTranslates).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreTranslates)
	}
}

func getCoreTranslate(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreTranslate CoreTranslate
	if err := g.Where("id = ?", id).First(&coreTranslate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, coreTranslate)
	}
}

func createCoreTranslate(c *gin.Context) {
	var coreTranslate CoreTranslate
	c.BindJSON(&coreTranslate)
	g.Create(&coreTranslate)
	c.JSON(200, coreTranslate)
}

func updateCoreTranslate(c *gin.Context) {
	var coreTranslate CoreTranslate
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&coreTranslate).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&coreTranslate)
	g.Save(&coreTranslate)
	c.JSON(200, coreTranslate)
}
func deleteCoreTranslate(c *gin.Context) {
	id := c.Params.ByName("id")
	var coreTranslate CoreTranslate
	d := g.Where("id = ?", id).Delete(&coreTranslate)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
