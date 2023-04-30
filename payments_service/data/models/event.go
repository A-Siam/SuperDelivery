package models

type Event struct {
	EventName string `json:"event_name"`
	Service   string `json:"service_name"`
}
