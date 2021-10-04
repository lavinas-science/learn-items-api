package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lavinas-science/learn-items-api/domain/items"
	"github.com/lavinas-science/learn-items-api/domain/queries"
	"github.com/lavinas-science/learn-items-api/services"
	"github.com/lavinas-science/learn-oauth-go/oauth"
	http_utils "github.com/lavinas-science/learn-utils-go/http"
	"github.com/lavinas-science/learn-utils-go/rest_errors"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}
	if oauth.GetCallerId(r) == 0 {
		rErr := rest_errors.NewUnauthorizedError("not authorized")
		http_utils.RespondError(w, rErr)
		return
	}
	rBody, rErr := ioutil.ReadAll(r.Body)
	if rErr != nil {
		restErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, restErr)
	}
	defer r.Body.Close()
	var it items.Item
	if err := json.Unmarshal(rBody, &it); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, restErr)
	}
	it.Seller = oauth.GetCallerId(r)
	ic, err := services.ItemsService.Create(it)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, ic)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	/*
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}
	if oauth.GetCallerId(r) == 0 {
		rErr := rest_errors.NewUnauthorizedError("not authorized")
		http_utils.RespondError(w, rErr)
		return
	}
	*/
	vars := mux.Vars(r)
	id := strings.TrimSpace(vars["id"])
	it, err := services.ItemsService.Get(id)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, it)
}


func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	/*
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}
	if oauth.GetCallerId(r) == 0 {
		rErr := rest_errors.NewUnauthorizedError("not authorized")
		http_utils.RespondError(w, rErr)
		return
	}
	*/
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, rErr)
		return
	}
	defer r.Body.Close()
	var q queries.EsQuery
	if err := json.Unmarshal(b, &q); err != nil {
		rErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, rErr)
		return
	}
	i, rErr :=  services.ItemsService.Search(q)
	if rErr != nil {
		http_utils.RespondError(w, rErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, i)
}