package auth

import (
	"encoding/json"
	"errors"
	"flutter_shop_app/api"
	"flutter_shop_app/app"
	"flutter_shop_app/orm"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NewHTTPHandler returns a new http authentication handler
func NewHTTPHandler(app app.State, dataStore *orm.UserDataStore) *HTTPHandler {
	return &HTTPHandler{
		app:        app,
		store:      dataStore,
		jwtManager: newJwtManager("secret_key"),
	}
}

// HTTPHandler manages HTTP requests
type HTTPHandler struct {
	app        app.State
	store      *orm.UserDataStore
	jwtManager *jwtManager
}

// SignUp registers a new user
func (handler *HTTPHandler) SignUp(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	newUser, err := api.ParseJSON(req.Body, &orm.User{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, api.NewErrorResponse(err.Error()).ToJSON())
		return
	}
	users, rowsAffected, err := handler.store.Add(newUser)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, api.NewErrorResponse(err.Error()).ToJSON())
		return
	}
	apiResponse, err := json.Marshal(api.NewSuccessResponse(users, rowsAffected))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, api.NewErrorResponse(err.Error()).ToJSON())
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(apiResponse))
}

// Login logs a user in if username and password match
func (handler *HTTPHandler) Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	data, err := api.ParseJSON(req.Body, &LoginData{})
	loginData := data.(*LoginData)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, api.NewErrorResponse(err.Error()).ToJSON())
		return
	}
	data, _, err = handler.store.GetByEmail(loginData.Email)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, api.NewErrorResponse(err.Error()).ToJSON())
		return
	}
	user := data.([]orm.User)[0]
	if user.Passwod != loginData.Password {
		err = errors.New("Invalid credentials")
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, api.NewErrorResponse(err.Error()).ToJSON())
		return
	}
	token, err := handler.jwtManager.createToken()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, api.NewErrorResponse(err.Error()).ToJSON())
		return
	}
	sessionData := SessionData{
		IDToken:   token,
		ExpiresIn: handler.jwtManager.expiredIn,
	}
	apiResponse, err := json.Marshal(api.NewSuccessResponse([]SessionData{sessionData}, 0))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, api.NewErrorResponse(err.Error()).ToJSON())
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(apiResponse))
}
