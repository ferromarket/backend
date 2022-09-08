package models

import (
	"github.com/ferromarket/backend/utils"
	"gorm.io/gorm"
)

type Vehiculo struct {
	gorm.Model
	ID                 uint64     `json:"ID" gorm:"primaryKey"`
	RepartidorID       uint64     `json:"RepartidorID"`
	Repartidor         Repartidor `json:"Repartidor"`
	Patente            string     `json:"Patente"`
	Marca              string     `json:"Marca"`
	Modelo             string     `json:"Modelo"`
	Ano                uint16     `json:"Ano"`
	PermisoCirculacion utils.Date `json:"PermisoCirculacion"`
	Seguro             utils.Date `json:"Seguro"`
}
