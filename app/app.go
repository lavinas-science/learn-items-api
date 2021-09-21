package app

import (
	"net/http"
	"time"
	
	"github.com/gorilla/mux"
)

var (
	r = mux.NewRouter()
)

func StartApp () {
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
