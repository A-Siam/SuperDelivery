package ports

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitializeEnv() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/payments_service/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error in reading config: %w", err))
	}
	// load development defaults
	viper.SetDefault("db.events.username", "admin")
	viper.SetDefault("db.events.password", "admin")
	viper.SetDefault("db.events.host", "localhost")
	viper.SetDefault("db.events.port", "5432")
	viper.SetDefault("db.events.name", "events")
	viper.SetDefault("mq.broker.host", "localhost")
	viper.SetDefault("mq.broker.port", "9092")
}
