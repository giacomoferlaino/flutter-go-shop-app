package orm

import (
	"encoding/json"
)

// MetaData contains a sql query metadata
type MetaData struct {
	Rows int64 `json:"rows"`
}

// Response is the ORM response data model
type Response struct {
	Meta MetaData      `json:"meta"`
	Data []interface{} `json:"data"`
}

// MarshalJSON returns the response in text format
func (response *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(*response)
}
