package services

import "github.com/cmani097/emp-service/models/entity"

type EmployeeService interface {
	CreateEmployee(employee *entity.Employee) error
	FindEmployeeById(id string) (*entity.Employee, error)
	FindAll() ([]*entity.Employee, error)
	UpdateEmployee(employee *entity.Employee) error
	DeleteEmployee(id string) error
}
