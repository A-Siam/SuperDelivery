package ports

import (
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func CreateSagaEventProducer() (*sarama.SyncProducer, error) {
	sagaQueueHost := viper.GetString("mq.broker.host" + ":" + "mq.broker.port")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{sagaQueueHost}, config)
	if err != nil {
		return nil, err
	}
	return &producer, err
}
