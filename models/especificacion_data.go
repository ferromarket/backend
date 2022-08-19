package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type EspecificacionData struct {
	gorm.Model
	ID    uint64 `json:"ID" gorm:"primaryKey"`
	Valor string `json:"Valor" gorm:"not null"`
}

func (EspecificacionData *EspecificacionData) Validate() error {
	if len(strings.TrimSpace(EspecificacionData.Valor)) == 0 {
		return errors.New("especificacion vacia")
	}
	return nil
}
