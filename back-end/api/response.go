package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendError(res http.ResponseWriter, err error) {
	res.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(res, NewErrorResponse(err.Error()).ToJSON())
	return
}

func sendSuccess(res http.ResponseWriter, data interface{}, rowsAffected int64) {
	apiResponse, err := json.Marshal(NewSuccessResponse(data, rowsAffected))
	if err != nil {
		sendError(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(apiResponse))
}

// NewSuccessResponse return a new http success response
func NewSuccessResponse(data interface{}, rows int64) *Response {
	return &Response{
		Meta: MetaData{Rows: rows},
		Data: data,
	}
}

// NewErrorResponse return a new http error response
func NewErrorResponse(error string) *Response {
	return &Response{
		Meta: MetaData{Error: error},
	}
}

// MetaData contains a sql query metadata
type MetaData struct {
	Rows  int64  `json:"rows"`
	Error string `json:"error"`
}

// Response is the ORM response data model
type Response struct {
	Meta MetaData    `json:"meta"`
	Data interface{} `json:"data"`
}

// ToJSON return the response in string format
func (response *Response) ToJSON() string {
	bs, _ := response.MarshalJSON()
	return string(bs)
}

// MarshalJSON is the custom implementation of the json marshaling for the response data model
func (response *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(*response)
}
