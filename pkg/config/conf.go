package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

// Config represents service config.
type Config struct {
	Log  *Log  `json:"log"`
	IP   *IP   `json:"ip"`
	Auth *Auth `json:"auth"`
}

type Log struct {
	Level string `json:"level"`
	JSON  bool   `json:"json"`
}

type IP struct {
	FromHeader bool   `json:"from_header"`
	HeaderKey  string `json:"header_key"`
}

type Auth struct {
	Tokens map[string]string `json:"tokens"`
}

// GetConfigFromFile reads app config from file.
func GetConfigFromFile(path string) (*Config, error) {
	if path == "" {
		return nil, errors.New("config value must be existing file path")
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, errors.Wrapf(err, "provided config file does not exists: %s", path)
	}

	j, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening config file: %s", path)
	}
	defer j.Close()

	b, err := io.ReadAll(j)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading config file %v", j)
	}

	var c Config
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, errors.Wrapf(err, "error unmarshalling config file %v", j)
	}
	return &c, nil
}
