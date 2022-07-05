package models

import "gorm.io/gorm"

type FerreteriaHorario struct {
	gorm.Model
	FerreteriaID uint64 `json:"FerreteriaID"`
	Ferreteria Ferreteria `json:"Ferreteria"`
	DiaID uint64 `json:"DiaID"`
	Dia Dia `json:"Dia"`
	AbrirID uint64 `json:"AbrirID"`
	Abrir Hora `json:"Abrir" gorm:"foreignKey:AbrirID"`
	CerrarID uint64 `json:"CerrarID"`
	Cerrar Hora `json:"Cerrar" gorm:"foreignKey:CerrarID"`
}
