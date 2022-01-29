package request

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mchmarny/restme/pkg/config"
	"github.com/mchmarny/restme/pkg/log"
)

var (
	ipRegExp = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
)

// NewRequestService creates new RequestService instance.
func NewService(logger *log.Logger, conf *config.Config) (*Service, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}
	if conf == nil {
		return nil, fmt.Errorf("config is nil")
	}

	return &Service{
		logger: logger,
		conf:   conf,
	}, nil
}

// Service provides object representing the inbound HTTP request.
type Service struct {
	logger *log.Logger
	conf   *config.Config
}

// RequestHandler handles the inbound requests.
func (s *Service) RequestHandler(c *gin.Context) {
	result := &Response{
		Request: s.getRequestMetadata(c),
		Headers: make(map[string]interface{}),
		EnvVars: make(map[string]interface{}),
	}

	// env vars
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		result.EnvVars[pair[0]] = pair[1]
	}

	// headers
	for name, headers := range c.Request.Header {
		name = strings.ToLower(name)
		if name == "authorization" {
			continue
		}
		for i, h := range headers {
			if len(headers) > 1 {
				result.Headers[fmt.Sprintf("%s[%d]", name, i)] = h
			} else {
				result.Headers[name] = h
			}
		}
	}

	c.IndentedJSON(http.StatusOK, result)
}

func (s *Service) getRequestMetadata(c *gin.Context) gin.H {
	id, err := uuid.NewUUID()
	if err != nil {
		s.logger.Errorf("Error while getting id: %v\n", err)
	}

	from := c.Request.RemoteAddr

	if s.conf.IP != nil && s.conf.IP.FromHeader {
		from = parseIPv4FromHeader(c.Request.Header.Get(s.conf.IP.HeaderKey))
	}

	return gin.H{
		"id":       id,
		"time":     time.Now().UTC(),
		"version":  s.logger.GetAppVersion(),
		"path":     c.Request.URL.EscapedPath(),
		"protocol": c.Request.Proto,
		"host":     c.Request.Host,
		"method":   c.Request.Method,
		"from":     from,
	}
}

func parseIPv4FromHeader(v string) string {
	if v == "" {
		return ""
	}

	ips := ipRegExp.FindAllString(v, -1)
	if len(ips) > 0 {
		return ips[0]
	}

	return ""
}
