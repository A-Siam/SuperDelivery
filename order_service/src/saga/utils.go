package saga

import (
	"github.com/A-Siam/super_delivery/order_service/src/models"
	"github.com/A-Siam/super_delivery/order_service/src/utils"
	"github.com/Shopify/sarama"
)

func CreateSagaEventProducer(event models.Event) (*sarama.SyncProducer, error) {
	sagaQueueHost, err := utils.GetEnvWithCheck("KAFKA_BROKER_HOST_URL")
	if err != nil {
		return nil, err
	}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{sagaQueueHost}, config)
	if err != nil {
		return nil, err
	}
	return &producer, err
}
