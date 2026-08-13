package projects

import (
	"context"
	"database/sql"

	"github.com/oklog/ulid/v2"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, name, description string) (Project, error) {
	//generate ULID
	newId := ulid.Make().String()
	query := `
		INSERT INTO projects(id, name, description, created_at, updated_at)
		VALUES($1,$2,$3,NOW(),NOW())
		RETURNING created_at, updated_at
	`
	p := Project{
		ID:          newId,
		Name:        name,
		Description: description,
	}
	if err := r.db.QueryRowContext(ctx, query, newId, name, description).Scan(&p.CreatedAt, &p.UpdatedAt); err != nil {
		return Project{}, err
	}
	return p, nil
}

/*
fun only for debug

	func (r *Repository) ListAll(ctx context.Context) ([]Project, error) {
		query := "SELECT id, name, description, created_at, updated_at FROM projects"
		rows, err := r.db.QueryContext(ctx, query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var projects []Project
		for rows.Next() {
			var p Project
			if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt); err != nil {
				return nil, err
			}
			projects = append(projects, p)
		}
		return projects, nil
	}
*/
func (r *Repository) FindById(ctx context.Context, id string) (Project, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM projects WHERE id = $1"
	var p Project
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return Project{}, err
	}
	return p, nil
}
func (r *Repository) Update(ctx context.Context, id, name, description string) (Project, error) {
	query := `
		UPDATE projects
		SET name = $1, description = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING id, name, description, created_at, updated_at
	`
	p := Project{
		ID:          id,
		Name:        name,
		Description: description,
	}
	if err := r.db.QueryRowContext(ctx, query, name, description, id).Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return Project{}, err
	}
	return p, nil
}
