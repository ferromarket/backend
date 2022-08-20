package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Hora struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	Hora string `json:"Hora" gorm:"unique;size:5;not null"`
}

func (hora *Hora) Validate() error {
	if (len(strings.TrimSpace(hora.Hora)) == 0) {
		return errors.New("nombre vacío")
	}
	if (len(hora.Hora) > 5) {
		return errors.New("hora está demasiado largo")
	}
	return nil
}
