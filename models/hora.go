package models

import "gorm.io/gorm"

type Hora struct {
	gorm.Model
	ID uint64 `json:"ID" gorm:"primaryKey"`
	Hora string `json:"Hora" gorm:"unique;size:5;not null"`
}
