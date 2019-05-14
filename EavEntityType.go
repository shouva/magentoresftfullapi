package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// EavEntityType :
type EavEntityType struct {
	EntityTypeId              uint16 `gorm:"column:entity_type_id;primary_key" form:"entity_type_id;primary_key" json:"entity_type_id;primary_key"`
	EntityTypeCode            string `gorm:"column:entity_type_code" form:"entity_type_code" json:"entity_type_code"`
	EntityModel               string `gorm:"column:entity_model" form:"entity_model" json:"entity_model"`
	AttributeModel            string `gorm:"column:attribute_model" form:"attribute_model" json:"attribute_model"`
	EntityTable               string `gorm:"column:entity_table" form:"entity_table" json:"entity_table"`
	ValueTablePrefix          string `gorm:"column:value_table_prefix" form:"value_table_prefix" json:"value_table_prefix"`
	EntityIdField             string `gorm:"column:entity_id_field" form:"entity_id_field" json:"entity_id_field"`
	IsDataSharing             uint16 `gorm:"column:is_data_sharing" form:"is_data_sharing" json:"is_data_sharing"`
	DataSharingKey            string `gorm:"column:data_sharing_key" form:"data_sharing_key" json:"data_sharing_key"`
	DefaultAttributeSetId     uint16 `gorm:"column:default_attribute_set_id" form:"default_attribute_set_id" json:"default_attribute_set_id"`
	IncrementModel            string `gorm:"column:increment_model" form:"increment_model" json:"increment_model"`
	IncrementPerStore         uint16 `gorm:"column:increment_per_store" form:"increment_per_store" json:"increment_per_store"`
	IncrementPadLength        uint16 `gorm:"column:increment_pad_length" form:"increment_pad_length" json:"increment_pad_length"`
	IncrementPadChar          string `gorm:"column:increment_pad_char" form:"increment_pad_char" json:"increment_pad_char"`
	AdditionalAttributeTable  string `gorm:"column:additional_attribute_table" form:"additional_attribute_table" json:"additional_attribute_table"`
	EntityAttributeCollection string `gorm:"column:entity_attribute_collection" form:"entity_attribute_collection" json:"entity_attribute_collection"`
}

// TableName :
func (*EavEntityType) TableName() string {
	return "eav_entity_type"
}

// handler create
func initRoutersEavEntityType(r *gin.Engine, eaventitytype string) {
	route := r.Group("eaventitytype")
	route.GET("/", getEavEntityTypes)
	route.GET("/:id", getEavEntityType)
	route.POST("/", createEavEntityType)
	route.PUT("/:id", updateEavEntityType)
	route.DELETE("/:id", deleteEavEntityType)
}

func getEavEntityTypes(c *gin.Context) {
	var eavEntityTypes []EavEntityType
	if err := g.Find(&eavEntityTypes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityTypes)
	}
}

func getEavEntityType(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityType EavEntityType
	if err := g.Where("id = ?", id).First(&eavEntityType).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, eavEntityType)
	}
}

func createEavEntityType(c *gin.Context) {
	var eavEntityType EavEntityType
	c.BindJSON(&eavEntityType)
	g.Create(&eavEntityType)
	c.JSON(200, eavEntityType)
}

func updateEavEntityType(c *gin.Context) {
	var eavEntityType EavEntityType
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&eavEntityType).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&eavEntityType)
	g.Save(&eavEntityType)
	c.JSON(200, eavEntityType)
}
func deleteEavEntityType(c *gin.Context) {
	id := c.Params.ByName("id")
	var eavEntityType EavEntityType
	d := g.Where("id = ?", id).Delete(&eavEntityType)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
