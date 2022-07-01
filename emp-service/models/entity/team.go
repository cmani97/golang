package entity

type Team struct {
	TeamID     uint   `json:"team_id" gorm:"primary_key"`
	Name       string `json:"name"`
	EmployeeID uint
}
