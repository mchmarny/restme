package cmd

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func MakeInvokeCmd(appName, cmdName, path string) *cli.Command {
	h := InvokeHandler{path: path}
	return &cli.Command{
		Name:   cmdName,
		Usage:  fmt.Sprintf("invoke %s service on %s", cmdName, appName),
		Flags:  []cli.Flag{addressFlag},
		Action: h.Action,
	}
}

// Service provides message echo service.
type InvokeHandler struct {
	path string
}

func (h *InvokeHandler) Action(c *cli.Context) error {
	address := c.String("address")
	if address == "" {
		return cli.Exit("address required", 1)
	}

	url := fmt.Sprintf("%s%s", address, h.path)
	result, err := httpInvoke(url, http.MethodGet, nil)
	if err != nil {
		return errors.Wrapf(err, "error invoking %s", url)
	}

	fmt.Println(string(result))

	return nil
}
