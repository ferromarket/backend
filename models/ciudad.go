package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Ciudad struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	RegionID uint64 `json:"RegionID"`
	Region Region `json:"Region"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
}

func (ciudad *Ciudad) Validate() error {
	if (len(strings.TrimSpace(ciudad.Nombre)) == 0) {
		return errors.New("nombre vacío")
	}
	if ciudad.ID <= 0 {
		return errors.New("ID invalido")
	}
	if ciudad.RegionID <= 0 {
		return errors.New("RegionID invalido")
	}
	err := ciudad.Region.Validate()
	if err != nil {
		return err
	}
	return nil
}
