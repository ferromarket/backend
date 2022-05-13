package models

import "gorm.io/gorm"

type Dia struct {
	gorm.Model
	ID uint64 `json:"id" gorm:"primaryKey"`
	Nombre string `json:"nombre" gorm:"unique;size:9;not null"`
}
