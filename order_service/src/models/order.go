package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Price float32 `json:"price"`
}
