package models

import "gorm.io/gorm"

type Pais struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`
	Nombre string `json:"nombre" gorm:"unique;not null"`
	Codigo string `json:"codigo" gorm:"unique;size:2;not null"`
}
