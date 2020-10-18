package db

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
	"log"
	"os"
)

var Db *gorm.DB

/*
getConfig reads in the config.dev.json file located in the db folder and generates a struct with the values necessary to connect to the database,
If the file is not according to format or the File is not found, an Error is thrown
 */
func getConfig() (*Configuration, error) {
	var cfg Configuration
	file, err := os.Open("config.dev.json")
	if err != nil {
		log.Printf("No Configurationfile found for Database: %s", err.Error())
	}
	bytesvalue, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(bytesvalue, &cfg)
	if err != nil {
		log.Printf("Configurationfile not following the expected format: %s", err.Error())
	}
	return &cfg, err
}

/*
OpenConnection calls getConfig and constructs a Connectionstring for the Database.
If the Connection fails, an Error is thrown
 */
func OpenConnection() error{
	cfg, err := getConfig()
	if err != nil {
		log.Fatalf("Error getting the Configuration: %s", err.Error())
	}

	connString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", cfg.Host, cfg.User, cfg.Databasename, cfg.Password, cfg.Portnumber)
	Db, err = gorm.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Error connecting to the Database: %s", err.Error())
	}
	return err
}
