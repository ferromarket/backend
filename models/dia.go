package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Dia struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	Nombre string `json:"Nombre" gorm:"unique;size:9;not null"`
}

func (dia *Dia) Validate() error {
	if (len(strings.TrimSpace(dia.Nombre)) == 0) {
		return errors.New("nombre vacío")
	}
	if (len(dia.Nombre) > 9) {
		return errors.New("nombre está demasiado largo")
	}
	if dia.ID <= 0 {
		return errors.New("ID invalido")
	}
	return nil
}
