package entity

import "time"

type Person struct {
	ID        uint64 `json:"id" gorm:"primary_key:auto_increment"`
	FirstName string `json:"firstName" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"lastName" binding:"required" gorm:"type:varchar(32)"`
	Age       int8   `json:"age" binding:"gte=1,lte=100"`
	Email     string `json:"email" validate:"required,email" gorm:"type:varchar(26)"`
}

type Video struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title       string    `json:"title" binding:"min=2,max=10" validate:"is-cool" gorm:"type:varchar(100)"`
	Description string    `json:"description" binding:"max=20" gorm:"type:varchar(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"type:foreignKey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"CURRENT_TIMESTAMP" json:"updated_at"`
}
