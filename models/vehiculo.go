package models

import (
	"github.com/ferromarket/backend/utils"
	"gorm.io/gorm"
)

type Vehiculo struct {
	gorm.Model
	ID                 uint64     `json:"ID" gorm:"primaryKey"`
	RepartidorID       uint64     `json:"RepartidorID" gorm:"unique;not null"`
	Repartidor         Repartidor `json:"Repartidor" gorm:"not null"`
	Patente            string     `json:"Patente" gorm:"unique;not null"`
	Marca              string     `json:"Marca" gorm:"not null"`
	Modelo             string     `json:"Modelo" gorm:"not null"`
	Ano                uint16     `json:"Ano" gorm:"not null"`
	PermisoCirculacion utils.Date `json:"PermisoCirculacion" gorm:"not null"`
	Seguro             utils.Date `json:"Seguro" gorm:"not null"`
}
