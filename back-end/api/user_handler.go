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
	productID, err := strconv.ParseUint(req.FormValue("productId"), 10, 64)
	if err != nil {
		sendError(res, errors.New("Invalid product id"))
		return
	}
	route := strings.Split(req.URL.Path, "/")[2]
	switch route {
	case userRoutes.Product:
		switch req.Method {
		case method.POST:
			handler.AddProduct(res, req, userID, uint(productID))
		case method.DELETE:
			handler.RemoveProduct(res, req, userID, uint(productID))
		}
	case userRoutes.FavoriteProduct:
		switch req.Method {
		case method.POST:
			handler.AddFavoriteProduct(res, req, userID, uint(productID))
		case method.DELETE:
			handler.RemoveFavoriteProduct(res, req, userID, uint(productID))
		}
	}
}

// AddProduct adds a new product to a given user
func (handler *UserHandler) AddProduct(res http.ResponseWriter, req *http.Request, userID uint, productID uint) {
	err := handler.store.AddProduct(userID, productID)
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, struct{}{}, 0)
}

// RemoveProduct removes a product from a give user
func (handler *UserHandler) RemoveProduct(res http.ResponseWriter, req *http.Request, userID uint, productID uint) {
	err := handler.store.RemoveProduct(userID, productID)
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, struct{}{}, 0)
}

// AddFavoriteProduct adds a new favorite product to a given user
func (handler *UserHandler) AddFavoriteProduct(res http.ResponseWriter, req *http.Request, userID uint, productID uint) {
	err := handler.store.AddFavoriteProduct(userID, productID)
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, struct{}{}, 0)
}

// RemoveFavoriteProduct removes a favorite product from a give user
func (handler *UserHandler) RemoveFavoriteProduct(res http.ResponseWriter, req *http.Request, userID uint, productID uint) {
	err := handler.store.RemoveFavoriteProduct(userID, productID)
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, struct{}{}, 0)
}
