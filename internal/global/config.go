package global

import (
	"flag"
	"fmt"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/internal/config"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/pkg/log"
)

// Version indicates the current version of the application.
var Version = "1.0.1"
var FlagConfig = flag.String("config", "./config/local.yml", "path to the config file")
var Logger log.Logger
var AppConfig *config.Config

var APIGroup = "v1"

func init() {
	flag.Parse()
	// create root logger tagged with server version
	Logger = log.New().With(nil, "version", Version)
	fmt.Printf("DEBUG: global/init FlagConfig: %+v\n", *FlagConfig)

	// load application configurations
	AppConfig, _ = config.Load(*FlagConfig, Logger)
	fmt.Printf("DEBUG: global/init AppConfig: %+v\n", *AppConfig)

}
