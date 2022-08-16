package project

import (
	"fmt"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/internal/errors"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/pkg/log"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/pkg/pagination"
	"github.com/go-ozzo/ozzo-routing/v2"
	"net/http"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{
		service,
		logger,
	}

	r.Get("/projects/<id>", res.get)
	r.Get("/projects", res.query)

	r.Use(authHandler)

	// the following endpoints require a valid JWT
	r.Post("/projects", res.create)
	r.Put("/projects/<id>", res.update)
	r.Delete("/projects/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	_project, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(_project)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	_projects, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	fmt.Printf("DEBUG: _projects: %+v \n", _projects)
	pages.Items = _projects
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateProjectRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	_project, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(_project, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateProjectRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	Project, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(Project)
}

func (r resource) delete(c *routing.Context) error {
	Project, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(Project)
}
