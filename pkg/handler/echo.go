package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func NewEchoHandler(logger *log.Logger) EchoHandler {
	return EchoHandler{
		logger: logger,
	}
}

type EchoHandler struct {
	logger *log.Logger
}

type message struct {
	On      int64  `json:"on"`
	Message string `json:"msg"`
}

func (h EchoHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.logger.Printf("serving: %+v", req)

	contentType := req.Header.Get("Content-Type")
	if contentType != contentTypeJSON {
		handleError(res, http.StatusUnsupportedMediaType, "Invalid content type, expected %s.", contentTypeJSON)
		return
	}

	var m message
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&m); err != nil {
		if errors.As(err, &unmarshalErr) {
			handleError(res, http.StatusBadRequest, "Bad request. Wrong type for field '%s'.", unmarshalErr.Field)
		} else {
			handleError(res, http.StatusBadRequest, "Bad Request: %v", err)
		}
		return
	}

	res.WriteHeader(http.StatusOK)
}
