package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type EspecificacionNombre struct {
	gorm.Model
	ID     uint64 `json:"ID" gorm:"primaryKey"`
	Nombre string `json:"Nombre" gorm:"not null"`
}

func (EspecificacionNombre *EspecificacionNombre) Validate() error {
	if len(strings.TrimSpace(EspecificacionNombre.Nombre)) == 0 {
		return errors.New("especificacion nombre")
	}
	return nil
}
