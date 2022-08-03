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
	ApellidoMaterno   string `json:"ApellidoMaterno" gorm:"not null"`
	Telefono          uint32 `json:"Telefono"`
	Direccion         string `json:"Direccion" gorm:"not null"`
	FechaNacimiento   utils.DateTime
	FechaRegistracion utils.DateTime `json:"FechaRegistracion"`
	TipoLicencia      uint8          `json:"TipoLicencia"`
	FechaLicencia     utils.DateTime `json:"FechaLicencia"`
}
