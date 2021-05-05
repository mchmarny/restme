package handler

import (
	"log"
	"net/http"
)

const (
	contentTypeJSON = "application/json"
)

func NewDefaultHandler(logger *log.Logger) DefaultHandler {
	return DefaultHandler{
		logger: logger,
	}
}

type DefaultHandler struct {
	logger *log.Logger
}

func (h DefaultHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.logger.Printf("serving: %+v", req)

	data := []byte("default handler")

	res.WriteHeader(http.StatusOK)

	if _, err := res.Write(data); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		h.logger.Printf("error writing data: %v", err)
	}
}
