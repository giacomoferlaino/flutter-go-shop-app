package api

import (
	"errors"
	"flutter_shop_app/api/method"
	"flutter_shop_app/app"
	"flutter_shop_app/orm"
	"io"
	"net/http"
)

// NewProductHandler returns a new http product handler
func NewProductHandler(app app.State) *ProductHandler {
	return &ProductHandler{
		app:   app,
		store: &orm.ProductDataStore{DB: app.Database},
	}
}

// ProductHandler manages HTTP requests for product APIs
type ProductHandler struct {
	app   app.State
	store *orm.ProductDataStore
}

func (handler *ProductHandler) parseJSON(reqBody io.ReadCloser) (interface{}, error) {
	return ParseJSON(reqBody, &orm.Product{})
}

// ServeHTTP implements the http.Handler interface
func (handler *ProductHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
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
			sendError(res, errors.New("Missing product id"))
			break
		}
		crudHandler.UpdateByID(res, req)
	case method.DELETE:
		if req.FormValue("id") == "" {
			sendError(res, errors.New("Missing product id"))
			break
		}
		crudHandler.DeleteByID(res, req)
	}
}
