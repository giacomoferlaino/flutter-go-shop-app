package api

import (
	"encoding/json"
)

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

// MarshalJSON returns the response in text format
func (response *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(*response)
}
