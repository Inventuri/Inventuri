package db

import (
	"testing"
)

func testConfigParsing(t *testing.T) {
	cfg, err := getConfig()
	if err != nil {
		t.Errorf("Error parsing the Configurationfile: %s", err.Error())
	}
	if cfg.Databasename != "inventuri" || cfg.Host != "localhost" {
		t.Errorf("Unexpected Values, Expected: inventuri, localhost. Got: %s, %s.", cfg.Databasename, cfg.Host)
	}
}

func testDBConnection(t *testing.T) {
	err := OpenConnection()
	if err != nil {
		t.Errorf("Error connecting to the Database: %s", err.Error())
	}
}
