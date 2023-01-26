package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string `json:"code"`
	Price uint   `json:"price"`
}

type User struct {
	gorm.Model
	NickName    string `json:"nickName"`
	Description string `json:"Description"`
}

type Blog struct {
	gorm.Model
	Title string `json:"Title"`
}

func main() {
	db := DBMigrate(DBConnect())

	r := gin.Default()

	r.GET("/product/:id", func(c *gin.Context) {
		var product Product
		db.First(&product, c.Param("id")).Scan(&product)
		c.JSON(200, gin.H{
			"id":    product.ID,
			"price": product.Price,
			"code":  product.Code,
		})
	})
	r.POST("/create", func(c *gin.Context) {
		var product Product
		c.BindJSON(&product)
		db.Create(&product)
		c.JSON(200, gin.H{
			"message": "登録成功!",
			"price":   product.Price,
			"code":    product.Code,
		})
	})
	r.GET("/404", func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "404やで",
		})
	})
	r.Run(":8080")
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Blog{})
	db.AutoMigrate(&Product{})
	return db
}

func create(db *gorm.DB, code string, price uint) *gorm.DB {
	db.Create(&Product{
		Code:  code,
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
