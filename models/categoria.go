package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Categoria struct {
	gorm.Model
	ID          uint64     `json:"ID" gorm:"primaryKey"`
	CategoriaID *uint64    `json:"CategoriaID"`
	Categoria   *Categoria `json:"Categoria"`
	Nombre      string     `json:"Nombre" gorm:"not null"`
}

func (Categoria *Categoria) Validate() error {
	if len(strings.TrimSpace(Categoria.Nombre)) == 0 {
		return errors.New("categoria nombre vacia")
	}
	return nil
}
