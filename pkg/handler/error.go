package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleError(res http.ResponseWriter, code int, msg string, args ...interface{}) {
	res.Header().Set("Content-Type", contentTypeJSON)
	res.WriteHeader(code)
	resp := make(map[string]string)
	resp["message"] = fmt.Sprintf(msg, args...)
	jsonResp, _ := json.Marshal(resp)
	if _, err := res.Write(jsonResp); err != nil {
		panic(err)
	}
}
