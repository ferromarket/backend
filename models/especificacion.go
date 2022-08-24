package models

import "gorm.io/gorm"

type Especificacion struct {
	gorm.Model
	ProductoID             uint64               `json:"ProductoID" gorm:"primaryKey;autoIncrement:false"`
	Producto               Producto             `json:"Producto"`
	EspecificacionNombreID uint64               `json:"EspecificacionNombreID" gorm:"primaryKey;autoIncrement:false"`
	EspecificacionNombre   EspecificacionNombre `json:"EspecificacionNombre"`
	EspecificacionDataID   uint64               `json:"EspecificacionDataID" gorm:"primaryKey;autoIncrement:false"`
	EspecificacionData     EspecificacionData   `json:"EspecificacionData"`
}
