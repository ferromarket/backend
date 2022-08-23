package models

import (
	"time"

	"gorm.io/gorm"
)

type FavProd struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	IDProd uint64 `json:"IDProd" gorm:"not null"`
	IDUsuario uint64 `json:"IDUsuario" gorm:"not null"`
	Usuario Usuario `json:"Usuario" gorm:"not null"`
	NombreProd string `json:"NombreProd" gorm:"not null"`
}
