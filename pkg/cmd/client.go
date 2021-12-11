package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
)

const (
	httpClientTimeoutSec = 300
)

func httpInvoke(address, method string, content io.Reader) ([]byte, error) {
	tokePath, err := getTokenFilePath()
	if err != nil {
		return nil, errors.Wrapf(err, "error saving JWT: %v", err)
	}

	token, err := os.ReadFile(tokePath)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading test key %s: %v", tokePath, err)
	}

	req, err := http.NewRequest(method, address, content)
	if err != nil {
		return nil, errors.Wrapf(err, "error creating request to %s", address)
	}

	req.Header.Set("User-Agent", httpClientName)
	req.Header.Set("Content-Type", jsonContentType)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{
		Timeout: time.Second * httpClientTimeoutSec,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "request error: %+v", req)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing body %s", string(body))
	}

	return body, nil
}
