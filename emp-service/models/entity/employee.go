package entity

type Employee struct {
	EmployeeID  uint   `json:"employee_id" gorm:"primary_key"`
	FirstName   string `json:"first_name" gorm:"type:varchar(32)"`
	LastName    string `json:"last_name" gorm:"type:varchar(32)"`
	Email       string `json:"email" gorm:"type:varchar(26)"`
	PhoneNumber string `json:"pone_number" gorm:"type:varchar(13)"`
	Salary      uint   `json:"salary"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
	// DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	AddressID uint    `json:"address_id"`
	Address   Address `gorm:"foreignKey:AddressID;references:AddressID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Teams     []Team  `json:"teams" gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
