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
}{
	Product:         "product",
	FavoriteProduct: "favoriteProduct",
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

func (handler *UserHandler) parseJSON(reqBody io.ReadCloser) (interface{}, error) {
	return ParseJSON(reqBody, &orm.Product{})
}

// ServeHTTP implements the http.Handler interface
func (handler *UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	userID, err := authenticate(req, handler.jwtManager)
	if err != nil {
		sendError(res, err)
		return
	}

	route := strings.Split(req.URL.Path, "/")[2]
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
	productID, err := strconv.ParseUint(req.FormValue("productId"), 10, 64)
	if err != nil {
		sendError(res, errors.New("Invalid product id"))
		return
	}
	err = handler.store.AddProduct(userID, uint(productID))
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, []struct{}{}, 0)
}

// RemoveProduct removes a product from a give user
func (handler *UserHandler) RemoveProduct(res http.ResponseWriter, req *http.Request, userID uint) {
	productID, err := strconv.ParseUint(req.FormValue("productId"), 10, 64)
	if err != nil {
		sendError(res, errors.New("Invalid product id"))
		return
	}
	err = handler.store.RemoveProduct(userID, uint(productID))
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, []struct{}{}, 0)
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
	productID, err := strconv.ParseUint(req.FormValue("productId"), 10, 64)
	if err != nil {
		sendError(res, errors.New("Invalid product id"))
		return
	}
	err = handler.store.AddFavoriteProduct(userID, uint(productID))
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, []struct{}{}, 0)
}

// RemoveFavoriteProduct removes a favorite product from a give user
func (handler *UserHandler) RemoveFavoriteProduct(res http.ResponseWriter, req *http.Request, userID uint) {
	productID, err := strconv.ParseUint(req.FormValue("productId"), 10, 64)
	if err != nil {
		sendError(res, errors.New("Invalid product id"))
		return
	}
	err = handler.store.RemoveFavoriteProduct(userID, uint(productID))
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, []struct{}{}, 0)
}
