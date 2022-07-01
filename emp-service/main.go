package main

import (
	"context"

	"github.com/cmani097/emp-service/cmd"
	"github.com/cmani097/emp-service/controllers"
	"github.com/cmani097/emp-service/services"
	"github.com/gin-gonic/gin"

	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
)

var (
	server             *gin.Engine
	database           *gorm.DB
	err                error
	ctx                context.Context
	employeecontroller controllers.EmployeeController
	employeeservice    services.EmployeeService
)

func init() {
	server = gin.Default()
	database = cmd.ConnectToPostgres()

	employeeservice = services.NewEmployeeService(database, ctx)
	employeecontroller = controllers.NewEmployeeController(employeeservice)
}

func main() {
	server.Use(func(c *gin.Context) {
		c.Set("db", database)
		c.Next()
	})

	server.GET("/test", employeecontroller.Test)
	server.POST("/employee", employeecontroller.NewEmployee)
	server.GET("/employee/:id", employeecontroller.FindEmployeeById)
	server.GET("/employees", employeecontroller.FindAll)

	server.Run(":8083")
}
