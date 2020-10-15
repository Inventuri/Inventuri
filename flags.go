package main

import cli "github.com/urfave/cli/v2"

var CommonFlags = []cli.Flag {
	&cli.StringFlag{
				EnvVars: []string{"ENV"},
				Name: "env",
				Usage: "env",
				Value: "local",
	}, &cli.StringFlag{
		EnvVars: []string{"PORT"},
		Name: "port",
		Usage: "bind port for the service",
		Value: "3000",
	},
}