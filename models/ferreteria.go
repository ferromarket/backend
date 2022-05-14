package models

import "gorm.io/gorm"

type Ferreteria struct {
	gorm.Model `json:"-"`
	ID uint64 `json:"ID" gorm:"primaryKey"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
	ComunaID uint64 `json:"-"`
	Comuna Comuna `json:"Comuna"`
	Direccion string `json:"Direccion"`
	Descripcion string `json:"Descripcion"`
	Dias []Dia `json:"-" gorm:"many2many:ferreteria_horario"`
	Horarios []FerreteriaHorario `json:"Horarios" gorm:"->"`
}
