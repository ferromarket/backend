package models

import "gorm.io/gorm"

type Comuna struct {
	gorm.Model
	ID uint64 `json:"id" gorm:"primaryKey"`
	CiudadID uint64 `json:"ciudad_id"`
	Ciudad Ciudad `json:"ciudad"`
	Nombre string `json:"nombre" gorm:"unique;not null"`
}
