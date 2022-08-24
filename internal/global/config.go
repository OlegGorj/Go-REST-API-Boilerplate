package global

import (
	"flag"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/internal/config"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/pkg/log"
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

// Version indicates the current version of the application.
var Version = "1.0.2"
var FlagConfig = flag.String("config", "./config/local.yml", "path to the config file")
var Logger log.Logger
var AppConfig *config.Config
var APIGroup = "v1"

// var DBConnection *dbx.DB = nil

func init() {
	flag.Parse()
	// create root logger tagged with server version
	Logger = log.New().With(nil, "version", Version)

	// load application configurations
	AppConfig, _ = config.Load(*FlagConfig, Logger)

}

func ConnectDB() (*dbx.DB, error) {
	// connect to the database
	DBConnection, err := dbx.MustOpen("postgres", AppConfig.DSN)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}

	return DBConnection, nil
}
