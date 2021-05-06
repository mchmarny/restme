package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const (
	echoHandlerPath = "/v1/echo"
)

func NewEchoHandler(logger *log.Logger) EchoHandler {
	return EchoHandler{
		logger: logger,
		Path:   echoHandlerPath,
	}
}

type EchoHandler struct {
	logger *log.Logger
	Path   string
}

type message struct {
	On      int64  `json:"on"`
	Message string `json:"msg"`
}

func (h EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("serving: %+v", r)

	if r.URL.Path != h.Path {
		handleError(w, http.StatusNotFound, "Expected: %s, got:%s", h.Path, r.URL.Path)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != contentTypeJSON {
		handleError(w, http.StatusUnsupportedMediaType, "Invalid content type, expected %s", contentTypeJSON)
		return
	}

	switch r.Method {
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		handleError(w, http.StatusMethodNotAllowed, "Supported methods: POST")
		return
	}
}

func (h EchoHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(w, http.StatusUnsupportedMediaType, "Invalid method, expected %s", http.MethodPost)
		return
	}

	var m message
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&m); err != nil {
		if errors.As(err, &unmarshalErr) {
			handleError(w, http.StatusBadRequest, "Wrong type for field: %s", unmarshalErr.Field)
		} else {
			handleError(w, http.StatusBadRequest, "Invalid message: %v", err)
		}
		return
	}

	h.logger.Printf("message: %v", m)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(m); err != nil {
		handleError(w, http.StatusInternalServerError, "Error encoding message: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
