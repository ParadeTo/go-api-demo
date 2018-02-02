package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"apidemo/db"
	"time"
)

type Template struct {
	ID int `gorm:"primary;AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(255)"`
	Content string `gorm:"type:longtext"`
	UserId int `gorm:"not null"`
	CreateTime time.Time
	UpdateTime time.Time
}

func (Template) TableName() string {
	return "t_content_template"
}

func (t Template) Get(c *gin.Context)  {
	conn := db.GetMysqlConn()
	conn.First(t)
	fmt.Println(t)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}


func init () {
	fmt.Println("template")
}