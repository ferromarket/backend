package models

import "gorm.io/gorm"

type Dia struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	Nombre string `json:"Nombre" gorm:"unique;size:9;not null"`
}
