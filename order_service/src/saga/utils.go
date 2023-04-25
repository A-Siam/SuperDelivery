package saga

import (
	"github.com/A-Siam/super_delivery/order_service/src/models"
	"github.com/Shopify/sarama"
)

func CreateSagaEventProducer(event models.Event) (*sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return nil, err
	}
	return &producer, err
}
