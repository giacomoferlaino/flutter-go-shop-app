package api

import (
	"errors"
	"flutter_shop_app/api/method"
	"flutter_shop_app/app"
	"flutter_shop_app/orm"
	"io"
	"net/http"
)

// NewOrderHandler returns a new http order handler
func NewOrderHandler(app app.State) *OrderHandler {
	return &OrderHandler{
		app:   app,
		store: &orm.OrderDataStore{DB: app.Database},
	}
}

// OrderHandler manages HTTP requests for order APIs
type OrderHandler struct {
	app   app.State
	store *orm.OrderDataStore
}

func (handler *OrderHandler) parseJSON(reqBody io.ReadCloser) (interface{}, error) {
	return ParseJSON(reqBody, &orm.Order{})
}

// ServeHTTP implements the http.Handler interface
func (handler *OrderHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	crudHandler := HTTPCRUDHandler{app: handler.app, store: handler.store, parseJSON: handler.parseJSON}
	switch req.Method {
	case method.GET:
		if req.FormValue("id") != "" {
			crudHandler.GetByID(res, req)
			break
		}
		crudHandler.Get(res, req)
	case method.POST:
		crudHandler.Post(res, req)
	case method.PUT:
		if req.FormValue("id") == "" {
			sendError(res, errors.New("Missing order id"))
			break
		}
		crudHandler.UpdateByID(res, req)
	case method.DELETE:
		if req.FormValue("id") == "" {
			sendError(res, errors.New("Missing order id"))
			break
		}
		crudHandler.DeleteByID(res, req)
	}
}
