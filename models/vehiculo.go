package models

import (
	"time"

	"gorm.io/gorm"
)

type Vehiculo struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	RepartidorID uint64 `json:"RepartidorID"`
	Repartidor Repartidor `json:"Repartidor"`
	Tipo uint8 `json:"Tipo"`
	Marca string `json:"Marca"`
	Modelo string `json:"Modelo"`
	Ano uint16 `json:"Ano"`
	PermisoCirculacion time.Time `json:"PermisoCirculacion"`
	Seguro time.Time `json:"Seguro"`
}
