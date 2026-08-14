package projects

import "time"

type Column struct {
	ID          string    `db:"id"`
	ProjectID   string    `db:"project_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Color       string    `db:"color"`
	Position    int       `db:"position"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
