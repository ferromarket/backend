package models

import "gorm.io/gorm"

type Dia struct {
	gorm.Model `json:"-"`
	ID uint64 `json:"ID" gorm:"primaryKey"`
	Nombre string `json:"Nombre" gorm:"unique;size:9;not null"`
	Ferreterias []Ferreteria `json:"-" gorm:"many2many:ferreteria_horario"`
}
