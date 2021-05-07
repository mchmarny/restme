package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mchmarny/restme/pkg/load"
)

// Load represents simple HTTP load rresult
type LoadResult struct {
	Request *RequestMetadata    `json:"request,omitempty"`
	Result  *load.CPULoadResult `json:"result,omitempty"`
}

func LoadHandler(w http.ResponseWriter, r *http.Request) {
	durStr := r.URL.Query().Get("duration")
	duration, err := time.ParseDuration(durStr)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid duration parameter: '%s'", durStr)
		return
	}

	result := &LoadResult{
		Request: getRequestMetadata(r),
		Result:  load.MakeCPULoad(duration),
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(result); err != nil {
		handleError(w, http.StatusInternalServerError, "Error processing request: %v", err)
		return
	}
}
