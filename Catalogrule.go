package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Catalogrule :
type Catalogrule struct {
	RuleId               uint16     `gorm:"column:rule_id;primary_key" form:"rule_id;primary_key" json:"rule_id;primary_key"`
	Name                 string     `gorm:"column:name" form:"name" json:"name"`
	Description          string     `gorm:"column:description" form:"description" json:"description"`
	FromDate             *time.Time `gorm:"column:from_date" form:"from_date" json:"from_date"`
	ToDate               *time.Time `gorm:"column:to_date" form:"to_date" json:"to_date"`
	IsActive             uint16     `gorm:"column:is_active" form:"is_active" json:"is_active"`
	ConditionsSerialized string     `gorm:"column:conditions_serialized" form:"conditions_serialized" json:"conditions_serialized"`
	ActionsSerialized    string     `gorm:"column:actions_serialized" form:"actions_serialized" json:"actions_serialized"`
	StopRulesProcessing  uint16     `gorm:"column:stop_rules_processing" form:"stop_rules_processing" json:"stop_rules_processing"`
	SortOrder            uint16     `gorm:"column:sort_order" form:"sort_order" json:"sort_order"`
	SimpleAction         string     `gorm:"column:simple_action" form:"simple_action" json:"simple_action"`
	DiscountAmount       float64    `gorm:"column:discount_amount" form:"discount_amount" json:"discount_amount"`
	SubIsEnable          uint16     `gorm:"column:sub_is_enable" form:"sub_is_enable" json:"sub_is_enable"`
	SubSimpleAction      string     `gorm:"column:sub_simple_action" form:"sub_simple_action" json:"sub_simple_action"`
	SubDiscountAmount    float64    `gorm:"column:sub_discount_amount" form:"sub_discount_amount" json:"sub_discount_amount"`
}

// TableName :
func (*Catalogrule) TableName() string {
	return "catalogrule"
}

// handler create
func initRoutersCatalogrule(r *gin.Engine, catalogrule string) {
	route := r.Group("catalogrule")
	route.GET("/", getCatalogrules)
	route.GET("/:id", getCatalogrule)
	route.POST("/", createCatalogrule)
	route.PUT("/:id", updateCatalogrule)
	route.DELETE("/:id", deleteCatalogrule)
}

func getCatalogrules(c *gin.Context) {
	var catalogrules []Catalogrule
	if err := g.Find(&catalogrules).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogrules)
	}
}

func getCatalogrule(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogrule Catalogrule
	if err := g.Where("id = ?", id).First(&catalogrule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, catalogrule)
	}
}

func createCatalogrule(c *gin.Context) {
	var catalogrule Catalogrule
	c.BindJSON(&catalogrule)
	g.Create(&catalogrule)
	c.JSON(200, catalogrule)
}

func updateCatalogrule(c *gin.Context) {
	var catalogrule Catalogrule
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&catalogrule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&catalogrule)
	g.Save(&catalogrule)
	c.JSON(200, catalogrule)
}
func deleteCatalogrule(c *gin.Context) {
	id := c.Params.ByName("id")
	var catalogrule Catalogrule
	d := g.Where("id = ?", id).Delete(&catalogrule)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
