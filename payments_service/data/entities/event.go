package entities

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	EventName string
	Service   string
}
