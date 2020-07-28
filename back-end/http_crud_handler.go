package main

import (
	"encoding/json"
	"flutter_shop_app/app"
	"flutter_shop_app/orm"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// NewHTTPCRUDHandler returns a new http Handler to manage CRUD request
func NewHTTPCRUDHandler(app app.State, dataStore orm.DataStore) *HTTPCRUDHandler {
	return &HTTPCRUDHandler{app: app, store: dataStore}
}

// HTTPCRUDHandler contains the HTTP endpoint CRUD handlers
type HTTPCRUDHandler struct {
	app   app.State
	store orm.DataStore
}

// Get returns all saved items
func (handler *HTTPCRUDHandler) Get(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	queryResponse, err := handler.store.GetAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	apiResponse, err := json.Marshal(queryResponse)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(apiResponse))
}

// Post saves a new item
func (handler *HTTPCRUDHandler) Post(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	newItem, err := handler.store.ParseJSON(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	queryResponse, err := handler.store.Add(newItem)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	apiResponse, err := json.Marshal(queryResponse)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(apiResponse))
}

// GetByID returns a item based on its ID
func (handler *HTTPCRUDHandler) GetByID(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	queryResponse, err := handler.store.GetByID(uint(id))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	apiResponse, err := json.Marshal(queryResponse)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(apiResponse))
}

// DeleteByID delete an item based on its ID
func (handler *HTTPCRUDHandler) DeleteByID(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	queryResponse, err := handler.store.DeleteByID(uint(id))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	apiResponse, err := json.Marshal(queryResponse)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(apiResponse))
}

// UpdateByID update an items based on its ID
func (handler *HTTPCRUDHandler) UpdateByID(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	item, err := handler.store.ParseJSON(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	queryResponse, err := handler.store.UpdateByID(uint(id), item)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	apiResponse, err := json.Marshal(queryResponse)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(apiResponse))
}
