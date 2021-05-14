package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mchmarny/restme/pkg/echo"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

const (
	httpClientName  = "restme"
	jsonContentType = "application/json"
)

var (
	addressFlag = &cli.StringFlag{
		Name:    "address",
		Aliases: []string{"a"},
		Usage:   "base url of the service (default: http://localhost:8080",
		Value:   "http://localhost:8080",
	}
)

func MakeEchoCmd(appName string) *cli.Command {
	return &cli.Command{
		Name:  "message",
		Usage: fmt.Sprintf("echos %s message", appName),
		Flags: []cli.Flag{
			addressFlag,
			&cli.StringFlag{
				Name:     "content",
				Aliases:  []string{"m"},
				Usage:    "message content",
				Required: true,
			},
		},
		Action: sendMessage,
	}
}

func sendMessage(c *cli.Context) error {
	address := c.String("address")
	if address == "" {
		return cli.Exit("address required", 1)
	}
	contentStr := c.String("content")
	if contentStr == "" {
		return cli.Exit("content required", 1)
	}

	m := &echo.Message{
		On:      time.Now().Unix(),
		Message: contentStr,
	}

	content, err := json.Marshal(m)
	if err != nil {
		return errors.Wrapf(err, "error marshaling %+v", m)
	}

	buf := bytes.NewBuffer(content)

	url := fmt.Sprintf("%s/v1/echo/message", address)
	if _, err := httpInvoke(url, http.MethodPost, buf); err != nil {
		return errors.Wrapf(err, "error invoking %s", address)
	}

	return nil
}
