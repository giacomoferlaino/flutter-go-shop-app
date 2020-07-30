package api

import (
	"errors"
	"flutter_shop_app/app"
	"flutter_shop_app/auth"
	"flutter_shop_app/orm"
	"net/http"
	"strings"
)

var authRoutes = struct {
	login  string
	signup string
}{
	login:  "login",
	signup: "signup",
}

// NewAuthHandler returns a new http authentication handler
func NewAuthHandler(app app.State, jwtManager *auth.JwtManager) *AuthHandler {
	return &AuthHandler{
		app:        app,
		store:      &orm.UserDataStore{DB: app.Database},
		jwtManager: jwtManager,
	}
}

// AuthHandler manages HTTP requests for authentication APIs
type AuthHandler struct {
	app        app.State
	store      *orm.UserDataStore
	jwtManager *auth.JwtManager
}

// ServeHTTP implements the http.Handler interface
func (handler *AuthHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	route := strings.Split(req.URL.Path, "/")[2]
	switch route {
	case authRoutes.login:
		handler.Login(res, req)
	case authRoutes.signup:
		handler.SignUp(res, req)
	}
}

// SignUp registers a new user
func (handler *AuthHandler) SignUp(res http.ResponseWriter, req *http.Request) {
	newUser, err := ParseJSON(req.Body, &orm.User{})
	if err != nil {
		sendError(res, err)
		return
	}
	users, rowsAffected, err := handler.store.Add(newUser)
	if err != nil {
		sendError(res, err)
		return
	}
	sendSuccess(res, users, rowsAffected)
}

// Login logs a user in if username and password match
func (handler *AuthHandler) Login(res http.ResponseWriter, req *http.Request) {
	data, err := ParseJSON(req.Body, &auth.LoginData{})
	loginData := data.(*auth.LoginData)
	if err != nil {
		sendError(res, err)
		return
	}
	data, _, err = handler.store.GetByEmail(loginData.Email)
	if err != nil {
		sendError(res, err)
		return
	}
	users, ok := data.([]orm.User)
	if !ok || users[0].Password != loginData.Password {
		err = errors.New("Invalid credentials")
		sendError(res, err)
		return
	}
	token, err := handler.jwtManager.CreateToken(users[0].ID)
	if err != nil {
		sendError(res, err)
		return
	}
	sessionData := auth.SessionData{
		IDToken:   token,
		ExpiresIn: handler.jwtManager.ExpiredIn,
	}
	sendSuccess(res, []auth.SessionData{sessionData}, 0)
}
