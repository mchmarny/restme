package handler

import (
	"log"
	"net/http"
)

const (
	contentTypeJSON    = "application/json"
	defaultHandlerPath = "/"
)

func NewDefaultHandler(logger *log.Logger) DefaultHandler {
	return DefaultHandler{
		logger: logger,
		Path:   defaultHandlerPath,
	}
}

type DefaultHandler struct {
	logger *log.Logger
	Path   string
}

func (h DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("serving: %+v", r)

	data := []byte("default handler")

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Printf("error writing data: %v", err)
	}
}
