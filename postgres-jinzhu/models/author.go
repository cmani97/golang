package models

type Author struct {
	ID       uint `json:"author_id" gorm:"primary_key"`
	Name     string
	Degree   string
	Contacts []Contact `gorm:"foreignKey:AuthID;references:AuthID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
