package controllers

import (
	http_utils "github.com/lavinas-science/learn-utils-go/http"
	"net/http"
)

const (
	pong = "pong"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingController struct{}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	http_utils.RespondJson(w, http.StatusCreated, pong)
}
