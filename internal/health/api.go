package health


mport routing "github.com/go-ozzo/ozzo-routing/v2"

// RegisterHandlers registers the handlers that perform health checks.
func RegisterHandlers(r *routing.Router, version string) {
	r.To("GET,HEAD", "/health", health(version))
}

// healthcheck responds to a healthcheck request.
func health(version string) routing.Handler {
	return func(c *routing.Context) error {
		return c.Write("OK " + version)
	}
}
