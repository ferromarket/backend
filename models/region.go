package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Region struct {
	gorm.Model `json:"-"`
	ID uint64 `json:"ID" gorm:"primaryKey"`
	PaisID uint64 `json:"PaisID"`
	Pais Pais `json:"Pais"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
	Codigo string `json:"Codigo" gorm:"unique;size:4;not null"`
}

func (region *Region) Validate() error {
	if (len(strings.TrimSpace(region.Nombre)) == 0) {
		return errors.New("nombre vacío")
	}
	if (len(strings.TrimSpace(region.Codigo)) == 0) {
		return errors.New("código vacío")
	}
	if (len([]rune(region.Codigo)) > 4) {
		return errors.New("código demasiado grande")
	}
	return nil
}
