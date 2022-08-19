package models

import (
	"errors"

	"gorm.io/gorm"
)

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

func (ferreteria_horario *FerreteriaHorario) Validate() error {
	// Don't try to validate the Ferreteria object, most likely would end up in an infinite loop
	if ferreteria_horario.FerreteriaID <= 0 {
		return errors.New("FerreteriaID invalido")
	}
	if ferreteria_horario.DiaID <= 0 {
		return errors.New("DiaID invalido")
	}
	if ferreteria_horario.AbrirID <= 0 {
		return errors.New("AbrirID invalido")
	}
	if ferreteria_horario.CerrarID <= 0 {
		return errors.New("CerrarID invalido")
	}
	err := ferreteria_horario.Dia.Validate()
	if err != nil {
		return err
	}
	err = ferreteria_horario.Abrir.Validate()
	if err != nil {
		return err
	}
	err = ferreteria_horario.Cerrar.Validate()
	if err != nil {
		return err
	}
	return nil
}
