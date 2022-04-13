package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Employee struct {
	emp_no     int
	birth_date time.Time
	first_name string
	last_name  string
	gender     byte
	hire_date  time.Time
}

func main() {
	dsn := "root:16/04/2002Farouk@tcp(127.0.0.1:3306)/employees?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Print("Database connection succeeded")
	// var employee Employee
	db.AutoMigrate(&Employee{})
	var employees []Employee
	result := db.Debug().Find(&employees)
	fmt.Print("Number of affected rows is: ")
	fmt.Print(result.RowsAffected)

	// router := gin.Default()

	// fmt.Print(router.AppEngine)
	// // This handler will match /user/john but will not match /user/ or /user
	// router.GET("/user/:name", func(ctx *gin.Context) {
	// 	name := ctx.Param("name")
	// 	ctx.String(http.StatusOK, "Hello %s", name)
	// })

	// router.GET("/user/:name/*action", func(ctx *gin.Context) {
	// 	name := ctx.Param("name")
	// 	action := ctx.Param("action")
	// 	message := name + " is " + action
	// 	ctx.String(http.StatusOK, message)
	// })

	// router.POST("/user/:name/*action", func(ctx *gin.Context) {
	// 	b := ctx.FullPath() == "/user/:name/*action" // true
	// 	ctx.String(http.StatusOK, "%t", b)
	// })

	// router.Run()

}
