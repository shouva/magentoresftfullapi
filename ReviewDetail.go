package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ReviewDetail :
type ReviewDetail struct {
	DetailId   uint64 `gorm:"column:detail_id;primary_key" form:"detail_id;primary_key" json:"detail_id;primary_key"`
	ReviewId   uint64 `gorm:"column:review_id" form:"review_id" json:"review_id"`
	StoreId    uint16 `gorm:"column:store_id" form:"store_id" json:"store_id"`
	Title      string `gorm:"column:title" form:"title" json:"title"`
	Detail     string `gorm:"column:detail" form:"detail" json:"detail"`
	Nickname   string `gorm:"column:nickname" form:"nickname" json:"nickname"`
	CustomerId uint16 `gorm:"column:customer_id" form:"customer_id" json:"customer_id"`
}

// TableName :
func (*ReviewDetail) TableName() string {
	return "review_detail"
}

// handler create
func initRoutersReviewDetail(r *gin.Engine, reviewdetail string) {
	route := r.Group("reviewdetail")
	route.GET("/", getReviewDetails)
	route.GET("/:id", getReviewDetail)
	route.POST("/", createReviewDetail)
	route.PUT("/:id", updateReviewDetail)
	route.DELETE("/:id", deleteReviewDetail)
}

func getReviewDetails(c *gin.Context) {
	var reviewDetails []ReviewDetail
	if err := g.Find(&reviewDetails).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reviewDetails)
	}
}

func getReviewDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var reviewDetail ReviewDetail
	if err := g.Where("id = ?", id).First(&reviewDetail).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, reviewDetail)
	}
}

func createReviewDetail(c *gin.Context) {
	var reviewDetail ReviewDetail
	c.BindJSON(&reviewDetail)
	g.Create(&reviewDetail)
	c.JSON(200, reviewDetail)
}

func updateReviewDetail(c *gin.Context) {
	var reviewDetail ReviewDetail
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&reviewDetail).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&reviewDetail)
	g.Save(&reviewDetail)
	c.JSON(200, reviewDetail)
}
func deleteReviewDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var reviewDetail ReviewDetail
	d := g.Where("id = ?", id).Delete(&reviewDetail)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
