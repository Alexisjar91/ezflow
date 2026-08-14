package projects

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/oklog/ulid/v2"
)

type ColumnRepository struct {
	db *sqlx.DB
}

func NewColumnRepository(db *sqlx.DB) *ColumnRepository {
	return &ColumnRepository{db: db}
}

func (r *ColumnRepository) Create(ctx context.Context, c *Column) (Column, error) {
	c.ID = ulid.Make().String()
	lastPosition, err := r.getLastPosition(ctx, c.ProjectID)
	if err != nil {
		return Column{}, err
	}

	c.Position = lastPosition + 1000

	query := `
		INSERT INTO columns(id, name, description, color, position, project_id, created_at, updated_at)
		VALUES(:id,:name,:description,:color,:position,:project_id,NOW(),NOW())
		RETURNING id, name, description, color, position, project_id, created_at, updated_at
	`
	rows, err := r.db.NamedQueryContext(ctx, query, c)
	if err != nil {
		return Column{}, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(c); err != nil {
			return Column{}, err
		}
	}
	return *c, nil
}

func (r *ColumnRepository) Update(ctx context.Context, c *Column) (Column, error) {
	query := `
		UPDATE columns
		SET name = :name, description = :description, color = :color, position = :position, updated_at = NOW()
		WHERE id = :id
		RETURNING id, name, description, color, position, project_id, created_at, updated_at
	`
	rows, err := r.db.NamedQueryContext(ctx, query, c)
	if err != nil {
		return Column{}, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(c); err != nil {
			return Column{}, err
		}
	}
	return *c, nil
}

func (r *ColumnRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM columns WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *ColumnRepository) getLastPosition(ctx context.Context, projectID string) (int, error) {
	query := "SELECT COALESCE(MAX(position), 0) FROM columns WHERE project_id = $1"
	var lastPosition int
	if err := r.db.QueryRowContext(ctx, query, projectID).Scan(&lastPosition); err != nil {
		return 0, err
	}
	return lastPosition, nil
}
