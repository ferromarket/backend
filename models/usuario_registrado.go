package models

import (
	"errors"
	"strings"

	"github.com/ferromarket/backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	ID              uint64     `json:"ID" gorm:"primaryKey"`
	RUT             string     `json:"RUT" gorm:"unique;size:9;not null"`
	Contrasena      string     `json:"Contrasena" gorm:"not null"`
	Email           string     `json:"Email" gorm:"not null"`
	Nombres         string     `json:"Nombres" gorm:"not null"`
	ApellidoPaterno string     `json:"ApellidoPaterno" gorm:"not null"`
	ApellidoMaterno string     `json:"ApellidoMaterno"`
	Telefono        uint32     `json:"Telefono" gorm:"not null"`
	Direccion       string     `json:"Direccion" gorm:"not null"`
	FechaNacimiento utils.Date `json:"FechaNacimiento" gorm:"not null"`
	Roles           []Rol      `json:"-" gorm:"many2many:usuario_rol"`
}

func (user *Usuario) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Contrasena = string(bytes)
	return nil
}

func (user *Usuario) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Contrasena), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (usuario *Usuario) Validate() error {
	if len(strings.TrimSpace(usuario.Nombres)) == 0 {
		return errors.New("Nombre vacío")
	}
	return nil
}
