package health

import (
	"github.com/OlegGorj/Go-REST-API-Boilerplate/internal/global"
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
	db, err := global.ConnectDB()
	if err == nil {
		db_connect = true
	}
	defer func() {
		if err := db.Close(); err != nil {
			global.Logger.Error(err)
		}
	}()

	return &Health{
		API_version: global.Version,
		API_health:  "true",
		DB_health:   strconv.FormatBool(db_connect),
	}
}

// health responds to a health request.
func health(version string) routing.Handler {
	return func(c *routing.Context) error {
		return c.Write(healthChecks())
	}
}
