package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CronSchedule :
type CronSchedule struct {
	ScheduleId  uint16     `gorm:"column:schedule_id;primary_key" form:"schedule_id;primary_key" json:"schedule_id;primary_key"`
	JobCode     string     `gorm:"column:job_code" form:"job_code" json:"job_code"`
	Status      string     `gorm:"column:status" form:"status" json:"status"`
	Messages    string     `gorm:"column:messages" form:"messages" json:"messages"`
	CreatedAt   *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	ScheduledAt *time.Time `gorm:"column:scheduled_at" form:"scheduled_at" json:"scheduled_at"`
	ExecutedAt  *time.Time `gorm:"column:executed_at" form:"executed_at" json:"executed_at"`
	FinishedAt  *time.Time `gorm:"column:finished_at" form:"finished_at" json:"finished_at"`
}

// TableName :
func (*CronSchedule) TableName() string {
	return "cron_schedule"
}

// handler create
func initRoutersCronSchedule(r *gin.Engine, cronschedule string) {
	route := r.Group("cronschedule")
	route.GET("/", getCronSchedules)
	route.GET("/:id", getCronSchedule)
	route.POST("/", createCronSchedule)
	route.PUT("/:id", updateCronSchedule)
	route.DELETE("/:id", deleteCronSchedule)
}

func getCronSchedules(c *gin.Context) {
	var cronSchedules []CronSchedule
	if err := g.Find(&cronSchedules).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cronSchedules)
	}
}

func getCronSchedule(c *gin.Context) {
	id := c.Params.ByName("id")
	var cronSchedule CronSchedule
	if err := g.Where("id = ?", id).First(&cronSchedule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cronSchedule)
	}
}

func createCronSchedule(c *gin.Context) {
	var cronSchedule CronSchedule
	c.BindJSON(&cronSchedule)
	g.Create(&cronSchedule)
	c.JSON(200, cronSchedule)
}

func updateCronSchedule(c *gin.Context) {
	var cronSchedule CronSchedule
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&cronSchedule).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&cronSchedule)
	g.Save(&cronSchedule)
	c.JSON(200, cronSchedule)
}
func deleteCronSchedule(c *gin.Context) {
	id := c.Params.ByName("id")
	var cronSchedule CronSchedule
	d := g.Where("id = ?", id).Delete(&cronSchedule)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
