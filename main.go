package main

import (
	"Inventuri/app/db"
	"Inventuri/app/models"
	"log"
	"os"

	fig "github.com/common-nighthawk/go-figure"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	cli "github.com/urfave/cli/v2"
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

func start(c *cli.Context) error {
	fig.NewFigure(c.App.Name, "standard", true).Print()

	//connect to Database
	err := db.OpenConnection()
	if err != nil {
		log.Fatalf("Error connecting to the Database: %s", err.Error())
	}

	//migrate Models to the postgre Database
	if err = db.Db.AutoMigrate(models.Product{}).Error; err != nil {
		log.Fatalf("Error migrating Product to Database: %s", err.Error())
	}
	if err = db.Db.AutoMigrate(models.User{}).Error; err != nil {
		log.Fatalf("Error migrating User to Database: %s", err.Error())
	}
	if err = db.Db.AutoMigrate(models.Business{}).Error; err != nil {
		log.Fatalf("Error migrating Business to Database: %s", err.Error())
	}

	return err
}
