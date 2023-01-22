package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := DBMigrate(DBConnect())

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		var result Product
		db.Last(&result).Scan(&result)
		fmt.Println(result)
		c.JSON(200, gin.H{
			"message": "成功",
			"test": "aaaaaaa",
		})
	})
	r.POST("/create", func(c *gin.Context) {
		create(db, "D11", 1000)
		c.JSON(200, gin.H{
			"message": "登録成功",
		})
	})
	r.GET("/404", func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "404やで",
		})
	})
	r.Run(":8080")
}

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

type User struct {
	gorm.Model
	NickName string `json:"nickName"`
	Description string `json:"Description"`
}

type Blog struct {
	gorm.Model
	Title string `json:"Title"`
}

// Array of User
type Users []User

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Blog{})
	db.AutoMigrate(&Product{})
	return db
}

func create(db *gorm.DB, code string, price uint) *gorm.DB {
	db.Create(&Product{
		Code: code,
		Price: price,
	})
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
