package entity

type Address struct {
	ID         uint   `json:"address_id" gorm:"primary_key"`
	HouseNo    string `json:"house_no"`
	Village    string `json:"village"`
	State      string `json:"state"`
	PostalCode int    `json:"postal_code"`
}
