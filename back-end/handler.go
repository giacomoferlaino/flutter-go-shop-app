package main

import (
	"encoding/json"
	"flutter_shop_app/app"
	"flutter_shop_app/orm"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// NewHandler returns a new Handler
func NewHandler(app app.State, dataStore orm.DataStore) *Handler {
	return &Handler{app: app, store: dataStore}
}

// Handler contains the HTTP endpoint handlers
type Handler struct {
	app   app.State
	store orm.DataStore
}

func (handler *Handler) parseProduct(reqBody io.ReadCloser) (*orm.Product, error) {
	body, err := ioutil.ReadAll(reqBody)
	defer reqBody.Close()
	if err != nil {
		return nil, err
	}
	product := &orm.Product{}
	err = json.Unmarshal(body, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Get returns all saved products
func (handler *Handler) Get(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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

// Post saves a new product
func (handler *Handler) Post(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	newProduct, err := handler.parseProduct(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	queryResponse, err := handler.store.Add(*newProduct)
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

// GetByID returns a product based on its ID
func (handler *Handler) GetByID(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
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

// DeleteByID returns the number of deleted rows
func (handler *Handler) DeleteByID(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
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

// UpdateByID returns the number of deleted rows
func (handler *Handler) UpdateByID(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	product, err := handler.parseProduct(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	queryResponse, err := handler.store.UpdateByID(uint(id), *product)
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
