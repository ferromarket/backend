package models

import "gorm.io/gorm"

type EspecificacionData struct {
	gorm.Model
	ID    uint64 `json:"ID" gorm:"primaryKey"`
	Valor string `json:"Valor"`
}
