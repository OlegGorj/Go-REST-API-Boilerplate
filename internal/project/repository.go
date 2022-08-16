package project

import (
	"context"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/internal/entity"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/pkg/dbcontext"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/pkg/log"
)

// Repository encapsulates the logic to access projects from the data source.
type Repository interface {
	// Get returns the Project with the specified Project ID.
	Get(ctx context.Context, id string) (entity.Project, error)
	// Count returns the number of projects.
	Count(ctx context.Context) (int, error)
	// Query returns the list of projects with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Project, error)
	// Create saves a new Project in the storage.
	Create(ctx context.Context, Project entity.Project) error
	// Update updates the Project with given ID in the storage.
	Update(ctx context.Context, Project entity.Project) error
	// Delete removes the Project with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists projects in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new Project repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the Project with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Project, error) {
	var _project entity.Project
	err := r.db.With(ctx).Select().Model(id, &_project)
	return _project, err
}

// Create saves a new Project record in the database.
// It returns the ID of the newly inserted Project record.
func (r repository) Create(ctx context.Context, Project entity.Project) error {
	return r.db.With(ctx).Model(&Project).Insert()
}

// Update saves the changes to an Project in the database.
func (r repository) Update(ctx context.Context, Project entity.Project) error {
	return r.db.With(ctx).Model(&Project).Update()
}

// Delete deletes an Project with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	_project, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&_project).Delete()
}

// Count returns the number of the Project records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("project").Row(&count)
	return count, err
}

// Query retrieves the Project records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Project, error) {
	var _projects []entity.Project
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&_projects)
	return _projects, err
}
