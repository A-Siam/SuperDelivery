package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Name  string
	Owner string
	Price float32
}
