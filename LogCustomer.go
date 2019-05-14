package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LogCustomer :
type LogCustomer struct {
	LogId      uint16     `gorm:"column:log_id;primary_key" form:"log_id;primary_key" json:"log_id;primary_key"`
	VisitorId  uint64     `gorm:"column:visitor_id" form:"visitor_id" json:"visitor_id"`
	CustomerId uint16     `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
	LoginAt    *time.Time `gorm:"column:login_at" form:"login_at" json:"login_at"`
	LogoutAt   *time.Time `gorm:"column:logout_at" form:"logout_at" json:"logout_at"`
	StoreId    uint16     `gorm:"column:store_id" form:"store_id" json:"store_id"`
}

// TableName :
func (*LogCustomer) TableName() string {
	return "log_customer"
}

// handler create
func initRoutersLogCustomer(r *gin.Engine, logcustomer string) {
	route := r.Group("logcustomer")
	route.GET("/", getLogCustomers)
	route.GET("/:id", getLogCustomer)
	route.POST("/", createLogCustomer)
	route.PUT("/:id", updateLogCustomer)
	route.DELETE("/:id", deleteLogCustomer)
}

func getLogCustomers(c *gin.Context) {
	var logCustomers []LogCustomer
	if err := g.Find(&logCustomers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logCustomers)
	}
}

func getLogCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var logCustomer LogCustomer
	if err := g.Where("id = ?", id).First(&logCustomer).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logCustomer)
	}
}

func createLogCustomer(c *gin.Context) {
	var logCustomer LogCustomer
	c.BindJSON(&logCustomer)
	g.Create(&logCustomer)
	c.JSON(200, logCustomer)
}

func updateLogCustomer(c *gin.Context) {
	var logCustomer LogCustomer
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&logCustomer).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&logCustomer)
	g.Save(&logCustomer)
	c.JSON(200, logCustomer)
}
func deleteLogCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var logCustomer LogCustomer
	d := g.Where("id = ?", id).Delete(&logCustomer)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
