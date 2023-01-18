package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	DBMigrate(DBConnect())
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "成功",
		})
	})
	r.Run(":8080")
}

type User struct {
	gorm.Model
	NickName string `json:"nickName"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}

func DBConnect() *gorm.DB {
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "sample"
	DSN := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
