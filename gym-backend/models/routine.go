package models

import "gorm.io/gorm"

type Routine struct {
	gorm.Model
	Name   string
	UserID uint
}