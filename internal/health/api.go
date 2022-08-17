package health

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"strconv"
)

// RegisterHandlers registers the handlers that perform health checks.
func RegisterHandlers(r *routing.Router, version string) {

	r.To("GET,HEAD", "/health", health(version))

}

type Health struct {
	API_version string `json:"api_version"`
	DB_health   string `json:"db"`
	API_health  string `json:"api"`
}

func healthChecks() *Health {
	var db_connect = false

	//var logger = log.New()
	//var _flagConfig = flag.String("config", "./config/local.yml", "path to the config file")
	//var _cfg, _ = conf.Load(*_flagConfig, logger)
	//flag.Parse()
	//
	//// connect to the database
	//db, err := dbx.MustOpen("postgres", _cfg.DSN)
	//if err != nil {
	//	logger.Error(err)
	//	db_connect = false
	//}
	//db_connect = true
	//
	//defer func() {
	//	if err := db.Close(); err != nil {
	//		logger.Error(err)
	//	}
	//}()

	return &Health{
		API_version: "1.0.0",
		API_health:  "ok",
		DB_health:   strconv.FormatBool(db_connect),
	}
}

// health responds to a health request.
func health(version string) routing.Handler {
	return func(c *routing.Context) error {
		return c.Write(healthChecks())
	}
}
