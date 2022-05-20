package models

import "gorm.io/gorm"

type FerreteriaHorario struct {
	gorm.Model
	FerreteriaID uint64 `json:"FerreteriaID" gorm:"primaryKey;autoIncrement:false"`
	DiaID uint64 `json:"DiaID" gorm:"primaryKey;autoIncrement:false"`
	Dia Dia `json:"Dia"`
	AbrirID uint64 `json:"AbrirID"`
	Abrir Hora `json:"Abrir" gorm:"foreignKey:AbrirID"`
	CerrarID uint64 `json:"CerrarID"`
	Cerrar Hora `json:"Cerrar" gorm:"foreignKey:CerrarID"`
}
