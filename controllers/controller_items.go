package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lavinas-science/learn-items-api/domain/items"
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
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
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

}
