package models

import "gorm.io/gorm"

type Region struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	PaisID uint64 `json:"PaisID"`
	Pais Pais `json:"Pais"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
	Codigo string `json:"Codigo" gorm:"unique;size:4;not null"`
}
