package models

import "gorm.io/gorm"

type Rol struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
	Descripcion string `json:"Descripcion" gorm:"unique;not null"`
	Usuarios []Usuario `json:"-" gorm:"many2many:usuario_rol"`
}
