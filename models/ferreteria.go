package models

import "gorm.io/gorm"

type Ferreteria struct {
	gorm.Model
	ID uint64 `json:"id" gorm:"primaryKey"`
	Nombre string `json:"nombre" gorm:"unique;not null"`
	ComunaID uint64 `json:"comuna_id"`
	Comuna Comuna `json:"comuna"`
	Direccion string `json:"direccion"`
	Descripcion string `json:"descripcion"`
	Horarios []Dia `json:"horarios" gorm:"many2many:ferreteria_horario"`
}
