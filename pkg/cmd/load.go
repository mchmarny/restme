package cmd

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func MakeLoadCmd(appName string) *cli.Command {
	return &cli.Command{
		Name:  "load",
		Usage: fmt.Sprintf("echos %s message", appName),
		Flags: []cli.Flag{
			addressFlag,
			&cli.StringFlag{
				Name:    "duration",
				Aliases: []string{"d"},
				Usage:   "load duration (default: 3s)",
				Value:   "3s",
			},
		},
		Action: sendLoad,
	}
}

func sendLoad(c *cli.Context) error {
	address := c.String("address")
	if address == "" {
		return cli.Exit("address required", 1)
	}
	duration := c.String("duration")
	if duration == "" {
		return cli.Exit("duration required", 1)
	}

	url := fmt.Sprintf("%s/v1/load/cpu/%s", address, duration)
	result, err := httpInvoke(url, http.MethodGet, nil)
	if err != nil {
		return errors.Wrapf(err, "error invoking %s", address)
	}

	fmt.Println(string(result))

	return nil
}
