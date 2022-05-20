package models

import "gorm.io/gorm"

type Comuna struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	CiudadID uint64 `json:"CiudadID"`
	Ciudad Ciudad `json:"Ciudad"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
}
