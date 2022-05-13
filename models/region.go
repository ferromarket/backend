package models

import "gorm.io/gorm"

type Region struct {
	gorm.Model
	ID uint64 `json:"id" gorm:"primaryKey"`
	PaisID uint64 `json:"pais_id"`
	Pais Pais `json:"pais"`
	Nombre string `json:"nombre" gorm:"unique;not null"`
	Codigo string `json:"codigo" gorm:"unique;size:4;not null"`
}
