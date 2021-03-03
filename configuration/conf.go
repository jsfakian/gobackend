package configuration

import (
	"log"

	"github.com/pelletier/go-toml"
)

//Configuration is the struct that describes the available configuration variables
type Configuration struct {
	ExpirationCookie int64 // In seconds
	Certificate      string
	ServerKey        string
}

//Conf is the variable that keeps the configuration of the user
var Conf Configuration

//LoadConfiguration reads the config.json of the repo for the global variables set by the user
func LoadConfiguration() {
	confFile := "configuration/server.conf"
	config, err := toml.LoadFile(confFile)
	if err != nil {
		log.Fatal(err)
	}

	Conf.ExpirationCookie = config.Get("expiration_cookie").(int64)

	Conf.Certificate = config.Get("certificate").(string)
	Conf.ServerKey = config.Get("server_key").(string)
}
