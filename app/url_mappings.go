package app

import (
	"net/http"
	"github.com/lavinas-science/learn-items-api/controllers"
)

func mapUrls() {
	r.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	r.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}