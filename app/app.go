package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lavinas-science/learn-items-api/clients/elasticsearch"
)

var (
	r = mux.NewRouter()
)

func StartApp() {
	elasticsearch.Init()
	mapUrls()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
