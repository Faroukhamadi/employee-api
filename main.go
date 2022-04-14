package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	Emp_no     int
	Birth_date time.Time
	First_name string
	Last_name  string
	Gender     byte
	Hire_date  time.Time
}

type Dept_emp struct {
	Emp_no    int
	Dept_no   int
	From_date time.Time
	To_date   time.Time
}

// taha hssin

type Department struct {
	Dept_no   int
	Dept_name string
}

type Result struct {
	Emp_no    int
	Dept_name string
}

func main() {
	dsn := "root:16/04/2002Farouk@tcp(127.0.0.1:3306)/employees?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Employee{})
	router := gin.Default()

	router.GET("/employees/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var employee Employee
		db.First(&employee, "emp_no = ?", id)
		ctx.JSON(200, gin.H{
			"employee": employee,
		})
	})

	router.GET("/employees/first_name/:first_name", func(ctx *gin.Context) {
		firstName := ctx.Param("first_name")
		var employee Employee
		db.First(&employee, "first_name = ?", firstName)
		ctx.JSON(200, gin.H{
			"employee": employee,
		})
	})

	router.GET("/employees/last_name/:last_name", func(ctx *gin.Context) {
		lastName := ctx.Param("last_name")
		var employee Employee
		db.First(&employee, "last_name = ?", lastName)
		ctx.JSON(200, gin.H{
			"employee": employee,
		})
	})

	// TODO: Implement the count for this later
	router.GET("/employees/department/:dep/count", func(ctx *gin.Context) {
		dep := ctx.Param("dep")
		var result []Result
		rows := db.Model(&Department{}).Select("departments.dept_name, dept_emp.emp_no").Joins("left join dept_emp on dept_emp.dept_no = departments.dept_no").Where("dept_name = ?", dep).Scan(&result)
		ctx.JSON(200, gin.H{
			"count_emp_" + dep: rows.RowsAffected,
		})
	})

	router.GET("/employees/count", func(ctx *gin.Context) {
		var employees []Employee
		result := db.Find(&employees)
		ctx.JSON(200, gin.H{
			"count": result.RowsAffected,
		})
	})

	router.GET("/employees/count/male", func(ctx *gin.Context) {
		var employees []Employee
		result := db.Find(&employees, "Gender = ?", "1")
		ctx.JSON(200, gin.H{
			"male_count": result.RowsAffected,
		})
	})

	router.GET("/employees/count/female", func(ctx *gin.Context) {
		var employees []Employee
		result := db.Find(&employees, "Gender = ?", "2")
		ctx.JSON(200, gin.H{
			"male_count": result.RowsAffected,
		})
	})

	router.Run()

}
