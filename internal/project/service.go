package project

import (
	"context"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/internal/entity"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/pkg/log"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

// Service encapsulates usecase logic for Projects.
type Service interface {
	Get(ctx context.Context, id string) (Project, error)
	Query(ctx context.Context, offset, limit int) ([]Project, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateProjectRequest) (Project, error)
	Update(ctx context.Context, id string, input UpdateProjectRequest) (Project, error)
	Delete(ctx context.Context, id string) (Project, error)
}

// Project represents the data about an Project.
type Project struct {
	entity.Project
}

// CreateProjectRequest represents an Project creation request.
type CreateProjectRequest struct {
	Name string `json:"name"`
}

// Validate validates the CreateProjectRequest fields.
func (m CreateProjectRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateProjectRequest represents an Project update request.
type UpdateProjectRequest struct {
	Name string `json:"name"`
}

// Validate validates the CreateProjectRequest fields.
func (m UpdateProjectRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new Project service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the Project with the specified the Project ID.
func (s service) Get(ctx context.Context, id string) (Project, error) {
	_project, err := s.repo.Get(ctx, id)
	if err != nil {
		return Project{}, err
	}
	return Project{_project}, nil
}

// Create creates a new Project.
func (s service) Create(ctx context.Context, req CreateProjectRequest) (Project, error) {
	if err := req.Validate(); err != nil {
		return Project{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.Project{
		ID:        id,
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return Project{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the Project with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateProjectRequest) (Project, error) {
	if err := req.Validate(); err != nil {
		return Project{}, err
	}

	_project, err := s.Get(ctx, id)
	if err != nil {
		return _project, err
	}
	_project.Name = req.Name
	_project.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, _project.Project); err != nil {
		return _project, err
	}
	return _project, nil
}

// Delete deletes the Project with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Project, error) {
	_project, err := s.Get(ctx, id)
	if err != nil {
		return Project{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Project{}, err
	}
	return _project, nil
}

// Count returns the number of Projects.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the Projects with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Project, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Project{}
	for _, item := range items {
		result = append(result, Project{item})
	}
	return result, nil
}
