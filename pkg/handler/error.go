package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleError(w http.ResponseWriter, code int, msg string, args ...interface{}) {
	w.Header().Set("Content-Type", contentTypeJSON)
	w.WriteHeader(code)
	resp := make(map[string]string)
	resp["message"] = fmt.Sprintf(msg, args...)
	jsonResp, _ := json.Marshal(resp)
	if _, err := w.Write(jsonResp); err != nil {
		panic(err)
	}
}
