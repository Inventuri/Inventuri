package main

import (
	fig "github.com/common-nighthawk/go-figure"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Inventuri"
	app.Usage = "Start the Backend + API"
	app.Version = "0.1"
	app.Flags = CommonFlags
	app.Action = start

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error starting the App %s", err.Error())
	}
}

func start(c *cli.Context) error{
	fig.NewFigure(c.App.Name, "standard", true).Print()

	return nil
}
