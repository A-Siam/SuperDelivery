package service

import (
	"github.com/A-Siam/super_delivery/order_service/src/data/db"
	"github.com/A-Siam/super_delivery/order_service/src/models"
)

func CreateOrder(orderName, owner string, price float32) (models.Order, error) {
	order := models.Order{
		Owner: owner,
		Name:  orderName,
		Price: price,
	}
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

	return order, nil
}
