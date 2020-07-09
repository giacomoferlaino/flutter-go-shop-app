package product

import (
	"encoding/json"
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
	query := "select * from product"
	rows, err := handler.app.Database.Query(query)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
	}
	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(res, err)
		}
		products = append(products, product)

	}
	data, err := json.Marshal(products)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
	}
	fmt.Fprintln(res, string(data))
}
