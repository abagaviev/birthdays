package models

type Birthday struct {
	ID          uint   `json:"id" gorm:"primary_key`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	DateOfBirth string `json:"dateOfBirth"`
}
