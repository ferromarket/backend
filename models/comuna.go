package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Comuna struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	CiudadID uint64 `json:"CiudadID"`
	Ciudad Ciudad `json:"Ciudad"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
}

func (comuna *Comuna) Validate() error {
	if (len(strings.TrimSpace(comuna.Nombre)) == 0) {
		return errors.New("nombre vacío")
	}
	if comuna.ID <= 0 {
		return errors.New("ID invalido")
	}
	if comuna.CiudadID <= 0 {
		return errors.New("CiudadID invalido")
	}
	err := comuna.Ciudad.Validate()
	if err != nil {
		return err
	}
	return nil
}
