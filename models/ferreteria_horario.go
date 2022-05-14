package models

import "gorm.io/gorm"

type FerreteriaHorario struct {
	gorm.Model `json:"-"`
	FerreteriaID uint64 `json:"-" gorm:"primaryKey;autoIncrement:false"`
	DiaID uint64 `json:"-" gorm:"primaryKey;autoIncrement:false"`
	Dia Dia `json:"Dia"`
	AbrirID uint64 `json:"-"`
	Abrir Hora `json:"Abrir" gorm:"foreignKey:AbrirID"`
	CerrarID uint64 `json:"-"`
	Cerrar Hora `json:"Cerrar" gorm:"foreignKey:CerrarID"`
}
