package models

import "gorm.io/gorm"

type Ciudad struct {
	gorm.Model
	ID uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	RegionID uint64 `json:"region_id"`
	Region Region `json:"region"`
	Nombre string `json:"nombre" gorm:"unique;not null"`
}
