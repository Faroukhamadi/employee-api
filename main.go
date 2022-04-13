package main

import (
	// "github.com/gin-gonic/gin"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:16/04/2002Farouk@tcp(127.0.0.1:3306)/employees?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Print("Database connection succeeded")
	db.AutoMigrate(&Product{})

	db.Unscoped().Delete(&Product{}, 4)

	// r := gin.Default()
	// 	r.GET("/ping", func(ctx *gin.Context) {
	// 		ctx.JSON(200, gin.H{
	// 			"message": "pong",
	// 		})
	// 	})
	// 	r.Run()
}
