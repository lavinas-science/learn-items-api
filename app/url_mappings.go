package app

import (
	"github.com/lavinas-science/learn-items-api/controllers"
	"net/http"
)

func mapUrls() {
	r.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	r.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	r.HandleFunc("items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
}
