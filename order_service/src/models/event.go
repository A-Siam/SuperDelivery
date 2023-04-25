package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	EventName string `json:"event_name"`
	Service   string `json:"service"`
}
