package cmd

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func MakeRuntimeCmd(appName string) *cli.Command {
	return &cli.Command{
		Name:   "runtime",
		Usage:  fmt.Sprintf("invoke %s runtime service", appName),
		Flags:  []cli.Flag{addressFlag},
		Action: invokeRuntime,
	}
}

func invokeRuntime(c *cli.Context) error {
	address := c.String("address")
	if address == "" {
		return cli.Exit("address required", 1)
	}

	url := fmt.Sprintf("%s/v1/runtime/info", address)
	if _, err := httpInvoke(url, http.MethodGet, nil); err != nil {
		return errors.Wrapf(err, "error invoking %s", address)
	}

	return nil
}
