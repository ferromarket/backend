package models

import "gorm.io/gorm"

type FerreteriaHorario struct {
	gorm.Model
	FerreteriaID uint64 `json:"ferreteria_id" gorm:"primaryKey"`
	DiaID uint64 `json:"dia_id" gorm:"primaryKey"`
	AbrirID uint64 `json:"abrir_id"`
	Abrir Hora `json:"abrir"`
	CerrarID uint64 `json:"cerrar_id"`
	Cerrar Hora `json:"cerrar"`
}
