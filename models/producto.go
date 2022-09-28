package models

import "gorm.io/gorm"

type Producto struct {
	gorm.Model
	ID             uint64           `json:"ID" gorm:"primaryKey"`
	CategoriaID    *uint64          `json:"CategoriaID" gorm:"not null"`
	Categoria      *Categoria       `json:"Categoria"`
	Nombre         string           `json:"Nombre" gorm:"not null"`
	Valor1         string           `json:"Valor1" gorm:"not null"`
	Valor2         string           `json:"Valor2" gorm:"not null"`
	Especificacion []Especificacion `json:"Especificacion"`
}
