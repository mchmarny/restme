package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mchmarny/restme/pkg/auth"
	"github.com/pkg/errors"
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
					makeTokenCmd(),
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func makeTokenCmd() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: fmt.Sprintf("creates %s token", appName),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "secret",
				Aliases:  []string{"s"},
				Usage:    "path to the shared secret that will be used by the service to parse the generated token",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "issuer",
				Aliases:  []string{"i"},
				Usage:    "identity of party generating the token",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "email",
				Aliases:  []string{"e"},
				Usage:    "email address for whom this token is generated",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "ttl",
				Usage: "duration how long this token wil be valid (e.g. 15m, 24h)",
				Value: "8760h",
			},
		},
		Action: makeTokenAction,
	}
}

func makeTokenAction(c *cli.Context) error {
	secretPath := c.String("secret")
	if secretPath == "" {
		return cli.Exit("secret required", 1)
	}
	issuer := c.String("issuer")
	if issuer == "" {
		return cli.Exit("issuer required", 1)
	}
	email := c.String("email")
	if email == "" {
		return cli.Exit("email required", 1)
	}
	ttl := c.String("ttl")
	if ttl == "" {
		return cli.Exit("ttl required", 1)
	}

	secret, err := ioutil.ReadFile(secretPath)
	if err != nil {
		return cli.Exit("error reading secret", 1)
	}

	tokenStr, err := auth.MakeJWT(secret, issuer, email, ttl)
	if err != nil {
		return errors.Wrapf(err, "error making JWT: %v", err)
	}

	fmt.Println(tokenStr)

	return nil
}
