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
		return errors.New("código está vacío")
	}
	if (len(region.Codigo) > 4) {
		return errors.New("código está demasiado grande")
	}
	if region.ID <= 0 {
		return errors.New("ID invalido")
	}
	if region.PaisID <= 0 {
		return errors.New("PaisID invalido")
	}
	err := region.Pais.Validate()
	if err != nil {
		return err
	}
	return nil
}
