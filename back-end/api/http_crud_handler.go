package api

import (
	"flutter_shop_app/app"
	"flutter_shop_app/orm"
	"io"
	"net/http"
	"strconv"
)

// HTTPCRUDHandler contains the HTTP endpoint CRUD handlers
type HTTPCRUDHandler struct {
	app       app.State
	store     orm.DataStore
	parseJSON func(reqBody io.ReadCloser) (interface{}, error)
}

// Get returns all saved items
func (handler *HTTPCRUDHandler) Get(res http.ResponseWriter, req *http.Request) {
	items, rowsAffected, err := handler.store.GetAll()
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, items, rowsAffected)
}

// Post saves a new item
func (handler *HTTPCRUDHandler) Post(res http.ResponseWriter, req *http.Request) {
	newItem, err := handler.parseJSON(req.Body)
	if err != nil {
		sendError(res, err)
		return
	}
	items, rowsAffected, err := handler.store.Add(newItem)
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, items, rowsAffected)
}

// GetByID returns a item based on its ID
func (handler *HTTPCRUDHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseUint(req.FormValue("id"), 10, 64)
	if err != nil {
		sendError(res, err)
		return
	}
	items, rowsAffected, err := handler.store.GetByID(uint(id))
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, items, rowsAffected)
}

// DeleteByID delete an item based on its ID
func (handler *HTTPCRUDHandler) DeleteByID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.FormValue("id"), 10, 64)
	if err != nil {
		sendError(res, err)
		return
	}
	rowsAffected, err := handler.store.DeleteByID(uint(id))
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, nil, rowsAffected)
}

// UpdateByID update an items based on its ID
func (handler *HTTPCRUDHandler) UpdateByID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.FormValue("id"), 10, 64)
	if err != nil {
		sendError(res, err)
		return
	}
	item, err := handler.parseJSON(req.Body)
	if err != nil {
		sendError(res, err)
		return
	}
	items, rowsAffected, err := handler.store.UpdateByID(uint(id), item)
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, items, rowsAffected)
}
