package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Pais struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
	Codigo string `json:"Codigo" gorm:"unique;size:2;not null"`
}

func (pais *Pais) Validate() error {
	if (len(strings.TrimSpace(pais.Nombre)) == 0) {
		return errors.New("nombre está vacío")
	}
	if (len(strings.TrimSpace(pais.Codigo)) == 0) {
		return errors.New("código está vacío")
	}
	if (len(pais.Codigo) > 2) {
		return errors.New("código está demasiado grande")
	}
	return nil
}
