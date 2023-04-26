package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/A-Siam/super_delivery/order_service/src/data/db"
	"github.com/A-Siam/super_delivery/order_service/src/models"
	"github.com/A-Siam/super_delivery/order_service/src/saga"
	"github.com/Shopify/sarama"
)

func CreateOrder(order models.Order) (models.Order, error) {
	ordersDB, err := db.InitOrdersDBConnection()
	eventsDB, err := db.InitEventsDBConnection()
	if err != nil {
		return models.Order{}, err
	}

	otx := ordersDB.Create(&order)
	event := models.Event{
		EventName: fmt.Sprintf("ORDER_CREATED_%d", order.ID),
		Service:   "ORDER_SERVICE",
	}
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

func GetAllOrders() ([]*models.Order, error) {
	var orders []*models.Order
	ordersDB, err := db.InitOrdersDBConnection()
	eventsDB, err := db.InitEventsDBConnection()
	if err != nil {
		return make([]*models.Order, 0), err
	}
	ordersDB.Find(&orders)
	for _, order := range orders {
		eventsDB.Create(models.Event{
			EventName: fmt.Sprintf("FIND_OPRDER_%d", (*order).ID),
			Service:   "ORDER_SERVICE",
		})
	}
	//  no need to call saga as there is no transaction
	return orders, nil
}

func RollbackOrderCreation(eventName string) error {
	tokens := strings.Split(eventName, "_")
	orderId, err := strconv.ParseUint(tokens[len(tokens)-1], 10, 64)
	ordersDB, err := db.InitOrdersDBConnection()
	eventsDB, err := db.InitEventsDBConnection()
	if err != nil {
		return err
	}
	ordersDB.Delete(&models.Order{}, uint(orderId))
	eventsDB.Create(models.Event{
		EventName: fmt.Sprintf("DELETE_ORDER_%d", orderId),
		Service:   "ORDER_SERVICE",
	})
	return nil
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
