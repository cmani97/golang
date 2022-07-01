package dto

import "github.com/cmani097/emp-service/models/entity"

type EmployeeDto struct {
	FirstName   string         `json:"first_name" binding:"required" `
	LastName    string         `json:"last_name" binding:"required"`
	Email       string         `json:"email" binding:"required"`
	PhoneNumber string         `json:"pone_number" binding:"required"`
	Salary      uint           `json:"salary" binding:"required"`
	Teams       []entity.Team  `json:"teams" binding:"required"`
	Address     entity.Address `json:"address" binding:"required"`
}
