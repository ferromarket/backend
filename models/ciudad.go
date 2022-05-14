package models

import "gorm.io/gorm"

type Ciudad struct {
	gorm.Model `json:"-"`
	ID uint64 `json:"ID" gorm:"primaryKey"`
	RegionID uint64 `json:"-"`
	Region Region `json:"Region"`
	Nombre string `json:"Nombre" gorm:"unique;not null"`
}
