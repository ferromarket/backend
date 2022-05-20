package models

import (
	"time"

	"gorm.io/gorm"
)

type Repartidor struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	UsuarioID uint64 `json:"UsuarioID"`
	Usuario Usuario `json:"Usuario"`
	FechaRegistracion time.Time `json:"FechaRegistracion"`
	TipoLicencia uint8 `json:"TipoLicencia"`
	FechaLicencia time.Time `json:"FechaLicencia"`
}
