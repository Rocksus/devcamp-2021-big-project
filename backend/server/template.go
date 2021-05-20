package server

import (
	"encoding/json"
	"net/http"
)

type headerData struct {
	ProcessTime  float32 `json:"process_time_ms"`
	ErrorMessage string  `json:"error_message,omitempty"`
}

type response struct {
	Header headerData  `json:"header"`
	Data   interface{} `json:"data"`
}

func RenderResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	resp := response{
		Header: headerData{},
		Data:   data,
	}

	w.WriteHeader(statusCode)
	d, _ := json.Marshal(resp)
	w.Write(d)
	return
}

func RenderError(w http.ResponseWriter, statusCode int, err error) {
	resp := response{
		Header: headerData{
			ErrorMessage: err.Error(),
		},
	}

	w.WriteHeader(statusCode)
	d, _ := json.Marshal(resp)
	w.Write(d)
	return
}
