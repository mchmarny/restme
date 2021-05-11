package request

import "github.com/gin-gonic/gin"

// Response represents simple HTTP resource response.
type Response struct {
	Request gin.H                  `json:"request,omitempty"`
	Headers map[string]interface{} `json:"headers"`
	EnvVars map[string]interface{} `json:"env_vars"`
}
