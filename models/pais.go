package models

import "gorm.io/gorm"

type Pais struct {
	gorm.Model `json:"-"`
	ID uint64 `json:"ID" gorm:"primaryKey"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
	Codigo string `json:"Codigo" gorm:"unique;size:2;not null"`
}
