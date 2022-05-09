package models

type Ferreteria struct {
	ID uint `json:"" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique"`
}
