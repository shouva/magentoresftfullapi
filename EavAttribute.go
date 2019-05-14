package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavAttribute :
type EavAttribute struct {
	AttributeId    uint16 `gorm:"column:attribute_id;primary_key" form:"attribute_id;primary_key" json:"attribute_id;primary_key"`
	EntityTypeId   uint16 `gorm:"column:entity_type_id" form:"entity_type_id" json:"entity_type_id"`
	AttributeCode  string `gorm:"column:attribute_code" form:"attribute_code" json:"attribute_code"`
	AttributeModel string `gorm:"column:attribute_model" form:"attribute_model" json:"attribute_model"`
	BackendModel   string `gorm:"column:backend_model" form:"backend_model" json:"backend_model"`
	BackendType    string `gorm:"column:backend_type" form:"backend_type" json:"backend_type"`
	BackendTable   string `gorm:"column:backend_table" form:"backend_table" json:"backend_table"`
	FrontendModel  string `gorm:"column:frontend_model" form:"frontend_model" json:"frontend_model"`
	FrontendInput  string `gorm:"column:frontend_input" form:"frontend_input" json:"frontend_input"`
	FrontendLabel  string `gorm:"column:frontend_label" form:"frontend_label" json:"frontend_label"`
	FrontendClass  string `gorm:"column:frontend_class" form:"frontend_class" json:"frontend_class"`
	SourceModel    string `gorm:"column:source_model" form:"source_model" json:"source_model"`
	IsRequired     uint16 `gorm:"column:is_required" form:"is_required" json:"is_required"`
	IsUserDefined  uint16 `gorm:"column:is_user_defined" form:"is_user_defined" json:"is_user_defined"`
	DefaultValue   string `gorm:"column:default_value" form:"default_value" json:"default_value"`
	IsUnique       uint16 `gorm:"column:is_unique" form:"is_unique" json:"is_unique"`
	Note           string `gorm:"column:note" form:"note" json:"note"`
}

// TableName :
func (*EavAttribute) TableName() string {
	return "eav_attribute"
}

// handler create
func initRoutersEavAttribute(r *gin.Engine, eavattribute string) {
	route := r.Group("eavattribute")
	route.GET("/", getEavAttributes)
	route.GET("/:id", getEavAttribute)
	route.POST("/", createEavAttribute)
	route.PUT("/:id", updateEavAttribute)
	route.DELETE("/:id", deleteEavAttribute)
}

func getEavAttributes(c *gin.Context) {
	var eavAttributes []EavAttribute
	if err := g.Find(&eavAttributes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavAttributes)
	}
}

func getEavAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavAttribute EavAttribute
	if err := g.Where("id = ?", id).First(&eavAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavAttribute)
	}
}

func createEavAttribute(c *gin.Context) {
	var eavAttribute EavAttribute
	c.BindJSON(&eavAttribute)
	g.Create(&eavAttribute)
	c.JSON(200, eavAttribute)
}

func updateEavAttribute(c *gin.Context) {
	var eavAttribute EavAttribute
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavAttribute).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavAttribute)
	g.Save(&eavAttribute)
	c.JSON(200, eavAttribute)
}
func deleteEavAttribute(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavAttribute EavAttribute
	d := g.Where("id = ?", id).Delete(&eavAttribute)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
