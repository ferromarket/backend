package models

import "gorm.io/gorm"

type Comuna struct {
	gorm.Model `json:"-"`
	ID uint64 `json:"ID" gorm:"primaryKey"`
	CiudadID uint64 `json:"-"`
	Ciudad Ciudad `json:"Ciudad"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
}
