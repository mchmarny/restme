package cmd

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func MakeRequestCmd(appName string) *cli.Command {
	return &cli.Command{
		Name:   "request",
		Usage:  fmt.Sprintf("invoke %s request service", appName),
		Flags:  []cli.Flag{addressFlag},
		Action: invokeRequest,
	}
}

func invokeRequest(c *cli.Context) error {
	address := c.String("address")
	if address == "" {
		return cli.Exit("address required", 1)
	}

	url := fmt.Sprintf("%s/v1/request/info", address)
	if _, err := httpInvoke(url, http.MethodGet, nil); err != nil {
		return errors.Wrapf(err, "error invoking %s", address)
	}

	return nil
}
