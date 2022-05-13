package models

import "gorm.io/gorm"

type Hora struct {
	gorm.Model
	ID uint64 `json:"id" gorm:"primaryKey"`
	Hora string `json:"hora" gorm:"unique;size:5;not null"`
}
