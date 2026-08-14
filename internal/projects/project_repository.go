package projects

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/oklog/ulid/v2"
)

type ProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(ctx context.Context, p *Project) (Project, error) {
	p.ID = ulid.Make().String()
	query := `
		INSERT INTO projects(id, name, description, color, created_at, updated_at)
		VALUES(:id,:name,:description,:color,NOW(),NOW())
		RETURNING id, name, description, color, created_at, updated_at
	`
	rows, err := r.db.NamedQueryContext(ctx, query, p)
	if err != nil {
		return Project{}, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(p); err != nil {
			return Project{}, err
		}
	}
	return *p, nil
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
func (r *ProjectRepository) FindById(ctx context.Context, id string) (Project, error) {
	query := "SELECT id, name, description, color, created_at, updated_at FROM projects WHERE id = $1"
	var p Project
	if err := r.db.QueryRowxContext(ctx, query, id).StructScan(&p); err != nil {
		return Project{}, err
	}
	return p, nil
}
func (r *ProjectRepository) Update(ctx context.Context, p *Project) (Project, error) {
	query := `
		UPDATE projects
		SET name = :name, description = :description, color = :color, updated_at = NOW()
		WHERE id = :id
		RETURNING id, name, description, color, created_at, updated_at
	`
	rows, err := r.db.NamedQueryContext(ctx, query, p)
	if err != nil {
		return Project{}, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(p); err != nil {
			return Project{}, err
		}
	}
	return *p, nil
}
