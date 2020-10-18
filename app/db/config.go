package db

type Configuration struct {
	User string `json:"user"`
	Password string `json:"password"`
	Databasename string `json:"databasename"`
	Portnumber string `json:"portnumber"`
	Host string `json:"host"`
}
