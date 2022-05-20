package models

import "gorm.io/gorm"

type Ciudad struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	RegionID uint64 `json:"RegionID"`
	Region Region `json:"Region"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
}
