package services

import (
	"context"

	"github.com/cmani097/emp-service/models/entity"
	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
)

type EmployeeServiceImpl struct {
	db  *gorm.DB
	ctx context.Context
}

func NewEmployeeService(db *gorm.DB, ctx context.Context) EmployeeService {
	return &EmployeeServiceImpl{
		db:  db,
		ctx: ctx,
	}
}

func (employeeservice *EmployeeServiceImpl) CreateEmployee(employee *entity.Employee) error {
	result := employeeservice.db.Create(&employee)
	if result.Error != nil {
		// log.Err(conErr).Str("request_id", request_id).Msg("Error occurred while getting a DB connection from the connection pool")
		return result.Error
	}
	return nil
}

func (employeeService *EmployeeServiceImpl) FindEmployeeById(id string) (*entity.Employee, error) {
	var employee *entity.Employee
	result := employeeService.db.Where("employee_id = ?", id).First(&employee)
	if result.Error != nil {
		return nil, result.Error
	}
	return employee, nil
}

func (employeeService *EmployeeServiceImpl) FindAll() ([]*entity.Employee, error) {
	var employees []*entity.Employee
	err := employeeService.db.Debug().Preload("Teams").Find(&employees)
	if err.Error != nil {
		return nil, err.Error
	}
	return employees, nil
}

func (employeeService *EmployeeServiceImpl) UpdateEmployee(employee *entity.Employee) error {
	result := employeeService.db.Update(employee)
	if result.Error != nil {
		return result.Error
	}
	// var emp *entity.Employee
	// employeeService.db.Debug().Where("employee_id = ?", employee.EmployeeID).Find(&emp)
	return nil
}

func (employeeService *EmployeeServiceImpl) DeleteEmployee(id string) error {
	// result := employeeService.db.Debug().Where("employee_id = ?", 1).Delete()
	return nil
}
