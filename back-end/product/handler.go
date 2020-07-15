package product

import (
	"encoding/json"
	"flutter_shop_app/app"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// NewHandler returns a new Handler
func NewHandler(app app.State) *Handler {
	store := dataStore{db: app.Database}
	return &Handler{app: app, store: store}
}

// Handler contains the HTTP endpoint handlers
type Handler struct {
	app   app.State
	store dataStore
}

func (handler *Handler) parseProduct(reqBody io.ReadCloser) (*Product, error) {
	body, err := ioutil.ReadAll(reqBody)
	defer reqBody.Close()
	if err != nil {
		return nil, err
	}
	product := &Product{}
	err = json.Unmarshal(body, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Get returns all saved products
func (handler *Handler) Get(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	products, err := handler.store.getAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	response, err := json.Marshal(products)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(response))
}

// Post saves a new product
func (handler *Handler) Post(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	newProduct, err := handler.parseProduct(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	createdProduct, err := handler.store.add(*newProduct)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	response, err := json.Marshal(createdProduct)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(response))
}

// GetByID returns a product based on its ID
func (handler *Handler) GetByID(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	products, err := handler.store.getByID(uint(id))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	response, err := json.Marshal(products)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(response))
}

// DeleteByID returns the number of deleted rows
func (handler *Handler) DeleteByID(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	deletedRows, err := handler.store.deleteByID(uint(id))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	data := []struct {
		DeletedRows int64 `json:"deletedRows"`
	}{{DeletedRows: deletedRows}}
	response, err := json.Marshal(data)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(response))
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
	updatedProduct, err := handler.store.updateByID(uint(id), *product)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	response, err := json.Marshal(updatedProduct)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(response))
}
