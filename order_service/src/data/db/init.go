package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/** init db connection */
func InitOrdersDBConnection() (db *gorm.DB, err error) {
	conStr, err := getOrdersConnectionString()
	db, err = initDBConnection(conStr)
	return
}

func InitEventsDBConnection() (db *gorm.DB, err error) {
	conStr, err := getEventsDBConnectionString()
	db, err = initDBConnection(conStr)
	return
}

func initDBConnection(connectionString string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	return
}

/** getting connection strings */
func getEventsDBConnectionString() (string, error) {
	eventsDBHost := os.Getenv("EVENTS_DB_HOST")
	eventsDBUsername := os.Getenv("EVENTS_DB_USERNAME")
	eventsDBPassword := os.Getenv("EVENTS_DB_PASSWORD")
	eventsDBName := os.Getenv("EVENTS_DB_DB_NAME")
	if eventsDBHost == "" || eventsDBUsername == "" || eventsDBPassword == "" || eventsDBName == "" {
		return "", &DBConnectionError{message: "one of the connection environment vairable is not set or doesn't have a proper value"}
	}
	return "", nil
}

func getOrdersConnectionString() (string, error) {
	ordersDBHost := os.Getenv("ORDERS_DB_HOST")
	ordersDBUsername := os.Getenv("ORDERS_DB_USERNAME")
	ordersDBPassword := os.Getenv("ORDERS_DB_PASSWORD")
	ordersDBName := os.Getenv("ORDERS_DB_DB_NAME")
	if ordersDBHost == "" || ordersDBUsername == "" || ordersDBPassword == "" || ordersDBName == "" {
		return "", &DBConnectionError{message: "one of the connection environment vairable is not set or doesn't have a proper value"}
	}
	return "", nil
}
