package models

import (
	"github.com/ferromarket/backend/utils"
	"gorm.io/gorm"
)

type Repartidor struct {
	gorm.Model
	ID                uint64 `json:"ID" gorm:"primaryKey"`
	RUT               string `json:"RUT" gorm:"unique;not null"`
	Contrasena        string `json:"Contrasena" gorm:"not null"`
	Email             string `json:"Email" gorm:"not null"`
	Nombres           string `json:"Nombres" gorm:"not null"`
	ApellidoPaterno   string `json:"ApellidoPaterno" gorm:"not null"`
	ApellidoMaterno   string `json:"ApellidoMaterno"`
	Telefono          uint32 `json:"Telefono" gorm:"not null"`
	Direccion         string `json:"Direccion" gorm:"not null"`
	FechaNacimiento   utils.Date `json:"FechaNacimiento" gorm:"not null"`
	FechaRegistracion utils.Date `json:"FechaRegistracion" gorm:"not null"`
	TipoLicencia      string          `json:"TipoLicencia" gorm:"not null"`
	FechaLicencia     utils.Date`json:"FechaLicencia" gorm:"not null"`
}
