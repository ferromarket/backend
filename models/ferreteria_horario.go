package models

import "gorm.io/gorm"

type FerreteriaHorario struct {
	gorm.Model
	FerreteriaID uint64 `json:"FerreteriaID" gorm:"primaryKey"`
	DiaID uint64 `json:"DiaID" gorm:"primaryKey"`
	AbrirID uint64 `json:"AbrirID"`
	Abrir Hora `json:"Abrir"`
	CerrarID uint64 `json:"CerrarID"`
	Cerrar Hora `json:"Cerrar"`
}
