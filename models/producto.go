package models

import "gorm.io/gorm"

type Producto struct {
	gorm.Model
	ID          uint64    `json:"ID" gorm:"primaryKey"`
	CategoriaID uint64    `json:"CategoriaID" gorm:"not null"`
	Categoria   Categoria `json:"Categoria"`
	Nombre      string    `json:"Producto" gorm:"not null"`
}
