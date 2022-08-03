package models

import (
	"github.com/ferromarket/backend/utils"
	"gorm.io/gorm"
)

type Repartidor struct {
	gorm.Model
	ID                uint64 `json:"ID" gorm:"primaryKey"`
	RUT               string `json:"RUT" gorm:"unique;size:9;not null"`
	Contrasena        string `json:"Contrasena" gorm:"not null"`
	Email             string `json:"Email" gorm:"not null"`
	Nombres           string `json:"Nombres" gorm:"not null"`
	ApellidoPaterno   string `json:"ApellidoPaterno" gorm:"not null"`
	ApellidoMaterno   string `json:"ApellidoMaterno"`
	Telefono          uint32 `json:"Telefono" gorm:"not null"`
	Direccion         string `json:"Direccion" gorm:"not null"`
	FechaNacimiento   utils.DateTime `json:"FechaNacimiento" gorm:"not null"`
	FechaRegistracion utils.DateTime `json:"FechaRegistracion" gorm:"not null"`
	TipoLicencia      uint8          `json:"TipoLicencia" gorm:"not null"`
	FechaLicencia     utils.DateTime `json:"FechaLicencia" gorm:"not null"`
}
