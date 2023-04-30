package ports

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var orderConnection *gorm.DB
var evnetsConnection *gorm.DB

func InitEventsDBConnection() (db *gorm.DB, err error) {
	if evnetsConnection != nil {
		return evnetsConnection, nil
	}
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
	eventsDBHost := viper.Get("db.events.host")
	eventsDBPort := viper.GetString("db.events.port")
	eventsDBUsername := viper.GetString("db.events.username")
	eventsDBPassword := viper.GetString("db.events.password")
	eventsDBName := viper.GetString("db.events.name")
	if eventsDBHost == "" ||
		eventsDBUsername == "" ||
		eventsDBPassword == "" ||
		eventsDBName == "" ||
		eventsDBPort == "" {
		return "", &DBConnectionError{Message: "one of the connection environment vairable is not set or doesn't have a proper value"}
	}
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		eventsDBHost, eventsDBUsername, eventsDBPassword, eventsDBName, eventsDBPort,
	), nil
}
