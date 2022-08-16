package health

import (
	"encoding/json"
	"fmt"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/pkg/log"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

var logger = log.New()

// var flagConfig = flag.String("config", "./config/local.yml", "path to the config file")
// var cfg, _ = config.Load(*flagConfig, logger)

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
	return &Health{
		API_version: "1.0.0",
		API_health:  "ok",
		DB_health:   "ok",
	}
}

// health responds to a health request.
func health(version string) routing.Handler {
	return func(c *routing.Context) error {
		marshalled, err := json.Marshal(healthChecks())
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Printf("marshalled: %s \n", string(marshalled))
		return c.WriteWithStatus(string(marshalled), 200)
	}
}
