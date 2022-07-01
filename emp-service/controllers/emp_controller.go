package controllers

import (
	"fmt"
	"net/http"

	"github.com/cmani097/emp-service/models/dto"
	"github.com/cmani097/emp-service/models/entity"
	"github.com/cmani097/emp-service/services"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeservice services.EmployeeService
}

func NewEmployeeController(employeeservice services.EmployeeService) EmployeeController {
	return EmployeeController{
		employeeservice: employeeservice,
	}
}

func (employeecontroller *EmployeeController) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "success"})
}

func (employeecontroller *EmployeeController) NewEmployee(ctx *gin.Context) {
	fmt.Println("NewEmployee called")
	var employeeInput dto.EmployeeDto
	if err := ctx.ShouldBindJSON(&employeeInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	employee := entity.Employee{FirstName: employeeInput.FirstName, LastName: employeeInput.LastName,
		Email: employeeInput.Email, PhoneNumber: employeeInput.PhoneNumber, Salary: employeeInput.Salary,
		Teams: employeeInput.Teams, Address: employeeInput.Address}

	err := employeecontroller.employeeservice.CreateEmployee(&employee)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (employeecontroller *EmployeeController) FindEmployeeById(ctx *gin.Context) {
	// id := ctx.Param("employee_id")
	employee, err := employeecontroller.employeeservice.FindEmployeeById("1")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": employee})
}

func (employeecontroller *EmployeeController) FindAll(ctx *gin.Context) {
	employee, err := employeecontroller.employeeservice.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": employee})
}
