package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ReviewStatu :
type ReviewStatu struct {
	StatusId   uint16 `gorm:"column:status_id;primary_key" form:"status_id;primary_key" json:"status_id;primary_key"`
	StatusCode string `gorm:"column:status_code" form:"status_code" json:"status_code"`
}

// TableName :
func (*ReviewStatu) TableName() string {
	return "review_status"
}

// handler create
func initRoutersReviewStatu(r *gin.Engine, reviewstatu string) {
	route := r.Group("reviewstatu")
	route.GET("/", getReviewStatus)
	route.GET("/:id", getReviewStatu)
	route.POST("/", createReviewStatu)
	route.PUT("/:id", updateReviewStatu)
	route.DELETE("/:id", deleteReviewStatu)
}

func getReviewStatus(c *gin.Context) {
	var reviewStatus []ReviewStatu
	if err := g.Find(&reviewStatus).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reviewStatus)
	}
}

func getReviewStatu(c *gin.Context) {
	id := c.Params.ByName("id")
	var reviewStatu ReviewStatu
	if err := g.Where("id = ?", id).First(&reviewStatu).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reviewStatu)
	}
}

func createReviewStatu(c *gin.Context) {
	var reviewStatu ReviewStatu
	c.BindJSON(&reviewStatu)
	g.Create(&reviewStatu)
	c.JSON(200, reviewStatu)
}

func updateReviewStatu(c *gin.Context) {
	var reviewStatu ReviewStatu
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&reviewStatu).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&reviewStatu)
	g.Save(&reviewStatu)
	c.JSON(200, reviewStatu)
}
func deleteReviewStatu(c *gin.Context) {
	id := c.Params.ByName("id")
	var reviewStatu ReviewStatu
	d := g.Where("id = ?", id).Delete(&reviewStatu)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
