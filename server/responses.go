package server

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Response struct {
	Write func(w http.ResponseWriter)
}

func NewErrorResponse(m string, code int, err error) *Response {
	log.Errorf("%v: %v", m, err)
	write := func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"message": m})
	}
	return &Response{Write: write}
}

func NewSuccessJsonResponse(d interface{}, code int) *Response {
	write := func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(d)
	}
	return &Response{Write: write}
}
