package product

import (
	"flutter_shop_app/app"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NewHandler returns a new Handler
func NewHandler(app app.State) *Handler {
	return &Handler{app: app}
}

// Handler contains the HTTP endpoint handlers
type Handler struct {
	app app.State
}

// Get handles the HTTP GET method
func (handler *Handler) Get(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(res, "Index page!")
}
