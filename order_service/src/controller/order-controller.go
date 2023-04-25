package controller

import (
	"net/http"

	"github.com/A-Siam/super_delivery/order_service/src/config"
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
	panic("unimplemented")
}

func createOrderHandler(j http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func revertOrderHandler(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
