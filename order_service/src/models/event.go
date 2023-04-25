package models

import (
	"time"
)

type Event struct {
	EventName string `json:"event_name"`
	Service   string `json:"service"`
	CreatedAt time.Time
}
