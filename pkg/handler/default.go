package handler

import (
	"net/http"
)

const (
	contentTypeJSON    = "application/json"
	defaultHandlerPath = "/"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	data := []byte("not implemented")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
