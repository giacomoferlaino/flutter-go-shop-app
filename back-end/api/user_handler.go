package api

import (
	"errors"
	"flutter_shop_app/api/method"
	"flutter_shop_app/app"
	"flutter_shop_app/auth"
	"flutter_shop_app/orm"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var userRoutes = struct {
	Product         string
	FavoriteProduct string
	Order           string
}{
	Product:         "product",
	FavoriteProduct: "product/favorite",
	Order:           "order",
}

// NewUserHandler returns a new http user handler
func NewUserHandler(app app.State, jwtManager *auth.JwtManager) *UserHandler {
	return &UserHandler{
		app:        app,
		store:      &orm.UserDataStore{DB: app.Database},
		jwtManager: jwtManager,
	}
}

// UserHandler manages HTTP requests for users APIs
type UserHandler struct {
	app        app.State
	store      *orm.UserDataStore
	jwtManager *auth.JwtManager
}

func (handler *UserHandler) parseProduct(reqBody io.ReadCloser) (interface{}, error) {
	return ParseJSON(reqBody, &orm.Product{})
}

func (handler *UserHandler) parseOrder(reqBody io.ReadCloser) (interface{}, error) {
	return ParseJSON(reqBody, &orm.Order{})
}

// ServeHTTP implements the http.Handler interface
func (handler *UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	userID, err := authenticate(req, handler.jwtManager)
	if err != nil {
		sendError(res, err)
		return
	}

	route := strings.Join(strings.Split(req.URL.Path, "/")[2:], "/")
	switch route {
	case userRoutes.Product:
		switch req.Method {
		case method.GET:
			handler.GetProducts(res, req, userID)
		case method.POST:
			handler.AddProduct(res, req, userID)
		case method.DELETE:
			handler.RemoveProduct(res, req, userID)
		}
	case userRoutes.FavoriteProduct:
		switch req.Method {
		case method.GET:
			handler.GetFavoriteProducts(res, req, userID)
		case method.POST:
			handler.AddFavoriteProduct(res, req, userID)
		case method.DELETE:
			handler.RemoveFavoriteProduct(res, req, userID)
		}
	case userRoutes.Order:
		switch req.Method {
		case method.GET:
			handler.GetOrders(res, req, userID)
		case method.POST:
			handler.AddOrder(res, req, userID)
		}
	}
}

// GetProducts reads all the product related to a specific user
func (handler *UserHandler) GetProducts(res http.ResponseWriter, req *http.Request, userID uint) {
	data, _, err := handler.store.GetByID(userID, true)
	if err != nil {
		sendError(res, err)
		return
	}
	users := data.([]orm.User)
	products := users[0].Products
	sendSuccess(res, products, int64(len(products)))
}

// AddProduct adds a new product to a given user
func (handler *UserHandler) AddProduct(res http.ResponseWriter, req *http.Request, userID uint) {
	data, err := handler.parseProduct(req.Body)
	newProduct, ok := data.(*orm.Product)
	if err != nil || !ok {
		sendError(res, errors.New("Invalid product"))
	}
	createdProduct, err := handler.store.AddProduct(userID, *newProduct)
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, []orm.Product{*createdProduct}, 1)
}

// RemoveProduct removes a product from a give user
func (handler *UserHandler) RemoveProduct(res http.ResponseWriter, req *http.Request, userID uint) {
	productID, err := strconv.ParseUint(req.FormValue("id"), 10, 64)
	if err != nil {
		sendError(res, errors.New("Invalid product id"))
		return
	}
	err = handler.store.RemoveProduct(userID, uint(productID))
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, []struct{}{}, 1)
}

// GetFavoriteProducts reads all the product related to a specific user
func (handler *UserHandler) GetFavoriteProducts(res http.ResponseWriter, req *http.Request, userID uint) {
	data, _, err := handler.store.GetByID(userID, true)
	if err != nil {
		sendError(res, err)
		return
	}
	users := data.([]orm.User)
	products := users[0].FavoriteProducts
	sendSuccess(res, products, int64(len(products)))
}

// AddFavoriteProduct adds a new favorite product to a given user
func (handler *UserHandler) AddFavoriteProduct(res http.ResponseWriter, req *http.Request, userID uint) {
	productID, err := strconv.ParseUint(req.FormValue("id"), 10, 64)
	if err != nil {
		sendError(res, errors.New("Invalid product id"))
		return
	}
	err = handler.store.AddFavoriteProduct(userID, uint(productID))
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, []struct{}{}, 1)
}

// RemoveFavoriteProduct removes a favorite product from a give user
func (handler *UserHandler) RemoveFavoriteProduct(res http.ResponseWriter, req *http.Request, userID uint) {
	productID, err := strconv.ParseUint(req.FormValue("id"), 10, 64)
	if err != nil {
		sendError(res, errors.New("Invalid product id"))
		return
	}
	err = handler.store.RemoveFavoriteProduct(userID, uint(productID))
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, []struct{}{}, 1)
}

// GetOrders returns a list of orders for a given user
func (handler *UserHandler) GetOrders(res http.ResponseWriter, req *http.Request, userID uint) {
	orders, err := handler.store.GetOrders(userID)
	if err != nil {
		sendError(res, err)
	}
	sendSuccess(res, orders, 1)
}

// AddOrder adds a new order for a given user
func (handler *UserHandler) AddOrder(res http.ResponseWriter, req *http.Request, userID uint) {
	data, err := handler.parseOrder(req.Body)
	newOrder, ok := data.(*orm.Order)
	if err != nil || !ok {
		sendError(res, errors.New("Invalid order"))
		return
	}
	createdOrder, err := handler.store.AddOrder(userID, *newOrder)
	if err != nil {
		sendError(res, err)
	}
	sendSuccess(res, []orm.Order{*createdOrder}, 1)
}
