package service

import (
	"encoding/json"
	"os"

	"github.com/A-Siam/super_delivery/order_service/src/data/db"
	"github.com/A-Siam/super_delivery/order_service/src/models"
	"github.com/A-Siam/super_delivery/order_service/src/saga"
	"github.com/Shopify/sarama"
)

func CreateOrder(order models.Order) (models.Order, error) {
	event := models.Event{
		EventName: "ORDER_CREATED",
		Service:   "ORDER_SERVICE",
	}
	ordersDB, err := db.InitOrdersDBConnection()
	eventsDB, err := db.InitEventsDBConnection()
	if err != nil {
		return models.Order{}, err
	}

	otx := ordersDB.Create(&order)
	etx := eventsDB.Create(&event)
	if otx.Error != nil {
		return models.Order{}, otx.Error
	}
	if etx.Error != nil {
		return models.Order{}, etx.Error
	}
	produceOrderCreatedEvent(event)
	return order, nil
}

func produceOrderCreatedEvent(event models.Event) error {
	msgBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}
	producerPtr, err := saga.CreateSagaEventProducer()
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: os.Getenv("EVENTS_TOPIC_NAME_URL"),
		Value: sarama.StringEncoder(string(msgBytes)),
	}
	producer := *producerPtr
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
