package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mchmarny/restme/pkg/cmd"
	"github.com/urfave/cli/v2"
)

const (
	appName = "restme"
)

var (
	appVersion = "v0.0.1-default"
)

func main() {
	app := &cli.App{
		Name: fmt.Sprintf("%s - %s", appName, appVersion),
		Commands: []*cli.Command{
			{
				Name:  "token",
				Usage: fmt.Sprintf("%s token tooling", appName),
				Subcommands: []*cli.Command{
					cmd.MakeTokenCmd(appName),
				},
			},
			{
				Name:  "echo",
				Usage: fmt.Sprintf("%s echo message", appName),
				Subcommands: []*cli.Command{
					cmd.MakeEchoCmd(appName),
				},
			},
			{
				Name:  "invoke",
				Usage: fmt.Sprintf("%s service invoke", appName),
				Subcommands: []*cli.Command{
					cmd.MakeInvokeCmd(appName, "request", "/v1/request/info"),
					cmd.MakeInvokeCmd(appName, "runtime", "/v1/runtime/info"),
					cmd.MakeLoadCmd(appName),
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
