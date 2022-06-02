package models

import "gorm.io/gorm"

type EspecificacionNombre struct {
	gorm.Model
	ID     uint64 `json:"ID" gorm:"primaryKey"`
	Nombre string `json:"Nombre"`
}
