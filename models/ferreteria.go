package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Ferreteria struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
	ComunaID uint64 `json:"ComunaID"`
	Comuna Comuna `json:"Comuna"`
	Direccion string `json:"Direccion"`
	Descripcion string `json:"Descripcion"`
	Dias []Dia `json:"-" gorm:"many2many:ferreteria_horario"`
	Horarios []FerreteriaHorario `json:"Horarios" gorm:"->"`
}

func (ferreteria *Ferreteria) Validate() error {
	if (len(strings.TrimSpace(ferreteria.Nombre)) == 0) {
		return errors.New("nombre vacío")
	}
	if (len(strings.TrimSpace(ferreteria.Direccion)) == 0) {
		return errors.New("dirección vacío")
	}
	if (len(strings.TrimSpace(ferreteria.Descripcion)) == 0) {
		return errors.New("descripción vacío")
	}
	return nil
}
