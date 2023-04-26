package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/A-Siam/super_delivery/order_service/src/config"
	"github.com/A-Siam/super_delivery/order_service/src/models"
	"github.com/A-Siam/super_delivery/order_service/src/service"
)

func StartOrderController(rootMux *http.ServeMux) {
	rootMux.HandleFunc(config.API_V1+"/orders", ordersHandler)
	rootMux.HandleFunc(config.API_V1+"/orders/revert", revertOrderHandler)
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		createOrderHandler(w, r)
	} else if r.Method == http.MethodGet {
		getOrdersHandler(w, r)
	} else {
		methodNotAllowedHandler(w, r)
	}
}

func getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := service.GetAllOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ordersJson, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(ordersJson)
}

func createOrderHandler(j http.ResponseWriter, r *http.Request) {
	order := &models.Order{}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		j.WriteHeader(http.StatusBadRequest)
		j.Write([]byte(fmt.Sprintf("reading error: %s", err.Error())))
		return
	}
	err = json.Unmarshal(data, order)
	if err != nil {
		j.WriteHeader(http.StatusBadRequest)
		j.Write([]byte(fmt.Sprintf("converting error: %s", err.Error())))
		return
	}
	// TODO
	storedOrder, err := service.CreateOrder(*order)
	if err != nil {
		http.Error(j, fmt.Sprintf("error while saving order: %s", err.Error()), http.StatusBadRequest)
		return
	}
	storedOrderJson, err := json.Marshal(storedOrder)
	if err != nil {
		http.Error(j, fmt.Sprintf("serialization error %s", err.Error()), http.StatusInternalServerError)
	}
	j.WriteHeader(http.StatusCreated)
	j.Write(storedOrderJson)
}

func revertOrderHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var eventDetails *models.Event
	err = json.Unmarshal(bodyBytes, eventDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.RollbackOrderCreation(eventDetails.EventName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte{})
}
