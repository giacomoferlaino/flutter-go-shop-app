package product

import (
	"encoding/json"
	"flutter_shop_app/app"
	"fmt"
	"io/ioutil"
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

// Get returns all saved products
func (handler *Handler) Get(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	query := "select * from product"
	rows, err := handler.app.Database.Query(query)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.ImageURL, &product.IsFavorite)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(res, err)
			return
		}
		products = append(products, product)

	}
	data, err := json.Marshal(products)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(data))
}

// Post saves a new product
func (handler *Handler) Post(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	reqBody, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	newProduct := &Product{}
	err = json.Unmarshal(reqBody, newProduct)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	query := `INSERT INTO product (title, description, price, imageUrl, isFavorite)
	VALUES(?, ?, ?, ?, ?)`
	result, err := handler.app.Database.Exec(query, newProduct.Title, newProduct.Description, newProduct.Price, newProduct.ImageURL, newProduct.IsFavorite)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	newID, err := result.LastInsertId()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	fmt.Fprintln(res, newID)
}
